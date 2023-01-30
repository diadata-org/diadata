package main

import (
	"context"
	"flag"
	"sync"
	"time"

	pairfilters "github.com/diadata-org/diadata/internal/pkg/pairFilterService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	testing           = flag.Bool("testing", false, "set true for testing environment")
	filtersBlockTopic int
	tradesBlockTopic  int
)

func init() {
	flag.Parse()

	filtersBlockTopic = kafkaHelper.TopicFiltersBlock
	tradesBlockTopic = kafkaHelper.TopicTradesBlock

	if *testing {
		filtersBlockTopic = kafkaHelper.TopicFiltersBlockTest
		tradesBlockTopic = kafkaHelper.TopicTradesBlockTest
	}
	// if *replica {
	// 	filtersBlockTopic = kafkaHelper.TopicFiltersBlockReplica
	// 	tradesBlockTopic = kafkaHelper.TopicTradesBlockReplica
	// }
}

func main() {

	s, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	channel := make(chan *dia.FiltersBlock)

	f := pairfilters.NewFiltersBlockService(s, channel)

	// Writer that sends filter blocks to the metaFilter service.
	w := kafkaHelper.NewSyncWriterWithCompression(filtersBlockTopic)
	defer func() {
		err := w.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// Handle filter blocks, i.e. send to metaFilter service
	wg := sync.WaitGroup{}
	go handler(channel, &wg, w)

	// Reader for blocks from tradesblockservice.
	r := kafkaHelper.NewReaderNextMessage(tradesBlockTopic)
	defer func() {
		err := r.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// Continuously read from stream of tradesblocks and process them.
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Info("get block from tradesBlock")
			var tb dia.TradesBlock
			err := tb.UnmarshalBinary(m.Value)
			if err != nil {
				log.Error("error unmarshalling trades block")
			}
			if err == nil {
				t0 := time.Now()
				log.Info("number of trades in received tradesblock: ", len(tb.TradesBlockData.Trades))
				f.ProcessTradesBlock(&tb)
				log.Info("time spent by filtersblockservice for processing tradesblock: ", time.Since(t0))
			}
		}
	}

}

func handler(channel chan *dia.FiltersBlock, wg *sync.WaitGroup, w *kafka.Writer) {
	var block int
	for {
		log.Info("enter filtersblock write loop...")
		filtersblock, ok := <-channel
		log.Info("...got filtersblock")
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
