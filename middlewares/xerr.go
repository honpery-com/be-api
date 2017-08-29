package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Xerr() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
