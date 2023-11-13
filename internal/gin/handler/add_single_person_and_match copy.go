package handler

import (
	"bito-assignment/domain"
	"bito-assignment/internal/provider"
	"log"

	"github.com/gin-gonic/gin"
)

type AddSinglePersonAndMatchRequest struct {
	Name        string `json:"name" binding:"required"`
	Height      int    `json:"height" binding:"required,gte=0"`
	Gender      string `json:"gender" binding:"required,oneof=female male"`
	WantedDates int    `json:"wanted_dates" binding:"required,gte=0"`
}

func AddSinglePersonAndMatch(c *gin.Context) {
	var req AddSinglePersonAndMatchRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println("AddSinglePersonAndMatch error: ", err)
		Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	newUser := domain.User{
		Name:           req.Name,
		Height:         req.Height,
		Gender:         domain.Gender(req.Gender),
		RemainingDates: req.WantedDates,
	}

	svc, err := provider.NewService()
	if err != nil {
		log.Println("AddSinglePersonAndMatch error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	matches, errFormat := svc.AddSinglePersonAndMatch(&newUser)
	if errFormat != nil {
		log.Println("AddSinglePersonAndMatch error: ", errFormat.Message)
		Failed(c, domain.ErrorServer, errFormat.Message)
		return
	}

	Success(c, matches)
}
