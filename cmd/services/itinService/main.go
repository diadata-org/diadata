package main

import (
	itinService2 "github.com/diadata-org/diadata/internal/itinService"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
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
		itins, err := itinService2.GetItins(itinUrl)
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
}
