package handlers

import (
	"fmt"
	"strconv"
)

type UserHandler interface {
	StartRegister(session *Session, chatId int64)
	StartLogin(session *Session, chatId int64)
	Logout(session *Session, chatId int64)
	ShowProfile(session *Session, chatId int64)
	StartBalanceChange(session *Session, chatId int64, topUp bool)
	HandleText(session *Session, chatId int64, text string)
}

type UserHandlerImpl struct {
	Deps
}

func NewUserHandlerImpl(d Deps) *UserHandlerImpl {
	return &UserHandlerImpl{Deps: d}
}

func (h *UserHandlerImpl) StartRegister(session *Session, chatId int64) {
	session.Reset()
	h.ask(session, chatId, StepRegisterName, "📝 <b>Регистрация</b>\n\nВведите ваше имя:")
}

func (h *UserHandlerImpl) StartLogin(session *Session, chatId int64) {
	session.Reset()
	h.ask(session, chatId, StepLoginEmail, "🔑 <b>Вход</b>\n\nВведите email:")
}

func (h *UserHandlerImpl) Logout(session *Session, chatId int64) {
	h.Sessions.Drop(chatId)
	h.send(chatId, "🚪 Вы вышли из аккаунта.", kbPtr(authMenu()))
}

func (h *UserHandlerImpl) ShowProfile(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}

	user, err := session.Client.GetUserById(session.UserId)
	if err != nil {
		h.fail(chatId, "не удалось получить профиль", err, mainMenu())
		return
	}

	text := fmt.Sprintf(
		"👤 <b>Профиль</b>\n\n📧 Email: %s\n💰 Баланс: <b>%.2f</b>\n📅 Создан: %s",
		user.Email, user.Balance, user.CreatedAt.Format("02.01.2006"),
	)
	h.send(chatId, text, kbPtr(profileMenu()))
}

func (h *UserHandlerImpl) StartBalanceChange(session *Session, chatId int64, topUp bool) {
	if !h.requireLogin(session, chatId) {
		return
	}

	if topUp {
		session.Data["balance_op"] = "topup"
		h.ask(session, chatId, StepBalanceAmount, "➕ На сколько пополнить баланс?")
	} else {
		session.Data["balance_op"] = "deduct"
		h.ask(session, chatId, StepBalanceAmount, "➖ Сколько списать с баланса?")
	}
}

func (h *UserHandlerImpl) HandleText(session *Session, chatId int64, text string) {
	switch session.Step {
	case StepRegisterName:
		session.Data["name"] = text
		h.ask(session, chatId, StepRegisterEmail, "Введите email:")
	case StepRegisterEmail:
		session.Data["email"] = text
		h.ask(session, chatId, StepRegisterPassword, "Введите пароль:")
	case StepRegisterPassword:
		session.Data["password"] = text
		h.finishRegister(session, chatId)
	case StepLoginEmail:
		session.Data["email"] = text
		h.ask(session, chatId, StepLoginPassword, "Введите пароль:")
	case StepLoginPassword:
		session.Data["password"] = text
		h.finishLogin(session, chatId)
	case StepBalanceAmount:
		h.finishBalanceChange(session, chatId, text)
	}
}

func (h *UserHandlerImpl) finishRegister(session *Session, chatId int64) {
	name := session.Data["name"]
	email := session.Data["email"]
	password := session.Data["password"]
	session.Reset()

	err := session.Client.Register(name, email, password)
	if err != nil {
		h.fail(chatId, "не удалось зарегистрироваться", err, authMenu())
		return
	}

	h.send(chatId, "✅ Регистрация прошла успешно! Теперь войдите в аккаунт.", kbPtr(authMenu()))
}

func (h *UserHandlerImpl) finishLogin(session *Session, chatId int64) {
	email := session.Data["email"]
	password := session.Data["password"]
	session.Reset()

	userId, err := session.Client.Login(email, password)
	if err != nil {
		h.fail(chatId, "не удалось войти", err, authMenu())
		return
	}

	session.UserId = userId
	h.send(chatId, "✅ Вход выполнен!", kbPtr(mainMenu()))
}

func (h *UserHandlerImpl) finishBalanceChange(session *Session, chatId int64, text string) {
	amount, err := strconv.ParseFloat(text, 64)
	if err != nil {
		h.send(chatId, "⚠️ Введите сумму числом, например 1500.50", nil)
		return
	}

	op := session.Data["balance_op"]
	session.Reset()

	user, err := session.Client.GetUserById(session.UserId)
	if err != nil {
		h.fail(chatId, "не удалось получить текущий баланс", err, mainMenu())
		return
	}

	newBalance := user.Balance
	if op == "topup" {
		newBalance += amount
	} else {
		newBalance -= amount
	}

	updated, err := session.Client.UpdateBalance(session.UserId, newBalance)
	if err != nil {
		h.fail(chatId, "не удалось обновить баланс", err, mainMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Баланс обновлён: <b>%.2f</b>", updated.Balance), kbPtr(profileMenu()))
}
