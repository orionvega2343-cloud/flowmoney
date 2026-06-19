package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockBudgetService struct {
	mock.Mock
}

func (r *MockBudgetService) CreateBudget(b models.Budget) (models.Budget, error) {
	args := r.Called(b)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (r *MockBudgetService) GetBudgetById(id int) (models.Budget, error) {
	args := r.Called(id)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (r *MockBudgetService) GetBudgetByCategoryId(catId int) (models.Budget, error) {
	args := r.Called(catId)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (r *MockBudgetService) GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error) {
	args := r.Called(userId, month, year)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (r *MockBudgetService) UpdateBudget(amount float64, id int) (models.Budget, error) {
	args := r.Called(amount, id)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (r *MockBudgetService) DeleteBudgetById(id int) error {
	args := r.Called(id)
	return args.Error(0)
}
