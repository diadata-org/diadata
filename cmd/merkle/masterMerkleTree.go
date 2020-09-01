package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
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

	// Get today's merkle root
	timestamp := time.Now()
	dailyTree, err := merklehashing.DailyTree(timestamp)
	if err != nil {
		log.Fatal(err)
	}
	dailyRootHash := dailyTree.MerkleRoot()
	fmt.Println("daily root hash: ", hex.EncodeToString(dailyRootHash))
	// Get last master tree
	lastID, err := ds.GetLastID("", level)
	if err != nil {
		log.Error(err)
	}
	ID := strconv.Itoa(int(lastID))
	masterTree, err := ds.GetDailyTreeByID("", level, ID)
	if err != nil {
		log.Error(err)
	}
	// Extend tree by today's merkle root
	newHash := merkletree.ByteContent(dailyRootHash)
	err = masterTree.ExtendTree([]merkletree.Content{newHash})
	if err != nil {
		log.Error(err)
	}
	// Save newMasterTree
	ds.SaveDailyTreeInflux(masterTree, "", level, time.Time{})
}
