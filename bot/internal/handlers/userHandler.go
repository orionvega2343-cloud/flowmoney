package handlers

import (
	"fmt"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

type UserHandlers struct{ Deps }

func NewUserHandlers(d Deps) *UserHandlers { return &UserHandlers{d} }

func (h *UserHandlers) Register(bot *tele.Bot) {
	bot.Handle("/start", h.start)
	bot.Handle("/register", h.register)
	bot.Handle("/login", h.login)
	bot.Handle("/balance", h.balance)
	bot.Handle(txtProfile, h.profile)
	bot.Handle(txtLogout, h.logout)
}

func (h *UserHandlers) start(c tele.Context) error {
	acc := h.account(c)
	if acc.LoggedIn() {
		return c.Send("👋 С возвращением!", mainKeyboard())
	}
	return c.Send(
		"👋 Привет! Я бот <b>FlowMoney</b> — помогу следить за бюджетом.\n\n"+
			"📝 Регистрация: /register Имя Email Пароль\n"+
			"🔑 Вход: /login Email Пароль",
		&tele.ReplyMarkup{RemoveKeyboard: true},
	)
}

func (h *UserHandlers) register(c tele.Context) error {
	args := c.Args()
	if len(args) < 3 {
		return c.Send("⚠️ Использование: /register Имя Email Пароль")
	}
	name, email, password := args[0], args[1], args[2]

	acc := h.account(c)
	if err := acc.Client.Register(name, email, password); err != nil {
		return h.fail(c, "не удалось зарегистрироваться", err)
	}
	return c.Send("✅ Регистрация прошла успешно! Теперь войдите: /login Email Пароль")
}

func (h *UserHandlers) login(c tele.Context) error {
	args := c.Args()
	if len(args) < 2 {
		return c.Send("⚠️ Использование: /login Email Пароль")
	}
	email, password := args[0], args[1]

	acc := h.account(c)
	userId, err := acc.Client.Login(email, password)
	if err != nil {
		return h.fail(c, "не удалось войти", err)
	}
	acc.UserId = userId
	return c.Send("✅ Вход выполнен!", mainKeyboard())
}

func (h *UserHandlers) logout(c tele.Context) error {
	h.Store.Drop(c.Chat().ID)
	return c.Send("🚪 Вы вышли из аккаунта.", &tele.ReplyMarkup{RemoveKeyboard: true})
}

func (h *UserHandlers) profile(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	user, err := acc.Client.GetUserById(acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось получить профиль", err)
	}

	text := fmt.Sprintf(
		"👤 <b>Профиль</b>\n\n📧 Email: %s\n💰 Баланс: <b>%.2f</b>\n📅 Создан: %s\n\nИзменить баланс: /balance ±Сумма",
		user.Email, user.Balance, user.CreatedAt.Format("02.01.2006"),
	)
	return c.Send(text)
}

func (h *UserHandlers) balance(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 1 {
		return c.Send("⚠️ Использование: /balance ±Сумма, например /balance 500 или /balance -200")
	}
	delta, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return c.Send("⚠️ Сумма должна быть числом, например /balance 500 или /balance -200")
	}

	user, err := acc.Client.GetUserById(acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось получить текущий баланс", err)
	}

	updated, err := acc.Client.UpdateBalance(acc.UserId, user.Balance+delta)
	if err != nil {
		return h.fail(c, "не удалось обновить баланс", err)
	}
	return c.Send(fmt.Sprintf("✅ Баланс обновлён: <b>%.2f</b>", updated.Balance))
}
