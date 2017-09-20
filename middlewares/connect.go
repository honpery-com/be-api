package middlewares

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() func(*gin.Context) {
	return func(ctx *gin.Context) {
		db, err := sql.Open("postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBPwd,
			config.DBName,
		))

		if err != nil {
			panic(err)
		}

		ctx.Set("DB", db)

		ctx.Next()

		db.Close()
	}
}
