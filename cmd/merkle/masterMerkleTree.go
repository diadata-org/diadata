package main

import (
	"fmt"
	"time"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}

	// Initialize process by setting the genesis master node.
	// This should eventually go into an init file or at least be executed only once.
	// var initialContainer merkletree.StorageBucket
	// initialContainer.Content = []byte("audit trail starts here")
	// genesisTree, err := merkletree.NewTree([]merkletree.Content{initialContainer})
	// if err != nil {
	// 	log.Error(err)
	// }
	// // Save genesis tree
	// ds.SaveDailyTreeInflux(*genesisTree, "", "0", []string{}, time.Time{})

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				masterTree, err := merklehashing.MasterTree(ds)
				if err != nil {
					log.Error(err)
				}
				fmt.Println("master tree saved: ", masterTree)

			}
		}
	}()
	select {}

	// fmt.Println("root of master tree: ", hex.EncodeToString(masterTree.MerkleRoot))
	// fmt.Println("master hash strategy: ", masterTree.HashStrategy)
	// // ds, err := models.NewRedisAuditStore()
	// // if err != nil {
	// // 	log.Fatal("NewInfluxDataStore: ", err)
	// // }
	// val, err := ds.GetPoolsParentID("12", "hash-trades")
	// if err == nil {
	// 	fmt.Println("val: ", val)
	// }
	// fmt.Println("error: ", err)

}
