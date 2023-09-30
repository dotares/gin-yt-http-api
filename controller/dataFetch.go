package dataFetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
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

	bodyString := string(body)

	re := regexp.MustCompile("(?i)\\{\"videoRenderer\":\\{\"videoId\":\"[^\"]*\"")
	fmt.Println("******************* regex match:", re.MatchString(bodyString));
}