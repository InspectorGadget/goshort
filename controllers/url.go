package controllers

import (
	"net/http"

	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/InspectorGadget/goshort/structs"
	"github.com/gin-gonic/gin"
)

func AddUrlToUser(c *gin.Context) {
	userId := c.Param("id")
	var urlObject structs.AddUrlRequest
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
	if err := c.ShouldBindJSON(&urlObject); err != nil {
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
	if err := initializers.DB.Model(&models.Url{}).Where("short = ?", &urlObject.Short).First(&models.Url{}).Error; err == nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"message": "URL short conflicts!",
			},
		)
		return
	}

	// Check if URL has http or https
	if urlObject.Url != "" && !(urlObject.Url[:4] == "http" || urlObject.Url[:5] == "https") {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "URL must start with http:// or https://",
			},
		)
		return
	}

	// Set user ID for the URL and add in DB
	newUrl := models.Url{
		UserID: existingUser.ID,
		Short:  urlObject.Short,
		Url:    urlObject.Url,
	}
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

func ListUrlByUser(c *gin.Context) {
	userId := c.Param("id")

	var urls []models.Url
	if err := initializers.DB.Model(&models.Url{}).Where("user_id = ?", userId).Find(&urls).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "Not urls for the user",
			},
		)
	}

	var urlsResponse []gin.H
	for _, url := range urls {
		urlsResponse = append(urlsResponse, url.Serialize())
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"urls": urlsResponse,
		},
	)
}

func DeleteUrlByUser(c *gin.Context) {
	userId := c.Param("id")
	urlId := c.Param("urlid")

	// Check if user exists
	if err := initializers.DB.Model(&models.User{}).Where("id = ?", userId).First(&models.User{}).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "User does not exist!",
			},
		)
		return
	}

	// Check if URL with short exists
	var existingUrl models.Url
	if err := initializers.DB.Model(&models.Url{}).Where("id = ? && user_id = ?", urlId, userId).First(&existingUrl).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "URL does not exist!",
			},
		)
		return
	}

	// Delete the URL
	if err := initializers.DB.Delete(&existingUrl).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "An error has occurred while deleting URL from user",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "URL successfully deleted.",
		},
	)
}
