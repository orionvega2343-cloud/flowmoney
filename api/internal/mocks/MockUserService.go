package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) Login(email string, password string, secret string) (string, error) {
	args := m.Called(email, password, secret)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) GetUserById(id int) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) UpdateBalance(id int, balance float64) (models.User, error) {
	args := m.Called(id, balance)
	return args.Get(0).(models.User), args.Error(1)
}
