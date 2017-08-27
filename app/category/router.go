package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/honpery-com/be-api/app/common"
)

func Router(router *gin.RouterGroup) {

	router.GET("/categories", func(c *gin.Context) {

	})

	router.GET("/categories/:category_id", func(c *gin.Context) {
		category_id := c.Param("category_id")
		c.JSON(http.StatusOK, gin.H{"id": category_id})
	})

	router.POST("/categories", func(c *gin.Context) {
		category := Category{}
		err := c.BindJSON(&category)

		if err != nil {
			c.Error(err)
			return
		}

		err, session, coll := Connect(CategoryColl)

		if err != nil {
			c.Error(err)
			return
		}

		defer session.Close()

		err = coll.Insert(category)

		if err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, category)
		}
		// err, result := XModel(CategoryColl).Create(category)

		// if err != nil {
		// 	c.Error(err)
		// } else {
		// 	c.JSON(http.StatusOK, result)
		// }
	})

	router.PATCH("/categories/:category_id", func(c *gin.Context) {
		category := Category{}
		err := c.BindJSON(&category)
		category_id := c.Param("category_id")

		if err != nil {
			c.Error(err)
			return
		}

		err, result := XModel(CategoryColl).Update(category_id, category)

		if err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})

	router.DELETE("/categories/:category_id", func(c *gin.Context) {

	})

}
