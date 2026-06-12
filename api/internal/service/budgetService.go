package service

import (
	"errors"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/repository"
)

var ErrBudgetExceeded = errors.New("бюджет превышен")

type BudgetService interface {
	CreateBudget(b models.Budget) (models.Budget, error)
	GetBudgetById(id int) (models.Budget, error)
	GetBudgetByCategoryId(catId int) (models.Budget, error)
	GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error)
	UpdateBudget(amount float64, id int) (models.Budget, error)
	DeleteBudgetById(id int) error
}

type BudgetServiceImpl struct {
	B  repository.BudgetRepo
	Tr repository.TransactionRepo
}

func NewBudgetService(b repository.BudgetRepo, tr repository.TransactionRepo) *BudgetServiceImpl {
	return &BudgetServiceImpl{B: b, Tr: tr}
}

func (r *BudgetServiceImpl) CreateBudget(b models.Budget) (models.Budget, error) {
	res, err := r.B.CreateBudget(b)
	if err != nil {
		return models.Budget{}, err
	}
	return res, nil
}

func (r *BudgetServiceImpl) GetBudgetById(id int) (models.Budget, error) {
	res, err := r.B.GetBudgetById(id)
	if err != nil {
		return models.Budget{}, err
	}
	return res, nil
}

func (r *BudgetServiceImpl) GetBudgetByCategoryId(catId int) (models.Budget, error) {
	res, err := r.B.GetBudgetByCategoryId(catId)
	if err != nil {
		return models.Budget{}, err
	}
	return res, nil
}

func (r *BudgetServiceImpl) GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error) {
	res, err := r.B.GetByUserIdAndMonth(userId, month, year)
	if err != nil {
		return models.Budget{}, err
	}

	tr, err := r.Tr.GetTransactionByUserId(userId)
	if err != nil {
		return models.Budget{}, err
	}

	sum := 0.0

	for _, v := range tr {
		if v.Type == "expense" && res.Month == int(v.Date.Month()) && res.Year == v.Date.Year() {
			sum += v.Amount
		}
	}

	if sum > res.Amount {
		return models.Budget{}, ErrBudgetExceeded
	}
	return res, nil
}

func (r *BudgetServiceImpl) UpdateBudget(amount float64, id int) (models.Budget, error) {
	res, err := r.B.UpdateBudget(amount, id)
	if err != nil {
		return models.Budget{}, err
	}
	return res, nil
}

func (r *BudgetServiceImpl) DeleteBudgetById(id int) error {
	err := r.B.DeleteBudgetById(id)
	if err != nil {
		return err
	}
	return nil
}
