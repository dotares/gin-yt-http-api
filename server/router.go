package router

import (
	dataFetch "api/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		dataFetch.DataFetch()
	})
	r.Run()
}