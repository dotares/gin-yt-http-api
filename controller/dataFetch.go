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

	idRegex := regexp.MustCompile("(?i)\\{\"videoRenderer\":\\{\"videoId\":\"[^\"]*\"")
	titleRegex := regexp.MustCompile("(?i),\"title\":\\{\"runs\":\\[\\{\"text\":\"[^\"]*\"\\}\\]")
	artistRegex := regexp.MustCompile("(?i),\"longBylineText\":\\{\"runs\":\\[\\{\"text\":\"[^\"]*\"")
	coverRegex := regexp.MustCompile("(?i)\\{\"videoId\":\"[^\"]*\",\"thumbnail\":\\{\"thumbnails\":\\[\\{\"url\":\"[^\"]*\",\"width\":360,\"height\":202\\}")

	idMatches := idRegex.FindAllString(bodyString, -1)
	titleMatches := titleRegex.FindAllString(bodyString, -1)
	artistMatches := artistRegex.FindAllString(bodyString, -1)
	coverMatches := coverRegex.FindAllString(bodyString, -1)

	for _, m := range idMatches {
		fmt.Println(m[len(m)-12:len(m)-1]) 
	} 

	for _, m := range titleMatches {
		fmt.Println(m[27:len(m)-3]) 
	} 

	for _, m := range artistMatches {
		fmt.Println(m[36:len(m)-1]) 
	} 

	for _, m := range coverMatches {
		fmt.Println(m[60:len(m)-27]) 
	} 
}