package helpers

import "github.com/gin-gonic/gin"

func GetUsernameFromHeader(c *gin.Context) string {
	for key, values := range c.Request.Header {
		if key == "Goshort-Username" {
			return values[0]
		}
	}

	return ""
}
