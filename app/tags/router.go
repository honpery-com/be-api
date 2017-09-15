package tags

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/tags", func(c *gin.Context) {

	})

	r.GET("/tags/:tag_id", func(c *gin.Context) {

	})

	r.POST("/tags", func(c *gin.Context) {

	})

	r.PUT("/tags/:tag_id", func(c *gin.Context) {

	})

	r.DELETE("/tags/:tag_id", func(c *gin.Context) {

	})
}
