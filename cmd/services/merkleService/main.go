package main

import (
	"fmt"
	"sync"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
)

func main() {
	topicMap := merklehashing.GetHashTopics()
	wg := sync.WaitGroup{}
	for key := range topicMap {
		fmt.Println("storing ", topicMap[key])
		wg.Add(1)
		go merklehashing.HashPoolLoop(topicMap[key])
	}
	defer wg.Wait()
}
