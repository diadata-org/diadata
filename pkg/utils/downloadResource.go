package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

var openseaKey string

// DownloadResource is a simple utility that downloads a resource
// from @url and stores it into @filepath.
func DownloadResource(path, url string) (err error) {

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
	out, err := os.Create(filepath.Clean(path))
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
	response, err := http.Get(url) //nolint:noctx,gosec

	// Check, whether the request was successful
	if err != nil {
		log.Error(err)
		return []byte{}, response.StatusCode, err
	}

	// Close response body after function
	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			log.Error("error closing body ", errClose)
		}
	}(response.Body)

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

func getOpenseaApiKey() string {
	if openseaKey != "" {
		return openseaKey
	}
	if Getenv("USE_ENV", "false") == "true" {
		openseaKey = Getenv("API_KEY_OPENSEA", "")
		return openseaKey
	}

	var lines []string
	var filename string
	if Getenv("EXEC_MODE", "debug") == "production" {
		filename = "/run/secrets/Opensea-API.key"
	} else {
		filename = "../../secrets/Opensea-API.key"
	}

	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			log.Error("failure closing opensea-api.key file ", errClose)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file for opensea API key should have exactly one line")
	}
	openseaKey = lines[0]
	return openseaKey
}

// OpenseaGetRequest returns the data for a get request on @url with an Opensea API key.
func OpenseaGetRequest(OpenseaURL string) ([]byte, int, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", OpenseaURL, nil)
	if err != nil {
		log.Print(err)
		return []byte{}, 0, err
	}
	apiKey := getOpenseaApiKey()
	q := url.Values{}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-API-KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			fmt.Println("Error closing body ", errClose)
		}
	}(resp.Body)
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read response body: ", err)
		return respData, resp.StatusCode, err
	}
	return respData, resp.StatusCode, nil
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
		Symbol             string    `json:"Symbol"`
		Name               string    `json:"Name"`
		Price              float64   `json:"Price"`
		PriceYesterday     *float64  `json:"PriceYesterday"`
		VolumeYesterdayUSD *float64  `json:"VolumeYesterdayUSD"`
		Source             string    `json:"Source"`
		Time               time.Time `json:"Time"`
		ITIN               string    `json:"ITIN"`
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
