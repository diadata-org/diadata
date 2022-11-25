package main

import (
	"context"
	"flag"
	"sync"
	"time"

	metafilters "github.com/diadata-org/diadata/internal/pkg/metaFilterService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	filtersBlockTopic  int
	metaFilterInTopic  int
	metaFilterOutTopic int
)

func init() {
	flag.Parse()

	filtersBlockTopic = kafkaHelper.TopicFiltersBlock
	metaFilterInTopic = kafkaHelper.TopicMetaFilterIn
	metaFilterOutTopic = kafkaHelper.TopicMetaFilterOut

}

func main() {

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	channel := make(chan *dia.FiltersBlock)

	f := metafilters.NewMetaFilterService(loadPairFilterPointsFromPreviousBlock(), datastore, channel)

	w := kafkaHelper.NewSyncWriterWithCompression(metaFilterOutTopic)

	defer func() {
		err := w.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	wg := sync.WaitGroup{}

	go handler(channel, &wg, w)

	fbsReader := kafkaHelper.NewReaderNextMessage(metaFilterInTopic)
	defer func() {
		err := fbsReader.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	for {
		m, err := fbsReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Info("get block from tradesBlock")
			var fb dia.FiltersBlock
			err := fb.UnmarshalBinary(m.Value)
			if err != nil {
				log.Error("error unmarshalling filters block")
			}
			if err == nil {
				t0 := time.Now()
				log.Info("number of trades in received filtersblock: ", len(fb.FiltersBlockData.FilterPoints))
				f.ProcessFiltersBlock(&fb)
				log.Info("time spent by filtersblockservice for processing filtersblock: ", time.Since(t0))
				// In historical mode, send timestamp of last trade as soon as fbs is done.

			}
		}
	}

}

func handler(channel chan *dia.FiltersBlock, wg *sync.WaitGroup, w *kafka.Writer) {
	var block int
	for {
		filtersblock, ok := <-channel
		if !ok {
			log.Printf("handler: finishing channel")
			wg.Done()
			return
		}
		block++
		log.Infoln("kafka: generated ", block, " blocks")
		err := kafkaHelper.WriteMessage(w, filtersblock)
		if err != nil {
			log.Errorln("kafka: handleBlocks", err)
		}
	}
}

func loadPairFilterPointsFromPreviousBlock() []dia.FilterPoint {
	// load the previous block points so that we have a value even if
	// there is no filter values
	lastFilterPoints := []dia.FilterPoint{}
	lastFilterBlock, err := kafkaHelper.GetLastElement(metaFilterOutTopic)
	if err == nil {
		lastFilterPoints = lastFilterBlock.(dia.FiltersBlock).FiltersBlockData.FilterPoints
	}
	return lastFilterPoints
}
