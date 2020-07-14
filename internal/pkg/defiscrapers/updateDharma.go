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

type DHARMAMArket struct {
	Data struct {
		DDai []DDai `json:"dDai"`
	} `json:"data"`
}

type DDai struct {
	ID                  string `json:"id"`
	LastUpdateTimestamp int    `json:"lastUpdateTimestamp"`
	LiquidityRate       string `json:"liquidityRate"`
	Name                string `json:"name"`
	Price               struct {
		ID string `json:"id"`
	} `json:"price"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalSupply        string `json:"totalSupply"`
}

func fetchDHARMAMarkets() (dharmaRate DHARMAMArket, err error) {
	jsonData := map[string]string{
		"query": `
		{
			dDai: checkpoints(first: 1, orderBy: blockNumber, orderDirection: desc) {
			  exchangeRate
			  supplyRatePerBlock
			  estimatedAPR
			  estimatedAPY
			  totalSupply
			  totalSupplyUnderlying
			  version
			}
		  }
		  
`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/0age/dharma-dai-subgraph", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &dharmaRate)
	log.Println(dharmaRate)
	return
}

func (s *DefiScraper) UpdateDHARMA(protocol dia.DefiProtocol) error {
	log.Printf("DHARMAScraper update")

	markets, err := fetchDHARMAMarkets()
	if err != nil {
		return err
	}

	for _, market := range markets.Data.DDai {

		totalSupplyAPR, err := strconv.ParseFloat(market.TotalSupply, 64)
		if err != nil {
			return err
		}
		totalBorrowAPR, err := strconv.ParseFloat(market.StableBorrowRate, 64)
		if err != nil {
			return err
		}
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         "DHARMA",
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
