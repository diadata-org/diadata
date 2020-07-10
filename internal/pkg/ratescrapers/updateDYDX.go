package ratescrapers

import (
	"encoding/json"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"fmt"
	"strconv"
	"time"
	log "github.com/sirupsen/logrus"

)

type DYDXRate struct {
	PBTCUSDC struct {
		Current struct {
			Market                     string    `json:"market"`
			EffectiveAt                time.Time `json:"effectiveAt"`
			FundingRate                string    `json:"fundingRate"`
			FundingRate8Hr             string    `json:"fundingRate8Hr"`
			AveragePremiumComponent    string    `json:"averagePremiumComponent"`
			AveragePremiumComponent8Hr string    `json:"averagePremiumComponent8Hr"`
		} `json:"current"`
	} `json:"PBTC-USDC"`
}

func fetchfundingrates()(dydxrate DYDXRate,err error){
 	JSONdata, _ := utils.GetRequest("https://api.dydx.exchange/v1/funding-rates")
	err = json.Unmarshal(JSONdata,&dydxrate)
	return
}

func (s *RateScraper) UpdateDYDX() error {
	log.Printf("DYDXScraper update")
	dydxrate, err := fetchfundingrates()
	if err !=nil{
		return err
	}

	// Convert interest rate from string to float64
	rate, err := strconv.ParseFloat(dydxrate.PBTCUSDC.Current.FundingRate, 64)
	if err != nil {
		fmt.Println(err)
	}
	t := &models.InterestRate{
		Symbol:          "DYDX",
		Value:           rate,
		PublicationTime: time.Now(),
		EffectiveDate:   dydxrate.PBTCUSDC.Current.EffectiveAt,
		Source:          "dydxAPI",
	}

	// Send new data through channel chanInterestRate
	log.Printf("writing interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t


	log.Info("Update complete")

	return err


	return nil

}
