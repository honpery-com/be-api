package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/config"
	mgo "gopkg.in/mgo.v2"
)

func Connect() func(c *gin.Context) {

	return func(c *gin.Context) {
		mongo_url := fmt.Sprintf("%s:%d", config.DBHost, config.DBPort)
		session, err := mgo.Dial(mongo_url)

		if err != nil {
			fmt.Println("mongo connect error.")
		}

		c.Set("mgo_session", *session)
		c.Set("mgo_db", session.DB(config.DBName))

		c.Next()

		session.Clone()
	}

}
