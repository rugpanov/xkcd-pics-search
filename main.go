package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName = "./pics.json"
	const urlBase = "https://xkcd.com/"
	const urlPostfix = "/info.0.json"

	if !IsFileExist(fileName) {
		pics := FetchComics(urlBase, urlPostfix)
		pics.SaveTo(fileName)
	}

	var cache = *initAndBuildCache(fileName)
	chatWithUser(cache)
}

func chatWithUser(cache map[string]Pics) {
	for {
		var keywords = *requestKeywords()
		var foundPicsSet = searchForPics(cache, keywords)
		for index, pic := range foundPicsSet {
			fmt.Printf("[%d] %s %s %s\n", index, pic.Title, pic.Day+"/"+pic.Month+"/"+pic.Year, pic.Img)
		}
	}
}

func searchForPics(cache map[string]Pics, keywords []string) Pics {
	var foundPicsSet = make(map[*Comics]struct{})
	for _, kword := range keywords {
		var foundPics = cache[strings.ToLower(kword)]
		for _, title := range foundPics {
			foundPicsSet[title] = struct{}{}
		}
	}

	var result Pics
	for pic := range foundPicsSet {
		result = append(result, pic)
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

func initAndBuildCache(fileName string) *map[string]Pics {
	fmt.Printf("Start reading file %s\n", fileName)
	pics := Read(fileName)
	fmt.Printf("Finish reading file %s\n", fileName)

	fmt.Printf("Start preparing cache\n")
	var cache = make(map[string]Pics)
	for _, pic := range pics {
		var compiled = regexp.MustCompile("[^a-zA-Z0-9-_]")
		var splited = compiled.Split(pic.Transcript, -1)
		for _, token := range splited {
			cache[strings.ToLower(token)] = append(cache[token], pic)
		}
	}
	fmt.Printf("Finish preparing cache\n")

	return &cache
}
