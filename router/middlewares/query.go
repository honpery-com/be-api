package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CoreQuery struct {
	skip     int
	limit    int
	sort     string
	datefrom string
	dateto   string
}

func Query() gin.HandlerFunc {
	return func(c *gin.Context) {
		core := CoreQuery{}
		core.skip = toInt(c.Query("skip"), 0)
		core.limit = toInt(c.Query("limit"), 20)
		core.sort = c.Query("sort")
		core.datefrom = c.Query("datefrom")
		core.dateto = c.Query("dateto")
		c.Set("Query", core)
		c.Next()
	}
}

func toInt(s string, defValue int) int {
	num, err := strconv.Atoi(s)

	if err != nil {
		return defValue
	}

	return num
}
