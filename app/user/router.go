package user

import (
	"github.com/gin-gonic/gin"
)

func Router(c *gin.RouterGroup) {

	c.GET("/users", func(c *gin.Context) {

	})

	c.GET("/users/:user_id", func(c *gin.Context) {

	})

	c.POST("/users", func(c *gin.Context) {

	})

	c.PATCH("/users/:user_id", func(c *gin.Context) {

	})

	c.DELETE("/users/:user_id", func(c *gin.Context) {

	})
}
