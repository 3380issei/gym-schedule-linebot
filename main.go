package main

import (
	"gym-schedule-linebot/controller"
	"gym-schedule-linebot/db"
	"gym-schedule-linebot/repository"
	"gym-schedule-linebot/router"
	"gym-schedule-linebot/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	gymRepository := repository.NewGymRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	gymUsecase := usecase.NewGymUsecase(gymRepository)
	linebotController := controller.NewLinebotController(userUsecase, gymUsecase)

	router := router.NewRouter(linebotController)

	router.Run()
}
