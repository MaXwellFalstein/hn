package hnapi

const (
	// AskHNStoriesURL is the URL used to retrieve up to 200 of the latest Ask HN
	// stories
	AskHNStoriesURL = "https://hacker-news.firebaseio.com/v0/askstories.json"
)

// AskHNItemNumbers retrieves up to 200 of the latest Ask HN stories at
// https://hacker-news.firebaseio.com/v0/askstories,
func AskHNItemNumbers() (*StoryItemNumbers, error) {
	var askHNIDs StoryItemNumbers
	err := getJSON(AskHNStoriesURL, &askHNIDs)
	if err != nil {
		return nil, err
	}

	return &askHNIDs, nil
}
