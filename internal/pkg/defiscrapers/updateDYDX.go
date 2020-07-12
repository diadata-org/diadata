package defiscrapers

import (
	"bytes"
	"encoding/json"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
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

func fetchmarkets()(dydxrate[] DYDXMarket,err error){
	var  response map[string][]DYDXMarket
	jsondata, err := utils.GetRequest("https://api.dydx.exchange/v1/markets")
	err = json.Unmarshal(jsondata,&response)
	return  response["markets"], err
}

func (s *DefiScraper) UpdateDYDX(protocol dia.DefiProtocol) error {
	log.Printf("DYDXScraper update")
	markets, err:= fetchmarkets()
	if err!=nil{
		return err
	}

	for _, market := range markets {
		totalSupplyAPR, err := strconv.ParseFloat(market.TotalSupplyAPR, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.TotalSupplyAPR, 64)
		if err != nil {
			return err
		}
		asset := &dia.DefiRate{
  			Timestamp: time.Now(),
  			Asset: market.Symbol,
  			Protocol: protocol,
			LendingRate: totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
	}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, s.chanDefiRate)
		s.chanDefiRate <- asset
	}
	log.Info("Update complete")
	return nil
}

func getMarketByID(marketID string)(dydxrate DYDXMarket,err error){
 		var  response map[string]DYDXMarket
		var url bytes.Buffer
		url.WriteString("https://api.dydx.exchange/v1/markets/")
		url.WriteString(marketID)
		jsonDATA, err := utils.GetRequest(url.String())
		err = json.Unmarshal(jsonDATA,&response)
		return  response["market"], err
	}



func (s *DefiScraper)UpdateDYDXState(protocolName string)error{
	log.Info("Updating DEFI state .. ")
	// Get Total USDC
	// Get Total ETH
	usdcMarket, err := getMarketByID("0")
	if err!=nil{
		return err
	}

	ethMarket, err := getMarketByID("0")
	if err!=nil{
		return err
	}
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.TotalSupplyPar, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.TotalSupplyPar, 64)
	if err != nil {
		return err
	}

	defistate :=  &dia.DefiProtocolState{
		TotalUSD: totalUSDCSupplyPAR,
		TotalETH: totalETHSupplyPAR,
		Protocol: protocolName,
	}
	s.chanDefiState <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, s.chanDefiRate)

	log.Info("Update State complete")

	return nil

}