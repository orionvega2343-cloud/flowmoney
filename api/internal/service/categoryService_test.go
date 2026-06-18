package service

import (
	"errors"
	"flowmoney/api/internal/mocks"
	"flowmoney/api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateCategorySuccess(t *testing.T) {
	//Arrange
	var (
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("CreateCategory", mock.Anything).Return(category, nil)
	//Act
	_, err := svc.CreateCategory(category)
	//Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_CreateCategoryFail(t *testing.T) {
	//Arrange
	var (
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("CreateCategory", mock.Anything).Return(models.Category{}, errors.New("some error"))
	//Act
	_, err := svc.CreateCategory(category)
	//Assert
	assert.Equal(t, errors.New("some error"), err)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_GetCategoryByIdSuccess(t *testing.T) {
	//Arrange
	var (
		id       = 1
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("GetCategoryById", id).Return(category, nil)
	//Act
	_, err := svc.GetCategoryById(id)
	//Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_GetCategoryByIdFail(t *testing.T) {
	//Arrange
	var (
		id       = 1
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("GetCategoryById", id).Return(category, errors.New("some error"))
	//Act
	_, err := svc.GetCategoryById(id)
	//Assert
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_GetByUserIdSuccess(t *testing.T) {
	//Arrange
	var (
		mockRepo   = new(mocks.MockCategoryRepo)
		svc        = NewCategoryService(mockRepo)
		id         = 1
		categories = []models.Category{}
	)
	mockRepo.On("GetByUserId", id).Return(categories, nil)
	//Act
	_, err := svc.GetByUserId(id)
	//Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_GetByUserIdFail(t *testing.T) {
	//Arrange
	var (
		id         = 1
		mockRepo   = new(mocks.MockCategoryRepo)
		svc        = NewCategoryService(mockRepo)
		categories = []models.Category{}
	)
	mockRepo.On("GetByUserId", id).Return(categories, errors.New("some error"))
	//Act
	_, err := svc.GetByUserId(id)
	//Assert
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_UpdateCategorySuccess(t *testing.T) {
	//Arrange
	var (
		id       = 1
		title    string
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("GetCategoryById", id).Return(category, nil)
	mockRepo.On("UpdateCategory", id, title).Return(category, nil)
	//Act
	_, err := svc.UpdateCategory(id, title)
	//Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_UpdateCategoryFail(t *testing.T) {
	//Arrange
	var (
		id       = 1
		title    string
		mockRepo = new(mocks.MockCategoryRepo)
		svc      = NewCategoryService(mockRepo)
		category = models.Category{}
	)
	mockRepo.On("GetCategoryById", id).Return(category, errors.New("some error"))
	//Act
	_, err := svc.UpdateCategory(id, title)
	//Assert
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
