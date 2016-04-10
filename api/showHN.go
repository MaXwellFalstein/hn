package hnapi

const (
	// ShowHNStoriesURL is the URL used to retrieve up to 200 of the latest Show
	// HN stories
	ShowHNStoriesURL = "https://hacker-news.firebaseio.com/v0/showstories.json"
)

// ShowHNItemNumbers retrieves up to 200 of the latest Show HN stories at
// https://hacker-news.firebaseio.com/v0/showstories,
func ShowHNItemNumbers() (*StoryItemNumbers, error) {
	var showHNIDs StoryItemNumbers
	err := getJSON(ShowHNStoriesURL, &showHNIDs)
	if err != nil {
		return nil, err
	}

	return &showHNIDs, nil
}
