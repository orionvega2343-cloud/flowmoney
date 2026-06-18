package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

// Category moocks
type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) GetCategoryById(id int) (models.Category, error) {
	args := m.Called(id)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryRepo) CreateCategory(c models.Category) (models.Category, error) {
	args := m.Called(c)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryRepo) GetByUserId(userId int) ([]models.Category, error) {
	args := m.Called(userId)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepo) UpdateCategory(id int, title string) (models.Category, error) {
	args := m.Called(id, title)
	return args.Get(0).(models.Category), args.Error(1)
}
