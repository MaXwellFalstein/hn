package hnapi

import (
	"encoding/json"
	"net/http"
)

// get makes an http(s) get request expecting
func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
