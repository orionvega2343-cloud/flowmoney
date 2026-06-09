package repository

import (
	"flowmoney/api/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(u models.User) (models.User, error) {
	err := r.db.QueryRow(`INSERT INTO users(balance,email,password,created_at) VALUES($1,$2,$3,$4) RETURNING id`, u.Balance, u.Email, u.Password, u.CreatedAt).Scan(&u.Id)
	if err != nil {
		return models.User{}, err
	}
	return u, err

}

func (r *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var u models.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepo) GetUserById(id int) (models.User, error) {
	var u models.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE id = $1`, id)

	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepo) UpdateBalance(id int, balance float64) (models.User, error) {
	_, err := r.db.Exec(`UPDATE users SET balance=$1 WHERE id=$2`, balance, id)
	res, err := r.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
