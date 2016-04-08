package hnapi

import "log"

// GetUser retrieves users which are identified by case-sensitive ids, and live
// under https://hacker-news.firebaseio.com/v0/user/. Only users that have
// public activity (comments or story submissions) on the site are available
// through the API.
func GetUser(username string) *HNUser {
	user := HNUser{}
	err := getJSON("https://hacker-news.firebaseio.com/v0/user/"+username+".json?print=pretty", &user)
	if err != nil {
		log.Panicln(err.Error())
	}

	return &user
}
