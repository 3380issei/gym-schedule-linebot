package model

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name"`
	GymID    string `json:"gym_id"`
	Gym      Gym    `json:"gym" gorm:"foreignkey:GymID"`
}
