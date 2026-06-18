package service

import (
	"errors"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateTransactionSuccess(t *testing.T) {
	//Arrange
	var (
		mockUserRepo   = new(mocks.MockUserRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		user           = models.User{}
		transaction    = models.Transaction{}
	)
	mockUserRepo.On("GetUserById", mock.Anything).Return(user, nil)
	mockUserRepo.On("UpdateBalance", mock.Anything).Return(user, nil)
	mockTxRepo.On("CreateTransaction", mock.Anything).Return(transaction, nil)

	//Act
	_, err := svc.CreateTransaction(transaction)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockTxRepo.AssertExpectations(t)
}

func Test_CreateTransactionFail(t *testing.T) {
	//Arrange
	var (
		mockUserRepo   = new(mocks.MockUserRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		user           = models.User{}
		transaction    = models.Transaction{}
	)
	mockUserRepo.On("GetUserById", mock.Anything).Return(user, errors.New("user not found"))

	//Act
	_, err := svc.CreateTransaction(transaction)

	//Assert
	assert.Equal(t, errors.New("user not found"), err)
	assert.Error(t, err)
	mockTxRepo.AssertExpectations(t)
}

func Test_getTransactionByIdSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockUserRepo   = new(mocks.MockUserRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		transaction    = models.Transaction{}
	)
	mockTxRepo.On("GetTransactionById", mock.Anything).Return(transaction, nil)

	//Act
	_, err := svc.GetTransactionById(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockTxRepo.AssertExpectations(t)
}

func Test_GetTransactionByIdFail(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockUserRepo   = new(mocks.MockUserRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		transaction    = models.Transaction{}
	)
	mockTxRepo.On("GetTransactionById", mock.Anything).Return(transaction, errors.New("user not found"))

	//Act
	_, err := svc.GetTransactionById(id)

	//Assert
	assert.Equal(t, errors.New("user not found"), err)
	assert.Error(t, err)
	mockTxRepo.AssertExpectations(t)
}

func Test_GetTransactionByUserIdSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockUserRepo   = new(mocks.MockUserRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		transaction    = []models.Transaction{}
	)
	mockTxRepo.On("GetTransactionByUserId", mock.Anything).Return(transaction, nil)

	//Act
	_, err := svc.GetTransactionByUserId(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockTxRepo.AssertExpectations(t)
}

func Test_GetTransactionByUserIdFail(t *testing.T) {
	//arrange
	var (
		id             = 1
		mockUserRepo   = new(mocks.MockUserRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		svc            = NewTransactionService(mockTxRepo, mockUserRepo, mockBudgetRepo)
		transaction    = []models.Transaction{}
	)
	mockTxRepo.On("GetTransactionByUserId", mock.Anything).Return(transaction, errors.New("user not found"))

	//Act
	_, err := svc.GetTransactionByUserId(id)

	//Assert
	assert.Equal(t, errors.New("user not found"), err)
	assert.Error(t, err)
	mockTxRepo.AssertExpectations(t)
}
