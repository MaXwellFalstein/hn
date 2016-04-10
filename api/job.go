package hnapi

const (
	// JobStoriesURL is the URL used to retrieve up to 200 of the latest Show
	// HN stories
	JobStoriesURL = "https://hacker-news.firebaseio.com/v0/jobstories.json"
)

// JobtemNumbers retrieves up to 200 of the latest Job stories at
// https://hacker-news.firebaseio.com/v0/jobstories,
func JobtemNumbers() (*StoryItemNumbers, error) {
	var jobsIDs StoryItemNumbers
	err := getJSON(ShowHNStoriesURL, &jobsIDs)
	if err != nil {
		return nil, err
	}

	return &jobsIDs, nil
}
