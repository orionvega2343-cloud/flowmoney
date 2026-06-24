package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Deps — общие зависимости, которые нужны всем хендлерам бота.
type Deps struct {
	Bot      *tgbotapi.BotAPI
	Sessions *SessionStore
}

func (d Deps) send(chatId int64, text string, kb *tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = tgbotapi.ModeHTML
	if kb != nil {
		msg.ReplyMarkup = *kb
	}
	if _, err := d.Bot.Send(msg); err != nil {
		log.Println("flowmoney bot: send failed:", err)
	}
}

func (d Deps) edit(chatId int64, messageId int, text string, kb tgbotapi.InlineKeyboardMarkup) {
	e := tgbotapi.NewEditMessageTextAndMarkup(chatId, messageId, text, kb)
	e.ParseMode = tgbotapi.ModeHTML
	if _, err := d.Bot.Send(e); err != nil {
		log.Println("flowmoney bot: edit failed:", err)
	}
}

// ask переводит сессию на следующий шаг формы и присылает её текст вопроса.
func (d Deps) ask(session *Session, chatId int64, step Step, text string) {
	session.Step = step
	d.send(chatId, text, nil)
}

// requireLogin отправляет приглашение войти, если у чата ещё нет активной сессии.
func (d Deps) requireLogin(session *Session, chatId int64) bool {
	if session.LoggedIn() {
		return true
	}
	d.send(chatId, "🔒 Сначала нужно войти в аккаунт.", kbPtr(authMenu()))
	return false
}

func (d Deps) fail(chatId int64, action string, err error, kb tgbotapi.InlineKeyboardMarkup) {
	d.send(chatId, "❌ "+action+": "+err.Error(), kbPtr(kb))
}
