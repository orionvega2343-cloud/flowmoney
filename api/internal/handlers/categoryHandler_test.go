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

func Test_CreateCategorySuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/category", body)
	)
	mockSvc.On("CreateCategory", mock.Anything).Return(models.Category{}, nil)

	//Act
	hndlr.CreateCategory(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_CreateCategoryFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/category", body)
	)
	mockSvc.On("CreateCategory", mock.Anything).Return(models.Category{}, errors.New("some error"))

	//Act
	hndlr.CreateCategory(w, r)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetCategoryByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/category/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetCategoryById", mock.Anything).Return(models.Category{}, nil)

	//Act
	hndlr.GetCategoryById(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetCategoryByIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/category/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetCategoryById", mock.Anything).Return(models.Category{}, errors.New("some error"))

	//Act
	hndlr.GetCategoryById(w, r)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetByUserIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/category/user/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetByUserId", mock.Anything).Return([]models.Category{}, nil)

	//Act
	hndlr.GetByUserId(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetByUserIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/category/user/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetByUserId", mock.Anything).Return([]models.Category{}, errors.New("some error"))

	//Act
	hndlr.GetByUserId(w, r)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateCategorySuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/category/update/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateCategory", mock.Anything, mock.Anything).Return(models.Category{}, nil)

	//Act
	hndlr.UpdateCategory(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateCategoryFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockCategoryService)
		hndlr   = NewCategoryHandler(mockSvc)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/category/update/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateCategory", mock.Anything, mock.Anything).Return(models.Category{}, errors.New("some error"))

	//Act
	hndlr.UpdateCategory(w, r)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}
