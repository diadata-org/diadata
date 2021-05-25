package optionscrapers

import (
	"encoding/json"
	Otoken "github.com/diadata-org/diadata/internal/pkg/option-scrapers/opyncontracts/OpynToken"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/option-scrapers/opyncontracts/OtokenFactory"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

const (
	wsDial   = "wss://eth-mainnet.ws.alchemyapi.io/v2/GwTfLo_j5DW_GjJSr_XiEwPiYaEmBJvX"
	restDial = "https://eth-mainnet.alchemyapi.io/v2/GwTfLo_j5DW_GjJSr_XiEwPiYaEmBJvX"
)

var OtokenFactoryAddress = common.HexToAddress("0x7C06792Af1632E77cb27a558Dc0885338F4Bdf8E")
var OtokenControllerAddress = common.HexToAddress("0x4ccc2339F87F6c59c6893E1A678c2266cA58dC72")

type OpynInstrumentsResponse struct {
	Bids struct {
		Total   int `json:"total"`
		Page    int `json:"page"`
		PerPage int `json:"perPage"`
		Records []struct {
			Order struct {
				Signature struct {
					SignatureType int    `json:"signatureType"`
					R             string `json:"r"`
					S             string `json:"s"`
					V             int    `json:"v"`
				} `json:"signature"`
				Sender              string `json:"sender"`
				Maker               string `json:"maker"`
				Taker               string `json:"taker"`
				TakerTokenFeeAmount string `json:"takerTokenFeeAmount"`
				MakerAmount         string `json:"makerAmount"`
				TakerAmount         string `json:"takerAmount"`
				MakerToken          string `json:"makerToken"`
				TakerToken          string `json:"takerToken"`
				Salt                string `json:"salt"`
				VerifyingContract   string `json:"verifyingContract"`
				FeeRecipient        string `json:"feeRecipient"`
				Expiry              string `json:"expiry"`
				ChainID             int    `json:"chainId"`
				Pool                string `json:"pool"`
			} `json:"order"`
			MetaData struct {
				OrderHash                    string    `json:"orderHash"`
				RemainingFillableTakerAmount string    `json:"remainingFillableTakerAmount"`
				CreatedAt                    time.Time `json:"createdAt"`
			} `json:"metaData"`
		} `json:"records"`
	} `json:"bids"`
	Asks struct {
		Total   int `json:"total"`
		Page    int `json:"page"`
		PerPage int `json:"perPage"`
		Records []struct {
			Order struct {
				Signature struct {
					SignatureType int    `json:"signatureType"`
					R             string `json:"r"`
					S             string `json:"s"`
					V             int    `json:"v"`
				} `json:"signature"`
				Sender              string `json:"sender"`
				Maker               string `json:"maker"`
				Taker               string `json:"taker"`
				TakerTokenFeeAmount string `json:"takerTokenFeeAmount"`
				MakerAmount         string `json:"makerAmount"`
				TakerAmount         string `json:"takerAmount"`
				MakerToken          string `json:"makerToken"`
				TakerToken          string `json:"takerToken"`
				Salt                string `json:"salt"`
				VerifyingContract   string `json:"verifyingContract"`
				FeeRecipient        string `json:"feeRecipient"`
				Expiry              string `json:"expiry"`
				ChainID             int    `json:"chainId"`
				Pool                string `json:"pool"`
			} `json:"order"`
			MetaData struct {
				OrderHash                    string    `json:"orderHash"`
				RemainingFillableTakerAmount string    `json:"remainingFillableTakerAmount"`
				CreatedAt                    time.Time `json:"createdAt"`
			} `json:"metaData"`
		} `json:"records"`
	} `json:"asks"`
}

type OpynOptionScraper struct {
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

func NewOpynETHOptionScraper() *OpynOptionScraper {

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

	s := &OpynOptionScraper{
		WsClient:      wsClient,
		RestClient:    restClient,
		chanOrderBook: make(chan *dia.OptionOrderbookDatum),
		DataStore:     ds,
	}
	s.GetAndStoreOptionsMeta()

	log.Errorln("Opyn scrapper created")

	return s

}



func (scraper *OpynOptionScraper) Scrape() {

	go func() {
		for _, market := range scraper.markets {
			log.Infoln("Token Address", market.id)

			var response OpynInstrumentsResponse

			b, err := utils.GetRequest("https://api.0x.org/sra/v4/orderbook?baseToken=" + market.id + "&quoteToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&perPage=100")
			if err != nil {
				log.Errorln("Error", err)
			}
			json.Unmarshal(b, &response)

			var (
				resolvedAskPX, resolvedAskSize, resolvedBidSize, resolvedBidPX float64
			)

			for _, askOut := range response.Asks.Records {
				log.Errorln("response.Asks.Records", len(response.Asks.Records))

				makeAmount, _ := strconv.ParseFloat(askOut.Order.MakerAmount, 64)
				takerAmount, _ := strconv.ParseFloat(askOut.Order.TakerAmount, 64)

				resolvedAskSize = makeAmount / 1e8
				resolvedAskPX = (takerAmount / 1e6) / (makeAmount / 1e8)

				var o = dia.OptionOrderbookDatum{
					InstrumentName:  market.id,
					ObservationTime: market.expiryDate,
					AskSize:         resolvedAskSize,
					AskPrice:        resolvedAskPX,
					BidPrice:        resolvedBidPX,
					BidSize:         resolvedBidSize,
				}
				log.Infoln("Got trade", o)
				scraper.chanOrderBook <- &o

			}
			for _, bidOut := range response.Bids.Records {
				makeAmount, _ := strconv.ParseFloat(bidOut.Order.MakerAmount, 64)
				takerAmount, _ := strconv.ParseFloat(bidOut.Order.TakerAmount, 64)

				resolvedBidSize = takerAmount / 1e8
				resolvedBidPX = (makeAmount / 1e6) / (takerAmount / 1e8)

				var o = dia.OptionOrderbookDatum{
					InstrumentName:  market.id,
					ObservationTime: market.expiryDate,
					BidPrice:        resolvedBidPX,
					BidSize:         resolvedBidSize,
				}
				log.Infoln("Got trade", o)
				scraper.chanOrderBook <- &o

			}

		}

	}()

}

func (scraper *OpynOptionScraper) getOPYNInstruments() (options []OptionAttrs) {

	optionFilterer, err := OtokenFactory.NewOtokenFactoryFilterer(OtokenFactoryAddress, scraper.WsClient)
	if err != nil {
		log.Fatal(err)
	}
	var block uint64

	block = uint64(11345363)
	tokesn, err := optionFilterer.FilterOtokenCreated(&bind.FilterOpts{Start: block}, []common.Address{}, []common.Address{}, []common.Address{})
	if err != nil {
		log.Error("error in get watching channel: ", err)
	}
	for tokesn.Next() {
		// Get Opyn Token Details
		// 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 get only ETH options
		if common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2") == tokesn.Event.Underlying {
			optionAttrs := scraper.getOtokenDetail(tokesn.Event.TokenAddress)
 			op := OptionAttrs{
				underLyingToken: tokesn.Event.Underlying.String(),
				id:              tokesn.Event.TokenAddress.String(),
				strikePrice:     optionAttrs.strikePrice,
				optionType:      optionAttrs.optionType,
				expiryDate:      optionAttrs.expiryDate,
 			}
			options = append(options, op)
 		}

	}
	return

}

func (scraper *OpynOptionScraper) FetchMarkets() {
	scraper.markets = scraper.getOPYNInstruments()
}

func (scraper *OpynOptionScraper) getOtokenDetail(otokenAddress common.Address) (optionAttrs OptionAttrs) {
	oTokenData, err := Otoken.NewOtoken(otokenAddress, scraper.WsClient)
	if err != nil {
		log.Errorln("Error getting oTokenData  ", err)
	}
	strikePrice, err := oTokenData.StrikePrice(&bind.CallOpts{})
	if err != nil {
		log.Errorln("Error getting strikePrice Detail ", err)
	}
	isPut, err := oTokenData.IsPut(&bind.CallOpts{})
	if err != nil {
		log.Errorln("Error getting isPut Detail ", err)
	}

	expiryTime, err := oTokenData.ExpiryTimestamp(&bind.CallOpts{})
	if err != nil {
		log.Errorln("Error getting expiryTime Detail ", err)
	}

	optionType := dia.CallOption

	if isPut {
		optionType = dia.PutOption
	}

	sp := float64(strikePrice.Int64())/ 1e8

	optionAttrs.strikePrice = sp
	optionAttrs.optionType = optionType
	optionAttrs.expiryDate = time.Unix(expiryTime.Int64(), 0)
	return

}

func (scraper *OpynOptionScraper) FetchInstruments() {
	scraper.FetchMarkets()
}

func (scraper *OpynOptionScraper) MetaOnOptionIsAvailable(option OptionAttrs) (available bool, err error) {
	available = false
	err = nil

	// TODO: can make this faster by specifying BaseCurrency/QuoteCurrency instead
	optionMetas, err := scraper.DataStore.GetOptionMeta(option.underLyingTokenName)
	if err != nil {
		return
	}

	for _, optionMeta := range optionMetas {
		if optionMeta.InstrumentName == option.id {
			return true, nil
		}
	}

	return
}

func (scraper *OpynOptionScraper) GetAndStoreOptionsMeta() (err error) {
	instruments := scraper.getOPYNInstruments()

	for _, instrument := range instruments {
		var available bool
		available, err = scraper.MetaOnOptionIsAvailable(instrument)

		if err != nil {
			return
		}

		if !available {
			optionMeta := dia.OptionMeta{
				InstrumentName: instrument.id,
				BaseCurrency:   "ETH",
				ExpirationTime: instrument.expiryDate,
				StrikePrice:     instrument.strikePrice,
				OptionType:     instrument.optionType,

			}

			scraper.DataStore.SetOptionMeta(&optionMeta)

		}
	}

	return nil
}

func (scraper *OpynOptionScraper) Channel() chan *dia.OptionOrderbookDatum {
	return scraper.chanOrderBook
}
