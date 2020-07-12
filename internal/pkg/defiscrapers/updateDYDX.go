package defiscrapers

import (
	"encoding/json"
	"fmt"
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
			fmt.Println(err)
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.TotalSupplyAPR, 64)
		if err != nil {
			fmt.Println(err)
		}
		asset := &dia.DefiRate{
  			Timestamp: time.Now(),
  			Asset: market.Symbol,
  			Protocol: protocol,
			LendingRate: totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
	}
		log.Printf("writing lending rate for  %#v in %v\n", asset, s.chanDefiRate)
		s.chanDefiRate <- asset
	}
	log.Info("Update complete")
	return nil



}
