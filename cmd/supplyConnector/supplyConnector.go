package main

import (
	supply "github.com/diadata-org/diadata/internal/pkg/supplyBlockService"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/segmentio/kafka-go"
	"sync"
)

// handleTrades delegates trade information to Kafka
func handle(s *supply.SupplyScraper, wg *sync.WaitGroup, w *kafka.Writer) {
	for {
		t, ok := <-s.Channel()
		if !ok {
			wg.Done()
			return
		}
		kafkaHelper.WriteMessage(w, t)
	}
}

func main() {
	w := kafkaHelper.NewWriter(kafkaHelper.TopicSuppliesBlock)
	defer w.Close()

	wg := sync.WaitGroup{}

	s := supply.NewSupplyScraper()

	go handle(s, &wg, w)
	wg.Add(1)
	defer wg.Wait()

}
