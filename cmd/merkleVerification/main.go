package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/cbergoon/merkletree"
	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// verifyContent verifies storage tree with Id @i of @topic.
func verifyContent(topic string, i int, ds models.AuditStore, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	tree, err := ds.GetStorageTreeByID(topic, strconv.Itoa(i))
	if err != nil {
		log.Fatal("error retrieving storage tree: ", err)
	}
	leafs := tree.Leafs
	// verify each leaf (content) of storage tree
	for _, leaf := range leafs {
		cnt := leaf.C.(merkletree.StorageBucket)
		verif, err := merklehashing.VerifyContent(cnt)
		if err != nil {
			log.Fatal("error verifying content: ", err)
		}
		if verif == false {
			log.Errorf("could not verify content with ID %v in tree with ID %v, topic: %s\n", cnt.ID, i, topic)
			break
		}
	}
}

// main retrieves all (storage) buckets from the storage table and verifies them with
// the actual master root hash.
func main() {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				var wg sync.WaitGroup
				topicMap := merklehashing.GetHashTopics()
				for key := range topicMap {
					log.Infof("verifying topic %s ... \n", topicMap[key])
					// Get highest ID from storage table which is already hashed in merkle table.
					// There may be higher IDs in the storage table which are not hashed into daily trees yet.
					// For this reason they cannot be verified (or falsified).
					lastID, err := ds.GetYoungestChildMerkle(topicMap[key])
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("number of storage trees to check for %s: %v \n", topicMap[key], lastID)

					// Iterate over storage trees in go routines
					for i := 0; i <= int(lastID); i++ {
						go verifyContent(topicMap[key], i, ds, &wg)
					}

					log.Infof("Successfully verified %v storage trees.", lastID)
				}
				log.Infof("%s -- All content in storage successfully verified. \n", time.Now().String())
			}
		}
	}()
	select {}

}
