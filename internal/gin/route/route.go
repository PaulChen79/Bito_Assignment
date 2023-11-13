package route

import (
	"bito-assignment/internal/gin/handler"
	"bito-assignment/internal/gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/system_matcher")
	{
		api.POST("/add_and_match", handler.AddSinglePersonAndMatch)
		api.DELETE("/:user_name/remove", handler.RemoveSinglePerson)
		api.GET("/query_single_person/:user_name/:number", handler.QuerySinglePerson)
	}

	return r
}
