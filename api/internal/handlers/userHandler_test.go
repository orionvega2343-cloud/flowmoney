package handlers

import (
	"errors"
	"flowmoney/api/internal/config"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateUserSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/tasks", body)
	)
	mockSvc.On("CreateUser", mock.Anything).Return(models.User{}, nil)

	//Act
	hndlr.CreateUser(w, r)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_CreateUserFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/tasks", body)
	)
	mockSvc.On("CreateUser", mock.Anything).Return(models.User{}, errors.New("some error"))

	//Act
	hndlr.CreateUser(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_LoginSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/login", body)
	)
	mockSvc.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("token-string", nil)

	//Act
	hndlr.Login(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_LoginFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("POST", "/login", body)
	)
	mockSvc.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("some error"))

	//Act
	hndlr.Login(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetUserByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/user/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetUserById", mock.Anything).Return(models.User{}, nil)

	//Act
	hndlr.GetUserById(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_GetUserByIdFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		r       = httptest.NewRequest("GET", "/user/1", nil)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("GetUserById", mock.Anything).Return(models.User{}, errors.New("some error"))

	//Act
	hndlr.GetUserById(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateBalanceSuccess(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/user/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateBalance", mock.Anything, mock.Anything).Return(models.User{}, nil)

	//Act
	hndlr.UpdateBalance(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func Test_UpdateBalanceFail(t *testing.T) {
	//Arrange
	var (
		mockSvc = new(mocks.MockUserService)
		cfg     = config.Jwt{}
		hndlr   = NewUserHandler(mockSvc, cfg)
		w       = httptest.NewRecorder()
		body    = strings.NewReader("{}")
		r       = httptest.NewRequest("PUT", "/user/1", body)
	)
	r.SetPathValue("id", "1")
	mockSvc.On("UpdateBalance", mock.Anything, mock.Anything).Return(models.User{}, errors.New("some error"))

	//Act
	hndlr.UpdateBalance(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}
