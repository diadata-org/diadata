package main

import (
	"context"
	"flag"
	"sync"

	"github.com/diadata-org/diadata/internal/pkg/tradesBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func handleBlocks(blockMaker *tradesBlockService.TradesBlockService, wg *sync.WaitGroup, w *kafka.Writer) {
	for {
		t, ok := <-blockMaker.Channel()
		if !ok {
			log.Printf("handleBlocks: finishing channel")
			wg.Done()
			return
		}
		err := kafkaHelper.WriteMessage(w, t)
		if err != nil {
			log.Errorln("handleBlocks", err)
		}
	}
}

func init() {
	flag.Parse()
	if !*historical {
		tradesBlockTopic = kafkaHelper.TopicTradesBlock
		tradesTopic = kafkaHelper.TopicTrades
	}
	if *testing {
		tradesBlockTopic = kafkaHelper.TopicTradesBlockTest
		tradesTopic = kafkaHelper.TopicTradesTest
	}
	if *replica {
		tradesBlockTopic = kafkaHelper.TopicTradesBlockReplica
		tradesTopic = kafkaHelper.TopicTradesReplica
	}
}

var (
	historical       = flag.Bool("historical", false, "digest current or historical trades.")
	testing          = flag.Bool("testing", false, "set true for testing environment.")
	replica          = flag.Bool("replica", false, "set true if trades should be fetched from and forwarded to replica topics.")
	tradesBlockTopic int
	tradesTopic      int
)

func main() {
	if *historical {
		log.Info("run tradesblock service in historical mode")
	}

	kafkaWriter := kafkaHelper.NewSyncWriterWithCompression(tradesBlockTopic)
	defer func() {
		err := kafkaWriter.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	kafkaReader := kafkaHelper.NewReaderNextMessage(tradesTopic)
	defer func() {
		err := kafkaReader.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	s, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Error("New relational datastore: ", err)
	}

	service := tradesBlockService.NewTradesBlockService(s, relDB, dia.BlockSizeSeconds, *historical)

	wg := sync.WaitGroup{}
	go handleBlocks(service, &wg, kafkaWriter)

	log.Printf("starting...")

	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var t dia.Trade
			err := t.UnmarshalBinary(m.Value)
			if err == nil {
				service.ProcessTrade(&t)
			} else {
				log.Printf("ignored message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			}
		}
	}
}
