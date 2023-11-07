package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

type ExternalDataSoure interface {
	Price(string) (float64, error)
}

type Coingecko struct {
	apiKey        string
	assetsMapping map[string]string // key blockchain+address value coingecko name
}

func NewCoinGeckoProvider(apiKey string, assetsMapping map[string]string) *Coingecko {
	return &Coingecko{apiKey: apiKey, assetsMapping: assetsMapping}
}

func (cg *Coingecko) Price(assetName string) (float64, error) {
	cgName := cg.assetsMapping[assetName]
	url := "https://pro-api.coingecko.com/api/v3/simple/price?ids=" + cgName + "&vs_currencies=usd&x_cg_pro_api_key=" + cg.apiKey
	response, err := http.Get(url)

	if err != nil {
		return 0.0, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on coingecko API call with return code %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}

	price := gjson.Get(string(contents), cgName+".usd").Float()

	return price, nil
}

// 5655df14-0213-429e-b375-51d338e8c79f

type CoinMarketCap struct {
	apiKey        string
	assetsMapping map[string]string // key blockchain+address value CoinMarketCap name
}

func NewCoinMarketCapProvider(apiKey string, assetsMapping map[string]string) *CoinMarketCap {
	return &CoinMarketCap{apiKey: apiKey, assetsMapping: assetsMapping}
}

func (cg *CoinMarketCap) Price(assetName string) (float64, error) {
	cmcName := cg.assetsMapping[assetName]
	urls := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest"
	req, err := http.NewRequestWithContext(context.Background(), "GET", urls, nil)

	if err != nil {
		return 0.0, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", cg.apiKey)

	q := url.Values{}
	q.Add("symbol", cmcName)

	req.URL.RawQuery = q.Encode()

	response, statusCode, err := HTTPRequest(req)

	if err != nil {
		return 0.0, fmt.Errorf("Error on coinmarketcap API call with return code %d", statusCode)

	}

	if 200 != statusCode {

		return 0.0, fmt.Errorf("Error on coinmarketcap API call with return code %d", statusCode)
	}

	price := gjson.Get(string(response), "data."+cmcName+".0.quote.USD.price").Float()

	return price, nil
}
