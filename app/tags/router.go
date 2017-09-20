package tags

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/tags", func(c *gin.Context) {
		tags, err := TagModel(c).List(c.MustGet("Query"))

		if err != nil {
			c.Error(err)
		}

		c.JSON(http.StatusOK, tags)
	})

	r.GET("/tags/:tag_id", func(c *gin.Context) {
		// tag, err := TagModel(c).Detail(c.Param("tag_id"))

		// if err != nil {
		// 	c.Error(err)
		// }

		// c.JSON(http.StatusOK, tag)
	})

	r.POST("/tags", func(c *gin.Context) {

	})

	r.PUT("/tags/:tag_id", func(c *gin.Context) {

	})

	r.DELETE("/tags/:tag_id", func(c *gin.Context) {

	})
}
