package model

import (
	"github.com/shopspring/decimal"
)

type Gym struct {
	ID        string          `json:"id" gorm:"primary_key"`
	Title     string          `json:"title" gorm:"not null"`
	Address   string          `json:"address"`
	Latitude  decimal.Decimal `json:"latitude" gorm:"not null"`
	Longitude decimal.Decimal `json:"longitude" gorm:"not null"`
}
