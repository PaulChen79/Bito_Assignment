package handler

import (
	"bito-assignment/domain"
	"bito-assignment/internal/provider"
	"bito-assignment/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func QuerySinglePerson(c *gin.Context) {
	number := c.Param("number")
	userName := c.Param("user_name")

	n, err := utils.StringToUint(number)
	if err != nil {
		log.Println("QuerySinglePerson error: ", err)
		Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	svc, err := provider.NewService()
	if err != nil {
		log.Println("QuerySinglePerson error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	matches, errFormat := svc.QuerySinglePerson(userName, n)
	if errFormat != nil {
		log.Println("QuerySinglePerson error: ", errFormat.Message)
		Failed(c, domain.ErrorServer, errFormat.Message)
		return
	}

	Success(c, matches)
}
