package handlers

import (
	"flowmoney/bot/internal/models"
	"fmt"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

const budgetHelp = "📊 <b>Бюджет</b>\n\n" +
	"Создать: /budget_new ID_категории Сумма Месяц Год\n" +
	"По ID: /budget_get ID\n" +
	"По категории: /budget_by_category ID_категории\n" +
	"За месяц: /budget_by_month Месяц Год\n" +
	"Изменить: /budget_update ID Сумма\n" +
	"Удалить: /budget_delete ID"

type BudgetHandlers struct{ Deps }

func NewBudgetHandlers(d Deps) *BudgetHandlers { return &BudgetHandlers{d} }

func (h *BudgetHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtBudget, h.help)
	bot.Handle("/budget_new", h.create)
	bot.Handle("/budget_get", h.get)
	bot.Handle("/budget_by_category", h.getByCategory)
	bot.Handle("/budget_by_month", h.getByMonth)
	bot.Handle("/budget_update", h.update)
	bot.Handle("/budget_delete", h.delete)
}

func (h *BudgetHandlers) help(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}
	return c.Send(budgetHelp)
}

func (h *BudgetHandlers) create(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 4 {
		return c.Send("⚠️ Использование: /budget_new ID_категории Сумма Месяц Год")
	}
	categoryId, err1 := strconv.Atoi(args[0])
	amount, err2 := strconv.ParseFloat(args[1], 64)
	month, err3 := strconv.Atoi(args[2])
	year, err4 := strconv.Atoi(args[3])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return c.Send("⚠️ Проверьте формат: /budget_new ID_категории Сумма Месяц Год")
	}

	budget, err := acc.Client.CreateBudget(acc.UserId, categoryId, amount, month, year)
	if err != nil {
		return h.fail(c, "не удалось создать бюджет", err)
	}
	return c.Send(fmt.Sprintf("✅ Бюджет <b>#%d</b> на <b>%.2f</b> создан (%02d.%d).", budget.Id, budget.Amount, budget.Month, budget.Year))
}

func (h *BudgetHandlers) get(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 1 {
		return c.Send("⚠️ Использование: /budget_get ID")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Send("⚠️ ID должен быть числом.")
	}

	budget, err := acc.Client.GetBudgetById(id)
	if err != nil {
		return h.fail(c, "не удалось найти бюджет", err)
	}
	return c.Send(formatBudget(budget))
}

func (h *BudgetHandlers) getByCategory(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 1 {
		return c.Send("⚠️ Использование: /budget_by_category ID_категории")
	}
	categoryId, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Send("⚠️ ID должен быть числом.")
	}

	budget, err := acc.Client.GetBudgetByCategoryId(categoryId)
	if err != nil {
		return h.fail(c, "не удалось найти бюджет по категории", err)
	}
	return c.Send(formatBudget(budget))
}

func (h *BudgetHandlers) getByMonth(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 2 {
		return c.Send("⚠️ Использование: /budget_by_month Месяц Год")
	}
	month, err1 := strconv.Atoi(args[0])
	year, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		return c.Send("⚠️ Месяц и год должны быть числами.")
	}

	budget, err := acc.Client.GetByUserIdAndMonth(acc.UserId, month, year)
	if err != nil {
		return h.fail(c, "не удалось получить бюджет за месяц", err)
	}
	return c.Send(formatBudget(budget))
}

func (h *BudgetHandlers) update(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 2 {
		return c.Send("⚠️ Использование: /budget_update ID Сумма")
	}
	id, err1 := strconv.Atoi(args[0])
	amount, err2 := strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		return c.Send("⚠️ Проверьте формат: /budget_update ID Сумма")
	}

	budget, err := acc.Client.UpdateBudget(amount, id)
	if err != nil {
		return h.fail(c, "не удалось изменить бюджет", err)
	}
	return c.Send(fmt.Sprintf("✅ Бюджет <b>#%d</b> обновлён: <b>%.2f</b>", budget.Id, budget.Amount))
}

func (h *BudgetHandlers) delete(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 1 {
		return c.Send("⚠️ Использование: /budget_delete ID")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Send("⚠️ ID должен быть числом.")
	}

	if err := acc.Client.DeleteBudgetById(id); err != nil {
		return h.fail(c, "не удалось удалить бюджет", err)
	}
	return c.Send(fmt.Sprintf("✅ Бюджет <b>#%d</b> удалён.", id))
}

func formatBudget(b models.Budget) string {
	return fmt.Sprintf(
		"📊 <b>Бюджет #%d</b>\n\n🏷 Категория: <b>%d</b>\n💰 Сумма: <b>%.2f</b>\n📅 Период: %02d.%d",
		b.Id, b.CategoryId, b.Amount, b.Month, b.Year,
	)
}
