package hnapi

const (
	// AskHNStoriesURL is the URL used to retrieve up to 200 of the latest Ask HN
	// stories
	AskHNStoriesURL = "https://hacker-news.firebaseio.com/v0/askstories.json"
)

// AskHNItemNumbers retrieves up to 200 of the latest Ask HN stories at
// https://hacker-news.firebaseio.com/v0/askstories,
func AskHNItemNumbers() (StoryItemNumbers, error) {
	var askHNIDs StoryItemNumbers
	err := getJSON(AskHNStoriesURL, &askHNIDs)
	if err != nil {
		return nil, err
	}

	return askHNIDs, nil
}

// RetrieveNAskHNStories returns the top n stories from Hacker News. Returns a
// pointer to a slice of Hacker News items.
func RetrieveNAskHNStories(n int, logger *Logger) (*[]HNItem, error) {
	var stories []HNItem
	logger.VerbosePrintln("Retrieving Ask Hacker News Story Item Numbers")
	askHNIDs, err := AskHNItemNumbers()
	if err != nil {
		return nil, err
	}

	logger.VerbosePrintfln("Retrieving %d Ask Hacker News Stories", n)
	for i, storyID := range askHNIDs {
		if i >= n {
			return &stories, nil
		}

		story := GetItem(storyID)
		stories = append(stories, *story)
	}

	return &stories, nil
}

// StreamNAskHNStories streams using the first n Ask Hacker News stories.
// Returns a channel over which to receive the stories.
func StreamNAskHNStories(n int, logger *Logger) (chan *HNItem, error) {
	c := make(chan *HNItem)
	logger.VerbosePrintln("Retrieving Ask Hacker News Story Item Numbers")
	askHNIDs, err := AskHNItemNumbers()
	if err != nil {
		return nil, err
	}

	go func(askHNIDs StoryItemNumbers, n int, logger *Logger, c chan *HNItem) {
		logger.VerbosePrintfln("Retrieving %d Ask Hacker News Stories", n)
		for i, storyID := range askHNIDs {
			if i >= n {
				close(c)
				return
			}

			story := GetItem(storyID)
			c <- story
		}

		close(c)
	}(askHNIDs, n, logger, c)

	return c, nil
}
