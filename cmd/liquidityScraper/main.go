package main

import (
	"flag"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	liquidityscraper "github.com/diadata-org/diadata/pkg/dia/scraper/liquidity-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/sirupsen/logrus"
)

var (
	exchangeName *string
	log          *logrus.Logger
)

func init() {
	exchangeName = flag.String("exchange", "Uniswap", "name of DEX.")
	flag.Parse()
	log = logrus.New()
}

func main() {

	log.Println("Liquidity Scraper: Start scraping liquidity")

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to postgres: ", err)
		return
	}

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("Error connecting to postgres: ", err)
		return
	}

	var feedselection1 dia.FeedSelection
	var eps dia.ExchangepairSelection
	feedselection1.Asset = dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN}
	eps.Exchange = dia.Exchange{Name: dia.BittrexExchange, Centralized: true}
	eps.Pairs = []dia.Pair{
		{
			BaseToken:  dia.Asset{Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Blockchain: dia.ETHEREUM},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
		{
			BaseToken:  dia.Asset{Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7", Blockchain: dia.ETHEREUM},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
		{
			BaseToken:  dia.Asset{Address: "840", Blockchain: dia.FIAT},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
	}
	feedselection1.Exchangepairs = append(feedselection1.Exchangepairs, eps)
	eps.Exchange = dia.Exchange{Name: dia.MEXCExchange, Centralized: true}
	eps.Pairs = []dia.Pair{
		{
			BaseToken:  dia.Asset{Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Blockchain: dia.ETHEREUM},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
		{
			BaseToken:  dia.Asset{Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7", Blockchain: dia.ETHEREUM},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
	}
	feedselection1.Exchangepairs = append(feedselection1.Exchangepairs, eps)

	var feedselection2 dia.FeedSelection
	feedselection2.Asset = dia.Asset{Address: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", Blockchain: dia.ETHEREUM}
	eps.Exchange = dia.Exchange{Name: dia.UniswapExchange, Centralized: false}
	eps.Pools = []dia.Pool{
		{Address: "0xBb2b8038a1640196FbE3e38816F3e67Cba72D940", Blockchain: dia.BlockChain{Name: dia.ETHEREUM}},
		{Address: "0xcD7989894bc033581532D2cd88Da5db0A4b12859", Blockchain: dia.BlockChain{Name: dia.ETHEREUM}},
	}
	feedselection2.Exchangepairs = append(feedselection2.Exchangepairs, eps)

	endtime := time.Now()
	starttime := endtime.Add(-time.Duration(40 * time.Minute))

	var (
		feedselection3 dia.FeedSelection
		feedselection4 dia.FeedSelection
		feedselection5 dia.FeedSelection
	)
	feedselection3.Asset = dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN}
	var eps3 []dia.ExchangepairSelection
	eps3 = append(eps3, dia.ExchangepairSelection{Exchange: dia.Exchange{Name: dia.BittrexExchange, Centralized: true}})
	eps3 = append(eps3, dia.ExchangepairSelection{Exchange: dia.Exchange{Name: dia.CoinBaseExchange, Centralized: true}})
	feedselection3.Exchangepairs = eps3

	feedselection4.Asset = dia.Asset{Address: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", Blockchain: dia.ETHEREUM}

	feedselection5.Asset = dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN}
	var eps5 dia.ExchangepairSelection
	eps5.Pairs = []dia.Pair{
		{
			BaseToken:  dia.Asset{Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Blockchain: dia.ETHEREUM},
			QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		},
		// {
		// 	BaseToken:  dia.Asset{Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7", Blockchain: dia.ETHEREUM},
		// 	QuoteToken: dia.Asset{Address: "0x0000000000000000000000000000000000000000", Blockchain: dia.BITCOIN},
		// },
	}
	feedselection5.Exchangepairs = append(feedselection5.Exchangepairs, eps5)

	s, err := datastore.GetTradesByFeedSelection([]dia.FeedSelection{feedselection5}, starttime, endtime)
	if err != nil {
		log.Error("get separator: ", err)
	} else {
		for _, tr := range s {
			if tr.Source == dia.UniswapExchange {
				log.Info("s: ", tr)
			}
		}
	}
	time.Sleep(10 * time.Minute)
	runLiquiditySource(relDB, datastore, *exchangeName)
	log.Infof("Successfully ran pool collector for %s", *exchangeName)

}

func runLiquiditySource(relDB *models.RelDB, datastore *models.DB, source string) {
	log.Info("Fetching pools from ", source)
	scraper := liquidityscraper.NewLiquidityScraper(source, relDB, datastore)

	for {
		select {
		case receivedPool := <-scraper.Pool():

			// Set to persistent DB.
			err := relDB.SetPool(receivedPool)
			if err != nil {
				log.Errorf("Error saving pool %v: %v", receivedPool, err)
			} else {
				log.Info("successfully set pool ", receivedPool)
			}

		case <-scraper.Done():
			return
		}
	}

}
