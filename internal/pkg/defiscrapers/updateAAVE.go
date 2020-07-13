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

type AAVEMarket struct {
	Data struct {
		Reserves []Reserve `json:"reserves"`
	} `json:"data"`
}

type Reserve struct {
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

func fetchAAVEMarkets() (aaverate AAVEMarket, err error) {
	jsonData := map[string]string{
		"query": `
          {
			reserves (where: {
				usageAsCollateralEnabled: true
			}) {
			id
			name
			price {
			  id
			}
			liquidityRate
			variableBorrowRate
			stableBorrowRate
			lastUpdateTimestamp
			totalLiquidity
		  }
		}
`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/aave/protocol-raw", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &aaverate)
	log.Println(aaverate)
	return
}

func (s *DefiScraper) UpdateAAVE(protocol dia.DefiProtocol) error {
	log.Printf("AAVEScraper update")

	markets, err := fetchAAVEMarkets()
	if err != nil {
		return err
	}

	for _, market := range markets.Data.Reserves {

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
			Asset:         market.Name,
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

func getAssetByAddress(address string) (reserve Reserve, err error) {
	markets, err := fetchAAVEMarkets()
	if err != nil {
		return
	}
	for _, market := range markets.Data.Reserves {
		if market.ID == address {
			reserve = market
		}
	}
	return
}

func (s *DefiScraper) UpdateAAVEState(protocolName string) error {
	log.Info("Updating DEFI state .. ")
	// Get Total USDC
	// Get Total ETH
	usdcMarket, err := getAssetByAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if err != nil {
		return err
	}
	ethMarket, err := getAssetByAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	if err != nil {
		return err
	}
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.LiquidityRate, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.LiquidityRate, 64)
	if err != nil {
		return err
	}
	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalUSDCSupplyPAR,
		TotalETH:  totalETHSupplyPAR,
		Protocol:  protocolName,
		Timestamp: time.Now(),
	}
	s.chanDefiState <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, s.chanDefiRate)

	log.Info("Update State complete")

	return nil

}
