package velarhelper

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

const (
	DeployerAddress  = "SP1Y5YSTAHZ88XYK1VPDH24GY0HPX5J4JECTMY4A1"
	VelarCoreAddress = DeployerAddress + ".univ2-core"
	VelarURL         = "https://api.velar.co"
)

type VelarClient struct {
	debug      bool
	httpClient *http.Client
	logger     *logrus.Entry
}

func NewVelarClient(logger *logrus.Entry, isDebug bool) *VelarClient {
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

	return &VelarClient{
		debug:      isDebug,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (c *VelarClient) GetAllTokens() ([]TokenMetadata, error) {
	req, err := http.NewRequest(http.MethodGet, VelarURL+"/tokens", http.NoBody)
	if err != nil {
		return nil, err
	}

	var tokens []TokenMetadata
	err = c.callVelarAPI(req, &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (c *VelarClient) GetAllTickers() ([]Ticker, error) {
	req, err := http.NewRequest(http.MethodGet, VelarURL+"/tickers", http.NoBody)
	if err != nil {
		return nil, err
	}

	var tickers []Ticker
	err = c.callVelarAPI(req, &tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}

func (c *VelarClient) callVelarAPI(request *http.Request, target interface{}) error {
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
	defer resp.Body.Close()

	if c.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.logger.WithError(err).Error("failed to dump response")
			return err
		}
		c.logger.Infof("\n%s\n", string(dump))
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to call Velar API, status code: %d", resp.StatusCode)
		c.logger.Error(err.Error())
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.WithError(err).Error("failed to read response body")
		return err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		c.logger.Error("failed to decode response body")
		return err
	}
	return nil
}
