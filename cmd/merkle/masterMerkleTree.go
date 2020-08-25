package main

import (
	"fmt"
	"time"

	merklehashing "github.com/diadata-org/diadata/internal/pkg/merkle-trees"
)

func main() {
	fmt.Println("here is going to be the master merkle root")
	// // Get today's merkle root
	timeFinal := time.Now()
	timeInit := timeFinal.AddDate(0, 0, -1)
	dailyTree, err := merklehashing.DailyTree(timeInit, timeFinal)
	fmt.Println(dailyTree, err)
	// // Get last master tree
	// masterTree := db.GetMasterTree(today)
	// // Extend tree by today's merkle root
	// newMasterTree := masterTree.append(mr)
	// // Save newMasterTree
	// db.SaveInfluxTree(newMasterTree)
}
