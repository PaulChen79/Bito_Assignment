package route

import (
	"bito-assignment/internal/gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())

	// api := r.Group("/system_matcher")
	// {
	// 	api.GET("ping", home.Ping)
	// }

	return r
}
