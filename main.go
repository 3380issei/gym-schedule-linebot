package main

import (
	"fmt"
	"gym-schedule-linebot/db"
)

func main() {
	db := db.NewDB()
	fmt.Println(db)
}
