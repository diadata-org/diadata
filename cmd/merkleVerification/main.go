package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cbergoon/merkletree"
	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// verifyContent verifies storage tree with Id @id of @topic.
func verifyContent(topic string, id string, ds models.AuditStore, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	tree, err := ds.GetStorageTreeByID(topic, id)
	if err != nil {
		log.Fatal("error retrieving storage tree: ", err)
	}
	leafs := tree.Leafs
	// verify each leaf (content) of storage tree
	for _, leaf := range leafs {
		cnt := leaf.C.(merkletree.StorageBucket)
		verif, err := merklehashing.VerifyContent(cnt, ds)
		if err != nil {
			log.Fatal("error verifying content: ", err)
		}
		if verif == false {
			log.Errorf("could not verify content with ID %v in tree with ID %s, topic: %s\n", cnt.ID, id, topic)
			break
		}
	}
}

func main() {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}
	time.Sleep(20 * time.Second)
	ticker := time.NewTicker(60 * 1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				// var wg sync.WaitGroup
				verificationTime := time.Now()
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
					fmt.Printf("youngest merkle child for %s: %d\n", topicMap[key], lastID)

					storageTreesToVerify, err := ds.GetStorageTreesInflux(topicMap[key], time.Time{}, time.Unix(0, lastID))
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("number of storage trees to check for %s: %v \n", topicMap[key], len(storageTreesToVerify))
					// Iterate over storage trees in go routines
					for _, val := range storageTreesToVerify {
						tree := merkletree.MerkleTree{}
						err = json.Unmarshal([]byte(val[4].(string)), &tree)
						tstamp, _ := time.Parse(time.RFC3339, val[0].(string))
						// Verify buckets in pool
						verif, err := merklehashing.VerifyBuckets(tree, topicMap[key], ds)
						if err != nil {
							log.Fatal(err)
						}
						if verif == false {
							log.Fatal("could not verify bucket in pool ")
						}
						// Verify pools in daily trees
						id := strconv.FormatInt(tstamp.UnixNano(), 10)
						verif, err = merklehashing.VerifyPool(tree, topicMap[key], id, ds)
						if err != nil {
							log.Fatal(err)
						}
						if verif == false {
							log.Fatal("could not verify bucket in pool ")
						}
					}
					log.Infof("successfully verified all pools for %s.", topicMap[key])

					// Verify daily trees for levels 2 and 1
					levels := []string{"2", "1"}
					var topic string
					for _, level := range levels {
						if level == "2" {
							topic = topicMap[key]
						}
						dailyTrees, err := ds.GetDailyTreesInflux(topic, level, time.Time{}, verificationTime)
						if err != nil {
							log.Fatal(err)
						}
						for _, val := range dailyTrees {
							dailyTree := merkletree.MerkleTree{}
							err = json.Unmarshal([]byte(val[6].(string)), &dailyTree)
							id := val[2].(string)
							verif, err := merklehashing.VerifyTree(dailyTree, level, id, ds)
							if err != nil {
								log.Fatal(err)
							}
							if verif == false {
								log.Fatalf("could not verify level %s tree \n", level)
							}
						}
					}

					log.Infof("Successfully verified %v storage trees from %s.", len(storageTreesToVerify), topicMap[key])
				}
				log.Infof("%s -- All content in storage successfully verified. \n", time.Now().String())
			}
		}
	}()
	select {}

}

// main retrieves all (storage) buckets from the storage table and verifies them with
// the actual master root hash.
func mainOld() {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}
	time.Sleep(20 * time.Second)
	ticker := time.NewTicker(60 * 1 * time.Second)
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
					fmt.Printf("youngest merkle child for %s: %d\n", topicMap[key], lastID)

					storageTreesToVerify, err := ds.GetStorageTreesInflux(topicMap[key], time.Time{}, time.Unix(0, lastID))
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("number of storage trees to check for %s: %v \n", topicMap[key], len(storageTreesToVerify))
					// Iterate over storage trees in go routines
					for _, val := range storageTreesToVerify {
						tree := merkletree.MerkleTree{}
						err = json.Unmarshal([]byte(val[4].(string)), &tree)
						tstamp, _ := time.Parse(time.RFC3339, val[0].(string))
						go verifyContent(topicMap[key], strconv.FormatInt(tstamp.UnixNano(), 10), ds, &wg)
					}

					log.Infof("Successfully verified %v storage trees.", len(storageTreesToVerify))
				}
				log.Infof("%s -- All content in storage successfully verified. \n", time.Now().String())
			}
		}
	}()
	select {}

}
