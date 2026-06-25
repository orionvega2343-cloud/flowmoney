package handlers

import (
	"errors"
	"flowmoney/bot/internal/models"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

var (
	btnBudgetNew      = tele.Btn{Unique: "bg_new"}
	btnBudgetGet      = tele.Btn{Unique: "bg_get"}
	btnBudgetByCat    = tele.Btn{Unique: "bg_by_cat"}
	btnBudgetByMonth  = tele.Btn{Unique: "bg_by_month"}
	btnBudgetUpdate   = tele.Btn{Unique: "bg_update"}
	btnBudgetDelete   = tele.Btn{Unique: "bg_delete"}
	btnBudgetNewCat   = tele.Btn{Unique: "bg_new_cat"}
	btnBudgetByCatPck = tele.Btn{Unique: "bg_by_cat_pick"}
)

type BudgetHandlers struct{ Deps }

func NewBudgetHandlers(d Deps) *BudgetHandlers { return &BudgetHandlers{d} }

func (h *BudgetHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtBudget, h.menu)
	bot.Handle(&btnBudgetNew, h.startNew)
	bot.Handle(&btnBudgetNewCat, h.pickNewCategory)
	bot.Handle(&btnBudgetGet, h.startGet)
	bot.Handle(&btnBudgetByCat, h.startByCategory)
	bot.Handle(&btnBudgetByCatPck, h.pickByCategory)
	bot.Handle(&btnBudgetByMonth, h.startByMonth)
	bot.Handle(&btnBudgetUpdate, h.startUpdate)
	bot.Handle(&btnBudgetDelete, h.startDelete)
}

func (h *BudgetHandlers) menu(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	kb := &tele.ReplyMarkup{}
	kb.Inline(
		tele.Row{{Unique: "bg_new", Text: "➕ Новый бюджет"}},
		tele.Row{
			{Unique: "bg_get", Text: "🔎 По ID"},
			{Unique: "bg_by_cat", Text: "🏷 По категории"},
		},
		tele.Row{{Unique: "bg_by_month", Text: "📅 За месяц"}},
		tele.Row{
			{Unique: "bg_update", Text: "✏️ Изменить"},
			{Unique: "bg_delete", Text: "🗑 Удалить"},
		},
	)
	return c.Send("📊 <b>Бюджет</b>", kb)
}

// categoryPickerMarkup строит инлайн-список категорий пользователя для
// выбора — нажатие на категорию продолжит сценарий с заданным unique.
func (h *BudgetHandlers) categoryPickerMarkup(acc *Account, unique string) (*tele.ReplyMarkup, error) {
	categories, err := acc.Client.GetByUserId(acc.UserId)
	if err != nil {
		return nil, err
	}
	rows := make([]tele.Row, 0, len(categories))
	for _, cat := range categories {
		rows = append(rows, tele.Row{{Unique: unique, Text: fmt.Sprintf("🏷 #%d %s", cat.Id, cat.Title), Data: strconv.Itoa(cat.Id)}})
	}
	kb := &tele.ReplyMarkup{}
	kb.Inline(rows...)
	return kb, nil
}

func (h *BudgetHandlers) startNew(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	kb, err := h.categoryPickerMarkup(acc, "bg_new_cat")
	if err != nil {
		return h.fail(c, "не удалось получить категории", err)
	}
	return c.Send("➕ <b>Новый бюджет</b>\n\nВыберите категорию:", kb)
}

func (h *BudgetHandlers) pickNewCategory(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	categoryId, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "Введите сумму бюджета:",
		Next: func(reply string) StepResult {
			amount, err := strconv.ParseFloat(strings.TrimSpace(reply), 64)
			if err != nil {
				return fail(errors.New("сумма должна быть числом"))
			}
			return ask("Введите месяц (1-12):", func(reply string) StepResult {
				month, err := strconv.Atoi(strings.TrimSpace(reply))
				if err != nil {
					return fail(errors.New("месяц должен быть числом от 1 до 12"))
				}
				return ask("Введите год:", func(reply string) StepResult {
					year, err := strconv.Atoi(strings.TrimSpace(reply))
					if err != nil {
						return fail(errors.New("год должен быть числом"))
					}
					budget, err := acc.Client.CreateBudget(acc.UserId, categoryId, amount, month, year)
					if err != nil {
						return fail(err)
					}
					return done(fmt.Sprintf("✅ Бюджет <b>#%d</b> на <b>%.2f</b> создан (%02d.%d).", budget.Id, budget.Amount, budget.Month, budget.Year))
				})
			})
		},
	})
}

func (h *BudgetHandlers) startGet(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "🔎 Введите ID бюджета:",
		Next: func(reply string) StepResult {
			id, err := strconv.Atoi(strings.TrimSpace(reply))
			if err != nil {
				return fail(errors.New("ID должен быть числом"))
			}
			budget, err := acc.Client.GetBudgetById(id)
			if err != nil {
				return fail(err)
			}
			return done(formatBudget(budget))
		},
	})
}

func (h *BudgetHandlers) startByCategory(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	kb, err := h.categoryPickerMarkup(acc, "bg_by_cat_pick")
	if err != nil {
		return h.fail(c, "не удалось получить категории", err)
	}
	return c.Send("🏷 Выберите категорию:", kb)
}

func (h *BudgetHandlers) pickByCategory(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	categoryId, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	budget, err := acc.Client.GetBudgetByCategoryId(categoryId)
	if err != nil {
		return h.fail(c, "не удалось найти бюджет по категории", err)
	}
	return c.Send(formatBudget(budget))
}

func (h *BudgetHandlers) startByMonth(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "📅 Введите месяц (1-12):",
		Next: func(reply string) StepResult {
			month, err := strconv.Atoi(strings.TrimSpace(reply))
			if err != nil {
				return fail(errors.New("месяц должен быть числом от 1 до 12"))
			}
			return ask("Введите год:", func(reply string) StepResult {
				year, err := strconv.Atoi(strings.TrimSpace(reply))
				if err != nil {
					return fail(errors.New("год должен быть числом"))
				}
				budget, err := acc.Client.GetByUserIdAndMonth(acc.UserId, month, year)
				if err != nil {
					return fail(err)
				}
				return done(formatBudget(budget))
			})
		},
	})
}

func (h *BudgetHandlers) startUpdate(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "✏️ Введите ID бюджета, который нужно изменить:",
		Next: func(reply string) StepResult {
			id, err := strconv.Atoi(strings.TrimSpace(reply))
			if err != nil {
				return fail(errors.New("ID должен быть числом"))
			}
			return ask("Введите новую сумму бюджета:", func(reply string) StepResult {
				amount, err := strconv.ParseFloat(strings.TrimSpace(reply), 64)
				if err != nil {
					return fail(errors.New("сумма должна быть числом"))
				}
				budget, err := acc.Client.UpdateBudget(amount, id)
				if err != nil {
					return fail(err)
				}
				return done(fmt.Sprintf("✅ Бюджет <b>#%d</b> обновлён: <b>%.2f</b>", budget.Id, budget.Amount))
			})
		},
	})
}

func (h *BudgetHandlers) startDelete(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "🗑 Введите ID бюджета, который нужно удалить:",
		Next: func(reply string) StepResult {
			id, err := strconv.Atoi(strings.TrimSpace(reply))
			if err != nil {
				return fail(errors.New("ID должен быть числом"))
			}
			if err := acc.Client.DeleteBudgetById(id); err != nil {
				return fail(err)
			}
			return done(fmt.Sprintf("✅ Бюджет <b>#%d</b> удалён.", id))
		},
	})
}

func formatBudget(b models.Budget) string {
	return fmt.Sprintf(
		"📊 <b>Бюджет #%d</b>\n\n🏷 Категория: <b>%d</b>\n💰 Сумма: <b>%.2f</b>\n📅 Период: %02d.%d",
		b.Id, b.CategoryId, b.Amount, b.Month, b.Year,
	)
}
