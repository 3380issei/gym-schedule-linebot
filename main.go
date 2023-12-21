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
	userUsecase := usecase.NewUserUsecase(userRepository)
	linebotController := controller.NewLinebotController(userUsecase)

	router := router.NewRouter(linebotController)

	router.Run()
}
