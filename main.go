package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	maxResults   = 25
	developerKey = os.Getenv("YOUTUBE_API_KEY")
)

func main() {
	getVideos("ギター")
	getVideos("ベース")
	getVideos("ドラム")
	getVideos("キーボード")
}

func getVideos(instrument string) {

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	query := "弾いてみた " + instrument

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(query).
		MaxResults(int64(maxResults))
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	printIDs(instrument, videos)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(instrument string, matches map[string]string) {
	fmt.Printf("%v\n", instrument)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Print("\n\n")
}
