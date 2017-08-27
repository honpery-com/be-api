package app

import (
	"./category"
	"github.com/gin-gonic/gin"
)

func V1(router *gin.Engine) *gin.RouterGroup {
	V1Group := router.Group("/v1")
	{
		category.Router(V1Group)
	}
	return V1Group
}
