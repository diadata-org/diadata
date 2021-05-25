package optionscrapers

import (
	"github.com/diadata-org/diadata/internal/pkg/option-scrapers/premiacontracts/PremiaMarket"
	"github.com/diadata-org/diadata/internal/pkg/option-scrapers/premiacontracts/PremiaOption"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

var PremiaMarketAddress = common.HexToAddress("0x45eBD0FC72E2056adb5c864Ea6F151ad943d94af")
var PremiaOptionAddress = common.HexToAddress("0x5920cb60B1c62dC69467bf7c6EDFcFb3f98548c0")

type OptionAttrs struct {
	underLyingToken     string
	underLyingTokenName string
	id                  string
	strikePrice         float64
	expiryDate          time.Time
	optionType          dia.OptionType
}

type PremiaScraper struct {
	markets    []OptionAttrs
	WsClient   *ethclient.Client
	RestClient *ethclient.Client

	PollFrequency      int8
	ScraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
	optionsWaitGroup   *sync.WaitGroup
	DataStore          *models.DB
	chanOrderBook      chan *dia.OptionOrderbookDatum
	Ratelimiter        *rate.Limiter
}

func NewPremiaETHOptionScraper() *PremiaScraper {
	var wsClient, restClient *ethclient.Client
	var err error

	wsClient, err = ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}

	restClient, err = ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("error getting modelstore")
	}

	s := &PremiaScraper{
		WsClient:      wsClient,
		RestClient:    restClient,
		chanOrderBook: make(chan *dia.OptionOrderbookDatum),
		DataStore:     ds,
	}
	//s.GetAndStoreOptionsMeta()

	log.Errorln("Premia scrapper created")

	return s

}

func (scrapper *PremiaScraper) subscribe() (chan *PremiaMarket.PremiaMarketOrderCreated, error) {

	sink := make(chan *PremiaMarket.PremiaMarketOrderCreated)

	PremiaMarket.NewPremiaMarketFilterer(PremiaMarketAddress, scrapper.WsClient)

	optionFilterer, err := PremiaMarket.NewPremiaMarketFilterer(PremiaMarketAddress, scrapper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
 	var hash [][32]byte


	subscribed, err := optionFilterer.WatchOrderCreated(&bind.WatchOpts{}, sink, hash, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get watching channel: ", err)
	}

	log.Infoln("subscribed", subscribed)

	return sink, nil
}

func (scrapper *PremiaScraper) Scrape() {

	log.Infoln("Scrape")

	sink, _ := scrapper.subscribe()


	go func() {

		for {
			orderCreated, ok := <-sink
			if ok {

				var (
					resolvedAskPX, resolvedAskSize, resolvedBidSize, resolvedBidPX float64
				)

				if orderCreated.Side == 0 {

					resolvedAskSize = float64(orderCreated.PricePerUnit.Int64())
					resolvedAskPX = float64(orderCreated.Amount.Int64()) / math.Exp(float64(orderCreated.Decimals))

					var o = dia.OptionOrderbookDatum{
						InstrumentName: orderCreated.OptionContract.String(),
						ExpirationTime: time.Unix(orderCreated.ExpirationTime.Int64(), 0),
						AskSize:        resolvedAskSize,
						AskPrice:       resolvedAskPX,
						BidPrice:       resolvedBidPX,
						BidSize:        resolvedBidSize,
					}
					log.Println("Got trade", o)
					scrapper.chanOrderBook <- &o

				} else {
					resolvedBidSize = float64(orderCreated.PricePerUnit.Int64())
					resolvedBidPX = float64(orderCreated.Amount.Int64()) / math.Exp(float64(orderCreated.Decimals))

					var o = dia.OptionOrderbookDatum{
						InstrumentName: orderCreated.OptionContract.String(),
						ExpirationTime: time.Unix(orderCreated.ExpirationTime.Int64(), 0),
						BidPrice:       resolvedBidPX,
						BidSize:        resolvedBidSize,
					}
					log.Println("Got trade", o)
					scrapper.chanOrderBook <- &o

				}

			}

		}

	}()

}

func (scrapper *PremiaScraper) FetchMarkets() {
	log.Infoln("FetchInstruments")

	optionFilterer, err := PremiaOption.NewPremiaOptionFilterer(PremiaOptionAddress, scrapper.WsClient)
	if err != nil {
		log.Fatal(err)
	}

	premiaOption, err := PremiaOption.NewPremiaOption(PremiaOptionAddress, scrapper.RestClient)
	if err != nil {
		log.Fatal(err)
	}
	var block uint64

	block = uint64(11345363)
	tokesn, err := optionFilterer.FilterOptionIdCreated(&bind.FilterOpts{Start: block}, []*big.Int{}, []common.Address{})
	if err != nil {
		log.Error("error in get watching channel: ", err)
	}
	for tokesn.Next() {

		response, err := premiaOption.OptionData(&bind.CallOpts{}, tokesn.Event.OptionId)
		if err != nil {
			log.Error("error in get getting option data: ", err)
		}
		optionType := dia.CallOption

		if !response.IsCall {
			optionType = dia.PutOption
		}

		market := OptionAttrs{
			underLyingToken: tokesn.Event.Token.String(),
			id:              tokesn.Event.Token.String(),
			optionType:      optionType,
		}
		scrapper.markets = append(scrapper.markets, market)

		// check only for eth options
		//if common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2") == tokesn.Event.Token {
		//	scrapper.markets = append(scrapper.markets, tokesn.Event.TokenAddress.String())
		//
		//}

	}

}

func (scrapper *PremiaScraper) FetchInstruments() {
	scrapper.FetchMarkets()
}

func (s *PremiaScraper) MetaOnOptionIsAvailable(option BinanceInstrument) (available bool, err error) {
	available = false
	err = nil

	// TODO: can make this faster by specifying BaseCurrency/QuoteCurrency instead
	optionMetas, err := s.DataStore.GetOptionMeta(option.Underlying)
	if err != nil {
		return
	}

	for _, optionMeta := range optionMetas {
		if optionMeta.InstrumentName == option.Symbol {
			return true, nil
		}
	}

	return
}


func (scrapper *PremiaScraper) Channel() chan *dia.OptionOrderbookDatum {
	return scrapper.chanOrderBook
}
