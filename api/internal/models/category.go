package models


type Category struct {
	Id int `db:"id" json:"id"`
	Title string `json:"title" db:"title"`
	UserId int `json:"user_id" db:"user_id"`

}