package router

import (
	"gym-schedule-linebot/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(lc controller.LinebotController) *gin.Engine {
	r := gin.Default()

	r.POST("/webhook", lc.CatchEvents)

	return r
}
