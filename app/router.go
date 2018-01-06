package app

import (
	"github.com/gin-gonic/gin"
)

// NewRouter registry router config
func NewRouter(e *gin.Engine) {

	// article router group
	article := e.Group("/articles")
	{
		article.GET("/", GetArticleList)
		article.GET("/:article_id", GetArticleDetail)
		article.POST("/", CreateArticle)
		article.PUT("/:article_id", UpdateArticle)
		article.PATCH("/:article_id", UpdateArticle)
	}

	// category router group
	category := e.Group("/categories")
	{
		category.GET("/", GetCategoryList)
		category.POST("/", CreateCategory)
		category.PUT("/:category_id", UpdateCategory)
		category.PATCH("/:category_id", UpdateCategory)
	}

	// tag router group
	tag := e.Group("/tags")
	{
		tag.GET("/", GetTagList)
		tag.POST("/", CreateTag)
		tag.PUT("/:tag_id", UpdateTag)
		tag.PATCH("/:tag_id", UpdateTag)
	}

	// user router group
	user := e.Group("/users")
	{
		user.GET("/", GetUserList)
		user.GET("/:user_id", GetUserDetail)
		user.PUT("/:user_id", UpdateUser)
		user.PATCH("/:user_id", UpdateUser)
	}

	// auth router group
	auth := e.Group("/auth")
	{
		auth.POST("/registry", Rigistry)
		auth.POST("/login", Login)
		auth.GET("/logout", Logout)
	}
}
