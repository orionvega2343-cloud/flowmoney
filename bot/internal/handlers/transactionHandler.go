package handlers

import (
	"fmt"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

// btnTransactionGet — шаблон инлайн-кнопки для конкретной транзакции в
// списке; нажатие сразу вызывает GetTransactionById с ID из callback-данных.
var btnTransactionGet = tele.Btn{Unique: "tx_get"}

type TransactionHandlers struct{ Deps }

func NewTransactionHandlers(d Deps) *TransactionHandlers { return &TransactionHandlers{d} }

func (h *TransactionHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtTransactions, h.list)
	bot.Handle("/income", h.income)
	bot.Handle("/expense", h.expense)
	bot.Handle(&btnTransactionGet, h.get)
}

func (h *TransactionHandlers) income(c tele.Context) error  { return h.create(c, "income") }
func (h *TransactionHandlers) expense(c tele.Context) error { return h.create(c, "expense") }

func (h *TransactionHandlers) create(c tele.Context, txType string) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 2 {
		return c.Send(fmt.Sprintf("⚠️ Использование: /%s Сумма ID_категории", txType))
	}
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return c.Send("⚠️ Сумма должна быть числом.")
	}
	categoryId, err := strconv.Atoi(args[1])
	if err != nil {
		return c.Send("⚠️ ID категории должен быть числом.")
	}

	transaction, err := acc.Client.CreateTransaction(acc.UserId, amount, txType, time.Now(), categoryId)
	if err != nil {
		return h.fail(c, "не удалось создать транзакцию", err)
	}

	return c.Send(fmt.Sprintf("✅ %s Транзакция <b>#%d</b> на сумму <b>%.2f</b> сохранена.", icon(txType), transaction.Id, transaction.Amount))
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
	if len(transactions) == 0 {
		return c.Send("💸 У вас пока нет транзакций.\n\nДоход: /income Сумма ID_категории\nРасход: /expense Сумма ID_категории")
	}

	rows := make([]tele.Row, 0, len(transactions))
	for _, t := range transactions {
		label := fmt.Sprintf("%s #%d %.2f — %s", icon(t.Type), t.Id, t.Amount, t.Date.Format("02.01.2006"))
		rows = append(rows, tele.Row{{Unique: "tx_get", Text: label, Data: strconv.Itoa(t.Id)}})
	}
	kb := &tele.ReplyMarkup{}
	kb.Inline(rows...)

	text := "📋 <b>Ваши транзакции</b>\n\nДоход: /income Сумма ID_категории\nРасход: /expense Сумма ID_категории"
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
