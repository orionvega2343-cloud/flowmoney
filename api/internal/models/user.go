package models

import "time"

type User struct {
	Id        int `json:"id" db:"id"`
	Balance   float64 `json:"balance" db:"balance"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
