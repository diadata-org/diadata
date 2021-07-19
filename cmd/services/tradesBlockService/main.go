package main

import (
	"context"
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

func main() {

	kafkaWriter := kafkaHelper.NewSyncWriter(kafkaHelper.TopicTradesBlock)
	defer func() {
		err := kafkaWriter.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	kafkaReader := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTrades)
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

	service := tradesBlockService.NewTradesBlockService(s, dia.BlockSizeSeconds)

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
