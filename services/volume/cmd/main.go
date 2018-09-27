package main

import (
	"context"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/kafkaHelper"
	"github.com/diadata-org/api-golang/services/volume"
	log "github.com/sirupsen/logrus"
)

func main() {

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTradesBlock)
	defer r.Close()

	f := volumeService.NewVolumeService()

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
