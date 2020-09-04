package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	// level := "0"

	// // Initialize process by setting the genesis master node
	// genesisMessage := merkletree.ByteContent([]byte("hashing starts here"))
	// genesisTree, err := merkletree.NewTree([]merkletree.Content{genesisMessage})
	// if err != nil {
	// 	log.Error(err)
	// }
	// // Save genesis tree
	// fmt.Println("tree before encoding: ", genesisTree.Root)
	// ds.SaveDailyTreeInflux(*genesisTree, "", level, time.Time{})

	// Get today's merkle root
	timestamp := time.Now()
	dailyTree, err := merklehashing.DailyTree(timestamp)
	if err != nil {
		log.Fatal(err)
	}
	dailyRootHash := dailyTree.MerkleRoot
	fmt.Println("daily root hash: ", hex.EncodeToString(dailyRootHash))

	vals, err := ds.GetDailyTreesInflux("", "1", time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		log.Error(err)
	}
	var auxTree merkletree.MerkleTree
	// fmt.Println(vals[0][5].(string))
	fmt.Println("unmarshal daily tree..")
	err = json.Unmarshal([]byte(vals[0][5].(string)), &auxTree)
	if err != nil {
		log.Error(err)
	}
	contents := []merkletree.Content{auxTree.Leafs[0].C.(merkletree.StorageBucket), auxTree.Leafs[1].C.(merkletree.StorageBucket)}
	rebuiltAuxTree, err := merkletree.NewTreeWithHashStrategy(contents, sha256.New)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("rr: ", rebuiltAuxTree.MerkleRoot)
	fmt.Println("rebuilt leafs: ", rebuiltAuxTree.Leafs[0].C)

	cont := rebuiltAuxTree.Leafs[0].C.(merkletree.StorageBucket)
	fmt.Println("leaf 0 before modif.: ", cont)
	cont.Content[0] = 221
	fmt.Println("leaf 0: ", rebuiltAuxTree.Leafs[0].C.(merkletree.StorageBucket))
	fmt.Println("leaf 1: ", rebuiltAuxTree.Leafs[1].C.(merkletree.StorageBucket))
	ver, err := rebuiltAuxTree.VerifyContent(cont)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("verification of tree: ", ver)

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
