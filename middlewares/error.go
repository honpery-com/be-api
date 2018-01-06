package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Error middleware
func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
