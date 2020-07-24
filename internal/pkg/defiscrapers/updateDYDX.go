package defiscrapers

import (
	"bytes"
	"encoding/json"
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type DYDXMarket struct {
	ID                        int         `json:"id"`
	Name                      string      `json:"name"`
	Symbol                    string      `json:"symbol"`
	SupplyIndex               string      `json:"supplyIndex"`
	BorrowIndex               string      `json:"borrowIndex"`
	SupplyInterestRateSeconds string      `json:"supplyInterestRateSeconds"`
	BorrowInterestRateSeconds string      `json:"borrowInterestRateSeconds"`
	TotalSupplyPar            string      `json:"totalSupplyPar"`
	TotalBorrowPar            string      `json:"totalBorrowPar"`
	LastIndexUpdateSeconds    string      `json:"lastIndexUpdateSeconds"`
	OraclePrice               string      `json:"oraclePrice"`
	CollateralRatio           string      `json:"collateralRatio"`
	MarginPremium             string      `json:"marginPremium"`
	SpreadPremium             string      `json:"spreadPremium"`
	CurrencyUUID              string      `json:"currencyUuid"`
	CreatedAt                 time.Time   `json:"createdAt"`
	UpdatedAt                 time.Time   `json:"updatedAt"`
	DeletedAt                 interface{} `json:"deletedAt"`
	Currency                  struct {
		UUID            string    `json:"uuid"`
		Symbol          string    `json:"symbol"`
		ContractAddress string    `json:"contractAddress"`
		Decimals        int       `json:"decimals"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
	} `json:"currency"`
	TotalSupplyAPR string `json:"totalSupplyAPR"`
	TotalBorrowAPR string `json:"totalBorrowAPR"`
	TotalSupplyAPY string `json:"totalSupplyAPY"`
	TotalBorrowAPY string `json:"totalBorrowAPY"`
	TotalSupplyWei string `json:"totalSupplyWei"`
	TotalBorrowWei string `json:"totalBorrowWei"`
}

type DYDXProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewDYDX(scraper *DefiScraper, protocol dia.DefiProtocol) *DYDXProtocol {
	return &DYDXProtocol{scraper: scraper, protocol: protocol}
}

func fetchDYDXMarkets() (dydxrate []DYDXMarket, err error) {
	var response map[string][]DYDXMarket
	jsondata, err := utils.GetRequest("https://api.dydx.exchange/v1/markets")
	err = json.Unmarshal(jsondata, &response)
	return response["markets"], err
}

// getTotalSupply returns the total supply in ETH locked in the protocol
func getDYDXTotalSupplyUSD() (float64, error) {
	markets, err := fetchDYDXMarkets()
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, market := range markets {
		marketSupply, err := strconv.ParseFloat(market.TotalSupplyWei, 64)
		if err != nil {
			return 0, err
		}
		marketBorrow, err := strconv.ParseFloat(market.TotalBorrowWei, 64)
		if err != nil {
			return 0, err
		}
		marketLiquidity := (marketSupply - marketBorrow) / math.Pow10(market.Currency.Decimals)
		// priceCoin, err := utils.GetCoinPrice(market.Symbol)
		// if err != nil {
		// 	return 0, err
		// }
		priceCoin := float64(0)
		if market.Symbol == "SAI" {
			// hotfix because coingecko is down. Change asap
			priceCoin = float64(1.42)
		} else {
			priceCoin, err = utils.GetCoinPrice(market.Symbol)
			if err != nil {
				return 0, err
			}
		}

		marketPrice := marketLiquidity * priceCoin
		sum += marketPrice
	}
	return sum, nil
}

func (proto *DYDXProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
	markets, err := fetchDYDXMarkets()
	if err != nil {
		return err
	}

	for _, market := range markets {
		totalSupplyAPR, err := strconv.ParseFloat(market.TotalSupplyAPR, 64)
		if err != nil {
			return err
		}
		// Return value per cent
		totalSupplyAPR *= 100
		totalBorrowAPR, err := strconv.ParseFloat(market.TotalBorrowAPR, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR *= 100
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Symbol,
			Protocol:      proto.protocol.Name,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

	}
	log.Info("Update complete")
	return nil
}

func getMarketByID(marketID string) (dydxrate DYDXMarket, err error) {
	var response map[string]DYDXMarket
	var url bytes.Buffer
	url.WriteString("https://api.dydx.exchange/v1/markets/")
	url.WriteString(marketID)
	jsonDATA, err := utils.GetRequest(url.String())
	err = json.Unmarshal(jsonDATA, &response)
	return response["market"], err
}

func (proto *DYDXProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	totalSupplyUSD, err := getDYDXTotalSupplyUSD()
	if err != nil {
		return err
	}
	priceETH, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return err
	}
	totalSupplyETH := totalSupplyUSD / priceETH
	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSD,
		TotalETH:  totalSupplyETH,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
