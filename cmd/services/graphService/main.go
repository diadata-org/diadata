package main

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/graph"
	"github.com/diadata-org/diadata/pkg/model"
	"log"
	"os"
	"time"
)

const (
	GRAPH_PATH = "/charts/"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dataStore, err := models.NewDataStore()
	if err != nil {
		log.Println(err)
		return
	}

	for {
		for _, symbol := range dia.SymbolsFrontPage() {
			points, err := dataStore.GetChartPoints(symbol)
			if err != nil {
				log.Println(err)
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
				log.Println(err)
			}

			err = graph.PriceGraph(pricePoints, timePoints, GRAPH_PATH+symbol+".png")
			if err != nil {
				log.Println(err)
			}
			err = os.Rename(GRAPH_PATH+symbol+".png", GRAPH_PATH+symbol)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(time.Minute * 2)
		}
	}
}