package controllers

import (
	"net/http"
	"time"

	"github.com/InspectorGadget/goshort/helpers"
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/InspectorGadget/goshort/structs"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	var authenticationObject structs.AuthenticationRequest
	if err := c.ShouldBind(&authenticationObject); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "One or more parameters are required.",
			},
		)
		return
	}

	// Check credentials
	var existingUser models.User
	if err := initializers.DB.Model(&models.User{}).Where("username = ? && password = ?", &authenticationObject.Username, &authenticationObject.Password).First(&existingUser).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Invalid credentials or user is not found",
			},
		)
		return
	}

	// Create token for user
	expiresAt := time.Now().Add(time.Hour * 24)
	tokenString, err := helpers.GenerateJWT(existingUser.Username, expiresAt)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	// Additional: Check if JWT token exist
	var existingToken models.Token
	if err := initializers.DB.Model(&models.Token{}).Where("token = ? && user_id = ?", tokenString, &existingUser.ID).First(&existingToken).Error; err == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Duplicate JWT token!",
			},
		)
		return
	}

	// Additional: Invalidate all existing Tokens (Only 1 token at a time)
	if err := initializers.DB.Where("user_id = ?", &existingUser.ID).Delete(&models.Token{}).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to invalidate existing tokens",
			},
		)
		return
	}

	// Add token to Database for User
	newTokenObject := models.Token{
		UserID:    existingUser.ID,
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}
	if err := initializers.DB.Create(&newTokenObject).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "An error has occurred while generating a JWT token",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		newTokenObject.Serialize(true),
	)
}
