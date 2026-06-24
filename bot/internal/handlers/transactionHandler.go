package handlers

import (
	"fmt"
	"strconv"
	"time"
)

type TransactionHandler interface {
	StartCreate(session *Session, chatId int64, txType string)
	StartGet(session *Session, chatId int64)
	List(session *Session, chatId int64)
	HandleText(session *Session, chatId int64, text string)
}

type TransactionHandlerImpl struct {
	Deps
}

func NewTransactionHandlerImpl(d Deps) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{Deps: d}
}

func (h *TransactionHandlerImpl) StartCreate(session *Session, chatId int64, txType string) {
	if !h.requireLogin(session, chatId) {
		return
	}
	session.Data["type"] = txType

	title := "💵 <b>Новый доход</b>"
	if txType == "expense" {
		title = "💳 <b>Новый расход</b>"
	}
	h.ask(session, chatId, StepTransactionAmount, title+"\n\nВведите сумму:")
}

func (h *TransactionHandlerImpl) StartGet(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepTransactionGetId, "🔎 Введите ID транзакции:")
}

func (h *TransactionHandlerImpl) List(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}

	transactions, err := session.Client.GetTransactionByUserId(session.UserId)
	if err != nil {
		h.fail(chatId, "не удалось получить транзакции", err, transactionMenu())
		return
	}

	if len(transactions) == 0 {
		h.send(chatId, "💸 У вас пока нет транзакций.", kbPtr(transactionMenu()))
		return
	}

	text := "📋 <b>Ваши транзакции:</b>\n\n"
	for _, t := range transactions {
		icon := "💵"
		if t.Type == "expense" {
			icon = "💳"
		}
		text += fmt.Sprintf("%s <b>#%d</b> %.2f — %s\n", icon, t.Id, t.Amount, t.Date.Format("02.01.2006"))
	}
	h.send(chatId, text, kbPtr(transactionMenu()))
}

func (h *TransactionHandlerImpl) HandleText(session *Session, chatId int64, text string) {
	switch session.Step {
	case StepTransactionAmount:
		if _, err := strconv.ParseFloat(text, 64); err != nil {
			h.send(chatId, "⚠️ Сумма должна быть числом, попробуйте ещё раз:", nil)
			return
		}
		session.Data["amount"] = text
		h.ask(session, chatId, StepTransactionCategoryId, "Введите ID категории:")
	case StepTransactionCategoryId:
		h.finishCreate(session, chatId, text)
	case StepTransactionGetId:
		h.finishGet(session, chatId, text)
	}
}

func (h *TransactionHandlerImpl) finishCreate(session *Session, chatId int64, categoryIdText string) {
	categoryId, err := strconv.Atoi(categoryIdText)
	if err != nil {
		h.send(chatId, "⚠️ ID категории должен быть числом, попробуйте ещё раз:", nil)
		return
	}

	amount, _ := strconv.ParseFloat(session.Data["amount"], 64)
	txType := session.Data["type"]
	session.Reset()

	transaction, err := session.Client.CreateTransaction(session.UserId, amount, txType, time.Now(), categoryId)
	if err != nil {
		h.fail(chatId, "не удалось создать транзакцию", err, transactionMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Транзакция <b>#%d</b> на сумму <b>%.2f</b> сохранена.", transaction.Id, transaction.Amount), kbPtr(transactionMenu()))
}

func (h *TransactionHandlerImpl) finishGet(session *Session, chatId int64, text string) {
	id, err := strconv.Atoi(text)
	if err != nil {
		h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
		return
	}
	session.Reset()

	transaction, err := session.Client.GetTransactionById(id)
	if err != nil {
		h.fail(chatId, "не удалось найти транзакцию", err, transactionMenu())
		return
	}

	icon := "💵"
	if transaction.Type == "expense" {
		icon = "💳"
	}
	h.send(chatId, fmt.Sprintf("%s <b>#%d</b> %.2f — %s", icon, transaction.Id, transaction.Amount, transaction.Date.Format("02.01.2006")), kbPtr(transactionMenu()))
}
