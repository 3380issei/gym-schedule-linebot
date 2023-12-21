package controller

import (
	"fmt"
	"gym-schedule-linebot/model"
	"gym-schedule-linebot/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type LinebotController interface {
	CatchEvents(c *gin.Context)
}

type linebotController struct {
	bot *linebot.Client
	uu  usecase.UserUsecase
}

func NewLinebotController(uu usecase.UserUsecase) *linebotController {
	secret := os.Getenv("LINE_CHANNEL_SECRET")
	token := os.Getenv("LINE_CHANNEL_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		log.Fatal(err)
	}

	return &linebotController{
		bot: bot,
		uu:  uu,
	}
}

func (lc *linebotController) CatchEvents(c *gin.Context) {

	cb, err := webhook.ParseRequest(os.Getenv("LINE_CHANNEL_SECRET"), c.Request)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.FollowEvent: // when user add bot as friend
			lc.createUser(e, c)
		}
	}
}

func (lc *linebotController) createUser(fe webhook.FollowEvent, c *gin.Context) {
	switch s := fe.Source.(type) {
	case webhook.UserSource:
		user := model.User{
			ID: s.UserId,
		}

		if err := lc.uu.CreateUser(&user); err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{"messsage": "success"})
	}
}
