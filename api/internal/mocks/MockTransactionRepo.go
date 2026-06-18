package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

// Transaction mock
type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) CreateTransaction(tr models.Transaction) (models.Transaction, error) {
	args := m.Called(tr)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) GetTransactionById(id int) (models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) GetTransactionByUserId(userId int) ([]models.Transaction, error) {
	args := m.Called(userId)
	return args.Get(0).([]models.Transaction), args.Error(1)
}
