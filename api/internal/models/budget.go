package models

type Budget struct {
	Id         int     `json:"id" db:"id"`
	UserId     int     `json:"user_id" db:"user_id"`
	CategoryId int     `json:"category_id" db:"category_id"`
	Amount     float64 `json:"amount" db:"amount"`
	Month      int     `json:"month" db:"month"`
	Year       int     `json:"year" db:"year"`
}
