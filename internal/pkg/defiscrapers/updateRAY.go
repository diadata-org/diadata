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

type RAYProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewRAY(scraper *DefiScraper, protocol dia.DefiProtocol) *RAYProtocol {
	return &RAYProtocol{scraper: scraper, protocol: protocol}
}

type RAYPortfolio struct {
	Data struct {
		Portfolios []Portfolio `json:"portfolios"`
	} `json:"data"`
}

type Portfolio struct {
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

func fetchRAYPortfolios() (rayPortfolio RAYPortfolio, err error) {
	jsonData := map[string]string{
		"query": `
		{
			portfolios {
				id
			  asset {
				name
			  }
			  raytokens {
				value
			  }
			}
		  }
		  
		  
`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/protofire/ray", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &rayPortfolio)
	log.Println(rayPortfolio)
	return
}

func (proto *RAYProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v \n ", proto.protocol.Name)

	markets, err := fetchRAYPortfolios()
	if err != nil {
		return err
	}

	for _, market := range markets.Data.Portfolios {

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
			Protocol:      proto.protocol.Name,
			LendingRate:   totalSupplyAPR,
			BorrowingRate: totalBorrowAPR,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

	}

	log.Info("Update complete")
	return nil
}

func getPortfolioByID(id string) (rayPortfolio Portfolio, err error) {
	rayPortfolios, err := fetchRAYPortfolios()
	if err != nil {
		return
	}
	for _, portfolio := range rayPortfolios.Data.Portfolios {
		if portfolio.ID == id {
			rayPortfolio = portfolio
		}
	}
	return
}

func (proto *RAYProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v \n ", proto.protocol)
	// Get Total USDC
	// Get Total ETH
	usdcMarket, err := getPortfolioByID("0x1e868d302424cfebaf2b757c06fdd1a32411fd445ebb51ffc433cc15bacfe3e3")
	if err != nil {
		return err
	}
	ethMarket, err := getPortfolioByID("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
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
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
