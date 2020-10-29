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

}
