package usecase

import (
	"gym-schedule-linebot/model"
	"gym-schedule-linebot/repository"
)

type GymUsecase interface {
	CreateGym(gym *model.Gym) error
}

type gymUsecase struct {
	gr repository.GymRepository
}

func NewGymUsecase(gr repository.GymRepository) *gymUsecase {
	return &gymUsecase{gr}
}

func (gu *gymUsecase) CreateGym(gym *model.Gym) error {
	if err := gu.gr.CreateGym(gym); err != nil {
		return err
	}
	return nil
}
