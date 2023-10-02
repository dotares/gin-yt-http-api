package dataFetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

// type video struct {
// 	videoId     string `json:"videoId"`
// 	videoTitle  string `json:"videoTitle"`
// 	videoArtist string `json:"VideoArtist"`
// 	videoCover  string `json:"VideoCover"`
// }

//	type videos struct {
//		videos []*video `json:"videos"`
//	}

type Match struct {
	videoId     string
	videoTitle  string
	videoArtist string
	videoCover  string
}

func DataFetch() {
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

	matches := make([]Match, len(idMatches))

	for i := range idMatches {
		matches[i] = Match{
			videoId:     idMatches[i][len(idMatches[i])-12 : len(idMatches[i])-1],
			videoTitle:  titleMatches[i][27 : len(titleMatches[i])-3],
			videoArtist: artistMatches[i][36 : len(artistMatches[i])-1],
			videoCover:  coverMatches[i][60 : len(coverMatches[i])-27],
		}
	}

	for _, m := range matches {
		fmt.Println(m.videoId, m.videoArtist, m.videoTitle, m.videoCover)
	}

	// for _, m := range idMatches {
	// 	fmt.Println(m[len(m)-12 : len(m)-1])
	// }

	// for _, m := range titleMatches {
	// 	fmt.Println(m[27 : len(m)-3])
	// }

	// for _, m := range artistMatches {
	// 	fmt.Println(m[36 : len(m)-1])
	// }

	// for _, m := range coverMatches {
	// 	fmt.Println(m[60 : len(m)-27])
	// }
}
