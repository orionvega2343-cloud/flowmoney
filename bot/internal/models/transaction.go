package models

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Amount     float64   `json:"amount"`
	Type       string    `json:"type"`
	Date       time.Time `json:"date"`
	CategoryId int       `json:"category_id"`
}
