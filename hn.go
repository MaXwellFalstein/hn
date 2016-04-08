package main

import (
	"fmt"

	"github.com/kkirsche/hn/api"
)

func main() {
	topStores := hnapi.RetrieveTopStoriesItemNumbers()

	fmt.Println(*topStores)
}
