package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

// Budget mock
type MockBudgetRepo struct {
	mock.Mock
}

func (m *MockBudgetRepo) CreateBudget(budget models.Budget) (models.Budget, error) {
	args := m.Called(budget)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (m *MockBudgetRepo) GetBudgetById(id int) (models.Budget, error) {
	args := m.Called(id)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (m *MockBudgetRepo) GetBudgetByCategoryId(catId int) (models.Budget, error) {
	args := m.Called(catId)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (m *MockBudgetRepo) GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error) {
	args := m.Called(userId, month, year)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (m *MockBudgetRepo) UpdateBudget(balance float64, id int) (models.Budget, error) {
	args := m.Called(balance, id)
	return args.Get(0).(models.Budget), args.Error(1)
}

func (m *MockBudgetRepo) DeleteBudgetById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
