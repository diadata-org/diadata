package main

import (
	"context"

	"github.com/diadata-org/diadata/internal/pkg/tradesEstimationService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {

	kafkaReader := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTradesEstimation)
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

	service := tradesEstimationService.NewTradesEstimationService(s)

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
