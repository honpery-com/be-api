package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Xauth() func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {

		}
		c.Next()
	}
}
