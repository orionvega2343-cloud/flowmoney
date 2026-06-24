package handlers

import (
	"flowmoney/bot/internal/models"
	"fmt"
	"strconv"
)

type BudgetHandler interface {
	StartCreate(session *Session, chatId int64)
	StartGet(session *Session, chatId int64)
	StartGetByCategory(session *Session, chatId int64)
	StartGetByMonth(session *Session, chatId int64)
	StartUpdate(session *Session, chatId int64)
	StartDelete(session *Session, chatId int64)
	HandleText(session *Session, chatId int64, text string)
}

type BudgetHandlerImpl struct {
	Deps
}

func NewBudgetHandlerImpl(d Deps) *BudgetHandlerImpl {
	return &BudgetHandlerImpl{Deps: d}
}

func (h *BudgetHandlerImpl) StartCreate(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetCreateCategoryId, "➕ <b>Новый бюджет</b>\n\nВведите ID категории:")
}

func (h *BudgetHandlerImpl) StartGet(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetGetId, "🔎 Введите ID бюджета:")
}

func (h *BudgetHandlerImpl) StartGetByCategory(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetCategoryId, "🏷 Введите ID категории:")
}

func (h *BudgetHandlerImpl) StartGetByMonth(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetMonthMonth, "📅 Введите месяц (1-12):")
}

func (h *BudgetHandlerImpl) StartUpdate(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetUpdateId, "✏️ Введите ID бюджета, который нужно изменить:")
}

func (h *BudgetHandlerImpl) StartDelete(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepBudgetDeleteId, "🗑 Введите ID бюджета, который нужно удалить:")
}

func (h *BudgetHandlerImpl) HandleText(session *Session, chatId int64, text string) {
	switch session.Step {
	case StepBudgetCreateCategoryId:
		if _, err := strconv.Atoi(text); err != nil {
			h.send(chatId, "⚠️ ID категории должен быть числом, попробуйте ещё раз:", nil)
			return
		}
		session.Data["category_id"] = text
		h.ask(session, chatId, StepBudgetCreateAmount, "Введите сумму бюджета:")
	case StepBudgetCreateAmount:
		if _, err := strconv.ParseFloat(text, 64); err != nil {
			h.send(chatId, "⚠️ Сумма должна быть числом, попробуйте ещё раз:", nil)
			return
		}
		session.Data["amount"] = text
		h.ask(session, chatId, StepBudgetCreateMonth, "Введите месяц (1-12):")
	case StepBudgetCreateMonth:
		if _, err := strconv.Atoi(text); err != nil {
			h.send(chatId, "⚠️ Месяц должен быть числом от 1 до 12, попробуйте ещё раз:", nil)
			return
		}
		session.Data["month"] = text
		h.ask(session, chatId, StepBudgetCreateYear, "Введите год:")
	case StepBudgetCreateYear:
		h.finishCreate(session, chatId, text)
	case StepBudgetGetId:
		h.finishGet(session, chatId, text)
	case StepBudgetCategoryId:
		h.finishGetByCategory(session, chatId, text)
	case StepBudgetMonthMonth:
		if _, err := strconv.Atoi(text); err != nil {
			h.send(chatId, "⚠️ Месяц должен быть числом от 1 до 12, попробуйте ещё раз:", nil)
			return
		}
		session.Data["month"] = text
		h.ask(session, chatId, StepBudgetMonthYear, "Введите год:")
	case StepBudgetMonthYear:
		h.finishGetByMonth(session, chatId, text)
	case StepBudgetUpdateId:
		if _, err := strconv.Atoi(text); err != nil {
			h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
			return
		}
		session.Data["id"] = text
		h.ask(session, chatId, StepBudgetUpdateAmount, "Введите новую сумму бюджета:")
	case StepBudgetUpdateAmount:
		h.finishUpdate(session, chatId, text)
	case StepBudgetDeleteId:
		h.finishDelete(session, chatId, text)
	}
}

func (h *BudgetHandlerImpl) finishCreate(session *Session, chatId int64, yearText string) {
	year, err := strconv.Atoi(yearText)
	if err != nil {
		h.send(chatId, "⚠️ Год должен быть числом, попробуйте ещё раз:", nil)
		return
	}

	categoryId, _ := strconv.Atoi(session.Data["category_id"])
	amount, _ := strconv.ParseFloat(session.Data["amount"], 64)
	month, _ := strconv.Atoi(session.Data["month"])
	session.Reset()

	budget, err := session.Client.CreateBudget(session.UserId, categoryId, amount, month, year)
	if err != nil {
		h.fail(chatId, "не удалось создать бюджет", err, budgetMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Бюджет <b>#%d</b> на <b>%.2f</b> создан (%02d.%d).", budget.Id, budget.Amount, budget.Month, budget.Year), kbPtr(budgetMenu()))
}

func (h *BudgetHandlerImpl) finishGet(session *Session, chatId int64, text string) {
	id, err := strconv.Atoi(text)
	if err != nil {
		h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
		return
	}
	session.Reset()

	budget, err := session.Client.GetBudgetById(id)
	if err != nil {
		h.fail(chatId, "не удалось найти бюджет", err, budgetMenu())
		return
	}

	h.send(chatId, formatBudget(budget), kbPtr(budgetMenu()))
}

func (h *BudgetHandlerImpl) finishGetByCategory(session *Session, chatId int64, text string) {
	categoryId, err := strconv.Atoi(text)
	if err != nil {
		h.send(chatId, "⚠️ ID категории должен быть числом, попробуйте ещё раз:", nil)
		return
	}
	session.Reset()

	budget, err := session.Client.GetBudgetByCategoryId(categoryId)
	if err != nil {
		h.fail(chatId, "не удалось найти бюджет по категории", err, budgetMenu())
		return
	}

	h.send(chatId, formatBudget(budget), kbPtr(budgetMenu()))
}

func (h *BudgetHandlerImpl) finishGetByMonth(session *Session, chatId int64, yearText string) {
	year, err := strconv.Atoi(yearText)
	if err != nil {
		h.send(chatId, "⚠️ Год должен быть числом, попробуйте ещё раз:", nil)
		return
	}

	month, _ := strconv.Atoi(session.Data["month"])
	session.Reset()

	budget, err := session.Client.GetByUserIdAndMonth(session.UserId, month, year)
	if err != nil {
		h.fail(chatId, "не удалось получить бюджет за месяц", err, budgetMenu())
		return
	}

	h.send(chatId, formatBudget(budget), kbPtr(budgetMenu()))
}

func (h *BudgetHandlerImpl) finishUpdate(session *Session, chatId int64, amountText string) {
	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		h.send(chatId, "⚠️ Сумма должна быть числом, попробуйте ещё раз:", nil)
		return
	}

	id, _ := strconv.Atoi(session.Data["id"])
	session.Reset()

	budget, err := session.Client.UpdateBudget(amount, id)
	if err != nil {
		h.fail(chatId, "не удалось изменить бюджет", err, budgetMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Бюджет <b>#%d</b> обновлён: <b>%.2f</b>", budget.Id, budget.Amount), kbPtr(budgetMenu()))
}

func (h *BudgetHandlerImpl) finishDelete(session *Session, chatId int64, text string) {
	id, err := strconv.Atoi(text)
	if err != nil {
		h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
		return
	}
	session.Reset()

	if err := session.Client.DeleteBudgetById(id); err != nil {
		h.fail(chatId, "не удалось удалить бюджет", err, budgetMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Бюджет <b>#%d</b> удалён.", id), kbPtr(budgetMenu()))
}

func formatBudget(b models.Budget) string {
	return fmt.Sprintf(
		"📊 <b>Бюджет #%d</b>\n\n🏷 Категория: <b>%d</b>\n💰 Сумма: <b>%.2f</b>\n📅 Период: %02d.%d",
		b.Id, b.CategoryId, b.Amount, b.Month, b.Year,
	)
}
