package users

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/users", func(c *gin.Context) {

	})

	r.GET("/users/:user_id", func(c *gin.Context) {

	})

	r.POST("/users", func(c *gin.Context) {

	})

	r.PUT("/users/:user_id", func(c *gin.Context) {

	})

	r.DELETE("/users/:user_id", func(c *gin.Context) {

	})
}
