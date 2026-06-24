package models

type Category struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}
