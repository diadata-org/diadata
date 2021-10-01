package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"

	defiscraper "github.com/diadata-org/diadata/internal/pkg/defiscrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, close := signal.NotifyContext(context.Background(), os.Interrupt)
	defer close()

	log := makeLogger(ctx)

	rateType := flag.String("type", "DYDX", "Type of Defi rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatalf("unable to create dataStore instance: %+v", err.Error())
	}

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Fatalf("unable to create relDB instance: %+v", err.Error())
	}

	sRate := defiscraper.SpawnDefiScraper(ctx, log, relDB, ds, *rateType)
	defer sRate.Close()

	// Send rates to the database while the scraper scrapes
	wg.Add(1)
	go handleDefiInterestRate(ctx, sRate.RateChannel(), &wg, ds)

	wg.Add(1)
	go handleDefiState(ctx, sRate.StateChannel(), &wg, ds)

	wg.Wait()
}

// handleDefiInterestRate delegates rate information to Kafka
func handleDefiInterestRate(ctx context.Context, c chan *dia.DefiRate, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()

	// Pull from channel as long as not empty
	for {
		select {
		case <-ctx.Done():
			return

		case t, ok := <-c:
			if ok {
				ds.SetDefiRateInflux(t)
			}
		}
	}
}

// handleDefiState delegates rate information to Kafka
func handleDefiState(ctx context.Context, c chan *dia.DefiProtocolState, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()

	// Pull from channel as long as not empty
	for {
		select {
		case <-ctx.Done():
			return

		case t, ok := <-c:
			if ok {
				ds.SetDefiStateInflux(t)
			}
		}
	}
}

func makeLogger(ctx context.Context) *logrus.Entry {
	logger := logrus.New()

	logLevel := logrus.InfoLevel

	log := logger.WithContext(ctx).WithFields(logrus.Fields{
		"proc": "defiscrapper",
	})

	if logLevelStr, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if err := (&logLevel).UnmarshalText([]byte(logLevelStr)); err != nil {
			log.Fatalf("unable to parse log level from environment variable LOG_LEVEL: %+v", err.Error())
		}
	}

	logger.SetLevel(logLevel)

	return log
}
