package controllers

import (
	"fmt"
	"net/http"

	"github.com/InspectorGadget/goshort/helpers"
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/InspectorGadget/goshort/structs"
	"github.com/gin-gonic/gin"
)

func AddRole(c *gin.Context) {
	username := helpers.GetUsernameFromHeader(c)

	var addRoleObject structs.AddRoleRequest
	if err := c.ShouldBind(&addRoleObject); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "One or more parameters required",
				"errors":  err.Error(),
			},
		)
		return
	}

	// Check if role name exists
	if err := initializers.DB.Model(&models.Role{}).Where("name = ?", addRoleObject.Name).First(&models.Role{}).Error; err == nil {
		c.AbortWithStatusJSON(
			http.StatusConflict,
			gin.H{
				"message": fmt.Sprintf("role with the name '%s' exists", addRoleObject.Name),
			},
		)
		return
	}

	// Find user
	existingUser := models.User{}
	if err := initializers.DB.Model(&models.User{}).Where("username = ?", username).First(&existingUser).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusConflict,
			gin.H{
				"message": "User does not exist",
			},
		)
		return
	}

	// Add new role
	newRole := models.Role{
		Name:    addRoleObject.Name,
		AddedBy: existingUser.ID,
	}
	if err := initializers.DB.Create(&newRole).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "An error has occurred while adding new role",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Role has been added.",
		},
	)
}
