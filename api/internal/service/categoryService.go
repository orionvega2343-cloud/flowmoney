package service

import (
	"flowmoney/api/internal/models"
)

type CategoryService interface {
	CreateCategory(c models.Category) (models.Category, error)
	GetCategoryById(id int) (models.Category, error)
	GetByUserId(id int) ([]models.Category, error)
	UpdateCategory(id int, title string) (models.Category, error)
}

type CategoryServiceImpl struct {
	Cr CategoryService
}

func NewCategoryService(cr CategoryService) *CategoryServiceImpl {
	return &CategoryServiceImpl{Cr: cr}
}

func (r *CategoryServiceImpl) CreateCategory(c models.Category) (models.Category, error) {
	res, err := r.Cr.CreateCategory(c)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func (r *CategoryServiceImpl) GetCategoryById(id int) (models.Category, error) {
	res, err := r.Cr.GetCategoryById(id)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func (r *CategoryServiceImpl) GetByUserId(id int) ([]models.Category, error) {
	res, err := r.Cr.GetByUserId(id)
	if err != nil {
		return []models.Category{}, err
	}
	return res, nil
}

func (r *CategoryServiceImpl) UpdateCategory(id int, title string) (models.Category, error) {
	_, err := r.Cr.GetCategoryById(id)
	if err != nil {
		return models.Category{}, err
	}
	res, err := r.Cr.UpdateCategory(id, title)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}
