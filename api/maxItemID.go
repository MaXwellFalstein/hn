package hnapi

import "log"

const (
	// MaxItemIDURL is the URL used to retrieve the current largest Hacker News
	// item id.
	MaxItemIDURL = "https://hacker-news.firebaseio.com/v0/maxitem.json"
)

// MaxItemID retrieves the current largest item id is at
// https://hacker-news.firebaseio.com/v0/maxitem. You can walk backward from
// here to discover all items.
func MaxItemID() int {
	var maxItemID int
	err := getJSON(MaxItemIDURL, &maxItemID)
	if err != nil {
		log.Panicln(err.Error())
	}

	return maxItemID
}
