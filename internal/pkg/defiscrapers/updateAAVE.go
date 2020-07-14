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

type AAVEProtocol struct {
	scrapper *DefiScraper
	protocol dia.DefiProtocol
}

func NewAAVE(scrapper *DefiScraper, protocol dia.DefiProtocol) *AAVEProtocol {
	return &AAVEProtocol{scrapper: scrapper, protocol: protocol}
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
			Protocol:      proto.protocol,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scrapper.RateChannel())
		proto.scrapper.RateChannel() <- asset

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
	log.Print("Updating DEFI state for %+v\\n ", proto.protocol.Name)
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
	proto.scrapper.StateChannel() <- deFIState
	log.Printf("writing DEFI state for  %#v in %v\n", deFIState, proto.scrapper.StateChannel())

	log.Info("Update State complete")

	return nil

}
