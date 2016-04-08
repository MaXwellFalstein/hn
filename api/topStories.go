package hnapi

import "log"

const (
	// TopStoriesURL is the URL used to retrieve the top hacker news stories item
	// numbers.
	TopStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
)

// RetrieveTopStoriesItemNumbers retrieves up to 500 top and new stories are at
// https://hacker-news.firebaseio.com/v0/topstories and
// https://hacker-news.firebaseio.com/v0/newstories.
//
// Example: https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty
func RetrieveTopStoriesItemNumbers() *TopStories {
	ts := TopStories{}
	err := getJSON(TopStoriesURL, &ts)
	if err != nil {
		log.Panicln(err.Error())
	}

	return &ts
}
