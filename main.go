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

func chatWithUser(cache map[string]Comics) {
	for {
		var keywords = *requestKeywords()
		var foundComics = searchForComics(cache, keywords)
		for index, comic := range foundComics {
			fmt.Printf("[%d] %s %s %s\n", index, comic.Title, comic.Day+"/"+comic.Month+"/"+comic.Year, comic.Img)
		}
	}
}

func searchForComics(cache map[string]Comics, keywords []string) Comics {
	var foundComics = make(map[*Comic]struct{})
	for _, kword := range keywords {
		var foundComicsCache = cache[strings.ToLower(kword)]
		for _, title := range foundComicsCache {
			foundComics[title] = struct{}{}
		}
	}

	var result Comics
	for comic := range foundComics {
		result = append(result, comic)
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

func initAndBuildCache(fileName string) *map[string]Comics {
	fmt.Printf("Start reading file %s\n", fileName)
	comics := Read(fileName)
	fmt.Printf("Finish reading file %s\n", fileName)

	fmt.Printf("Start preparing cache\n")
	var cache = make(map[string]Comics)
	for _, comic := range comics {
		var compiled = regexp.MustCompile("[^a-zA-Z0-9-_]")
		var splited = compiled.Split(comic.Transcript, -1)
		for _, token := range splited {
			cache[strings.ToLower(token)] = append(cache[token], comic)
		}
	}
	fmt.Printf("Finish preparing cache\n")

	return &cache
}
