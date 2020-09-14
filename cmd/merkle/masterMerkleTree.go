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
	ds, err := models.NewInfluxAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	level := "0"

	// Initialize process by setting the genesis master node
	var initialContainer merkletree.StorageBucket
	initialContainer.Content = []byte("hashing starts here")
	genesisTree, err := merkletree.NewTree([]merkletree.Content{initialContainer})
	if err != nil {
		log.Error(err)
	}
	// Save genesis tree
	ds.SaveDailyTreeInflux(*genesisTree, "", level, []string{}, time.Time{})

	// // Get today's merkle root
	// timestamp := time.Now()
	// dailyTree, err := merklehashing.DailyTree(timestamp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// dailyRootHash := dailyTree.MerkleRoot
	// fmt.Println("daily root hash: ", hex.EncodeToString(dailyRootHash))

	masterTree, err := merklehashing.MasterTree()
	if err != nil {
		log.Error(err)
	}
	fmt.Println("root of master tree: ", hex.EncodeToString(masterTree.MerkleRoot))

	db, err := models.NewRedisAuditStore()
	if err != nil {
		log.Fatal("NewInfluxDataStore: ", err)
	}
	db.GetPoolID("12", "hash-trades")

	// // Get last master tree
	// var masterTree merkletree.MerkleTree
	// lastID, err := ds.GetLastID("", level)
	// if err != nil {
	// 	log.Error(err)
	// }
	// ID := strconv.Itoa(int(lastID))
	// if ID != "0" {
	// 	masterTree, err = ds.GetDailyTreeByID("", level, ID)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}
	// }

	// // Extend tree by today's merkle root
	// newHash := merkletree.ByteContent(dailyRootHash)
	// err = masterTree.ExtendTree([]merkletree.Content{newHash})
	// if err != nil {
	// 	log.Error(err)
	// }
	// // Save newMasterTree
	// ds.SaveDailyTreeInflux(masterTree, "", level, time.Time{})

	// vals, err := ds.GetMerkletreesInflux("hash-interestrates", time.Now().AddDate(0, 0, -2), time.Now())
	// if err != nil {
	// 	log.Error(err)
	// }
	// var testTree merkletree.MerkleTree
	// err = json.Unmarshal([]byte(vals[0][3].(string)), &testTree)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println("testtree: ", testTree.Root.Left)
}
