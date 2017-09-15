package articles

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.GET("/articles", func(c *gin.Context) {

	})

	r.GET("/articles/:article_id", func(c *gin.Context) {

	})

	r.POST("/articles", func(c *gin.Context) {

	})

	r.PUT("/articles/:article_id", func(c *gin.Context) {

	})

	r.DELETE("/articles/:article_id", func(c *gin.Context) {

	})

}
