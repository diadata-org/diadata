package main

import (
	"flag"
	"fmt"
	"sync"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (
	key = flag.Int("topic-key", 0, "which topic to hash")
)

func init() {
	flag.Parse()
}

func main() {
	// TO DO: fetch topics from postgres

	topics := merklehashing.TopicInfo
	fmt.Println("flagged key: ", *key)
	wg := sync.WaitGroup{}
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}

	log.Info("Beginning hashing for topic " + topics[*key].Name + "...")
	wg.Add(1)
	go merklehashing.HashPoolLoop(topics[*key].Name, topics[*key].SizePool, topics[*key].SizeBucket, ds)
	defer wg.Wait()
}
