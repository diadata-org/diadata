package main

import (
	"context"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/kafkaHelper"
	"github.com/diadata-org/api-golang/services/tradesBlockService/tradesBlock"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"sync"
	"github.com/diadata-org/api-golang/services/model"
)

func handleBlocks(blockMaker *tradesBlockService.TradesBlockService, wg *sync.WaitGroup, w *kafka.Writer) {
	for {
		t, ok := <-blockMaker.Channel()
		if !ok {
			log.Printf("handleBlocks: finishing channel")
			wg.Done()
			return
		}
		kafkaHelper.WriteMessage(w, t)
	}
}

func main() {

	w := kafkaHelper.NewWriter(kafkaHelper.TopicTradesBlock)
	defer w.Close()

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTrades)
	defer r.Close()

	s, err := models.NewDataStore()
	if err != nil {
		log.Error(err)
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
