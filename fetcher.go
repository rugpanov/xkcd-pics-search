package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func FetchComics(urlBase string, urlPostfix string) *Comics {
	var comics Comics
	for i := 1; ; i++ {
		if i%25 == 0 {
			fmt.Printf("Fetching comics #%s\n", strconv.Itoa(i))
		}

		var url = urlBase + strconv.Itoa(i) + urlPostfix
		cur, shouldContinue := readComic(url)
		if !shouldContinue {
			break
		}
		comics = append(comics, cur)
	}

	return &comics
}

func readComic(url string) (comic *Comic, shouldContinue bool) {
	body, err := http.Get(url)
	HandleError(err, "Cannot read url: "+url)
	defer body.Body.Close()

	if body.StatusCode == http.StatusNotFound {
		return nil, false
	}

	comic = &Comic{}
	err = json.NewDecoder(body.Body).Decode(comic)
	HandleError(err, "Cannot decode: ")
	return comic, true
}
