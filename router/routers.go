package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "user list")
		})

		user.PUT("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "user delete")
		})
	}
	return r
}
