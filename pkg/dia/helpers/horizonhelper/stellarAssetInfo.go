package horizonhelper

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
)

const tomlReaderTimeout = 5

// StellarToml represents the structure of a Stellar.toml file
type StellarToml struct {
	CURRENCIES []struct {
		Code            string `toml:"code"`
		Issuer          string `toml:"issuer"`
		Name            string `toml:"name"`
		DisplayDecimals int64  `toml:"display_decimals"`
	} `toml:"CURRENCIES"`
}

func GetStellarAssetInfo(client *horizonclient.Client, assetCode, assetIssuer, blockchain string) (dia.Asset, error) {
	assetRequest := horizonclient.AssetRequest{
		ForAssetIssuer: assetIssuer,
		ForAssetCode:   assetCode,
	}

	assetUrl, _ := assetRequest.BuildURL()
	log := logrus.New().
		WithField("context", "StellarTomlReader").
		WithField("assetCode", assetCode).
		WithField("assetIssuer", assetIssuer)

	log.WithField("assetUrl", assetUrl).Info("input assetCode")

	page, err := client.Assets(assetRequest)
	if err != nil {
		log.WithError(err).Error("failed to fetch asset info from horizon")
		return dia.Asset{}, err
	}

	if len(page.Embedded.Records) == 0 {
		err = errors.New("no toml file")
		log.WithError(err).Error("failed to fetch asset info from horizon")
		return dia.Asset{}, err
	}

	tomlURL := page.Embedded.Records[0].Links.Toml.Href
	if tomlURL == "" {
		err = errors.New("empty asset toml url from horizon")
		log.WithError(err).Error("failed to fetch asset info with empty toml url")
		return dia.Asset{}, err
	}

	log = log.WithField("tomlURL", tomlURL)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	tomlClient := &http.Client{
		Transport: tr,
		Timeout:   tomlReaderTimeout * time.Second,
	}

	resp, err := tomlClient.Get(tomlURL)
	if err != nil {
		log.WithError(err).Error("failed to fetch asset toml:" + tomlURL)
		return dia.Asset{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err = errors.New("not 200 http response code from toml")
		log.WithError(err).Error("failed to fetch asset toml:" + tomlURL)
		return dia.Asset{}, err
	}
	var stellarToml StellarToml

	if _, err := toml.NewDecoder(resp.Body).Decode(&stellarToml); err != nil {
		log.WithError(err).Error("failed to decode asset toml")
		return dia.Asset{}, err
	}

	asset := dia.Asset{
		Address:    GetAssetAddress(assetCode, assetIssuer),
		Blockchain: blockchain,
		Symbol:     assetCode,
	}

	for _, currency := range stellarToml.CURRENCIES {
		if currency.Code == assetCode && currency.Issuer == assetIssuer {
			asset.Name = currency.Name
			asset.Decimals = uint8(currency.DisplayDecimals)
			if asset.Name == "" {
				asset.Name = assetCode
			}
			log.Info("asset found")
			return asset, nil
		}
	}

	err = errors.New("asset info not found")
	log.WithError(err).Error("failed to find asset info in asset toml")

	return dia.Asset{}, err
}

func concatStrings(bs, cs string) string {
	result := fmt.Sprintf("%s-%s", bs, cs)
	return result
}

func GetAssetAddress(assetCode, assetIssuer string) string {
	return concatStrings(assetCode, assetIssuer)
}

func GetAssetSymbolPair(baseSymbol, quoteSymbol string) string {
	return concatStrings(baseSymbol, quoteSymbol)
}
