package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Query() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
