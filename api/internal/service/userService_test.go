package service

import (
	"errors"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func Test_CreateUserSuccess(t *testing.T) {
	//Arrange
	mockRepo := new(mocks.MockUserRepo)
	svc := NewUserService(mockRepo)
	expectedValue := models.User{}
	mockRepo.On("CreateUser", mock.Anything).Return(expectedValue, nil)

	//Act
	_, err := svc.CreateUser(expectedValue)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_CreateUserFail(t *testing.T) {
	//Arrange
	mockRepo := new(mocks.MockUserRepo)
	svc := NewUserService(mockRepo)
	expectedValue := models.User{}
	mockRepo.On("CreateUser", mock.Anything).Return(expectedValue, errors.New("some error"))

	//Act
	_, err := svc.CreateUser(expectedValue)

	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)

}

func Test_LoginSuccess(t *testing.T) {
	//Arrange
	var (
		email    string
		password = "12345"
		secret   string
	)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	mockRepo := new(mocks.MockUserRepo)
	svc := NewUserService(mockRepo)
	mockRepo.On("GetUserByEmail", mock.Anything).Return(models.User{Password: string(hash)}, nil)

	//Act
	_, err = svc.Login(email, password, secret)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_LoginFail(t *testing.T) {
	//Arrange
	var (
		email    string
		password string
		secret   string
	)
	mockRepo := new(mocks.MockUserRepo)
	svc := NewUserService(mockRepo)
	mockRepo.On("GetUserByEmail", mock.Anything).Return(models.User{}, errors.New("some error"))

	//Act
	_, err := svc.Login(email, password, secret)

	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)

}

func Test_GetUserByIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockRepo = new(mocks.MockUserRepo)
		svc      = NewUserService(mockRepo)
		user     = models.User{}
		id       = 1
	)
	mockRepo.On("GetUserById", mock.Anything).Return(user, nil)

	//Act
	_, err := svc.GetUserById(id)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_GetUserByIdFail(t *testing.T) {
	//Arrange
	var (
		id       = 1
		mockRepo = new(mocks.MockUserRepo)
		svc      = NewUserService(mockRepo)
		user     = models.User{}
	)
	mockRepo.On("GetUserById", mock.Anything).Return(user, errors.New("some error"))

	//Act
	_, err := svc.GetUserById(id)

	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_UpdateBalanceSuccess(t *testing.T) {
	//Arrange
	var (
		id       = 1
		balance  float64
		mockRepo = new(mocks.MockUserRepo)
		svc      = NewUserService(mockRepo)
		user     = models.User{}
	)
	mockRepo.On("UpdateBalance", mock.Anything).Return(user, nil)

	//Act
	_, err := svc.UpdateBalance(id, balance)

	//Assert
	assert.Equal(t, nil, err)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func Test_UpdateBalanceFail(t *testing.T) {
	//Arrange
	var (
		id       = 1
		balance  float64
		mockRepo = new(mocks.MockUserRepo)
		svc      = NewUserService(mockRepo)
		user     = models.User{}
	)
	mockRepo.On("UpdateBalance", mock.Anything).Return(user, errors.New("some error"))

	//Act
	_, err := svc.UpdateBalance(id, balance)

	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
