package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var (
	db models.Datastore
)

type Pairs struct {
	Coins []dia.Pair
}

func getPairsFromConfig(exchange string) ([]dia.Pair, error) {
	configFileAPI := configCollectors.ConfigFileConnectors(exchange)
	// configFileAPI := "config/" + exchange + ".json"
	var coins Pairs
	err := gonfig.GetConf(configFileAPI, &coins)
	return coins.Coins, err
}

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			updateExchangePairs()
		}
	}
}

func (t *Task) stop() {
	log.Println("Stoping exchange pair update thread...")
	close(t.closed)
	t.wg.Wait()
	log.Println("Thread stopped, cleaning...")
	// Clean if required
	log.Println("Done")
}

func savePairsToFile(exchange string, pairs []dia.Pair) {
	log.Info("savePairsToFile:", exchange)
	b, e := json.Marshal(&Pairs{pairs})
	if e == nil {
		log.Info("marshalled ")
	} else {
		log.Error("erreur save", e)
	}
	e = ioutil.WriteFile("/tmp/"+exchange+".json", b, 0644)
}

func updateExchangePairs() {
	t, err := db.GetConfigTogglePairDiscovery()
	if err != nil {
		log.Error("updateExchangePairs GetConfigTogglePairDiscovery", err.Error())
		return
	}
	if t == false {
		log.Info("GetConfigTogglePairDiscovery = false, using default values")
		getInitialExchangePairs()
	} else {
		for _, e := range dia.Exchanges() {
			if e == "CoinBase" || e == "Huobi" || e == "Unknown" {
				continue
			}
			log.Println("Updating", e)
			c, err := dia.GetConfig(e)
			var s scrapers.APIScraper
			if err == nil {
				s = scrapers.NewAPIScraper(e, c.ApiKey, c.SecretKey)
			} else {
				log.Error("Error processing config for exchange:"+e+" error:", err.Error())
				s = scrapers.NewAPIScraper(e, "", "")
			}
			if s != nil {
				p, err := s.FetchAvailablePairs()
				if err == nil {
					addLocalPairs(e, p)
					err := db.SetAvailablePairsForExchange(e, p)
					if err == nil {
						log.Println("Exchange :" + e + " updated")
					} else {
						log.Error("Error adding pairs  to redis for exchange:"+e+" error:", err.Error())
					}
				} else {
					//	log.Info("locale ", err.Error())
					log.Error("Error fetching pairs for exchange:"+e+" error:", err.Error())
				}
				go func(s scrapers.APIScraper, e string) {
					time.Sleep(5 * time.Second)
					log.Error("Closing scrapper: " + e)
					s.Close()
				}(s, e)
			} else {
				log.Error("Error creating APIScraper forexchange:" + e)
			}
		}
		log.Println("Update complete.")
	}
}

func addLocalPairs(exchange string, remotePairs []dia.Pair) []dia.Pair {
	pLocales, _ := getPairsFromConfig(exchange)
	log.Info(exchange, " nb remote:", len(remotePairs), " nb pLocales:", len(pLocales))
	for i := range remotePairs {
		remotePairs[i].Ignore = true
	}
	for i, e := range remotePairs {
		for _, a := range pLocales {
			if a.Exchange == e.Exchange && a.Symbol == e.Symbol && e.ForeignName == a.ForeignName {
				remotePairs[i].Ignore = false
			}
		}
	}
	savePairsToFile(exchange, remotePairs)
	return remotePairs
}

func getInitialExchangePairs() {
	log.Println("Loading pairs from config...")
	for _, e := range dia.Exchanges() {
		if e == "Unknown" {
			continue
		}
		p, err := getPairsFromConfig(e)
		if err == nil {
			pairsToSave := []dia.Pair{}
			for _, pp := range p {
				if !pp.Ignore {
					pairsToSave = append(pairsToSave, pp)
				} else {
					log.Debug("ignoring", pp)
				}
			}
			// savePairsToFile(e, p)
			err := db.SetAvailablePairsForExchange(e, pairsToSave)
			if err == nil {
				log.Println("Exchange :" + e + " set")
			} else {
				log.Error("Error setting pairs for exchange:"+e+" error:", err.Error())
			}
		} else {
			log.Error("Error processing config for exchange:"+e+" error:", err.Error())
		}
	}
	log.Println("Update complete.")
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		/// Retrieve every hour
		ticker: time.NewTicker(time.Second * 60 * 60),
	}
	var e error
	db, e = models.NewDataStore()
	if e != nil {
		panic("Can not initialize db error:" + e.Error())
	}
	updateExchangePairs()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run() }()
	select {
	case <-c:
		log.Println("Got signal.")
		task.stop()
	}
}
