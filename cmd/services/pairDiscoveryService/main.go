package main

import (
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"sync"
	"time"
)

var (
	db models.Datastore
)

func getConfig(exchange string) (*dia.ConfigApi, error) {
	var configApi dia.ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}

func getPairsFromConfig(exchange string) ([]dia.Pair, error) {
	type Pairs struct {
		Coins []dia.Pair
	}
	configFileAPI := "config/" + exchange + ".json"
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

func updateExchangePairs() {
	log.Println("Updating exchanges pairs...")
	for _, e := range dia.Exchanges() {
		c, err := getConfig(e)
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
				err := db.SetAvailablePairsForExchange(e, p)
				if err == nil {
					log.Println("Exchange :" + e + " updated")
				} else {
					log.Error("Error adding pairs  to redis for exchange:"+e+" error:", err.Error())
				}
			} else {
				log.Error("Error fetching pairs for exchange:"+e+" error:", err.Error())
			}
			go func(s scrapers.APIScraper, e string) {
				time.Sleep(5 * time.Second)
				log.Error("Closing scrapper: " + e)

			}(s, e)
		} else {
			log.Error("Error creating APIScraper forexchange:" + e)
		}
	}
	log.Println("Update complete.")
}

func getInitialExchangePairs() {
	log.Println("Loading pairs from config...")
	for _, e := range dia.Exchanges() {
		p, err := getPairsFromConfig(e)
		if err == nil {
			err := db.SetAvailablePairsForExchange(e, p)
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
	getInitialExchangePairs()
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
