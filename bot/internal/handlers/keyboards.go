package handlers

import tele "gopkg.in/telebot.v3"

const (
	txtProfile      = "👤 Профиль"
	txtCategories   = "📁 Категории"
	txtTransactions = "💸 Транзакции"
	txtBudget       = "📊 Бюджет"
	txtLogout       = "🚪 Выйти"
)

// mainKeyboard — реплай-клавиатура с разделами бота, показывается после входа.
func mainKeyboard() *tele.ReplyMarkup {
	kb := &tele.ReplyMarkup{ResizeKeyboard: true}
	kb.Reply(
		tele.Row{{Text: txtProfile}, {Text: txtCategories}},
		tele.Row{{Text: txtTransactions}, {Text: txtBudget}},
		tele.Row{{Text: txtLogout}},
	)
	return kb
}
