package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Balance   float64   `json:"balance"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
