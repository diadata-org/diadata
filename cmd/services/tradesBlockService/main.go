package main

import (
	"context"
	"github.com/diadata-org/diadata/internal/pkg/tradesBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"sync"
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

	w := kafkaHelper.NewSyncWriter(kafkaHelper.TopicTradesBlock)
	defer w.Close()

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTrades)
	defer r.Close()

	s, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	tradesBlockService := tradesBlockService.NewTradesBlockService(s, dia.BlockSizeSeconds)

	wg := sync.WaitGroup{}
	go handleBlocks(tradesBlockService, &wg, w)

	log.Printf("starting...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var t dia.Trade
			err := t.UnmarshalBinary(m.Value)
			if err == nil {
				tradesBlockService.ProcessTrade(&t)
			} else {
				log.Printf("ignored message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			}
		}
	}
}
