package controllers

import (
	"net/http"

	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Welcome to GoShort API",
		},
	)
}

func RedirectClient(c *gin.Context) {
	short := c.Param("short")

	// Check if short exists in the DB
	var existingUrl models.Url
	if err := initializers.DB.Model(&models.Url{}).Where("short = ?", short).First(&existingUrl).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "URL not found.",
			},
		)
		return
	}

	// Replace address with the original URL
	c.Redirect(http.StatusFound, existingUrl.Url)
}
