package model

import "time"

type Schedule struct {
	ID       string    `json:"id" gorm:"primary_key"`
	Deadline time.Time `json:"deadline" gorm:"not null"`

	GymID string `json:"gym_id" gorm:"not null"`
	Gym   Gym    `gorm:"foreignkey:GymID"`

	UserID string `json:"user_id" gorm:"not null"`
	User   User   `gorm:"foreignkey:UserID"`
}
