package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// DownloadResource is a simple utility that downloads a resource
// from @url and stores it into @filepath.
func DownloadResource(filepath, url string) (err error) {

	fmt.Println("url: ", url)
	resp, err := http.Get(url) //nolint:noctx,gosec
	if err != nil {
		return
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return
}

// GetRequest performs a get request on @url and returns the response body
// as a slice of byte data.
func GetRequest(url string) ([]byte, int, error) {

	response, err := http.Get(url) //nolint:noctx,gosec
	if err != nil {
		log.Error("get request: ", err)
		return []byte{}, 0, err
	}

	// Close response body after function
	defer func() {
		cerr := response.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		return []byte{}, response.StatusCode, fmt.Errorf("HTTP Response Error %d", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	return XMLdata, response.StatusCode, err
}

// GetRequest performs a get request on @url and returns the response body
// as a slice of byte data.
func GetRequestWithStatus(url string) ([]byte, int, error) {

	// Get url
	response, err := http.Get(url)

	// Check, whether the request was successful
	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	// Close response body after function
	defer response.Body.Close()

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	return XMLdata, response.StatusCode, err
}

// PostRequest performs a POST request on @url and returns the response body
// as a slice of byte data.
func PostRequest(url string, body io.Reader) ([]byte, error) {

	// Get url
	response, err := http.Post(url, "", body) //nolint:noctx,gosec

	// Check, whether the request was successful
	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	// Close response body after function
	defer func() {
		cerr := response.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		log.Error("HTTP Response Error: ", response.StatusCode)
		return []byte{}, fmt.Errorf("HTTP Response Error %d", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	return XMLdata, err
}

// HTTPRequest returns the request body and defers the closing compliant
// to linting.
func HTTPRequest(request *http.Request) (body []byte, statusCode int, err error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer func() {
		cerr := resp.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	statusCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// CloseHTTPResp is a wrapper for closing http response bodies
// while complying with the linter.
func CloseHTTPResp(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Error(err)
	}
}

// GraphQLGet returns the body of the result of a graphQL GET query.
// @url is the base url of the graphQL API
// @query is a byte slice representing the graphQL query message
// @bearer contains the API key if present
func GraphQLGet(url string, query []byte, bearer string) ([]byte, int, error) {

	// Form post request with graphQL query
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query)) //nolint:noctx
	if err != nil {
		return []byte{}, 0, err
	}

	// Add authorization bearer to header
	req.Header.Add("Authorization", bearer)

	return HTTPRequest(req)
}

// GetCoinPrice Gets the price in USD of coin through our API.
// Looks it up on coingecko in case it doesn't find it there.
func GetCoinPrice(coin string) (float64, error) {
	// TODO Add UNIV2DAIETH quotation
	// log.Info("Get price for ", coin)
	switch coin {
	case "WETH":
		coin = "ETH"
	case "HBTC":
		coin = "BTC"
	}

	type Quotation struct {
		Symbol             string
		Name               string
		Price              float64
		PriceYesterday     *float64
		VolumeYesterdayUSD *float64
		Source             string
		Time               time.Time
		ITIN               string
	}
	/*
		type QuotationGecko struct {
			ID struct {
				Price string `json:"vs_currencies"`
			} `json:"ids"`
		}
	*/
	type QuotationCrptcmp struct {
		Price float64 `json:"USD"`
	}
	url := "https://api.diadata.org/v1/quotation/" + coin
	data, _, err := GetRequest(url)
	if err == nil {
		Quot := Quotation{}
		err = json.Unmarshal(data, &Quot)
		if err != nil {
			return 0, err
		}
		return Quot.Price, nil
	}
	// Try to get price from cryptocompare in case we don't have it in our API yet.
	data, _, err = GetRequest("https://min-api.cryptocompare.com/data/price?fsym=" + coin + "&tsyms=USD")
	if err != nil {
		log.Error("Could not get price")
		return 0, err
	}
	Quot := QuotationCrptcmp{}
	err = json.Unmarshal(data, &Quot)
	if err != nil {
		return 0, err
	}
	price := Quot.Price

	// // Retrieval through coingecko. Is down at the moment.
	// data, err = GetRequest("https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=USD")
	// if err != nil {
	// 	return 0, err
	// }
	// Quot := QuotationGecko{}
	// err = json.Unmarshal(data, &Quot)
	// price, err := strconv.ParseFloat(Quot.ID.Price, 64)
	// if err != nil {
	// 	return 0, err
	// }
	return price, nil

}
