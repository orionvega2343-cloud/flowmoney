package models

import "time"

type Transaction struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"user_id"`
	Amount     float64   `json:"amount" db:"amount"`
	Type       string    `json:"type" db:"type"`
	Date       time.Time `json:"date" db:"date"`
	CategoryId int       `json:"category_id" db:"category_id"`
}
