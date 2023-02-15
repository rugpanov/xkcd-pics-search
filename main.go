package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName = "./xkcd.json"
	const urlBase = "https://xkcd.com/"
	const urlPostfix = "/info.0.json"

	if !IsFileExist(fileName) {
		comics := FetchComics(urlBase, urlPostfix)
		comics.SaveTo(fileName)
	}

	var cache = *initAndBuildCache(fileName)
	chatWithUser(cache)
}

func chatWithUser(cache map[string]ComicSet) {
	for {
		var keywords = *requestKeywords()
		var foundComics = searchForComics(cache, keywords)
		for _, comic := range foundComics {
			fmt.Printf("%s %s/%s/%s %q\n", comic.Img, comic.Day, comic.Month, comic.Year, comic.Title)
		}
		fmt.Printf("found %d comics\n", len(foundComics))
	}
}

type ComicSet map[Comic]struct{}

func searchForComics(cache map[string]ComicSet, keywords []string) Comics {
	var result Comics

	var firstKword = keywords[0]
	var comicSet = cache[firstKword]
	for comic := range comicSet {
		var transcript = strings.ToLower(comic.Transcript)
		var title = strings.ToLower(comic.Title)

		var hasAll = true
		for _, kword := range keywords {
			if !strings.Contains(transcript, kword) && !strings.Contains(title, kword) {
				hasAll = false
				break
			}
		}
		if hasAll {
			var copyComic = comic
			result = append(result, &copyComic)
		}
	}

	return result
}

func requestKeywords() *[]string {
	fmt.Printf("Provide whitespace separated keywords. Type `exit` to leave:\n")
	var keywords string

	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	HandleError(err, "reader doesn't work: ")
	keywords = string(line)
	kwords := strings.Fields(strings.TrimSpace(keywords))
	if keywords == "exit" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
	return &kwords
}

func initAndBuildCache(fileName string) *map[string]ComicSet {
	fmt.Printf("Start reading file %s\n", fileName)
	comics := Read(fileName)
	fmt.Printf("Finish reading file %s. Read %d comics\n", fileName, len(comics))

	fmt.Printf("Start preparing cache\n")
	var cache = make(map[string]ComicSet)
	for _, comic := range comics {
		var compiled = regexp.MustCompile("[^a-zA-Z0-9-_]")
		var splited = compiled.Split(comic.Transcript, -1)
		for _, token := range splited {
			if cache[strings.ToLower(token)] == nil {
				cache[strings.ToLower(token)] = make(map[Comic]struct{})
			}
			var tokenCache = cache[strings.ToLower(token)]
			tokenCache[*comic] = struct{}{}
		}
	}
	fmt.Printf("Finish preparing cache. Cache size: %d\n", len(cache))

	return &cache
}
