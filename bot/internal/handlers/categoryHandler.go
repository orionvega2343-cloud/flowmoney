package handlers

import (
	"fmt"
	"strconv"
)

type CategoryHandler interface {
	StartCreate(session *Session, chatId int64)
	StartGet(session *Session, chatId int64)
	StartUpdate(session *Session, chatId int64)
	List(session *Session, chatId int64)
	HandleText(session *Session, chatId int64, text string)
}

type CategoryHandlerImpl struct {
	Deps
}

func NewCategoryHandlerImpl(d Deps) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{Deps: d}
}

func (h *CategoryHandlerImpl) StartCreate(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepCategoryCreateTitle, "➕ <b>Новая категория</b>\n\nВведите название:")
}

func (h *CategoryHandlerImpl) StartGet(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepCategoryGetId, "🔎 Введите ID категории:")
}

func (h *CategoryHandlerImpl) StartUpdate(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}
	h.ask(session, chatId, StepCategoryUpdateId, "✏️ Введите ID категории, которую нужно изменить:")
}

func (h *CategoryHandlerImpl) List(session *Session, chatId int64) {
	if !h.requireLogin(session, chatId) {
		return
	}

	categories, err := session.Client.GetByUserId(session.UserId)
	if err != nil {
		h.fail(chatId, "не удалось получить категории", err, categoryMenu())
		return
	}

	if len(categories) == 0 {
		h.send(chatId, "📁 У вас пока нет категорий.", kbPtr(categoryMenu()))
		return
	}

	text := "📋 <b>Ваши категории:</b>\n\n"
	for _, c := range categories {
		text += fmt.Sprintf("• <b>#%d</b> %s\n", c.Id, c.Title)
	}
	h.send(chatId, text, kbPtr(categoryMenu()))
}

func (h *CategoryHandlerImpl) HandleText(session *Session, chatId int64, text string) {
	switch session.Step {
	case StepCategoryCreateTitle:
		h.finishCreate(session, chatId, text)
	case StepCategoryGetId:
		h.finishGet(session, chatId, text)
	case StepCategoryUpdateId:
		id, err := strconv.Atoi(text)
		if err != nil {
			h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
			return
		}
		session.Data["id"] = strconv.Itoa(id)
		h.ask(session, chatId, StepCategoryUpdateTitle, "Введите новое название:")
	case StepCategoryUpdateTitle:
		h.finishUpdate(session, chatId, text)
	}
}

func (h *CategoryHandlerImpl) finishCreate(session *Session, chatId int64, title string) {
	session.Reset()

	category, err := session.Client.CreateCategory(title, session.UserId)
	if err != nil {
		h.fail(chatId, "не удалось создать категорию", err, categoryMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Категория <b>%s</b> создана (ID <b>%d</b>).", category.Title, category.Id), kbPtr(categoryMenu()))
}

func (h *CategoryHandlerImpl) finishGet(session *Session, chatId int64, text string) {
	id, err := strconv.Atoi(text)
	if err != nil {
		h.send(chatId, "⚠️ ID должен быть числом, попробуйте ещё раз:", nil)
		return
	}
	session.Reset()

	category, err := session.Client.GetCategoryById(id)
	if err != nil {
		h.fail(chatId, "не удалось найти категорию", err, categoryMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("📁 <b>#%d</b> %s", category.Id, category.Title), kbPtr(categoryMenu()))
}

func (h *CategoryHandlerImpl) finishUpdate(session *Session, chatId int64, title string) {
	id, _ := strconv.Atoi(session.Data["id"])
	session.Reset()

	category, err := session.Client.UpdateCategory(id, title)
	if err != nil {
		h.fail(chatId, "не удалось изменить категорию", err, categoryMenu())
		return
	}

	h.send(chatId, fmt.Sprintf("✅ Категория <b>#%d</b> обновлена: %s", category.Id, category.Title), kbPtr(categoryMenu()))
}
