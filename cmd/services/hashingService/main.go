package main

import (
	"sync"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	// topicMap := merklehashing.GetHashTopics()
	topics := merklehashing.TopicInfo
	wg := sync.WaitGroup{}
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	for key := range topics {
		log.Info("Beginning hashing for topic " + topics[key].Name + "...")
		wg.Add(1)
		go merklehashing.HashPoolLoop(topics[key].Name, topics[key].SizePool, topics[key].SizeBucket, ds)
	}
	defer wg.Wait()
}
