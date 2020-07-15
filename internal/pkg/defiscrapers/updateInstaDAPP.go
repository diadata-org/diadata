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

type InstaDAPPProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewInstaDAPP(scraper *DefiScraper, protocol dia.DefiProtocol) *InstaDAPPProtocol {
	return &InstaDAPPProtocol{scraper: scraper, protocol: protocol}
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

func getInstaDappAsset(address string) (instaDappMarket MakerToCompound, err error) {
	markets, err := fetchInstaDAPPMarkets()
	if err != nil {
		return
	}
	for _, market := range markets.Data.MakerToCompounds {
		if market.ID == address {
			instaDappMarket = market
		}
	}
	return
}

func (proto *InstaDAPPProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %v \n ", proto.protocol.Name)
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
			Protocol:      proto.protocol,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

	}
	log.Info("Update complete")
	return nil
}

func (proto *InstaDAPPProtocol) UpdateState() error {
	log.Print("Updating DEFI state for %+v\\n ", proto.protocol)
	usdcMarket, err := getCompoundAssetByAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if err != nil {
		return err
	}
	ethMarket, err := getCompoundAssetByAddress("0x0000000000000000000000000000000000000000")
	if err != nil {
		return err
	}
	totalUSDCSupplyPAR, err := strconv.ParseFloat(usdcMarket.TotalSupply, 64)
	if err != nil {
		return err
	}
	totalETHSupplyPAR, err := strconv.ParseFloat(ethMarket.TotalSupply, 64)
	if err != nil {
		return err
	}

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalUSDCSupplyPAR,
		TotalETH:  totalETHSupplyPAR,
		Protocol:  proto.protocol.Name,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
