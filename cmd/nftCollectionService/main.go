package main

import (
	"errors"
	"flag"
	"time"

	"github.com/jackc/pgconn"

	nftsource "github.com/diadata-org/diadata/internal/pkg/nftService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

var (
	log       = logrus.New()
	nftSource *string
	secret    *string
)

const (
	fetchPeriodMinutes = 24 * 60
)

func init() {
	nftSource = flag.String("source", "Opensea", "Data source for nft data collection")
	secret = flag.String("secret", "", "secret for asset source")
	flag.Parse()
}

// NewAssetScraper returns a scraper for assets on @exchange.
// For NewJSONReader @exchange is the folder in the config folder and @secret the filename.
func NewNFTClassCollector(nftSource string, secret string) nftsource.NFTClassSource {
	switch nftSource {
	case "Opensea":
		return nftsource.NewOpenseaNFTSource(secret)
	// case "nftlists":
	// 	return nftsource.NewJSONReader(nftSource)
	default:
		return nil
	}
}

func main() {

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	// Initial run:
	runNFTSource(relDB, *nftSource, *secret)
	// Afterwards, run every @fetchPeriodMinutes
	ticker := time.NewTicker(fetchPeriodMinutes * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				runNFTSource(relDB, *nftSource, *secret)
			}
		}
	}()
	select {}
}

func runNFTSource(relDB *models.RelDB, source string, secret string) error {

	log.Println("Fetching asset from ", source)
	nftCollector := NewNFTClassCollector(source, secret)
	for {
		select {
		case receivedClass := <-nftCollector.NFTClass():
			// Set to persistent DB
			err := relDB.SetNFTClass(receivedClass)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code == "23505" {
						log.Infof("asset %v already in db. continue.", receivedClass)
						continue
					} else {
						log.Errorf("postgres error saving asset %v: %v", receivedClass, err)
					}
				} else {
					log.Errorf("Error saving asset %v: %v", receivedClass, err)
				}
			} else {
				log.Info("successfully set asset ", receivedClass)
			}

		case <-nftCollector.Close():
			return nil
		}
	}
}
