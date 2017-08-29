package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.GET("/categories", func(c *gin.Context) {
		categories, err := Model(c).List(c.MustGet("xquery"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, categories)
	})

	router.GET("/categories/:category_id", func(c *gin.Context) {
		category, err := Model(c).Detail(c.Param("category_id"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, category)
	})

	router.POST("/categories", func(c *gin.Context) {
		category := Category{}
		err := c.BindJSON(&category)

		if err != nil {
			c.Error(err)
			return
		}

		category, err = Model(c).Create(category)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, category)
	})

	router.PATCH("/categories/:category_id", func(c *gin.Context) {
		category := Category{}
		err := c.BindJSON(&category)

		if err != nil {
			c.Error(err)
			return
		}

		category, err = Model(c).Update(c.Param("category_id"), category)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, category)
	})

	router.DELETE("/categories/:category_id", func(c *gin.Context) {
		category, err := Model(c).Delete(c.Param("category_id"))

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, category)
	})
}
