package middlewares

import (
	"net/http"
	"strings"

	"github.com/InspectorGadget/goshort/helpers"
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	for key, values := range c.Request.Header {
		if key == "Authorization" {
			bearerToken := strings.Split(values[0], "Bearer ")[1]

			// Check if Bearer token is within DB
			if err := initializers.DB.Model(&models.Token{}).Where("token = ?", bearerToken).First(&models.Token{}).Error; err != nil {
				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{
						"message": "Invalid or expired token",
					},
				)
				return
			}

			username, err := helpers.VerifyJWT(bearerToken)
			if err != nil {
				c.AbortWithStatusJSON(
					http.StatusBadRequest,
					gin.H{
						"message": err.Error(),
					},
				)
			}

			c.Request.Header.Set("Goshort-Username", username)
			c.Next()
		}
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}
