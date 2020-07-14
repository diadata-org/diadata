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

type InstaDAPPMarket struct {
	Data struct {
		MakerToCompounds []MakerToCompound `json:"makerToCompounds"`
	} `json:"data"`
}

type MakerToCompound struct {
	ID                  string `json:"id"`
	LastUpdateTimestamp int    `json:"lastUpdateTimestamp"`
	LiquidityRate       string `json:"liquidityRate"`
	Name                string `json:"name"`
	Price               struct {
		ID string `json:"id"`
	} `json:"price"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalLiquidity     string `json:"totalLiquidity"`
}

func fetchInstaDAPPMarkets() (instadapprate InstaDAPPMarket, err error) {
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
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/pierrickgt/instadapp", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &instadapprate)
	log.Println(instadapprate)
	return
}

func (s *DefiScraper) UpdateInstaDAPP(protocol dia.DefiProtocol) error {
	log.Printf("InstaDAPPScraper update")

	markets, err := fetchInstaDAPPMarkets()
	if err != nil {
		return err
	}

	for _, market := range markets.Data.MakerToCompounds {

		totalSupplyAPR, err := strconv.ParseFloat(market.LiquidityRate, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.StableBorrowRate, 64)
		if err != nil {
			return err
		}
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         "InstaDAPP",
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
