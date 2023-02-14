package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func FetchComics(urlBase string, urlPostfix string) *Pics {
	var pics Pics
	for i := 1; ; i++ {
		if i%25 == 0 {
			fmt.Printf("Fetching image #%s\n", strconv.Itoa(i))
		}

		var url = urlBase + strconv.Itoa(i) + urlPostfix
		cur, shouldContinue := readPicDescription(url)
		if !shouldContinue {
			break
		}
		pics = append(pics, cur)
	}

	return &pics
}

func readPicDescription(url string) (pics *Comics, shouldContinue bool) {
	body, err := http.Get(url)
	HandleError(err, "Cannot read url: "+url)
	defer body.Body.Close()

	if body.StatusCode == http.StatusNotFound {
		return nil, false
	}

	pics = &Comics{}
	err = json.NewDecoder(body.Body).Decode(pics)
	HandleError(err, "Cannot decode: ")
	return pics, true
}
