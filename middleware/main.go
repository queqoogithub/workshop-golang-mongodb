package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckAuthorization(c *gin.Context) {
	apiKey := ""
	if values, _ := c.Request.Header["Api-Key"]; len(values) > 0 {
		apiKey = values[0]
	}
	if apiKey != "my-gosoft-password" {
		c.JSON(401, gin.H{
			"err": "invalid api-key",
		})
		c.Abort()
		return
	}

	c.Next()

}
