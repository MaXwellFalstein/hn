package hnapi

import (
	"log"
	"strings"
)

// GetUser retrieves users which are identified by case-sensitive ids, and live
// under https://hacker-news.firebaseio.com/v0/user/. Only users that have
// public activity (comments or story submissions) on the site are available
// through the API.
func GetUser(username string) *HNUser {
	user := HNUser{}
	url := "https://hacker-news.firebaseio.com/v0/user/" + strings.TrimSpace(username) + ".json?print=pretty"
	err := getJSON(url, &user)
	if err != nil {
		log.Panicln(err.Error())
	}

	return &user
}
