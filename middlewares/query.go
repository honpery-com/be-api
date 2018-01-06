package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// CoreQuery build in query conditions
type CoreQuery struct {
	skip     int
	limit    int
	sort     string
	datefrom string
	dateto   string
}

// Query middleware
func Query() gin.HandlerFunc {
	return func(c *gin.Context) {
		q := CoreQuery{}
		q.skip = toInt(c.Query("skip"), 0)
		q.limit = toInt(c.Query("limit"), 20)
		q.sort = c.Query("sort")
		q.datefrom = c.Query("datefrom")
		q.dateto = c.Query("dateto")
		c.Set("Query", q)
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
