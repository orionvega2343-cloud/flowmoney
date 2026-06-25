package handlers

import tele "gopkg.in/telebot.v3"

// Register регистрирует все хендлеры бота на переданном экземпляре telebot.
func Register(bot *tele.Bot, store *Store) {
	d := Deps{Store: store}

	bot.Use(ackCallbacks)

	NewUserHandlers(d).Register(bot)
	NewCategoryHandlers(d).Register(bot)
	NewTransactionHandlers(d).Register(bot)
	NewBudgetHandlers(d).Register(bot)

	bot.Handle(tele.OnText, func(c tele.Context) error {
		acc := store.Get(c.Chat().ID)
		if acc.Step != nil {
			return continueDialog(c, acc)
		}
		return c.Send("Не понимаю 🤔 Воспользуйтесь клавиатурой ниже или командой /start.")
	})
}

// ackCallbacks сразу гасит "часики" на инлайн-кнопке, чтобы не ждать
// ответа хендлера — так делал прежний роутер перед разбором callback.Data.
func ackCallbacks(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Callback() != nil {
			_ = c.Respond()
		}
		return next(c)
	}
}
