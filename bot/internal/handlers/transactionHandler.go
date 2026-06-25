package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	// btnTransactionGet — нажатие на транзакцию в списке вызывает
	// GetTransactionById с ID из callback-данных.
	btnTransactionGet = tele.Btn{Unique: "tx_get"}
	btnTxIncome       = tele.Btn{Unique: "tx_income"}
	btnTxExpense      = tele.Btn{Unique: "tx_expense"}
	// btnTxIncomeCategory / btnTxExpenseCategory — выбор категории из списка
	// для новой транзакции; ID категории приходит в callback-данных.
	btnTxIncomeCategory  = tele.Btn{Unique: "tx_income_cat"}
	btnTxExpenseCategory = tele.Btn{Unique: "tx_expense_cat"}
)

type TransactionHandlers struct{ Deps }

func NewTransactionHandlers(d Deps) *TransactionHandlers { return &TransactionHandlers{d} }

func (h *TransactionHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtTransactions, h.list)
	bot.Handle(&btnTxIncome, h.startIncome)
	bot.Handle(&btnTxExpense, h.startExpense)
	bot.Handle(&btnTxIncomeCategory, h.pickIncomeCategory)
	bot.Handle(&btnTxExpenseCategory, h.pickExpenseCategory)
	bot.Handle(&btnTransactionGet, h.get)
}

func (h *TransactionHandlers) startIncome(c tele.Context) error  { return h.startCreate(c, "income") }
func (h *TransactionHandlers) startExpense(c tele.Context) error { return h.startCreate(c, "expense") }

// startCreate показывает список категорий, чтобы выбрать, к какой отнести
// новую транзакцию - нажатие на категорию продолжит создание.
func (h *TransactionHandlers) startCreate(c tele.Context, txType string) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	categories, err := acc.Client.GetByUserId(acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось получить категории", err)
	}
	if len(categories) == 0 {
		return c.Send("📁 Сначала создайте категорию в разделе «Категории».")
	}

	unique := "tx_income_cat"
	if txType == "expense" {
		unique = "tx_expense_cat"
	}
	kb := &tele.ReplyMarkup{}
	rows := make([]tele.Row, 0, len(categories))
	for _, cat := range categories {
		rows = append(rows, tele.Row{{Unique: unique, Text: fmt.Sprintf("🏷 #%d %s", cat.Id, cat.Title), Data: strconv.Itoa(cat.Id)}})
	}
	kb.Inline(rows...)

	return c.Send("Выберите категорию:", kb)
}

func (h *TransactionHandlers) pickIncomeCategory(c tele.Context) error {
	return h.startAmount(c, "income")
}

func (h *TransactionHandlers) pickExpenseCategory(c tele.Context) error {
	return h.startAmount(c, "expense")
}

func (h *TransactionHandlers) startAmount(c tele.Context, txType string) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	categoryId, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: fmt.Sprintf("%s Введите сумму:", icon(txType)),
		Next: func(reply string) StepResult {
			amount, err := strconv.ParseFloat(strings.TrimSpace(reply), 64)
			if err != nil {
				return fail(errors.New("сумма должна быть числом"))
			}

			transaction, err := acc.Client.CreateTransaction(acc.UserId, amount, txType, time.Now(), categoryId)
			if err != nil {
				return fail(err)
			}
			return done(fmt.Sprintf("✅ %s Транзакция <b>#%d</b> на сумму <b>%.2f</b> сохранена.", icon(txType), transaction.Id, transaction.Amount))
		},
	})
}

func (h *TransactionHandlers) list(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	transactions, err := acc.Client.GetTransactionByUserId(acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось получить транзакции", err)
	}

	kb := &tele.ReplyMarkup{}
	rows := []tele.Row{{
		{Unique: "tx_income", Text: "💵 Новый доход"},
		{Unique: "tx_expense", Text: "💳 Новый расход"},
	}}
	for _, t := range transactions {
		label := fmt.Sprintf("%s #%d %.2f — %s", icon(t.Type), t.Id, t.Amount, t.Date.Format("02.01.2006"))
		rows = append(rows, tele.Row{{Unique: "tx_get", Text: label, Data: strconv.Itoa(t.Id)}})
	}
	kb.Inline(rows...)

	text := "📋 <b>Ваши транзакции</b>"
	if len(transactions) == 0 {
		text = "💸 У вас пока нет транзакций."
	}
	return c.Send(text, kb)
}

func (h *TransactionHandlers) get(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	id, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	transaction, err := acc.Client.GetTransactionById(id)
	if err != nil {
		return h.fail(c, "не удалось найти транзакцию", err)
	}
	return c.Send(fmt.Sprintf("%s <b>#%d</b> %.2f — %s", icon(transaction.Type), transaction.Id, transaction.Amount, transaction.Date.Format("02.01.2006")))
}

func icon(txType string) string {
	if txType == "expense" {
		return "💳"
	}
	return "💵"
}
