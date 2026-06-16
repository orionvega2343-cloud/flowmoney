package repository

import (
	"flowmoney/api/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	CreateTransaction(tr models.Transaction) (models.Transaction, error)
	GetTransactionById(id int) (models.Transaction, error)
	GetTransactionByUserId(userId int) ([]models.Transaction, error)
}
type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) CreateTransaction(tr models.Transaction) (models.Transaction, error) {
	err := r.db.QueryRow(`INSERT INTO transactions(user_id,amount,type,date,category_id) VALUES($1,$2,$3,$4,$5) RETURNING id`, tr.UserId, tr.Amount, tr.Type, tr.Date, tr.CategoryId).Scan(&tr.Id)
	if err != nil {
		return models.Transaction{}, err
	}
	return tr, nil
}

func (r *TransactionRepo) GetTransactionById(id int) (models.Transaction, error) {
	var tr models.Transaction
	err := r.db.Get(&tr, "SELECT * FROM transactions WHERE id = $1", id)
	if err != nil {
		return models.Transaction{}, err
	}
	return tr, nil
}

func (r *TransactionRepo) GetTransactionByUserId(userId int) ([]models.Transaction, error) {
	var tr []models.Transaction
	err := r.db.Select(&tr, "SELECT * FROM transactions WHERE user_id = $1", userId)
	if err != nil {
		return []models.Transaction{}, err
	}
	return tr, nil
}
