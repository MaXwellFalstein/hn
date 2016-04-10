package hnapi

// HNItem represents stories, comments, jobs, Ask HNs and even polls . They're
// identified by their ids, which are unique integers, and live under
// https://hacker-news.firebaseio.com/v0/item/.
type HNItem struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Parent      int    `json:"parent"`
	Parts       []int  `json:"parts"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}
