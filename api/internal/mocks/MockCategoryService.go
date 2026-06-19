package mocks

import (
	"flowmoney/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) CreateCategory(c models.Category) (models.Category, error) {
	args := m.Called(c)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryService) GetCategoryById(id int) (models.Category, error) {
	args := m.Called(id)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryService) GetByUserId(id int) ([]models.Category, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) UpdateCategory(id int, title string) (models.Category, error) {
	args := m.Called(id, title)
	return args.Get(0).(models.Category), args.Error(1)
}
