package main

import (
	"github.com/diadata-org/diadata/internal/pkg/graphService"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	GRAPH_PATH = "/charts/"
	EXTENSION  = ".png"
)

func main() {
	dataStore, err := models.NewDataStore()
	if err != nil {
		log.Error(err)
		return
	}

	for {
		symbols, err := dataStore.GetAllSymbols()
		if err != nil {
			log.Error(err)
			return
		}

		for _, symbol := range symbols {

			log.Info("sleeping...")
			time.Sleep(1 * time.Second)

			points, err := dataStore.GetChartPoints7Days(symbol)
			if err != nil {
				log.Error(err)
				continue
			}

			if len(points) <= 0 {
				log.Println("No datapoints for symbol", symbol)
				continue
			}

			log.Println("Producing chart for", symbol, "with", len(points), "source datapoints")
			timePoints := make([]int64, len(points))
			pricePoints := make([]float64, len(points))

			for i := 0; i < len(points); i++ {
				timePoints[i] = points[i].UnixTime
				pricePoints[i] = points[i].Value
			}

			if _, err := os.Stat(GRAPH_PATH); os.IsNotExist(err) {
				err = os.MkdirAll(GRAPH_PATH, os.ModeDir|os.ModePerm)
				log.Error(err)
				continue
			}

			err = graph.PriceGraph(pricePoints, timePoints, GRAPH_PATH+symbol+EXTENSION)
			if err != nil {
				log.Error(err)
				continue
			}
			err = os.Rename(GRAPH_PATH+symbol+EXTENSION, GRAPH_PATH+symbol)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}
