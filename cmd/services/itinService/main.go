package main

import (
	"github.com/diadata-org/diadata/internal/pkg/itinService"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	itinUrl = "https://api.itsa.anyblock.tools/v1/list-tokens"
)

func main() {
	dataStore, err := models.NewDataStore()
	if err != nil {
		log.Error(err)
		return
	}

	for {
		itins, err := itinService.GetItins(itinUrl)
		if err != nil {
			log.Error(err)
			return
		}

		for _, itinToken := range itins {
			err := dataStore.SetItinData(itinToken)
			if err != nil {
				log.Error(err)
				return
			}
		}
		time.Sleep(24 * time.Hour)
	}
	return
}
