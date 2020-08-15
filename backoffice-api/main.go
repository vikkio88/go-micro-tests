package main

import (
	"backoffice-api/actions"
	"backoffice-api/libs/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	appConfig := config.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(func(c *gin.Context) {
		c.Set("config", appConfig)
	})

	router.GET("/ping", actions.Pong)

	router.Run()
}
