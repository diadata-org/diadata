package main

import (
	"context"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/kafkaHelper"
	"github.com/diadata-org/api-golang/internal/pkg/filtersBlockService"
	"github.com/diadata-org/api-golang/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"sync"
)

func handler(channel chan *dia.FiltersBlock, wg *sync.WaitGroup, w *kafka.Writer) {
	for {
		t, ok := <-channel
		if !ok {
			log.Printf("handler: finishing channel")
			wg.Done()
			return
		}
		kafkaHelper.WriteMessage(w, t)
	}
}

func loadFilterPointsFromPreviousBlock() []dia.FilterPoint {
	// load the previous block points so that we have a value even if
	// there is no trades
	lastFilterPoints := []dia.FilterPoint{}
	lastFilterBlock, err := kafkaHelper.GetLastElement(kafkaHelper.TopicFiltersBlock)
	if err == nil {
		lastFilterPoints = lastFilterBlock.(dia.FiltersBlock).FiltersBlockData.FilterPoints
	}
	return lastFilterPoints
}

func main() {

	w := kafkaHelper.NewWriter(kafkaHelper.TopicFiltersBlock)
	defer w.Close()

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTradesBlock)
	defer r.Close()

	channel := make(chan *dia.FiltersBlock)
	wg := sync.WaitGroup{}

	s, err := models.NewDataStore()
	if err != nil {
		log.Error(err)
	}

	f := filters.NewFiltersBlockService(loadFilterPointsFromPreviousBlock(), s, channel)

	go handler(channel, &wg, w)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var tb dia.TradesBlock
			err := tb.UnmarshalBinary(m.Value)
			if err == nil {
				f.ProcessTradesBlock(&tb)
			}
		}
	}
}
