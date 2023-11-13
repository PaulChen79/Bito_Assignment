package handler

import (
	"bito-assignment/domain"
	"bito-assignment/internal/provider"
	"log"

	"github.com/gin-gonic/gin"
)

func RemoveSinglePerson(c *gin.Context) {
	userName := c.Param("user_name")

	svc, err := provider.NewService()
	if err != nil {
		log.Println("RemoveSinglePerson error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	errFormat := svc.RemoveSinglePerson(userName)
	if errFormat != nil {
		log.Println("RemoveSinglePerson error: ", errFormat.Message)
		Failed(c, domain.ErrorServer, errFormat.Message)
		return
	}

	Success(c, nil)
}
