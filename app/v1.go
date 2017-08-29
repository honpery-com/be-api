package app

import (
	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app/auth"
	"github.com/honpery-com/be-api/app/category"
	"github.com/honpery-com/be-api/app/tag"
	"github.com/honpery-com/be-api/app/user"
)

func V1(router *gin.Engine) *gin.RouterGroup {
	V1Group := router.Group("/v1")
	{
		category.Router(V1Group)
		tag.Router(V1Group)
		article.Router(V1Group)
		user.Router(V1Group)
		auth.Router(V1Group)
	}

	return V1Group
}
