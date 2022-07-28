package defiscrapers

import (
	"errors"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const (
	// Determine frequency of scraping
	refreshRateDelay  = time.Second * 60 * 5
	refreshStateDelay = time.Second * 60 * 5
)

type nothing struct{}

// SpawnDefiScraper returns a new DefiScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func SpawnDefiScraper(datastore models.Datastore, rateType string) *DefiScraper {
	s := &DefiScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		tickerRate:    time.NewTicker(refreshRateDelay),
		tickerState:   time.NewTicker(refreshStateDelay),
		datastore:     datastore,
		chanDefiRate:  make(chan *dia.DefiRate),
		chanDefiState: make(chan *dia.DefiProtocolState),
	}

	log.Info("Defi scraper is built and triggered")
	go s.mainLoop(rateType)
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *DefiScraper) mainLoop(rateType string) {
	for {
		select {
		case <-s.tickerRate.C:
			s.UpdateRates(rateType)
		case <-s.tickerState.C:
			s.UpdateState(rateType)

		case <-s.shutdown: // user requested shutdown
			log.Println("DefiScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected Scrapers. Must only be called from mainLoop
func (s *DefiScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.tickerRate.Stop()
	s.tickerState.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (s *DefiScraper) Close() error {
	if s.closed {
		return errors.New("DefiScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns a channel that can be used to receive rate information
func (s *DefiScraper) RateChannel() chan *dia.DefiRate {
	return s.chanDefiRate
}

func (s *DefiScraper) StateChannel() chan *dia.DefiProtocolState {
	return s.chanDefiState
}

// UpdateRates calls the appropriate function corresponding to the rate type.
func (s *DefiScraper) UpdateRates(defiType string) error {
	var (
		helper   DeFIHelper
		protocol dia.DefiProtocol
	)

	switch defiType {
	case "DYDX":
		{

			protocol = dia.DefiProtocol{
				Name:                 "DYDX",
				Address:              "0x1e0447b19bb6ecfdae1e4ae1694b0c3659614e4e",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewDYDX(s, protocol)

		}
	case "AAVE":
		{

			protocol = dia.DefiProtocol{
				Name:                 "AAVE",
				Address:              "0x3dfd23A6c5E8BbcFc9581d2E864a68feb6a076d3",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewAAVE(s, protocol)
		}
	case "AAVEv2":
		{

			protocol = dia.DefiProtocol{
				Name:                 "AAVEv2",
				Address:              "0x52D306e36E3B6B02c153d0266ff0f85d18BCD413",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewAAVEv2(s, protocol)
		}
	case "DDEX":
		{

			protocol = dia.DefiProtocol{
				Name:                 "DDEX",
				Address:              "0x241e82C79452F51fbfc89Fac6d912e021dB1a3B7",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewDDEX(s, protocol)
		}
	// case "RAY":
	// 	{
	// 		protocol = dia.DefiProtocol{
	// 			Name:                 "RAY",
	// 			Address:              "0xE215e8160a5e0A03f2D6c7900b050F2f04eA5Cbb",
	// 			UnderlyingBlockchain: "Ethereum",
	// 			Token:                "",
	// 		}
	// 		helper = NewRAY(s, protocol)
	// 	}
	// case "DHARMA":
	// 	{

	// 		protocol = dia.DefiProtocol{
	// 			Name:                 "DHARMA",
	// 			Address:              "0x3f320a0B08B93D7562c1f2d008d8154c44147620",
	// 			UnderlyingBlockchain: "Ethereum",
	// 			Token:                "",
	// 		}
	// 		helper = NewDHARMA(s, protocol)
	// 	}
	case "COMPOUND":
		{

			protocol = dia.DefiProtocol{
				Name:                 "COMPOUND",
				Address:              "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewCompound(s, protocol)
		}
	case "CREAM":
		{

			protocol = dia.DefiProtocol{
				Name:                 "CREAM",
				Address:              "",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewCreamFinance(s, protocol)
		}
	case "BZX":
		{

			protocol = dia.DefiProtocol{
				Name:                 "BZX",
				Address:              "0x493c57c4763932315a328269e1adad09653b9081",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewBZX(s, protocol)
		}
	case "NUO":
		{

			protocol = dia.DefiProtocol{
				Name:                 "NUO",
				Address:              "0x64d14595152b430cf6940da15c6e39545c7c5b7e",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewNuo(s, protocol)
		}
	case "FORTUBE":
		{

			protocol = dia.DefiProtocol{
				Name:                 "FORTUBE",
				Address:              "0x936E6490eD786FD0e0f0C1b1e4E1540b9D41F9eF",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewForTube(s, protocol)
		}
	case "BITFINEX":
		{

			protocol = dia.DefiProtocol{
				Name:                 "BITFINEX",
				Address:              "",
				UnderlyingBlockchain: "Ethereum",
				Token:                "",
			}
			helper = NewBitfinex(s, protocol)
		}
	case "MAKERDAO":
		{

			protocol = dia.DefiProtocol{
				Name:                 "MAKERDAO",
				Address:              "",
				UnderlyingBlockchain: "Ethereum",
				Token:                "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
			}
			helper = NewMakerdao(s, protocol)
		}

	default:
		return errors.New("Error: " + defiType + " does not exist in database")

	}

	s.datastore.SetDefiProtocol(protocol)
	return helper.UpdateRate()
}

func (s *DefiScraper) UpdateState(defiType string) error {
	var helper DeFIHelper
	protocol, err := s.datastore.GetDefiProtocol(defiType)
	if err != nil {
		return err
	}

	switch defiType {
	case "DYDX":
		{
			helper = NewDYDX(s, protocol)
		}
	case "AAVE":
		{
			helper = NewAAVE(s, protocol)
		}
	case "AAVEv2":
		{
			helper = NewAAVEv2(s, protocol)
		}
	case "RAY":
		{
			helper = NewRAY(s, protocol)
		}
	case "DDEX":
		{
			helper = NewDDEX(s, protocol)
		}
	case "COMPOUND":
		{
			helper = NewCompound(s, protocol)
		}
	case "BZX":
		{
			helper = NewBZX(s, protocol)
		}
	case "NUO":
		{
			helper = NewNuo(s, protocol)
		}
	case "FORTUBE":
		{
			helper = NewForTube(s, protocol)
		}
	case "CREAM":
		{
			helper = NewCompound(s, protocol)
		}
	case "BITFINEX":
		{
			helper = NewBitfinex(s, protocol)
		}
	case "MAKERDAO":
		{
			helper = NewMakerdao(s, protocol)
		}
	default:
		return errors.New("Error: " + defiType + " does not exist in database")
	}
	return helper.UpdateState()
}
