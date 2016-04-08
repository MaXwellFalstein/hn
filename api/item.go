package hnapi

import (
	"log"
	"strconv"
)

// GetItem gets an item from Hacker News. This may be a story, comment, ask,
// job, poll, and one of a poll's parts
func GetItem(itemNumber int) *HNItem {
	item := HNItem{}
	itemString := strconv.Itoa(itemNumber)
	err := getJSON("https://hacker-news.firebaseio.com/v0/item/"+itemString+".json?print=pretty", &item)
	if err != nil {
		log.Panicln(err.Error())
	}

	return &item
}
