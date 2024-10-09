package bitflowhelper

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/sirupsen/logrus"
)

const DeployerAddress = "SPQC38PW542EQJ5M11CR25P7BS1CA6QT4TBXGB3M"

var StableSwapContracts = [...]string{
	"stableswap-stx-ststx-v-1-2",
	"stableswap-usda-susdt-v-1-2",
	"stableswap-aeusdc-susdt-v-1-2",
	"stableswap-usda-aeusdc-v-1-2",
	"stableswap-usda-aeusdc-v-1-4",
	"stableswap-abtc-xbtc-v-1-2",
}

type BitflowClient struct {
	debug      bool
	httpClient *http.Client
	logger     *logrus.Entry
	apiHost    string
	apiKey     string
}

func NewBitflowClient(apiHost, apiKey string, logger *logrus.Entry, isDebug bool) *BitflowClient {
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

	return &BitflowClient{
		debug:      isDebug,
		httpClient: httpClient,
		logger:     logger,
		apiHost:    apiHost,
		apiKey:     apiKey,
	}
}

func (c *BitflowClient) GetAllTokens() ([]TokenMetadata, error) {
	url := fmt.Sprintf("%s/getAllTokensAndPools?key=%s", c.apiHost, c.apiKey)
	resp, err := c.httpClient.Get(url)

	if err != nil {
		c.logger.WithError(err).Error("failed to fetch tokens")
		return nil, err
	}
	defer resp.Body.Close()

	if c.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.logger.WithError(err).Error("failed to dump response")
			return nil, err
		}
		c.logger.Debugf("\n%s\n", string(dump))
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch tokens, status code: %d", resp.StatusCode)
		c.logger.Error(err.Error())
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.WithError(err).Error("failed to read response body")
		return nil, err
	}

	var result GetAllTokensResponse

	err = json.Unmarshal(data, &result)
	if err != nil {
		c.logger.Error("failed to decode tokens response body")
		return nil, err
	}

	return result.Tokens, nil
}
