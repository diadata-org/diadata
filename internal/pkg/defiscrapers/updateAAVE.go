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
	Symbol              string `json:"symbol"`
	AToken              struct {
		UnderlyingAssetDecimals int `json:"underlyingAssetDecimals"`
	}
	Price struct {
		ID         string `json:"id"`
		PriceInEth string `json:"priceInEth"`
	} `json:"price"`
	StableBorrowRate   string `json:"stableBorrowRate"`
	VariableBorrowRate string `json:"variableBorrowRate"`
	TotalLiquidity     string `json:"totalLiquidity"`
	Decimals           int    `json:decimals`
}

func fetchAAVEMarkets() (aaverate AAVEMarket, err error) {
	jsonData := map[string]string{
		"query": `
          {
			reserves {
			id
			symbol
			aToken {
				underlyingAssetDecimals
			}
			price {
			  id
			  priceInEth
			}
			liquidityRate
			variableBorrowRate
			stableBorrowRate
			lastUpdateTimestamp
			totalLiquidity
			decimals
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

// getTotalSupply returns the total supply in ETH locked in the protocol
func getTotalSupply() (float64, error) {
	markets, err := fetchAAVEMarkets()
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, market := range markets.Data.Reserves {
		marketLiquidity, err := strconv.ParseFloat(market.TotalLiquidity, 64)
		if err != nil {
			return 0, err
		}
		marketLiquidity /= math.Pow10(market.Decimals)
		marketPrice, err := strconv.ParseFloat(market.Price.PriceInEth, 64)
		if err != nil {
			return 0, err
		}
		marketPrice /= math.Pow10(18)

		marketInEth := marketLiquidity * marketPrice
		sum += marketInEth
	}
	return sum, nil
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
			Asset:         market.Symbol,
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
	totalSupplyETH, err := getTotalSupply()
	if err != nil {
		return err
	}

	PriceETH, err := utils.GetCoinPrice("ETH")
	totalSupplyUSD := totalSupplyETH * PriceETH

	// ethMarket, err := getAssetByAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	// if err != nil {
	// 	return err
	// }

	deFIState := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSD,
		TotalETH:  totalSupplyETH,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- deFIState
	log.Printf("writing DEFI state for  %#v in %v\n", deFIState, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
