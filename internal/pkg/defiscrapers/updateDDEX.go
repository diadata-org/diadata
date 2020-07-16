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
			Protocol:      proto.protocol,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
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

func (proto *DDEXProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	usdcMarket, err := getDDEXAssetByAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if err != nil {
		return err
	}
	ethMarket, err := getDDEXAssetByAddress("0x000000000000000000000000000000000000000e")
	if err != nil {
		return err
	}
	totalSupplyUSDC, err := strconv.ParseFloat(usdcMarket.TotalSupplyAmount, 64)
	if err != nil {
		return err
	}
	totalBorrowUSDC, err := strconv.ParseFloat(usdcMarket.TotalBorrowAmount, 64)
	if err != nil {
		return err
	}
	totalBorrowETH, err := strconv.ParseFloat(ethMarket.TotalBorrowAmount, 64)
	if err != nil {
		return err
	}
	totalSupplyETH, err := strconv.ParseFloat(ethMarket.TotalSupplyAmount, 64)
	if err != nil {
		return err
	}

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSDC + totalBorrowUSDC,
		TotalETH:  totalBorrowETH + totalSupplyETH,
		Protocol:  proto.protocol.Name,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
