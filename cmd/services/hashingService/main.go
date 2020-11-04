package main

import (
	"fmt"
	"log"
	"sync"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
)

func main() {
	topicMap := merklehashing.GetHashTopics()
	wg := sync.WaitGroup{}
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	for key := range topicMap {
		fmt.Println("storing ", topicMap[key])
		wg.Add(1)
		go merklehashing.HashPoolLoop(topicMap[key], ds)
	}
	defer wg.Wait()
}
