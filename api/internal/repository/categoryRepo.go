package repository

import (
	"flowmoney/api/internal/models"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) CreateCategory(c models.Category) (models.Category, error) {
	err := r.db.QueryRow(`INSERT INTO categories(title,user_id) VALUES($1,$2) RETURNING id`, c.Title, c.UserId).Scan(&c.Id)
	if err != nil {
		return models.Category{}, err
	}
	return c, nil

}

func (r *CategoryRepo) GetCategoryById(id int) (models.Category, error) {
	var c models.Category
	err := r.db.Get(&c, "SELECT * FROM categories WHERE id=$1", id)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *CategoryRepo) GetByUserId(id int) ([]models.Category, error) {
	var c []models.Category
	err := r.db.Select(&c, "SELECT * FROM categories WHERE user_id=$1", id)
	if err != nil {
		return c, err
	}
	return c, nil

}

func (r *CategoryRepo) UpdateCategory(id int, title string) (models.Category, error) {
	_, err := r.db.Exec(`UPDATE categories SET title=$1 WHERE id=$2`, title, id)
	if err != nil {
		return models.Category{}, err
	}
	return r.GetCategoryById(id)
}
