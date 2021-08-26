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
	"time"

	log "github.com/sirupsen/logrus"
)

// DownloadResource is a simple utility that downloads a resource
// from @url and stores it into @filepath.
func DownloadResource(filepath, url string) error {

	log.Printf("Downloading data")

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// GetRequest performs a get request on @url and returns the response body
// as a slice of byte data.
func GetRequest(url string) ([]byte, error) {

	// Get url
	response, err := http.Get(url)

	// Check, whether the request was successful
	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	// Close response body after function
	defer response.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	return XMLdata, err
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
	response, err := http.Post(url, "", body)

	// Check, whether the request was successful
	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	// Close response body after function
	defer response.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		log.Error("HTTP Response Error: ", response.StatusCode)
		return []byte{}, fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	return XMLdata, err
}

// GraphQLGet returns the body of the result of a graphQL GET query.
// @url is the base url of the graphQL API
// @query is a byte slice representing the graphQL query message
// @bearer contains the API key if present
func GraphQLGet(url string, query []byte, bearer string) ([]byte, error) {

	// Form post request with graphQL query
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	if err != nil {
		return []byte{}, err
	}

	// Add authorization bearer to header
	req.Header.Add("Authorization", bearer)

	// Send request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

// OpenseaGetRequest returns the data for a get request on @url with an Opensea API key.
func OpenseaGetRequest(OpenseaURL string) ([]byte, int, error) {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/Opensea-API.key")
	} else {
		file, err = os.Open("../../secrets/Opensea-API.key")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
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
	apiKey := lines[0]
	client := &http.Client{}
	req, err := http.NewRequest("GET", OpenseaURL, nil)
	if err != nil {
		log.Print(err)
		return []byte{}, 0, err
	}

	q := url.Values{}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-API-KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	defer resp.Body.Close()
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
		Symbol             string
		Name               string
		Price              float64
		PriceYesterday     *float64
		VolumeYesterdayUSD *float64
		Source             string
		Time               time.Time
		ITIN               string
	}

	type QuotationGecko struct {
		ID struct {
			Price string `json:"vs_currencies"`
		} `json:"ids"`
	}

	type QuotationCrptcmp struct {
		Price float64 `json:"USD"`
	}
	url := "https://api.diadata.org/v1/quotation/" + coin
	data, err := GetRequest(url)
	if err == nil {
		Quot := Quotation{}
		err = json.Unmarshal(data, &Quot)
		if err != nil {
			return 0, err
		}
		return Quot.Price, nil
	}
	// Try to get price from cryptocompare in case we don't have it in our API yet.
	data, err = GetRequest("https://min-api.cryptocompare.com/data/price?fsym=" + coin + "&tsyms=USD")
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
