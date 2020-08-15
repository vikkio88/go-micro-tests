package actions

import (
	"backoffice-api/libs/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Pong action
func Pong(c *gin.Context) {
	config := c.MustGet("config").(config.Config)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"service": config.Get("APP_NAME"),
	})
}
