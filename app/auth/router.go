package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/auth/login", func(c *gin.Context) {

	})

	r.POST("/auth/register", func(c *gin.Context) {

	})
}
