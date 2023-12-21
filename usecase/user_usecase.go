package usecase

import (
	"gym-schedule-linebot/model"
	"gym-schedule-linebot/repository"
)

type UserUsecase interface {
	CreateUser(user *model.User) error
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *userUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) CreateUser(user *model.User) error {
	if err := uu.ur.CreateUser(user); err != nil {
		return err
	}
	return nil
}
