package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

var (
	btnAuthRegister   = tele.Btn{Unique: "auth_register"}
	btnAuthLogin      = tele.Btn{Unique: "auth_login"}
	btnProfileBalance = tele.Btn{Unique: "user_balance"}
)

func authMarkup() *tele.ReplyMarkup {
	kb := &tele.ReplyMarkup{}
	kb.Inline(tele.Row{
		{Unique: "auth_register", Text: "📝 Регистрация"},
		{Unique: "auth_login", Text: "🔑 Войти"},
	})
	return kb
}

func profileMarkup() *tele.ReplyMarkup {
	kb := &tele.ReplyMarkup{}
	kb.Inline(tele.Row{{Unique: "user_balance", Text: "💰 Изменить баланс"}})
	return kb
}

type UserHandlers struct{ Deps }

func NewUserHandlers(d Deps) *UserHandlers { return &UserHandlers{d} }

func (h *UserHandlers) Register(bot *tele.Bot) {
	bot.Handle("/start", h.start)
	bot.Handle(&btnAuthRegister, h.startRegister)
	bot.Handle(&btnAuthLogin, h.startLogin)
	bot.Handle(&btnProfileBalance, h.startBalance)
	bot.Handle(txtProfile, h.profile)
	bot.Handle(txtLogout, h.logout)
}

func (h *UserHandlers) start(c tele.Context) error {
	acc := h.account(c)
	if acc.LoggedIn() {
		return c.Send("👋 С возвращением!", mainKeyboard())
	}
	return c.Send(
		"👋 Привет! Я бот <b>FlowMoney</b> — помогу следить за бюджетом.\n\nЗарегистрируйтесь или войдите, чтобы продолжить.",
		authMarkup(),
	)
}

func (h *UserHandlers) startRegister(c tele.Context) error {
	acc := h.account(c)
	var name, email string

	return startDialog(c, acc, &Step{
		Prompt: "📝 <b>Регистрация</b>\n\nВведите ваше имя:",
		Next: func(reply string) StepResult {
			name = strings.TrimSpace(reply)
			return ask("Введите email:", func(reply string) StepResult {
				email = strings.TrimSpace(reply)
				return ask("Введите пароль:", func(reply string) StepResult {
					password := strings.TrimSpace(reply)
					if err := acc.Client.Register(name, email, password); err != nil {
						return fail(err)
					}
					return done("✅ Регистрация прошла успешно! Теперь войдите в аккаунт.", authMarkup())
				})
			})
		},
	})
}

func (h *UserHandlers) startLogin(c tele.Context) error {
	acc := h.account(c)
	var email string

	return startDialog(c, acc, &Step{
		Prompt: "🔑 <b>Вход</b>\n\nВведите email:",
		Next: func(reply string) StepResult {
			email = strings.TrimSpace(reply)
			return ask("Введите пароль:", func(reply string) StepResult {
				password := strings.TrimSpace(reply)
				userId, err := acc.Client.Login(email, password)
				if err != nil {
					return fail(err)
				}
				acc.UserId = userId
				return done("✅ Вход выполнен!", mainKeyboard())
			})
		},
	})
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
		"👤 <b>Профиль</b>\n\n📧 Email: %s\n💰 Баланс: <b>%.2f</b>\n📅 Создан: %s",
		user.Email, user.Balance, user.CreatedAt.Format("02.01.2006"),
	)
	return c.Send(text, profileMarkup())
}

func (h *UserHandlers) startBalance(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "💰 На сколько изменить баланс? Введите сумму, можно со знаком «-», например 500 или -200:",
		Next: func(reply string) StepResult {
			delta, err := strconv.ParseFloat(strings.TrimSpace(reply), 64)
			if err != nil {
				return fail(errors.New("сумма должна быть числом"))
			}

			user, err := acc.Client.GetUserById(acc.UserId)
			if err != nil {
				return fail(err)
			}

			updated, err := acc.Client.UpdateBalance(acc.UserId, user.Balance+delta)
			if err != nil {
				return fail(err)
			}
			return done(fmt.Sprintf("✅ Баланс обновлён: <b>%.2f</b>", updated.Balance))
		},
	})
}
