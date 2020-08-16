package actions

import (
	"backoffice-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PagesGetAll action to get all pages filtering them
func PagesGetAll(c *gin.Context) {
	pages := models.GetAllPages()
	c.JSON(http.StatusOK, gin.H{"data": pages})
}

// PageCreate create a page
func PageCreate(c *gin.Context) {
	var page models.Page
	c.BindJSON(&page)
	valid, errors := page.IsValid()

	if !valid {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"accepted": true})
}
