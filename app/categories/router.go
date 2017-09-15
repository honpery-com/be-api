package categories

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/categories", func(c *gin.Context) {

	})

	r.GET("/categories/:category_id", func(c *gin.Context) {

	})

	r.POST("/categories", func(c *gin.Context) {

	})

	r.PUT("/categories/:category_id", func(c *gin.Context) {

	})

	r.DELETE("/categories/:category_id", func(c *gin.Context) {

	})
}
