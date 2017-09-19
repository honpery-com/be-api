package app

import (
	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/articles"
	"github.com/honpery-com/be-api/app/categories"
	"github.com/honpery-com/be-api/app/tags"
	"github.com/honpery-com/be-api/app/users"
)

func RegisterRouter(router *gin.Engine) *gin.RouterGroup {

	V1Group := router.Group("/v1")
	{
		tags.Router(V1Group)
		categories.Router(V1Group)
		users.Router(V1Group)
		articles.Router(V1Group)
	}

	return V1Group
}
