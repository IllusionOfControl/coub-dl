package main

import (
	"coub-dl/api"
	"coub-dl/coub"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getCoubID(url string) string {
	re := regexp.MustCompile(`https://coub\.com/view/(\w+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) == 0 {
		fmt.Println("Invalid Coub URL.")
		os.Exit(1)
	}

	return matches[1]
}

func main() {
	loop := flag.Int("loop", 1, "loop the video N times (-1 for max loop)")
	coubUrl := flag.String("url", "", "downloadable Coub url (required)")

	flag.Parse()

	if *coubUrl == "" {
		fmt.Println("Error: --coubUrl argument is required")
		os.Exit(1)
	}

	coubId := getCoubID(*coubUrl)

	log.Printf("Get coub metadata from \"%s\"", *coubUrl)
	metadata, err := api.FetchCoubMetadata(coubId)
	if err != nil {
		log.Fatalln(err)
	}
	media := coub.NewCoubMedia(metadata)
	log.Printf("Founded coub media with title \"%s\"", media.Title)
	media.Save(coub.SaveOptions{Loop: strconv.Itoa(*loop)})
}
