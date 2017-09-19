package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app"
	"github.com/honpery-com/be-api/config"
)

func main() {
	router := gin.Default()

	// query handler.
	router.Use(middlewares.Query())

	// error handler.
	router.Use(middlewares.Error())

	// v1 router.
	app.RegisterRouter(router)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
