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

type NuoMarket struct {
	Data struct {
		Users []User `json:"users"`
	} `json:"data"`
}

type User struct {
	ID                  string `json:"id"`
	LastUpdateTimestamp int    `json:"lastUpdateTimestamp"`
	LiquidityRate       string `json:"liquidityRate"`
	Name                string `json:"name"`
	Price               struct {
		ID string `json:"id"`
	} `json:"price"`
	TotalOrdersLiquidated string `json:"totalOrdersLiquidated"`
	TotalOrdersSettled    string `json:"totalOrdersSettled"`
}

func fetchNuoMarkets() (nuorate NuoMarket, err error) {
	jsonData := map[string]string{
		"query": `
		{
			makerToCompounds {
			  cdpNumber
			  owner
			  daiAmount
			  ethAmount
			  fees
			  transactionHash
			  timestamp
			}
		  }
		  
`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/sudeepb02/nuonetwork", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &nuorate)
	log.Println(nuorate)
	return
}

func (s *DefiScraper) UpdateNuo(protocol dia.DefiProtocol) error {
	log.Printf("NuoScraper update")

	markets, err := fetchNuoMarkets()
	if err != nil {
		return err
	}

	for _, market := range markets.Data.Users {

		totalSupplyAPR, err := strconv.ParseFloat(market.TotalOrdersLiquidated, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.TotalOrdersSettled, 64)
		if err != nil {
			return err
		}
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         "Nuo",
			Protocol:      protocol,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, s.chanDefiRate)
		s.chanDefiRate <- asset

	}

	log.Info("Update complete")
	return nil
}
