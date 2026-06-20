package handlers

import (
	"encoding/json"
	"errors"
	"flowmoney/api/internal/errs"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/service"
	"net/http"
	"strconv"
)

type BudgetHandler interface {
	CreateBudget(w http.ResponseWriter, r *http.Request)
	GetBudgetById(w http.ResponseWriter, r *http.Request)
	GetBudgetByCategoryId(w http.ResponseWriter, r *http.Request)
	GetByUserIdAndMonth(w http.ResponseWriter, r *http.Request)
	UpdateBudget(w http.ResponseWriter, r *http.Request)
	DeleteBudgetById(w http.ResponseWriter, r *http.Request)
}

type BudgetHandlerImpl struct {
	Bs service.BudgetService
}

func NewBudgetHandlerImpl(bs service.BudgetService) *BudgetHandlerImpl {
	return &BudgetHandlerImpl{Bs: bs}
}

func (h *BudgetHandlerImpl) CreateBudget(w http.ResponseWriter, r *http.Request) {
	var b models.Budget

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	budget, err := h.Bs.CreateBudget(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *BudgetHandlerImpl) GetBudgetById(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	id, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	budget, err := h.Bs.GetBudgetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *BudgetHandlerImpl) GetBudgetByCategoryId(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	id, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	budget, err := h.Bs.GetBudgetByCategoryId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *BudgetHandlerImpl) GetByUserIdAndMonth(w http.ResponseWriter, r *http.Request) {
	var b models.Budget
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := r.PathValue("id")
	id, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	budget, err := h.Bs.GetByUserIdAndMonth(id, b.Month, b.Year)
	if errors.Is(err, errs.ErrBudgetExceeded) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *BudgetHandlerImpl) UpdateBudget(w http.ResponseWriter, r *http.Request) {
	var b models.Budget

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := r.PathValue("id")
	id, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	budget, err := h.Bs.UpdateBudget(b.Amount, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *BudgetHandlerImpl) DeleteBudgetById(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	id, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Bs.DeleteBudgetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
