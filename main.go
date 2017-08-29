package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app"
	"github.com/honpery-com/be-api/config"
	"github.com/honpery-com/be-api/middlewares"
)

func main() {
	router := gin.Default()

	// auth handler.
	router.Use(middlewares.Xauth())

	// connect db.
	router.Use(middlewares.Connect())

	// query handler.
	router.Use(middlewares.Xquery())

	// error handler.
	router.Use(middlewares.Xerr())

	// v1 router.
	app.V1(router)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
