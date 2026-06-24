package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Router разбирает входящие Update от Telegram и раздаёт их по хендлерам.
type Router struct {
	Deps
	User        *UserHandlerImpl
	Category    *CategoryHandlerImpl
	Transaction *TransactionHandlerImpl
	Budget      *BudgetHandlerImpl
}

func NewRouter(d Deps) *Router {
	return &Router{
		Deps:        d,
		User:        NewUserHandlerImpl(d),
		Category:    NewCategoryHandlerImpl(d),
		Transaction: NewTransactionHandlerImpl(d),
		Budget:      NewBudgetHandlerImpl(d),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	chatId := msg.Chat.ID
	session := r.Sessions.Get(chatId)

	if msg.IsCommand() && msg.Command() == "start" {
		session.Reset()
		r.greet(session, chatId)
		return
	}

	if session.Step == StepNone {
		r.send(chatId, "Выберите действие на клавиатуре ниже 👇", nil)
		return
	}

	r.routeText(session, chatId, msg.Text)
}

func (r *Router) greet(session *Session, chatId int64) {
	if session.LoggedIn() {
		r.send(chatId, "👋 С возвращением! Выберите действие:", kbPtr(mainMenu()))
		return
	}
	r.send(chatId, "👋 Привет! Я бот <b>FlowMoney</b> — помогу следить за бюджетом.\n\nЗарегистрируйтесь или войдите, чтобы продолжить.", kbPtr(authMenu()))
}

func (r *Router) routeText(session *Session, chatId int64, text string) {
	switch session.Step {
	case StepRegisterName, StepRegisterEmail, StepRegisterPassword,
		StepLoginEmail, StepLoginPassword, StepBalanceAmount:
		r.User.HandleText(session, chatId, text)
	case StepCategoryCreateTitle, StepCategoryGetId, StepCategoryUpdateId, StepCategoryUpdateTitle:
		r.Category.HandleText(session, chatId, text)
	case StepTransactionAmount, StepTransactionCategoryId, StepTransactionGetId:
		r.Transaction.HandleText(session, chatId, text)
	case StepBudgetCreateCategoryId, StepBudgetCreateAmount, StepBudgetCreateMonth, StepBudgetCreateYear,
		StepBudgetGetId, StepBudgetCategoryId, StepBudgetMonthMonth, StepBudgetMonthYear,
		StepBudgetUpdateId, StepBudgetUpdateAmount, StepBudgetDeleteId:
		r.Budget.HandleText(session, chatId, text)
	}
}

func (r *Router) handleCallback(cb *tgbotapi.CallbackQuery) {
	chatId := cb.Message.Chat.ID
	session := r.Sessions.Get(chatId)

	if _, err := r.Bot.Request(tgbotapi.NewCallback(cb.ID, "")); err != nil {
		log.Println("flowmoney bot: callback ack failed:", err)
	}

	switch cb.Data {
	case cbMenuMain:
		r.edit(chatId, cb.Message.MessageID, "📍 <b>Главное меню</b>", mainMenu())
	case cbMenuCategory:
		r.edit(chatId, cb.Message.MessageID, "📁 <b>Категории</b>", categoryMenu())
	case cbMenuTransaction:
		r.edit(chatId, cb.Message.MessageID, "💸 <b>Транзакции</b>", transactionMenu())
	case cbMenuBudget:
		r.edit(chatId, cb.Message.MessageID, "📊 <b>Бюджет</b>", budgetMenu())

	case cbRegister:
		r.User.StartRegister(session, chatId)
	case cbLogin:
		r.User.StartLogin(session, chatId)
	case cbLogout:
		r.User.Logout(session, chatId)
	case cbProfile:
		r.User.ShowProfile(session, chatId)
	case cbBalanceTopUp:
		r.User.StartBalanceChange(session, chatId, true)
	case cbBalanceDeduct:
		r.User.StartBalanceChange(session, chatId, false)

	case cbCategoryCreate:
		r.Category.StartCreate(session, chatId)
	case cbCategoryList:
		r.Category.List(session, chatId)
	case cbCategoryGet:
		r.Category.StartGet(session, chatId)
	case cbCategoryUpdate:
		r.Category.StartUpdate(session, chatId)

	case cbTransactionCreateIncome:
		r.Transaction.StartCreate(session, chatId, "income")
	case cbTransactionCreateExpense:
		r.Transaction.StartCreate(session, chatId, "expense")
	case cbTransactionList:
		r.Transaction.List(session, chatId)
	case cbTransactionGet:
		r.Transaction.StartGet(session, chatId)

	case cbBudgetCreate:
		r.Budget.StartCreate(session, chatId)
	case cbBudgetGet:
		r.Budget.StartGet(session, chatId)
	case cbBudgetByCategory:
		r.Budget.StartGetByCategory(session, chatId)
	case cbBudgetByMonth:
		r.Budget.StartGetByMonth(session, chatId)
	case cbBudgetUpdate:
		r.Budget.StartUpdate(session, chatId)
	case cbBudgetDelete:
		r.Budget.StartDelete(session, chatId)
	}
}
