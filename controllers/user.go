package controllers

import (
	"fmt"
	"net/http"

	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var userObject models.User

	if err := c.ShouldBindJSON(&userObject); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "One or more parameters are missing!",
			},
		)
		return
	}

	// Check if a user exist with the Username
	var existingUser models.User
	if err := initializers.DB.Model(&models.User{}).Where("username = ?", userObject.Username).First(&existingUser).Error; err == nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"message": "An user with the username exists!",
			},
		)
		return
	}

	// Add user
	if err := initializers.DB.Create(&userObject).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "An error has occured while adding user",
				"error":   err.Error(),
			},
		)
		return
	}

	// User has been added
	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "User successfully created",
			"user":    userObject.Values(),
		},
	)
}

func ListUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Model(&models.User{}).Preload("Urls").Find(&users)

	// Create a clean response without passwords
	var usersResponse []gin.H
	for _, user := range users {
		userValues := user.Values() // Use the Values method which should exclude sensitive data
		usersResponse = append(usersResponse, userValues)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"users": usersResponse,
		},
	)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	var existingUser models.User

	if err := initializers.DB.Model(&models.User{}).Where("id = ?", userId).First(&existingUser).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("User with ID '%s' not found.", userId),
			},
		)
		return
	}

	if err := initializers.DB.Delete(&existingUser).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Unable to delete user.",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "User deleted!",
		},
	)
}
