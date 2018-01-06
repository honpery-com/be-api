package main

import (
	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/app"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	app.NewRouter(r)

	r.Run()
}
