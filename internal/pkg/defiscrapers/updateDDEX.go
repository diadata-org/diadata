package defiscrapers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type DDEXMarket struct {
	Status   int         `json:"status"`
	Desc     string      `json:"desc"`
	Template string      `json:"template"`
	Params   interface{} `json:"params"`
	Data     struct {
		LendingPoolStats []LendingPool `json:"lendingPoolStats"`
	} `json:"data"`
}

type LendingPool struct {
	ID                             int       `json:"id"`
	BlockNumber                    int       `json:"blockNumber"`
	DateTime                       time.Time `json:"dateTime"`
	AssetAddress                   string    `json:"assetAddress"`
	Symbol                         string    `json:"symbol"`
	BorrowRate                     string    `json:"borrowRate"`
	SupplyRate                     string    `json:"supplyRate"`
	TotalBorrowAmount              string    `json:"totalBorrowAmount"`
	TotalSupplyAmount              string    `json:"totalSupplyAmount"`
	UtilizationRate                string    `json:"utilizationRate"`
	Kind                           string    `json:"kind"`
	CreatedAt                      time.Time `json:"createdAt"`
	UpdatedAt                      time.Time `json:"updatedAt"`
	SevenDayAnnualizedRateOfReturn string    `json:"7dayAnnualizedRateOfReturn"`
	InsuranceAmount                string    `json:"InsuranceAmount"`
	LoansOriginatedAmount          string    `json:"loansOriginatedAmount"`
	SupplyAddedAmount              string    `json:"supplyAddedAmount"`
	OraclePrice                    string    `json:"oraclePrice"`
}

type DDEXProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewDDEX(scraper *DefiScraper, protocol dia.DefiProtocol) *DDEXProtocol {
	return &DDEXProtocol{scraper: scraper, protocol: protocol}
}

func fetchddexmarkets() (ddexMarket DDEXMarket, err error) {
	jsondata, err := utils.GetRequest("https://api.ddex.io/v4/lending_pool_stats")
	err = json.Unmarshal(jsondata, &ddexMarket)
	return
}

func (proto *DDEXProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v \n ", proto.protocol.Name)
	markets, err := fetchddexmarkets()
	if err != nil {
		return err
	}
	log.Info("Total Markets in DDEX ", len(markets.Data.LendingPoolStats))
	for _, market := range markets.Data.LendingPoolStats {
		totalSupplyAPR, err := strconv.ParseFloat(market.SupplyRate, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.BorrowRate, 64)
		if err != nil {
			return err
		}
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Symbol,
			Protocol:      proto.protocol.Name,
			LendingRate:   totalSupplyAPR * 100,
			BorrowingRate: totalBorrowAPR * 100,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

	}
	log.Info("Update complete")
	return nil
}

func getDDEXAssetByAddress(address string) (reserve LendingPool, err error) {
	markets, err := fetchddexmarkets()
	if err != nil {
		return
	}
	for _, market := range markets.Data.LendingPoolStats {
		if market.AssetAddress == address {
			reserve = market
		}
	}
	return
}

func (proto *DDEXProtocol) getTotalSupply() (sum float64, err error) {
	markets, err := fetchddexmarkets()
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(markets.Data.LendingPoolStats); i++ {
		coinPrice, err := utils.GetCoinPrice(markets.Data.LendingPoolStats[i].Symbol)
		if err != nil {
			return 0, err
		}
		totalSupplyAsset, err := strconv.ParseFloat(markets.Data.LendingPoolStats[i].TotalSupplyAmount, 64)
		if err != nil {
			return 0, err
		}
		sum += coinPrice * totalSupplyAsset
	}
	return
}

func (proto *DDEXProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	totalSupplyUSD, err := proto.getTotalSupply()
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
