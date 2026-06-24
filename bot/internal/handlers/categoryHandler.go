package handlers

import (
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

// btnCategoryGet — шаблон инлайн-кнопки для конкретной категории в списке;
// нажатие сразу вызывает GetCategoryById с ID категории из callback-данных.
var btnCategoryGet = tele.Btn{Unique: "cat_get"}

type CategoryHandlers struct{ Deps }

func NewCategoryHandlers(d Deps) *CategoryHandlers { return &CategoryHandlers{d} }

func (h *CategoryHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtCategories, h.list)
	bot.Handle("/category_new", h.create)
	bot.Handle("/category_update", h.update)
	bot.Handle(&btnCategoryGet, h.get)
}

func (h *CategoryHandlers) list(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	categories, err := acc.Client.GetByUserId(acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось получить категории", err)
	}
	if len(categories) == 0 {
		return c.Send("📁 У вас пока нет категорий.\n\nСоздать: /category_new Название")
	}

	rows := make([]tele.Row, 0, len(categories))
	for _, cat := range categories {
		label := fmt.Sprintf("🏷 #%d %s", cat.Id, cat.Title)
		rows = append(rows, tele.Row{{Unique: "cat_get", Text: label, Data: strconv.Itoa(cat.Id)}})
	}
	kb := &tele.ReplyMarkup{}
	kb.Inline(rows...)

	text := "📋 <b>Ваши категории</b>\n\nСоздать: /category_new Название\nИзменить: /category_update ID Название"
	return c.Send(text, kb)
}

func (h *CategoryHandlers) get(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	id, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	category, err := acc.Client.GetCategoryById(id)
	if err != nil {
		return h.fail(c, "не удалось найти категорию", err)
	}
	return c.Send(fmt.Sprintf("📁 <b>#%d</b> %s", category.Id, category.Title))
}

func (h *CategoryHandlers) create(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	title := strings.TrimSpace(c.Message().Payload)
	if title == "" {
		return c.Send("⚠️ Использование: /category_new Название")
	}

	category, err := acc.Client.CreateCategory(title, acc.UserId)
	if err != nil {
		return h.fail(c, "не удалось создать категорию", err)
	}
	return c.Send(fmt.Sprintf("✅ Категория <b>%s</b> создана (ID <b>%d</b>).", category.Title, category.Id))
}

func (h *CategoryHandlers) update(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	args := c.Args()
	if len(args) < 2 {
		return c.Send("⚠️ Использование: /category_update ID Название")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Send("⚠️ ID должен быть числом.")
	}
	title := strings.Join(args[1:], " ")

	category, err := acc.Client.UpdateCategory(id, title)
	if err != nil {
		return h.fail(c, "не удалось изменить категорию", err)
	}
	return c.Send(fmt.Sprintf("✅ Категория <b>#%d</b> обновлена: %s", category.Id, category.Title))
}
