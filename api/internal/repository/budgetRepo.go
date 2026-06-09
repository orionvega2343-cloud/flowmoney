package repository

import (
	"flowmoney/api/internal/models"

	"github.com/jmoiron/sqlx"
)

type BudgetRepo struct {
	db *sqlx.DB
}

func NewBudgetRepo(db *sqlx.DB) *BudgetRepo {
	return &BudgetRepo{db: db}
}

func (r *BudgetRepo) CreateBudget(b models.Budget) (models.Budget, error) {
	err := r.db.QueryRow(`INSERT INTO budgets(user_id,category_id,amount,month,year) VALUES($1,$2,$3,$4,$5) RETURNING id`, b.UserId, b.CategoryId, b.Amount, b.Month, b.Year).Scan(&b.Id)
	if err != nil {
		return models.Budget{}, err
	}
	return b, nil
}

func (r *BudgetRepo) GetBudgetById(id int) (models.Budget, error) {
	var budget models.Budget
	err := r.db.Get(&budget, "SELECT * FROM budgets WHERE id = $1", id)
	if err != nil {
		return models.Budget{}, err
	}
	return budget, nil
}

func (r *BudgetRepo) GetBudgetByCategoryId(catId int) (models.Budget, error) {
	var budget models.Budget
	err := r.db.Get(&budget, "SELECT * FROM budgets WHERE category_id = $1", catId)
	if err != nil {
		return models.Budget{}, err
	}
	return budget, nil
}
func (r *BudgetRepo) GetByUserIdAndMonth(userId int, month int, year int) (models.Budget, error) {
	var budget models.Budget
	err := r.db.Get(&budget, "SELECT * FROM budgets WHERE user_id = $1 AND month = $2 AND year = $3", userId, month, year)
	if err != nil {
		return models.Budget{}, err
	}
	return budget, nil
}

func (r *BudgetRepo) UpdateBudget(amount float64, id int) (models.Budget, error) {
	_, err := r.db.Exec(`UPDATE budgets SET amount=$1 WHERE id=$2`, amount, id)
	res, err := r.GetBudgetById(id)
	if err != nil {
		return models.Budget{}, err
	}
	return res, nil
}

func (r *BudgetRepo) DeleteBudgetById(id int) error {
	_, err := r.db.Exec(`DELETE FROM budgets WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
