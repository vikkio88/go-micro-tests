package actions

import (
	"backoffice-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PagesGetAll action to get all pages filtering them
func PagesGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"a": "a"})
}

// PageCreate create a page
func PageCreate(c *gin.Context) {
	var page models.Page
	c.BindJSON(&page)
	valid, validation := page.IsValid()

	if !valid {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validation.Errors})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"accepted": true})
}
