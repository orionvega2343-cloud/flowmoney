package handlers

import tele "gopkg.in/telebot.v3"

// Deps — общие зависимости всех хендлеров бота.
type Deps struct {
	Store *Store
}

// account возвращает аккаунт чата и сбрасывает незавершённый диалог —
// любое нажатие кнопки или команда считается новым намерением пользователя.
// Сам диалог (continueDialog) получает аккаунт напрямую через Store, минуя
// этот сброс.
func (d Deps) account(c tele.Context) *Account {
	acc := d.Store.Get(c.Chat().ID)
	acc.Step = nil
	return acc
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
