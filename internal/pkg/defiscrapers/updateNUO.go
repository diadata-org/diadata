package defiscrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type NuoResponse struct {
	Status string `json:"status"`
	Data   struct {
		Reserves []NuoRate `json:"reserves"`
	} `json:"data"`
}

type NuoRate struct {
	Currency struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		ShortName        string `json:"short_name"`
		ImageURL         string `json:"image_url"`
		BaseAddress      string `json:"base_address"`
		MaxTradeLeverage int    `json:"max_trade_leverage"`
		TradeBuffer      int    `json:"trade_buffer"`
		DecimalCount     int    `json:"decimal_count"`
	} `json:"currency"`
	TotalBalance               float64 `json:"total_balance"`
	TotalBalancePrimary        float64 `json:"total_balance_primary"`
	ActualBalance              float64 `json:"actual_balance"`
	ActualBalancePrimary       float64 `json:"actual_balance_primary"`
	ActiveReserveCount         string  `json:"active_reserve_count"`
	Last365DReturn             float64 `json:"last_365d_return"`
	Last30DReturnValue         float64 `json:"last_30d_return_value"`
	Last30DReturnValuePrimary  float64 `json:"last_30d_return_value_primary"`
	PotentialReturn            float64 `json:"potential_return"`
	PotentialReturnPrimary     float64 `json:"potential_return_primary"`
	LendRate                   float64 `json:"lend_rate"`
	ActiveLoanAmountSum        float64 `json:"active_loan_amount_sum"`
	ActiveLoanAmountSumPrimary float64 `json:"active_loan_amount_sum_primary"`
	BorrowRate                 float64 `json:"borrow_rate"`
}

type NuoProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
}

func NewNuo(scraper *DefiScraper, protocol dia.DefiProtocol) *NuoProtocol {
	return &NuoProtocol{scraper: scraper, protocol: protocol}
}

func (proto *NuoProtocol) fetch(asset string) (rate NuoRate, err error) {
	res, _ := proto.fetchALL()
	for _, v := range res.Data.Reserves {

		if v.Currency.ShortName == asset {
			return v, nil
		}
	}
	err = errors.New("asset not found")
	return

}

func (proto *NuoProtocol) fetchALL() (rate NuoResponse, err error) {
	//curl 'https://api.nuo.network/reserves/all?primary_currency_short_name=USD&markets=USDC' \
	//-H 'authority: api.nuo.network' \
	//-H 'accept: application/json, text/plain, */*' \
	//-H 'authorization: Bearer null' \
	//-H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36' \
	//-H 'origin: https://app.nuo.network' \
	//-H 'sec-fetch-site: same-site' \
	//-H 'sec-fetch-mode: cors' \
	//-H 'sec-fetch-dest: empty' \
	//-H 'referer: https://app.nuo.network/lend' \
	//-H 'accept-language: en-US,en;q=0.9,kn;q=0.8' \
	//-H 'if-none-match: W/"5473-u4bZ7v7BG290uuR5jiFHcU30CRw"' \
	//--compressed

	req, err := http.NewRequest("GET", "https://api.nuo.network/reserves/all?primary_currency_short_name=USD&markets=USDC", nil)
	if err != nil {
		log.Error("error getting req ", err)
		return

	}
	//req.Header.Set("Authority", "api.nuo.network")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", "Bearer null")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
	req.Header.Set("Origin", "https://app.nuo.network")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://app.nuo.network/lend")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,kn;q=0.8")
	req.Header.Set("If-None-Match", "W/\"5473-u4bZ7v7BG290uuR5jiFHcU30CRw\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Error("error getting req ", err)
		return

	}
	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return
	}
	err = json.Unmarshal(jsonData, &rate)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (proto *NuoProtocol) getTotalSupply() (float64, error) {
	markets, err := proto.fetchALL()
	if err != nil {
		return 0, err
	}

	sum := float64(0)
	for _, market := range markets.Data.Reserves {
		totalSupplyAsset := market.TotalBalance
		priceCoin, err := utils.GetCoinPrice(market.Currency.ShortName)
		if err != nil {
			return 0, err
		}
		marketPrice := totalSupplyAsset * priceCoin
		sum += marketPrice
		fmt.Printf("market %s holds %v worth in USD \n", market.Currency.ShortName, marketPrice)
	}
	return sum, nil
}

func (proto *NuoProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
	markets, err := proto.fetchALL()
	if err != nil {
		log.Error("error fetching rates %+v\n ", err)
		return err
	}

	for _, market := range markets.Data.Reserves {
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Currency.ShortName,
			Protocol:      proto.protocol.Name,
			LendingRate:   market.LendRate,
			BorrowingRate: market.BorrowRate,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset

	}
	log.Info("Update complete")
	return nil
}

func (proto *NuoProtocol) UpdateState() error {
	log.Print("Updating DEFI state for %+v\n ", proto.protocol)
	totalSupplyUSD, err := proto.getTotalSupply()
	if err != nil {
		return err
	}
	priceETH, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return err
	}
	totalSupplyETH := totalSupplyUSD / priceETH

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSD,
		TotalETH:  totalSupplyETH,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
