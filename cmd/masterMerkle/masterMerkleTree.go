package main

import (
	"encoding/hex"
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

	ticker := time.NewTicker(60 * 1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				masterTree, err := merklehashing.MasterTree(ds)
				if err != nil {
					log.Error(err)
				}
				if masterTree.Isempty() {
					log.Infof("no new data available during the last hashing period. master tree and root remain the same.")
				} else {
					log.Infof("%v -- master tree saved with root: %s", time.Now(), hex.EncodeToString(masterTree.MerkleRoot))
				}
			}
		}
	}()
	select {}

}
