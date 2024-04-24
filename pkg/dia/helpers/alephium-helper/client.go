package alephiumhelper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	BackendURL              = "https://backend.mainnet.alephium.org"
	NodeURL                 = "https://node.mainnet.alephium.org"
	AYINPairContractAddress = "vyrkJHG49TXss6pGAz2dVxq5o7mBXNNXAV18nAeqVT1R"
)

const (
	SymbolMethod = iota
	NameMethod
	DecimalsMethod
	TokenPairMethod = 7
)

type AlephiumClient struct {
	Debug      bool
	HTTPClient *http.Client
	logger     *logrus.Entry
}

func NewAlephiumClient(logger *logrus.Entry, debug bool) *AlephiumClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: 0,
		},
	}
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	result := &AlephiumClient{
		HTTPClient: httpClient,
		Debug:      debug,
		logger:     logger,
	}

	return result
}

func (c *AlephiumClient) callAPI(request *http.Request, target interface{}) error {
	if c.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return err
		}
		log.Printf("DumpRequestOut: \n%s\n", string(dump))
	}

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if c.Debug && resp != nil {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
		c.logger.Printf("\n%s\n", string(dump))
	}
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		err = errors.New("not 200 http response code from api")
		c.logger.
			WithError(err).
			WithField("body", string(data)).
			WithField("url", request.URL).
			Error("failed to call api")
		return err
	}

	err = json.Unmarshal(data, &target)

	return err
}

func (c *AlephiumClient) GetSwapPairsContractAddresses(defaultContractsLimit int) (SubContractResponse, error) {
	var contractResponse SubContractResponse

	url := fmt.Sprintf("%s/contracts/%s/sub-contracts?limit=%d", BackendURL, AYINPairContractAddress, defaultContractsLimit)
	request, _ := http.NewRequest("GET", url, nil)

	err := c.callAPI(request, &contractResponse)

	return contractResponse, err
}

func (c *AlephiumClient) GetTokenPairAddresses(contractAddress string) ([]string, error) {
	group, err := groupOfAddress(contractAddress)
	if err != nil {
		return nil, err
	}
	inputData := CallContractRequest{
		Group:       int(group),
		Address:     contractAddress,
		MethodIndex: TokenPairMethod,
	}
	logger := c.logger.WithField("function", "GetTokenPairAddresses")

	jsonData, err := json.Marshal(inputData)

	if err != nil {
		logger.Fatalf("failed to marshal input data: %v", err)
		return nil, err
	}
	url := fmt.Sprintf("%s/contracts/call-contract", NodeURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatalf("failed to create request: %v", err)
		return nil, err
	}
	var response CallContractResult
	err = c.callAPI(req, &response)

	if err != nil {
		logger.WithError(err).Error("failed to callApi")
		return nil, err
	}
	if response.Error != nil {
		err = errors.New(*response.Error)
		logger.
			WithError(err).
			WithField("jsonData", string(jsonData)).
			WithField("contractAddress", contractAddress).
			Error("failed to get token pair")
		return nil, err
	}
	logger.
		WithField("response", response).
		Info("GetTokenPairAddresses")

	address1, err := AddressFromTokenId(response.Returns[0].Value)
	if err != nil {
		logger.WithError(err).Error("failed to calculate address1")
		return nil, err
	}
	address2, err := AddressFromTokenId(response.Returns[1].Value)
	if err != nil {
		logger.WithError(err).Error("failed to calculate address2")
		return nil, err
	}

	output := []string{address1, address2}
	return output, nil
}

func (c *AlephiumClient) GetTokenInfoForContract(contractAddress string) (*OutputResult, error) {
	inputData := make([]CallContractRequest, 0)
	logger := c.logger.WithField("function", "GetTokenInfoForContract")

	for i := 0; i < 3; i++ {
		group, err := groupOfAddress(contractAddress)
		if err != nil {
			return nil, err
		}
		row := CallContractRequest{
			Group:       int(group),
			Address:     contractAddress,
			MethodIndex: i,
		}
		inputData = append(inputData, row)
	}

	calls := Calls{Calls: inputData}
	jsonData, err := json.Marshal(calls)

	if err != nil {
		logger.Fatalf("failed to marshal input data: %v", err)
		return nil, err
	}
	url := fmt.Sprintf("%s/contracts/multicall-contract", NodeURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		logger.Fatalf("failed to create request: %v", err)
		return nil, err
	}

	var response MulticallContractResponse
	err = c.callAPI(req, &response)

	if err != nil {
		logger.WithError(err).Error("failed to callApi")
		return nil, err
	}
	output := OutputResult{
		Address: contractAddress,
		Results: []OutputField{},
	}
	for _, row := range response.Results {
		if row.Error != nil {
			err = errors.New(*row.Error)
			logger.
				WithError(err).
				WithField("jsonData", string(jsonData)).
				WithField("contractAddress", contractAddress).
				Error("failed to get token info")
			return nil, err
		}
		result := OutputField{
			ResponseResult: row.Type,
			Field:          row.Returns[0],
		}
		output.Results = append(output.Results, result)
	}
	return &output, nil
}
