package bifrosthelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"time"
	"unicode/utf8"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	AssetAddressURI = "AssetRegistry:Assets"
	Blockchain      = "bifrost"
	GetAssetsPath   = "assets"
	ExchangeName    = "Bifrost"	
)

type BifrostClient struct {
	logger            *logrus.Entry
	sleepBetweenCalls time.Duration
	debug             bool
}

func NewBifrostClient(logger *logrus.Entry, sleepBetweenCalls time.Duration, isDebug bool) *BifrostClient {
	return &BifrostClient{
		logger:            logger,
		sleepBetweenCalls: sleepBetweenCalls,
		debug:             isDebug,
	}
}

func (c *BifrostClient) waiting() {
	time.Sleep(c.sleepBetweenCalls)
}

func (c *BifrostClient) GetAssetAllAssets() ([]BifrostAssetMetadata, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	
	apiURL := utils.Getenv(strings.ToUpper(ExchangeName)+"_API_URL", "http://localhost:3000/bifrost/v1")
	getAllAssetsURI := fmt.Sprintf("%s/%s", apiURL, GetAssetsPath)

	c.logger.Infof("Getting assets from: %s", getAllAssetsURI)

	var assets []BifrostAssetMetadata
	var err error

	go func() {
		defer wg.Done()

		response, err := http.Get(getAllAssetsURI)
		if err != nil {
			c.logger.WithError(err).Error("Failed to get assets")
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.logger.Errorf("Failed to get assets, status code: %d", response.StatusCode)
			err = fmt.Errorf("failed to get assets, status code: %d", response.StatusCode)
			return
		}

		if json.NewDecoder(response.Body).Decode(&assets) != nil {
			c.logger.Error("Failed to decode assets")
			err = fmt.Errorf("failed to decode assets")
		}
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (c *BifrostClient) GetAllPoolAssets() ([]BifrostPoolMetadata, error) {
	
	var wg sync.WaitGroup
	wg.Add(1)
	
	apiURL := utils.Getenv(strings.ToUpper(ExchangeName)+"_API_URL", "http://localhost:3000/bifrost/v1")
	getAllPoolAssetsURI := fmt.Sprintf("%s/%s", apiURL, "pools")

	c.logger.Infof("Getting pool assets from: %s", getAllPoolAssetsURI)

	var pools []BifrostPoolMetadata
	var err error

	go func() {
		defer wg.Done()

		response, err := http.Get(getAllPoolAssetsURI)
		if err != nil {
			c.logger.WithError(err).Error("Failed to get token pools")
			return
		}
		defer response.Body.Close()

		c.logger.Infof("Response: %d", response.Body)

		if response.StatusCode != http.StatusOK {
			c.logger.Errorf("Failed to get token pools, status code: %d", response.StatusCode)
			err = fmt.Errorf("failed to get token pools, status code: %d", response.StatusCode)
			return
		}

		if json.NewDecoder(response.Body).Decode(&pools) != nil {
			c.logger.Error("Failed to decode token pools")
			err = fmt.Errorf("failed to decode token pools")
		}
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return pools, nil
}

func (c *BifrostClient) ScrapAssets() ([]*dia.Asset, error) {
	bifrostAssets, err := c.GetAssetAllAssets()
	if err != nil {
		c.logger.WithError(err).Error("Failed to get assets")
		return nil, err
	}

	diaAssets := c.parseAssets(bifrostAssets)
	c.logger.Infof("Scraped (%d) assets.", len(diaAssets))

	return diaAssets, nil
}

func (c *BifrostClient) parseAsset(bifrostAsset BifrostAssetMetadata) *dia.Asset {
	decimals, err := strconv.ParseUint(bifrostAsset.Decimals, 10, 8)
	if err != nil {
		c.logger.WithError(err).Errorf("Failed to parse decimals: %s", bifrostAsset.Decimals)
		return nil
	}

	return &dia.Asset{
		Name:       bifrostAsset.Name,
		Symbol:     bifrostAsset.Symbol,
		Decimals:   uint8(decimals),
		Blockchain: Blockchain,
		Address:    "Bifrost:Asset:" + strings.ToLower(bifrostAsset.AssetKey),
	}
}

func (c *BifrostClient) parseAssets(bifrostAssets []BifrostAssetMetadata) []*dia.Asset {
	diaAssets := make([]*dia.Asset, 0, len(bifrostAssets))
	for _, asset := range bifrostAssets {
		diaAssets = append(diaAssets, c.parseAsset(asset))
	}

	return diaAssets
}

func (c *BifrostClient) sanitizeToUTF8(data []byte) string {
	var buf bytes.Buffer
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			data = data[size:]
			continue
		}
		if r == 0 {
			data = data[size:]
			continue
		}
		buf.WriteRune(r)
		data = data[size:]
	}
	return buf.String()
}
