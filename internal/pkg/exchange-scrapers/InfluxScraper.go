package scrapers

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/segmentio/kafka-go"
)

var (
	filtersblockDoneTopic int
)

const (
	// One hour batches (notation in nanoseconds)
	batchDuration         = "121000000000"
	tradesReadMeasurement = "tradesOld"
)

type InfluxScraper struct {
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	run       bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers  map[string]*InfluxPairScraper
	exchangeName  string
	chanTrades    chan *dia.Trade
	measurement   string
	batchDuration int64
	db            *models.DB
	fbsDoneReader *kafka.Reader
}

// NewGateIOScraper returns a new GateIOScraper for the given pair
func NewInfluxScraper(scrape bool) *InfluxScraper {

	log.Info("make new Influx scraper...")
	db, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}
	influxURL := utils.Getenv("INFLUXURL", "")
	measurement := utils.Getenv("INFLUX_TRADES_MEASUREMENT", tradesReadMeasurement)
	batchDurationEnv := utils.Getenv("BATCH_DURATION", batchDuration)
	batchDurationInt, err := strconv.ParseInt(batchDurationEnv, 10, 64)
	if err != nil {
		log.Fatal("parse batch duration ", err)
	}
	db.SetInfluxClient(influxURL)

	// Make a kafka reader that listens to ok from the filtersblockservice
	filtersblockDoneTopic = kafkaHelper.TopicFiltersBlockDone
	fbsDoneReader := kafkaHelper.NewReaderNextMessage(filtersblockDoneTopic)

	s := &InfluxScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		pairScrapers:  make(map[string]*InfluxPairScraper),
		exchangeName:  "Influx",
		error:         nil,
		chanTrades:    make(chan *dia.Trade),
		measurement:   measurement,
		db:            db,
		batchDuration: batchDurationInt,
		fbsDoneReader: fbsDoneReader,
	}
	if scrape {
		go s.mainLoop()
	}

	return s
}

// runs in a goroutine until s is closed
func (s *InfluxScraper) mainLoop() {
	log.Info("enter main loop")
	var timeInit time.Time
	var timeInitInt int64
	var err error
	// Either take first timestamp from env var or take first trade time from DB.
	timeInitString := utils.Getenv("TIME_INIT", "")
	if timeInitString == "" {
		log.Info("get first trade date...")
		timeInit, err = s.db.GetFirstTradeDate(s.measurement)
		if err != nil {
			log.Error("get first trade date: ", err)
		}
		log.Info("got first trade date: ", timeInit)
	} else {
		timeInitInt, err = strconv.ParseInt(timeInitString, 10, 64)
		if err != nil {
			log.Fatal("parse init time: ", err)
		}
		timeInit = time.Unix(timeInitInt, 0)
	}

	// After @waitForFBSSeconds new trades are collected in order to successfully produce a new filtersblock.
	waitForFBSSecondsString := utils.Getenv("WAIT_FOR_FBS_SECONDS", "30")
	waitForFBSSeconds, err := strconv.ParseInt(waitForFBSSecondsString, 10, 64)
	if err != nil {
		log.Error("parse batch processing time string: ", err)
	}

	// final time is the last timestamp of trades exported from d2.
	timeFinalString := utils.Getenv("TIME_FINAL", "1636618800")
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Fatal("parse final time: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	starttime := timeInit
	endtime := starttime.Add(time.Duration(s.batchDuration))

	// Initial run.
	starttime, endtime = s.collectTrades(starttime, endtime)

	// Context initiates the collection of new trades if the fbs does not
	// return a signal after @waitForFBSSeconds seconds.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitForFBSSeconds)*time.Second)

	// Iterate until @timeFinal is reached.
	for {
		_, err := s.fbsDoneReader.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, ctx.Err()) {
				log.Info("system stalled. collect new trades.")
				starttime, endtime = s.collectTrades(starttime, endtime)
				cancel()
				ctx, cancel = context.WithTimeout(context.Background(), time.Duration(waitForFBSSeconds)*time.Second)
			} else {
				log.Error("read ok message from filtersblockservice: ", err.Error())
			}
		} else if starttime.Before(timeFinal) {
			starttime, endtime = s.collectTrades(starttime, endtime)
			cancel()
			ctx, cancel = context.WithTimeout(context.Background(), time.Duration(waitForFBSSeconds)*time.Second)
		} else {
			log.Info("done with iteration through trades. last timestamp: ", endtime)
			time.Sleep(120 * time.Hour)
		}
	}
}

func (s *InfluxScraper) collectTrades(starttime time.Time, endtime time.Time) (newstarttime time.Time, newendtime time.Time) {
	t0 := time.Now()
	trades, err := s.db.GetOldTradesFromInflux(s.measurement, "", true, starttime, endtime)
	if err != nil {
		if strings.Contains(err.Error(), "no trades in time range") {
			log.Warnf("%v: %v -- %v", err, starttime, endtime)
		} else {
			log.Error("get trades from influx: ", err)
			return
		}
	}
	log.Infof("got %d trades in time range %v -- %v", len(trades), starttime, endtime)
	log.Info("time passed for get old trades: ", time.Since(t0))
	for i := range trades {
		s.chanTrades <- &trades[i]
		// log.Info("got trade: ", trades[i])
	}
	log.Infof("fetched %d trades from influx.", len(trades))
	return endtime, endtime.Add(time.Duration(s.batchDuration))
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *InfluxScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	ps := &InfluxPairScraper{
		parent: s,
		pair:   pair,
	}
	return ps, nil
}

func (s *InfluxScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

// FetchTickerData collects all available information on an asset traded on GateIO
func (s *InfluxScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	return asset, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *InfluxScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *InfluxScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *InfluxScraper) Close() error {
	// close the pair scraper channels
	s.run = false
	for _, pairScraper := range s.pairScrapers {
		pairScraper.closed = true
	}

	close(s.shutdown)
	<-s.shutdownDone
	return nil
}

// GateIOPairScraper implements PairScraper for GateIO
type InfluxPairScraper struct {
	parent *InfluxScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *InfluxPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *InfluxPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *InfluxPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
