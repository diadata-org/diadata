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
	log.Println(aaverate)
	return
}

func (proto *AAVEProtocol) UpdateRate() error {
	log.Printf("Updating DEFI rate for %+v\n ", proto.protocol)

	markets, err := fetchAAVEMarkets()
	if err != nil {
		return err
	}


	// https://docs.aave.com/developers/integrating-aave/using-graphql#common-gotchas
	// as per the DOc value of rate need to be converted to power of 27
	for _, market := range markets.Data.Reserves {
		totalSupplyAPR := new(big.Float)
		totalSupplyAPR.SetString(market.LiquidityRate)
		totalSupplyAPR = new(big.Float).Quo(totalSupplyAPR, big.NewFloat(math.Pow10(27)))
		totalSupplyAPRPOW27,_ := totalSupplyAPR.Float64()

		totalBorrowAPR := new(big.Float)
		totalBorrowAPR.SetString(market.StableBorrowRate)
		totalBorrowAPR = new(big.Float).Quo(totalSupplyAPR, big.NewFloat(math.Pow10(27)))
		totalBorrowAPRPOW27,_ := totalSupplyAPR.Float64()


		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Name,
			Protocol:      proto.protocol,
			LendingRate:   totalSupplyAPRPOW27,
			BorrowingRate: totalBorrowAPRPOW27,
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
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.LiquidityRate, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.LiquidityRate, 64)
	if err != nil {
		return err
	}
	deFIState := &dia.DefiProtocolState{
		TotalUSD:  totalUSDCSupplyPAR,
		TotalETH:  totalETHSupplyPAR,
		Protocol:  proto.protocol.Name,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- deFIState
	log.Printf("writing DEFI state for  %#v in %v\n", deFIState, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
