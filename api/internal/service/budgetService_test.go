package service

import (
	"errors"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateBudgetSuccess(t *testing.T) {
	//Arrange
	var (
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("CreateBudget", mock.Anything).Return(budget, nil)

	//Act
	_, err := svc.CreateBudget(budget)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_CreateBudgetFail(t *testing.T) {
	//Arrange
	var (
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("CreateBudget", mock.Anything).Return(budget, errors.New("some error"))

	//Act
	_, err := svc.CreateBudget(budget)

	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_GetBudgetByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
		id             = 1
	)
	mockBudgetRepo.On("GetBudgetById", mock.Anything).Return(budget, nil)

	//Act
	_, err := svc.GetBudgetById(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_GetBudgetByIdFail(t *testing.T) {
	//Arrange
	var (
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
		id             = 1
	)
	mockBudgetRepo.On("GetBudgetById", mock.Anything).Return(budget, errors.New("budget not found"))

	//Act
	_, err := svc.GetBudgetById(id)

	//Assert
	assert.Equal(t, errors.New("budget not found"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_GetBudgetByCategoryIdSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("GetBudgetByCategoryId", mock.Anything).Return(budget, nil)

	//Act
	_, err := svc.GetBudgetByCategoryId(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_GetBudgetByCategoryIdFail(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("GetBudgetByCategoryId", mock.Anything).Return(budget, errors.New("category not found"))

	//Act
	_, err := svc.GetBudgetByCategoryId(id)

	//Assert
	assert.Equal(t, errors.New("category not found"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_GetByUserIdAndMonthFail(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{Amount: 100}
	)
	mockBudgetRepo.On("GetByUserIdAndMonth", mock.Anything, mock.Anything, mock.Anything).Return(budget, errors.New("budget not found"))

	//Act
	_, err := svc.GetByUserIdAndMonth(id, 2, 3)

	assert.Equal(t, errors.New("budget not found"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_GetByUserIdAndMonthExceeded(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{Amount: 100}
		transaction    = []models.Transaction{{Amount: 200, Type: "expense"}}
	)
	mockBudgetRepo.On("GetByUserIdAndMonth", mock.Anything, mock.Anything, mock.Anything).Return(budget, nil)
	mockTxRepo.On("GetTransactionByUserId", mock.Anything).Return(transaction, nil)

	//Act
	_, err := svc.GetByUserIdAndMonth(id, 2, 3)

	//Assert
	assert.Equal(t, ErrBudgetExceeded, err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_GetByUserIdAndMonthSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{Amount: 300}
		transaction    = []models.Transaction{{Amount: 200, Type: "expense"}}
	)
	mockBudgetRepo.On("GetByUserIdAndMonth", mock.Anything, mock.Anything, mock.Anything).Return(budget, nil)
	mockTxRepo.On("GetTransactionByUserId", mock.Anything).Return(transaction, nil)

	//Act
	_, err := svc.GetByUserIdAndMonth(id, 2, 3)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_UpdateBudgetSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		amount         = 1000.0
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("UpdateBudget", mock.Anything, mock.Anything).Return(budget, nil)

	//Act
	_, err := svc.UpdateBudget(amount, id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_UpdateBudgetFail(t *testing.T) {
	//Arrange
	var (
		id             = 1
		amount         = 1000.0
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
		budget         = models.Budget{}
	)
	mockBudgetRepo.On("UpdateBudget", mock.Anything, mock.Anything).Return(budget, errors.New("budget not found"))

	//Act
	_, err := svc.UpdateBudget(amount, id)

	//Assert
	assert.Equal(t, errors.New("budget not found"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)

}

func Test_DeleteBudgetByIdSuccess(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
	)
	mockBudgetRepo.On("DeleteBudgetById", id).Return(nil)

	//Act
	err := svc.DeleteBudgetById(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockBudgetRepo.AssertExpectations(t)
}

func Test_DeleteBudgetByIdFail(t *testing.T) {
	//Arrange
	var (
		id             = 1
		mockBudgetRepo = new(mocks.MockBudgetRepo)
		mockTxRepo     = new(mocks.MockTransactionRepo)
		svc            = NewBudgetService(mockBudgetRepo, mockTxRepo)
	)
	mockBudgetRepo.On("DeleteBudgetById", id).Return(errors.New("budget not found"))

	//Act
	err := svc.DeleteBudgetById(id)

	//Assert
	assert.Equal(t, errors.New("budget not found"), err)
	assert.Error(t, err)
	mockBudgetRepo.AssertExpectations(t)
}
