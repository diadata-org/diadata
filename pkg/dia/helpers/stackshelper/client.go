package stackshelper

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	StacksURL                = "https://api.mainnet.hiro.so"
	DefaultSleepBetweenCalls = 1000 // ms
	MaxPageLimit             = 50
)

type StacksClient struct {
	debug             bool
	httpClient        *http.Client
	logger            *logrus.Entry
	sleepBetweenCalls time.Duration
	apiKey            string
}

func NewStacksClient(logger *logrus.Entry, sleepBetweenCalls time.Duration, hiroAPIKey string, isDebug bool) *StacksClient {
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

	c := &StacksClient{
		httpClient:        httpClient,
		debug:             isDebug,
		logger:            logger,
		sleepBetweenCalls: sleepBetweenCalls,
		apiKey:            hiroAPIKey,
	}

	if hiroAPIKey != "" {
		logger.Info("found hiro stacks API key, decreasing client request timeout")
		c.sleepBetweenCalls = 120 * time.Millisecond
	}

	return c
}

func (c *StacksClient) GetLatestBlock() (Block, error) {
	var block Block

	url := fmt.Sprintf("%s/extended/v2/blocks/latest", StacksURL)
	req, _ := http.NewRequest("GET", url, http.NoBody)

	err := c.callStacksAPI(req, &block)
	if err != nil {
		return block, err
	}
	return block, nil
}

func (c *StacksClient) GetTransactionAt(txID string) (Transaction, error) {
	var transaction Transaction

	url := fmt.Sprintf("%s/extended/v1/tx/%s", StacksURL, txID)
	req, _ := http.NewRequest("GET", url, http.NoBody)

	err := c.callStacksAPI(req, &transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (c *StacksClient) GetAllBlockTransactions(height int) ([]Transaction, error) {
	var (
		resp    GetBlockTransactionsResponse
		txs     = make([]Transaction, 0)
		total   = MaxPageLimit
		baseURL = fmt.Sprintf("%s/extended/v2/blocks/%d/transactions", StacksURL, height)
	)

	for offset := 0; offset < total; offset += MaxPageLimit {
		url := fmt.Sprintf("%s?limit=%d&offset=%d", baseURL, MaxPageLimit, offset)
		req, _ := http.NewRequest("GET", url, http.NoBody)

		err := c.callStacksAPI(req, &resp)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return txs, nil
			}
			return nil, err
		}

		total = resp.Total
		txs = append(txs, resp.Results...)
	}

	return txs, nil
}

func (c *StacksClient) GetAddressTransactions(address string, limit, offset int) (GetAddressTransactionsResponse, error) {
	var resp GetAddressTransactionsResponse

	url := fmt.Sprintf(
		"%s/extended/v2/addresses/%s/transactions?limit=%d&offset=%d",
		StacksURL,
		address,
		limit,
		offset,
	)

	req, _ := http.NewRequest("GET", url, http.NoBody)
	err := c.callStacksAPI(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *StacksClient) GetDataMapEntry(contractId, mapName, key string) ([]byte, error) {
	address := strings.Split(contractId, ".")

	url := fmt.Sprintf("%s/v2/map_entry/%s/%s/%s", StacksURL, address[0], address[1], mapName)
	body := []byte(fmt.Sprintf(`"%s"`, key))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	var entry DataMapEntry
	err := c.callStacksAPI(req, &entry)
	if err != nil {
		return nil, err
	}

	entryBytes, err := hex.DecodeString(entry.Data[2:])
	if err != nil {
		c.logger.WithError(err).Error("failed to decode data-map entry")
		return nil, err
	}

	result, ok := deserializeCVOption(entryBytes)
	if !ok {
		err = errors.New("data-map entry not found")
		return nil, err
	}
	return result, nil
}

func (c *StacksClient) callStacksAPI(request *http.Request, target interface{}) error {
	if len(c.apiKey) > 0 {
		request.Header.Add("X-API-Key", c.apiKey)
	}

	if c.debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			c.logger.WithError(err).Error("failed to dump request out")
			return err
		}
		c.logger.Infof("\n%s", string(dump))
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	if c.debug && resp != nil {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.logger.WithError(err).Error("failed to dump response")
			return err
		}
		c.logger.Infof("\n%s", string(dump))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.WithError(err).Error("failed to read response body")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to call Hiro API, status code: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusNotFound {
			c.logger.
				WithField("status", resp.StatusCode).
				WithField("body", string(data)).
				WithField("url", request.URL).
				Error(err.Error())
		}
		return err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return err
	}

	c.waiting()

	return resp.Body.Close()
}

func (c *StacksClient) waiting() {
	time.Sleep(c.sleepBetweenCalls)
}
