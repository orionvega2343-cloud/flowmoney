package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

// User mocks
type MockUserRepo struct {
	mock.Mock
}

func (r *MockUserRepo) GetUserById(userId int) (models.User, error) {
	args := r.Called(userId)
	return args.Get(0).(models.User), args.Error(1)
}

func (r *MockUserRepo) GetUserByEmail(email string) (models.User, error) {
	args := r.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

func (r *MockUserRepo) CreateUser(user models.User) (models.User, error) {
	args := r.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (r *MockUserRepo) UpdateBalance(id int, balance float64) (models.User, error) {
	args := r.Called(id, balance)
	return args.Get(0).(models.User), args.Error(1)
}
