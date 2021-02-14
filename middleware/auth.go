package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//MyInterceptor ...
func MyInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}
