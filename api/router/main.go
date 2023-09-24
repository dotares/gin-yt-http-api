package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVideos(c *gin.Context) {
	c.String(http.StatusOK, "Hello this is a video file")
	
}

func main() {
	router := gin.Default()
	router.GET("/listen", getVideos)

	router.Run("localhost:5050")
}