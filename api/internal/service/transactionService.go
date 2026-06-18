package service

import (
	"errors"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/repository"
)

type TransactionService interface {
	CreateTransaction(tr models.Transaction) (models.Transaction, error)
	GetTransactionById(id int) (models.Transaction, error)
	GetTransactionByUserId(userId int) ([]models.Transaction, error)
}

type TransactionServiceImpl struct {
	Tr repository.TransactionRepository
	Ur repository.UserRepository
	Br repository.BudgetRepository
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func NewTransactionService(Tr repository.TransactionRepository, Ur repository.UserRepository, Br repository.BudgetRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{Tr: Tr, Ur: Ur, Br: Br}
}

func (r *TransactionServiceImpl) CreateTransaction(tr models.Transaction) (models.Transaction, error) {

	user, err := r.Ur.GetUserById(tr.UserId)

	if err != nil {
		return models.Transaction{}, err
	}

	currentBal := user.Balance

	if tr.Type == "income" {
		currentBal += tr.Amount

	} else if tr.Type == "expense" {
		currentBal -= tr.Amount
		if currentBal < 0 {
			return models.Transaction{}, ErrInsufficientFunds
		}
	}
	_, err = r.Tr.CreateTransaction(tr)

	if err != nil {
		return models.Transaction{}, err
	}

	_, err = r.Ur.UpdateBalance(tr.UserId, currentBal)

	if err != nil {
		return models.Transaction{}, err
	}

	return tr, nil
}

func (r *TransactionServiceImpl) GetTransactionById(id int) (models.Transaction, error) {
	res, err := r.Tr.GetTransactionById(id)
	if err != nil {
		return models.Transaction{}, err
	}
	return res, nil
}

func (r *TransactionServiceImpl) GetTransactionByUserId(userId int) ([]models.Transaction, error) {
	res, err := r.Tr.GetTransactionByUserId(userId)
	if err != nil {
		return []models.Transaction{}, err
	}
	return res, nil
}
