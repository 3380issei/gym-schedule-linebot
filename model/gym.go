package model

type Gym struct {
	ID        string  `json:"id" gorm:"primary_key"`
	Title     string  `json:"title" gorm:"not null"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude" gorm:"not null"`
	Longitude float64 `json:"longitude" gorm:"not null"`

	UserID string `json:"user_id" gorm:"not null"`
	User   User   `gorm:"foreignkey:UserID"`
}
