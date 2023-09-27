package dataFetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DataFetch () {
	ytURL := "https://www.youtube.com/results?search_query=grenade+audio" 

	res, err := http.Get(ytURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}