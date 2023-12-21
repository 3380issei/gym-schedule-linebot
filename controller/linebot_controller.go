package controller

import (
	"fmt"
	"gym-schedule-linebot/model"
	"gym-schedule-linebot/usecase"
	"log"
	"net/http"
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
			lc.handleFollowEvent(e, c)
		}
	}
}

func (lc *linebotController) handleFollowEvent(fe webhook.FollowEvent, c *gin.Context) {
	switch s := fe.Source.(type) {
	case webhook.UserSource:
		userID := s.UserId
		profile, err := lc.fetchUserProfileByID(userID)
		if err != nil {
			fmt.Printf("failed to get user profile: %v", err)
		}

		user := &model.User{
			ID:       userID,
			UserName: profile.DisplayName,
		}

		if err := lc.uu.CreateUser(user); err != nil {
			fmt.Printf("failed to create user: %v", err)
		}

		c.JSON(http.StatusCreated, user)
	}
}

func (lc *linebotController) fetchUserProfileByID(userID string) (*linebot.UserProfileResponse, error) {
	res, err := lc.bot.GetProfile(userID).Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}
