package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cbergoon/merkletree"
	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	ds, err := models.NewAuditStore()
	if err != nil {
		log.Fatal("NewAuditStore: ", err)
	}

	// Initialize process by setting the genesis master node
	var initialContainer merkletree.StorageBucket
	initialContainer.Content = []byte("audit trail starts here")
	genesisTree, err := merkletree.NewTree([]merkletree.Content{initialContainer})
	if err != nil {
		log.Error(err)
	}
	// Save genesis tree
	ds.SaveDailyTreeInflux(*genesisTree, "", "0", []string{}, time.Time{})

	masterTree, err := merklehashing.MasterTree(ds)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("root of master tree: ", hex.EncodeToString(masterTree.MerkleRoot))
	fmt.Println("master hash strategy: ", masterTree.HashStrategy)
	// ds, err := models.NewRedisAuditStore()
	// if err != nil {
	// 	log.Fatal("NewInfluxDataStore: ", err)
	// }
	val, err := ds.GetPoolID("12", "hash-trades")
	if err == nil {
		fmt.Println("val: ", val)
	}
	fmt.Println("error: ", err)

}
