package handlers

import (
	"encoding/json"
	"errors"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/service"
	"net/http"
	"strconv"
)

type TransactionHandler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionById(w http.ResponseWriter, r *http.Request)
	GetTransactionByUserId(w http.ResponseWriter, r *http.Request)
}

type TransactionHandlerImpl struct {
	Ts service.TransactionService
}

func NewTransactionHandlerImpl(ts service.TransactionService) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{Ts: ts}
}

func (t *TransactionHandlerImpl) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transaction, err := t.Ts.CreateTransaction(tx)
	if errors.Is(err, service.ErrInsufficientFunds) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TransactionHandlerImpl) GetTransactionById(w http.ResponseWriter, r *http.Request) {

	url := r.PathValue("id")
	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transaction, err := t.Ts.GetTransactionById(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TransactionHandlerImpl) GetTransactionByUserId(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transaction, err := t.Ts.GetTransactionByUserId(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
