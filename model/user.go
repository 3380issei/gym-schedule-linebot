package model

type User struct {
	ID string `json:"id" gorm:"primary_key"`
}