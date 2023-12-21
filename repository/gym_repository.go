package repository

import (
	"gym-schedule-linebot/model"

	"gorm.io/gorm"
)

type GymRepository interface {
	CreateGym(gym *model.Gym) error
}

type gymRepository struct {
	db *gorm.DB
}

func NewGymRepository(db *gorm.DB) *gymRepository {
	return &gymRepository{db}
}

func (gr *gymRepository) CreateGym(gym *model.Gym) error {
	result := gr.db.Create(gym)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
