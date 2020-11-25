package main

import (
	"context"
	"github.com/diadata-org/diadata/internal/pkg/indexBlockService"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/getsentry/raven-go"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"sync"
)

func init() {
	raven.SetDSN("https://b64e05c72835491a96f5ce3c16d17d6d:9e2eebf1e7074dd9acbc2833e1d618b5@sentry.io/1332578")
}

func handler(i chan *dia.IndexBlock, wg *sync.WaitGroup, w *kafka.Writer) {
	for {
		t, ok := <-i
		if !ok {
			log.Printf("handler: finishing channel")
			wg.Done()
			return
		}
		err := kafkaHelper.WriteMessage(w, t)
		if err != nil {
			log.Errorln("kafka: handleBlocks", err)
		}
	}
}

func listenForSuppliesUpdates(indexBlockService *indexBlockService.IndexBlockService) {
	log.Printf("listenForSuppliesUpdates")
	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicSuppliesBlock)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var b dia.SuppliesBlock
			err := b.UnmarshalBinary(m.Value)
			if err == nil {
				indexBlockService.ProcessSuppliesBlock(&b)
			}
		}
	}
}

/*func listenForVolatilityUpdates(indexBlockService *indexBlockService.IndexBlockService) {
	log.Printf("listenForSuppliesUpdates")
	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicVolatilityBlock)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var b priceCollection.VolatilityBlock
			err := b.UnmarshalBinary(m.Value)
			if err == nil {
				indexBlockService.ProcessVolatilityBlock(&b)
			}
		}
	}
}*/

func main() {

	w := kafkaHelper.NewSyncWriter(kafkaHelper.TopicIndexBlock)
	defer w.Close()

	w2 := kafkaHelper.NewSyncWriter(kafkaHelper.TopicIndexBlock2)
	defer w2.Close()

	wDaily := kafkaHelper.NewSyncWriter(kafkaHelper.TopicIndexBlockDaily)
	defer wDaily.Close()

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicFiltersBlock)
	defer r.Close()

	wg := sync.WaitGroup{}

	/*var vb *priceCollection.VolatilityBlock

	m, err := kafkaHelper.GetLastElement(kafkaHelper.TopicVolatilityBlock)
	if err == nil {
		vb2 := m.(priceCollection.VolatilityBlock)
		vb = &vb2
	} else {
		log.Error("couldnt get last element on TopicVolatilityBlock", err)
	}*/

	var sb *dia.SuppliesBlock
	m2, err := kafkaHelper.GetLastElement(kafkaHelper.TopicSuppliesBlock)
	if err == nil {
		sb2 := m2.(dia.SuppliesBlock)
		sb = &sb2
	} else {
		log.Error("couldnt get last element on TopicSuppliesBlock", err)
	}

	calculator := indexBlockService.NewIndexBlockService(sb)

	go handler(calculator.ChanIndexBlock(), &wg, w)
	go handler(calculator.ChanIndexBlock2(), &wg, w2)
	go handler(calculator.ChanIndexBlockDaily(), &wg, wDaily)

	go listenForSuppliesUpdates(calculator)
	//go listenForVolatilityUpdates(calculator)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var b dia.FiltersBlock
			err := b.UnmarshalBinary(m.Value)
			if err == nil {
				calculator.ProcessFiltersBlock(&b)
			} else {
				log.Printf("ignored message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			}
		}
	}
}
