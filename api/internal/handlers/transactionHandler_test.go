package handlers

import (
	"errors"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateTransactionSuccess(t *testing.T) {
	//Arranga
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/transactions", body)
	)
	mockSvc.On("CreateTransaction", mock.Anything).Return(models.Transaction{}, nil)

	//Act
	hndlr.CreateTransaction(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)

}

func Test_CreateTransactionFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/transactions", body)
	)
	mockSvc.On("CreateTransaction", mock.Anything).Return(models.Transaction{}, errors.New("some error"))

	//Act
	hndlr.CreateTransaction(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)

}

func Test_GetTransactionByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/transactions/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetTransactionById", mock.Anything).Return(models.Transaction{}, nil)

	//Act
	hndlr.GetTransactionById(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetTransactionByIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/transactions/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetTransactionById", mock.Anything).Return(models.Transaction{}, errors.New("some error"))

	//Act
	hndlr.GetTransactionById(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetTransactionByUserIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/transactions/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetTransactionByUserId", mock.Anything).Return([]models.Transaction{}, nil)

	//Act
	hndlr.GetTransactionByUserId(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)

}

func Test_GetTransactionByUserIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockTransactionService)
		hndlr   = NewTransactionHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/transactions/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetTransactionByUserId", mock.Anything).Return([]models.Transaction{}, errors.New("some error"))

	//Act
	hndlr.GetTransactionByUserId(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)

}
