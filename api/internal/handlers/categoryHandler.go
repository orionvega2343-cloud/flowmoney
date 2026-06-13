package handlers

import (
	"encoding/json"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/service"
	"net/http"
	"strconv"
)

type CategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryById(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
}

type CategoryHandlerImpl struct {
	Cs service.CategoryService
}

func NewCategoryHandler(cs service.CategoryService) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{Cs: cs}
}

func (h *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category

	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := h.Cs.CreateCategory(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *CategoryHandlerImpl) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")

	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	caetgory, err := h.Cs.GetCategoryById(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(caetgory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *CategoryHandlerImpl) GetByUserId(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")

	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	caetgory, err := h.Cs.GetByUserId(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(caetgory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *CategoryHandlerImpl) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := r.PathValue("id")

	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := h.Cs.UpdateCategory(parsed, cat.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
