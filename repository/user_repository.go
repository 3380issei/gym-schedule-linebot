package repository

import (
	"gym-schedule-linebot/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	result := ur.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
