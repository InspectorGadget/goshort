package controllers

import (
	"net/http"

	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/gin-gonic/gin"
)

func AddUrlToUser(c *gin.Context) {
	userId := c.Param("id")
	var newUrl models.Url
	var existingUser models.User

	// Check for existing user
	if err := initializers.DB.Model(&models.User{}).Where("id = ?", userId).First(&existingUser).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "User not found!",
			},
		)
		return
	}

	// Check for POST data
	if err := c.ShouldBindJSON(&newUrl); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "One or more parameters are missing!",
				"error":   err.Error(),
			},
		)
		return
	}

	// Check if Url short is unique
	var existingUrlCheck models.Url
	if err := initializers.DB.Model(&models.Url{}).Where("short = ?", newUrl.Short).First(&existingUrlCheck).Error; err == nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"message": "URL short conflicts!",
			},
		)
		return
	}

	// Set user ID for the URL
	newUrl.UserID = existingUser.ID

	// Update in DB
	if err := initializers.DB.Create(&newUrl).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "An error has occurred while adding URL",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "URL successfully created",
		},
	)
}
