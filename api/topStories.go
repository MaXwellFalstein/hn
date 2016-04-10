package hnapi

const (
	// TopStoriesURL is the URL used to retrieve the top hacker news stories item
	// numbers.
	TopStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json"
)

// TopStoriesItemNumbers retrieves up to 500 top and new stories are at
// https://hacker-news.firebaseio.com/v0/topstories and
// https://hacker-news.firebaseio.com/v0/newstories.
//
// Example: https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty
func TopStoriesItemNumbers() (StoryItemNumbers, error) {
	ts := StoryItemNumbers{}
	err := getJSON(TopStoriesURL, &ts)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

// RetrieveTopNStories returns the top n stories from Hacker News. Returns a
// pointer to a slice of Hacker News items.
func RetrieveTopNStories(n int, logger *Logger) (*[]HNItem, error) {
	var stories []HNItem
	logger.VerbosePrintln("Retrieving Top Story Item Numbers")
	topStoriesIDs, err := TopStoriesItemNumbers()
	if err != nil {
		return nil, err
	}

	logger.VerbosePrintfln("Retrieving Top %d Stories", n)
	for i, storyID := range topStoriesIDs {
		if i >= n {
			return &stories, nil
		}

		story := GetItem(storyID)
		stories = append(stories, *story)
	}

	return &stories, nil
}

// StreamTopNStories streams  the top n stories from Hacker News.
// Returns a channel over which to receive the stories.
func StreamTopNStories(n int, logger *Logger) (chan *HNItem, error) {
	c := make(chan *HNItem)
	logger.VerbosePrintln("Retrieving Top Story Item Numbers")
	topStoriesIDs, err := TopStoriesItemNumbers()
	if err != nil {
		logger.Printfln("Failed to retrieve Top Stories Item Numbers with error:\n\n%s.", err.Error())
		return nil, err
	}

	go func(topStoriesIDs StoryItemNumbers, n int, logger *Logger, c chan *HNItem) {
		logger.VerbosePrintfln("Retrieving Top %d Stories", n)
		for i, storyID := range topStoriesIDs {
			if i >= n {
				close(c)
				return
			}

			story := GetItem(storyID)
			c <- story
		}

		close(c)
	}(topStoriesIDs, n, logger, c)

	return c, nil
}
