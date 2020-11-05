package main

import (
	"sync"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	topicMap := merklehashing.GetHashTopics()
	wg := sync.WaitGroup{}
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	for key := range topicMap {
		log.Info("Beginning hashing for topic " + topicMap[key] + "...")
		wg.Add(1)
		go merklehashing.HashPoolLoop(topicMap[key], ds)
	}
	defer wg.Wait()
}
