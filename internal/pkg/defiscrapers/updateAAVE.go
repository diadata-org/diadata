package defiscrapers

import (
	"bytes"
	"encoding/json"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type AAVEMarket struct {
	Data struct {
		Reserves []Reserve `json:"reserves"`
	} `json:"data"`
}

type AAVEProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewAAVE(scraper *DefiScraper, protocol dia.DefiProtocol) *AAVEProtocol {
	return &AAVEProtocol{scraper: scraper, protocol: protocol}
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
	// log.Println(aaverate)
	return
}

func (proto *AAVEProtocol) UpdateRate() error {
	log.Printf("Updating DEFI rate for %+v\n ", proto.protocol)

	markets, err := fetchAAVEMarkets()
	if err != nil {
		return err
	}

	// https://docs.aave.com/developers/integrating-aave/using-graphql#common-gotchas
	// as per the DOC value of rate need to be converted to power of 27
	// As we compute rates in per cent and they're given as absolute here, factor is 10^25
	for _, market := range markets.Data.Reserves {
		totalSupplyAPR := new(big.Float)
		totalSupplyAPR.SetString(market.LiquidityRate)
		totalSupplyAPR = new(big.Float).Quo(totalSupplyAPR, big.NewFloat(math.Pow10(25)))
		totalSupplyAPRPOW25, _ := totalSupplyAPR.Float64()

		totalBorrowAPR := new(big.Float)
		totalBorrowAPR.SetString(market.StableBorrowRate)
		totalBorrowAPR = new(big.Float).Quo(totalBorrowAPR, big.NewFloat(math.Pow10(25)))
		totalBorrowAPRPOW25, _ := totalBorrowAPR.Float64()

		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Name,
			Protocol:      proto.protocol.Name,
			LendingRate:   totalSupplyAPRPOW25,
			BorrowingRate: totalBorrowAPRPOW25,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

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

func (proto *AAVEProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol.Name)
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
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.TotalLiquidity, 64)
	if err != nil {
		return err
	}
	totalUSDCSupplyPAR /= math.Pow(10, 18)
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.TotalLiquidity, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR /= math.Pow(10, 18)
	deFIState := &dia.DefiProtocolState{
		TotalUSD:  totalUSDCSupplyPAR,
		TotalETH:  totalETHSupplyPAR,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- deFIState
	log.Printf("writing DEFI state for  %#v in %v\n", deFIState, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
