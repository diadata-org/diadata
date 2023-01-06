package main

import (
	"context"
	"flag"
	"time"

	metafilters "github.com/diadata-org/diadata/internal/pkg/metaFilterService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (
	filtersBlockTopic int
)

func init() {
	flag.Parse()
	filtersBlockTopic = kafkaHelper.TopicFiltersBlock
}

func main() {

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	channel := make(chan *dia.FiltersBlock)

	f := metafilters.NewMetaFilterService(datastore, channel)

	// Reader for blocks from filtersblockservice.
	fbsReader := kafkaHelper.NewReaderNextMessage(filtersBlockTopic)
	defer func() {
		err := fbsReader.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// Continuously read from stream of filtersblocks and process them.
	for {
		m, err := fbsReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Info("get block from filtersBlock")
			var fb dia.FiltersBlock
			err := fb.UnmarshalBinary(m.Value)
			if err != nil {
				log.Error("error unmarshalling filters block")
			}
			if err == nil {
				t0 := time.Now()
				log.Info("number of filters in received metafiltersblock: ", len(fb.FiltersBlockData.FilterPoints))
				f.ProcessFiltersBlock(&fb)
				log.Info("time spent by filtersblockservice for processing filtersblock: ", time.Since(t0))
				// In historical mode, send timestamp of last trade as soon as fbs is done.

			}
		}
	}
}
