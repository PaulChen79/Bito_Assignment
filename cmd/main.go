package main

import (
	"bito-assignment/internal/gin/route"
	"bito-assignment/internal/provider"
)

func main() {
	config := provider.NewConfig()

	router := route.SetupRouter()
	router.Run(config.Gin.Host + ":" + config.Gin.Port)
}
