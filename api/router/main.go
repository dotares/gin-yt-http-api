package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVideos(c *gin.Context) {
	response, err := http.Get("testURL")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response)
}

func main() {
	router := gin.Default()
	router.GET("/listen", getVideos)

	router.Run("localhost:5050")
}