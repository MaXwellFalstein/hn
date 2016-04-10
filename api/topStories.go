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

// RetrieveTopNStories returns the top n stories from Hacker News. Returns a
// pointer to a slice of Hacker News items.
func RetrieveTopNStories(n int, logger *Logger) *[]HNItem {
	var stories []HNItem
	logger.VerbosePrintln("Retrieving Top Story Item Numbers")
	topStoriesIDs := RetrieveTopStoriesItemNumbers()

	logger.VerbosePrintfln("Retrieving Top %d Stories", n)
	for i, storyID := range *topStoriesIDs {
		if i >= n {
			return &stories
		}

		story := GetItem(storyID)
		stories = append(stories, *story)
	}

	return &stories
}

// StreamTopNStories streams using the top n stories from Hacker News. Returns a
// pointer to a slice of Hacker News items.
func StreamTopNStories(n int, logger *Logger) chan *HNItem {
	c := make(chan *HNItem)
	go func(n int, logger *Logger, c chan *HNItem) {
		logger.VerbosePrintln("Retrieving Top Story Item Numbers")
		topStoriesIDs := RetrieveTopStoriesItemNumbers()

		logger.VerbosePrintfln("Retrieving Top %d Stories", n)
		for i, storyID := range *topStoriesIDs {
			if i >= n {
				close(c)
				return
			}

			story := GetItem(storyID)
			c <- story
		}

		close(c)
	}(n, logger, c)

	return c
}
