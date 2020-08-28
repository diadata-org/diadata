package main

import (
	"fmt"
	"time"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("here is going to be the master merkle root")

	// Get today's merkle root
	timeStamp := time.Now()
	dailyTree, err := merklehashing.DailyTree(timeStamp)
	if err != nil {
		log.Fatal(err)
	}
	dailyRootHash := dailyTree.MerkleRoot()
	fmt.Println(dailyRootHash)
	// Get last master tree
	// masterTree := db.GetMasterTree(today)
	// // Extend tree by today's merkle root
	// newMasterTree := masterTree.append(mr)
	// // Save newMasterTree
	// db.SaveInfluxTree(newMasterTree)
}
