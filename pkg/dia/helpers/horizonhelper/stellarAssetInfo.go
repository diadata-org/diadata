package horizonhelper

import (
	"crypto/tls"
	"errors"
	"fmt"
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

type StellarAssetInfo struct {
	Logger *logrus.Entry
}

func (s *StellarAssetInfo) GetStellarAssetInfo(client *horizonclient.Client, assetCode, issuer, chainName string) (dia.Asset, error) {
	assetRequest := horizonclient.AssetRequest{
		ForAssetIssuer: issuer,
		ForAssetCode:   assetCode,
	}
	// for log messages only
	assetUrl, _ := assetRequest.BuildURL()

	s.Logger.
		WithField("assetUrl", assetUrl).
		WithField("assetCode", assetCode).
		WithField("issuer", issuer).
		Info("input assetCode")

	asset, err := client.Assets(assetRequest)
	if err != nil {
		s.Logger.
			WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
			}).Error("failed to fetch stellar assets from horizon")
		return dia.Asset{}, err
	}

	// s.Logger.Infof("asset.Embedded.Records %# v", pretty.Formatter(asset.Embedded.Records))
	if len(asset.Embedded.Records) == 0 {
		err = errors.New("no toml file")
		s.Logger.
			WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
			}).
			Error("failed to fetch stellar assets from horizon")
		return dia.Asset{}, err
	}
	tomlURL := asset.Embedded.Records[0].Links.Toml.Href
	if tomlURL == "" {
		err = errors.New("empty toml url from stellar")
		s.Logger.WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
			}).
			Error("failed to fetch stellar assets with empty toml url")
		return dia.Asset{}, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	tomlClient := &http.Client{
		Transport: tr,
		Timeout:   tomlReaderTimeout * time.Second,
	}

	resp, err := tomlClient.Get(tomlURL)
	if err != nil {
		s.Logger.
			WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
				"tomlURL":   tomlURL,
			}).
			Error("failed to fetch stellar tomlURL")
		return dia.Asset{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("not 200 http response code from toml")
		s.Logger.
			WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
				"tomlURL":   tomlURL,
			}).
			Error("failed to fetch stellar tomlURL:" + tomlURL)
		return dia.Asset{}, err
	}
	var stellarToml StellarToml
	if _, err := toml.DecodeReader(resp.Body, &stellarToml); err != nil {
		s.Logger.
			WithError(err).
			WithFields(logrus.Fields{
				"assetCode": assetCode,
				"issuer":    issuer,
				"tomlURL":   tomlURL,
			}).
			Error("failed to decode data from tomlURL")
		return dia.Asset{}, err
	}
	resultAsset := dia.Asset{
		Blockchain: chainName,
	}
	for _, currency := range stellarToml.CURRENCIES {
		if currency.Code == assetCode && currency.Issuer == issuer {
			resultAsset.Address = assetCode + "-" + issuer
			resultAsset.Name = currency.Name
			resultAsset.Symbol = assetCode
			resultAsset.Decimals = uint8(currency.DisplayDecimals)
			if resultAsset.Name == "" {
				resultAsset.Name = assetCode
			}
			s.Logger.
				WithFields(logrus.Fields{
					"assetCode": assetCode,
					"issuer":    issuer,
					"tomlURL":   tomlURL,
				}).
				Info("asset found")
			return resultAsset, nil
		}
	}

	err = errors.New("empty data from toml")
	s.Logger.
		WithError(err).
		WithFields(logrus.Fields{
			"assetCode": assetCode,
			"issuer":    issuer,
			"tomlURL":   tomlURL,
		}).
		Error("failed to find data for tomlURL")

	return dia.Asset{}, err
}

func (s *StellarAssetInfo) ConcatStrings(bs, cs string) string {
	result := fmt.Sprintf("%s-%s", bs, cs)
	return result
}

func (s *StellarAssetInfo) GetAddress(bs, cs string) string {
	return s.ConcatStrings(bs, cs)
}
