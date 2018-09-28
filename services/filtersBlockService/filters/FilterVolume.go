package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/services/model"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"time"
)

type FilterVolume struct {
	symbol      string
	currentTime time.Time
	volumeUSD   float64
	param       int
	lastTrade   *dia.Trade
}

func NewFilterVolume(symbol string, currentTime time.Time, param int) *FilterVolume {
	s := &FilterVolume{
		symbol:      symbol,
		param:       param, // TOFIX currently hardcoded on blocksize
		volumeUSD:   0.0,
		currentTime: currentTime,
	}
	return s
}

func (s *FilterVolume) filterPoint(time time.Time) dia.FilterPoint {
	r := dia.FilterPoint{
		Symbol: s.symbol,
		Value:  s.volumeUSD,
		Name:   "VOL" + strconv.Itoa(s.param),
		Time:   time,
	}
	return r
}

func (s *FilterVolume) compute(trade *dia.Trade) {
	s.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
}

func (s *FilterVolume) save(ds models.Datastore) error {
	err := ds.SetVolume(s.symbol, s.volumeUSD)
	if err != nil {
		log.Error("Error: %v\n", err)
	}
	s.volumeUSD = 0.0 //TOFIX
	return err
}

func (s *FilterVolume) copyToFilterBlock() bool {
	return false
}
