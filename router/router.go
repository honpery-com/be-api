package router

import (
	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app"
)

// New return router
func New() *gin.Engine {
	r := gin.New()

	articles := r.Group("/articles")
	{
		articles.GET("", app.GetArticles)
		articles.GET("/:article_id", app.GetArticle)
		articles.POST("", app.PostArticle)
		articles.PUT("/:article_id", app.PutArticle)
		articles.DELETE("/:article_id", app.DeleteArticle)
	}

	tags := r.Group("/tags")
	{
		tags.GET("", app.GetTags)
		tags.GET("/:tag_id", app.GetTag)
		tags.POST("", app.PostTag)
		tags.PUT("/:tag_id", app.PutTag)
		tags.DELETE("/:tag_id", app.DeleteTag)
	}

	return r
}
