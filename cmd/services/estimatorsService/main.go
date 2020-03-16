package main

import (
	"context"
	"github.com/diadata-org/diadata/internal/pkg/estimators"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"sync"
)

const (
	//maxHistoryBlocks number of data points to retrieve from db
	maxHistoryBlocks = 10000
)

func handler(channel chan estimators.PDF, wg *sync.WaitGroup, w *kafka.Writer) {
	var block int
	for {
		t, ok := <-channel
		if !ok {
			log.Printf("handler: finishing channel")
			wg.Done()
			return
		}
		block++
		log.Infoln("generated ", block, " blocks")
		kafkaHelper.WriteMessage(w, t)
	}
}

type estCfg struct {
	Exchange string
	Symbol   string
	Name     string
}

func parseConfig() ([]estCfg, error) {
	configFileAPI := "config/estimators.json"
	type Cfgs struct {
		Estimators []estCfg
	}
	var cfgs Cfgs
	err := gonfig.GetConf(configFileAPI, &cfgs)
	return cfgs.Estimators, err
}

func main() {
	channel := make(chan estimators.PDF)
	e := estimators.NewEstimatorsService(channel)
	if cfg, err := parseConfig(); err == nil {
		for _, est := range cfg {
			log.Info("Adding estimator:" + est.Name + " For symbol:" + est.Symbol + " in exchange:" + est.Exchange)
			e.CreateEstimator(est.Name, est.Symbol, est.Exchange)
		}
	} else {
		panic(err)
	}
	w := kafkaHelper.NewWriter(kafkaHelper.TopicPDFEstimators)
	defer w.Close()

	wg := sync.WaitGroup{}

	go handler(channel, &wg, w)

	r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTradesBlock)
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf(err.Error())
		} else {
			var tb dia.TradesBlock
			err := tb.UnmarshalBinary(m.Value)
			if err == nil {
				e.ProcessTradesBlock(&tb)
			}
		}
	}
}
