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

func Test_CreateBudgetSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/budgets", body)
	)
	mockSvc.On("CreateBudget", mock.Anything).Return(models.Budget{}, nil)

	//Act
	hndlr.CreateBudget(w, r)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_CreateBudgetFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/budgets", body)
	)
	mockSvc.On("CreateBudget", mock.Anything).Return(models.Budget{}, errors.New("some error"))

	//Act
	hndlr.CreateBudget(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetBudgetByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetBudgetById", mock.Anything).Return(models.Budget{}, nil)

	//Act
	hndlr.GetBudgetById(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetBudgetByIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetBudgetById", mock.Anything).Return(models.Budget{}, errors.New("some error"))

	//Act
	hndlr.GetBudgetById(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetBudgetByCategoryIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/budgets/1/categories", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetBudgetByCategoryId", mock.Anything).Return(models.Budget{}, nil)

	//Act
	hndlr.GetBudgetByCategoryId(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetBudgetByCategoryIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("GET", "/budgets/1/categories", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetBudgetByCategoryId", mock.Anything).Return(models.Budget{}, errors.New("some error"))

	//Act
	hndlr.GetBudgetByCategoryId(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetByUserIdAndMonthSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{\"month\": 5, \"year\": 2026}")
		r       = httptest.NewRequest("GET", "/budgets/1/categories", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetByUserIdAndMonth", mock.Anything, mock.Anything, mock.Anything).Return(models.Budget{}, nil)

	//Act
	hndlr.GetByUserIdAndMonth(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetByUserIdAndMonthFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{\"month\": 5, \"year\": 2026}")
		r       = httptest.NewRequest("GET", "/budgets/1/categories", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetByUserIdAndMonth", mock.Anything, mock.Anything, mock.Anything).Return(models.Budget{}, errors.New("some error"))

	//Act
	hndlr.GetByUserIdAndMonth(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateBudgetSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateBudget", mock.Anything, mock.Anything).Return(models.Budget{}, nil)

	//Act
	hndlr.UpdateBudget(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateBudgetFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateBudget", mock.Anything, mock.Anything).Return(models.Budget{}, errors.New("some error"))

	//Act
	hndlr.UpdateBudget(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_DeleteBudgetByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("DELETE", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("DeleteBudgetById", mock.Anything).Return(nil)

	//Act
	hndlr.DeleteBudgetById(w, r)

	//Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_DeleteBudgetByIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockBudgetService)
		hndlr   = NewBudgetHandlerImpl(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("DELETE", "/budgets/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("DeleteBudgetById", mock.Anything).Return(errors.New("some error"))

	//Act
	hndlr.DeleteBudgetById(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}
