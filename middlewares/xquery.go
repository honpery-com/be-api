package middlewares

import (
	"github.com/gin-gonic/gin"
)

type Query struct {
	limit int
	skip  int
}

func Xquery() func(c *gin.Context) {
	return func(c *gin.Context) {
		query := Query{}
		c.BindQuery(&query)
		c.Set("xquery", query)
		c.Next()
	}
}
