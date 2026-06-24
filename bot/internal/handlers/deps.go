package handlers

import tele "gopkg.in/telebot.v3"

// Deps — общие зависимости всех хендлеров бота.
type Deps struct {
	Store *Store
}

func (d Deps) account(c tele.Context) *Account {
	return d.Store.Get(c.Chat().ID)
}

// requireLogin отправляет приглашение войти, если у чата ещё нет аккаунта.
func (d Deps) requireLogin(c tele.Context, acc *Account) bool {
	if acc.LoggedIn() {
		return true
	}
	_ = c.Send("🔒 Сначала войдите: /login Email Пароль")
	return false
}

func (d Deps) fail(c tele.Context, action string, err error) error {
	return c.Send("❌ " + action + ": " + err.Error())
}
