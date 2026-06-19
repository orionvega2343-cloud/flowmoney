package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) CreateTransaction(tr models.Transaction) (models.Transaction, error) {
	args := m.Called(tr)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetTransactionById(id int) (models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetTransactionByUserId(userId int) ([]models.Transaction, error) {
	args := m.Called(userId)
	return args.Get(0).([]models.Transaction), args.Error(1)
}
