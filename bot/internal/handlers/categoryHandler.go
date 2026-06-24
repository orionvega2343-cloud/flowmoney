package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

var (
	// btnCategoryGet — нажатие на категорию в списке вызывает GetCategoryById
	// с ID из callback-данных.
	btnCategoryGet = tele.Btn{Unique: "cat_get"}
	// btnCategoryEdit запускает диалог изменения названия для категории,
	// ID которой пришёл в callback-данных.
	btnCategoryEdit = tele.Btn{Unique: "cat_edit"}
	btnCategoryNew  = tele.Btn{Unique: "cat_new"}
)

type CategoryHandlers struct{ Deps }

func NewCategoryHandlers(d Deps) *CategoryHandlers { return &CategoryHandlers{d} }

func (h *CategoryHandlers) Register(bot *tele.Bot) {
	bot.Handle(txtCategories, h.list)
	bot.Handle(&btnCategoryNew, h.startCreate)
	bot.Handle(&btnCategoryGet, h.get)
	bot.Handle(&btnCategoryEdit, h.startUpdate)
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

	kb := &tele.ReplyMarkup{}
	rows := []tele.Row{{{Unique: "cat_new", Text: "➕ Новая категория"}}}
	for _, cat := range categories {
		id := strconv.Itoa(cat.Id)
		rows = append(rows, tele.Row{
			{Unique: "cat_get", Text: fmt.Sprintf("🏷 #%d %s", cat.Id, cat.Title), Data: id},
			{Unique: "cat_edit", Text: "✏️", Data: id},
		})
	}
	kb.Inline(rows...)

	text := "📋 <b>Ваши категории</b>"
	if len(categories) == 0 {
		text = "📁 У вас пока нет категорий."
	}
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

func (h *CategoryHandlers) startCreate(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: "➕ <b>Новая категория</b>\n\nВведите название:",
		Next: func(reply string) StepResult {
			title := strings.TrimSpace(reply)
			if title == "" {
				return fail(errors.New("название не может быть пустым"))
			}
			category, err := acc.Client.CreateCategory(title, acc.UserId)
			if err != nil {
				return fail(err)
			}
			return done(fmt.Sprintf("✅ Категория <b>%s</b> создана (ID <b>%d</b>).", category.Title, category.Id))
		},
	})
}

func (h *CategoryHandlers) startUpdate(c tele.Context) error {
	acc := h.account(c)
	if !h.requireLogin(c, acc) {
		return nil
	}

	id, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return nil
	}

	return startDialog(c, acc, &Step{
		Prompt: fmt.Sprintf("✏️ Введите новое название для категории <b>#%d</b>:", id),
		Next: func(reply string) StepResult {
			title := strings.TrimSpace(reply)
			if title == "" {
				return fail(errors.New("название не может быть пустым"))
			}
			category, err := acc.Client.UpdateCategory(id, title)
			if err != nil {
				return fail(err)
			}
			return done(fmt.Sprintf("✅ Категория <b>#%d</b> обновлена: %s", category.Id, category.Title))
		},
	})
}
