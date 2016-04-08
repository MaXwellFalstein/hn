package hnapi

import "log"

const (
	// TopStoriesURL is the URL used to retrieve the top hacker news stories item
	// numbers.
	TopStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
)

// RetrieveTopStoriesItemNumbers retrieves from the hacker-news API the top
// Hacker News stories item numbers.
func RetrieveTopStoriesItemNumbers() *TopStories {
	ts := TopStories{}
	err := getJSON(TopStoriesURL, &ts)
	if err != nil {
		log.Panicln(err.Error())
	}

	return &ts
}
