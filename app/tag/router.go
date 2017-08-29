package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.RouterGroup) {

	c.GET("/tags", func(c *gin.Context) {
		result, err := Model(c).List(c.MustGet("xquery"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

	c.GET("/tags/:tag_id", func(c *gin.Context) {
		result, err := Model(c).Detail(c.Param("tag_id"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

	c.POST("/tags", func(c *gin.Context) {
		newTag := Tag{}

		if err := c.BindJSON(newTag); err != nil {
			c.Error(err)
			return
		}

		result, err := Model(c).Create(newTag)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

	c.PATCH("/tags/:tag_id", func(c *gin.Context) {
		newTag := Tag{}

		if err := c.BindJSON(newTag); err != nil {
			c.Error(err)
			return
		}

		result, err := Model(c).Update(c.Param("tag_id"), newTag)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

	c.DELETE("/tags/:tag_id", func(c *gin.Context) {
		result, err := Model(c).Delete(c.Param("tag_id"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

}
