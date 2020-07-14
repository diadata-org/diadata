package defiscrapers

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type CompoundMarket struct {
	Data struct {
		Markets []CompoundAsset `json:"markets"`
	} `json:"data"`
}

type CompoundAsset struct {
	BorrowRate               string `json:"borrowRate"`
	Cash                     string `json:"cash"`
	CollateralFactor         string `json:"collateralFactor"`
	ExchangeRate             string `json:"exchangeRate"`
	InterestRateModelAddress string `json:"interestRateModelAddress"`
	Name                     string `json:"name"`
	Reserves                 string `json:"reserves"`
	SupplyRate               string `json:"supplyRate"`
	Symbol                   string `json:"symbol"`
	ID                       string `json:"id"`
	TotalBorrows             string `json:"totalBorrows"`
	TotalSupply              string `json:"totalSupply"`
	UnderlyingAddress        string `json:"underlyingAddress"`
	UnderlyingName           string `json:"underlyingName"`
	UnderlyingPrice          string `json:"underlyingPrice"`
	UnderlyingSymbol         string `json:"underlyingSymbol"`
	ReserveFactor            string `json:"reserveFactor"`
	UnderlyingPriceUSD       string `json:"underlyingPriceUSD"`
}

type CompoundProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewCompound(scraper *DefiScraper, protocol dia.DefiProtocol) *CompoundProtocol {
	return &CompoundProtocol{scraper: scraper, protocol: protocol}
}

func fetchCompoundMarkets() (compoundrate CompoundMarket, err error) {
	jsonData := map[string]string{
		"query": `
          {
		  markets(first: 7) {
			borrowRate
			cash
			collateralFactor
			exchangeRate
			interestRateModelAddress
			name
			reserves
			supplyRate
			symbol
			id
			totalBorrows
			totalSupply
			underlyingAddress
			underlyingName
			underlyingPrice
			underlyingSymbol
			reserveFactor
			underlyingPriceUSD
		  }
}
`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/graphprotocol/compound-v2", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &compoundrate)
	log.Println(compoundrate)
	return
}

func getCompoundAssetByAddress(address string) (asset CompoundAsset, err error) {
	markets, err := fetchCompoundMarkets()
	if err != nil {
		return
	}
	for _, market := range markets.Data.Markets {
		if market.ID == address {
			asset = market
		}
	}
	return
}

func (proto *CompoundProtocol) UpdateRate() error {
	log.Print("Updating DEFI Rate for %+v\\n ", proto.protocol.Name)
	markets, err := fetchmarkets()
	if err != nil {
		return err
	}

	for _, market := range markets {
		totalSupplyAPR, err := strconv.ParseFloat(market.TotalSupplyAPR, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.TotalBorrowAPR, 64)
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

func (proto *CompoundProtocol) UpdateState() error {
	log.Print("Updating DEFI state for %+v\\n ", proto.protocol)
	usdcMarket, err := getCompoundAssetByAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if err != nil {
		return err
	}
	ethMarket, err := getCompoundAssetByAddress("0x0000000000000000000000000000000000000000")
	if err != nil {
		return err
	}
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.TotalSupply, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.TotalSupply, 64)
	if err != nil {
		return err
	}

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalUSDCSupplyPAR,
		TotalETH:  totalETHSupplyPAR,
		Protocol:  proto.protocol.Name,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
