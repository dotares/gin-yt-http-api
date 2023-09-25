package router

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200 ,"Hello")
	})
	r.Run()
}