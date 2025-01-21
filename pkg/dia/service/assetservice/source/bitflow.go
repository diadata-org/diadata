package source

import (
	"context"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/bitflowhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// BitflowAssetSource asset collector object - which serves assetCollector command
type BitflowAssetSource struct {
	bitflowClient *bitflowhelper.BitflowClient
	assetChannel  chan dia.Asset
	doneChannel   chan bool
	blockchain    string
	relDB         *models.RelDB
	logger        *logrus.Entry
}

// NewBitflowAssetSource creates object to fetch bitflow assets
// ENV values:
//
//	BITFLOW_API_HOST - (required, string), bitflow private REST API host, required to fetch assets
//	BITFLOW_API_KEY - (required, string), bitflow private API key, required to acess the REST API
//	BIFROST_DEBUG - (optional, bool), make stdout output with bitflow client http call, default = false
func NewBitflowAssetSource(exchange dia.Exchange, relDB *models.RelDB) *BitflowAssetSource {
	envPrefix := strings.ToUpper(exchange.Name)

	apiHost := utils.Getenv(envPrefix+"_API_HOST", "")
	apiKey := utils.Getenv(envPrefix+"_API_KEY", "")
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	bitflowClient := bitflowhelper.NewBitflowClient(
		apiHost,
		apiKey,
		log.WithContext(context.Background()).WithField("context", "BitflowClient"),
		isDebug,
	)

	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &BitflowAssetSource{
		bitflowClient: bitflowClient,
		assetChannel:  make(chan dia.Asset),
		doneChannel:   make(chan bool),
		blockchain:    exchange.BlockChain.Name,
		relDB:         relDB,
		logger:        logger,
	}

	go scraper.fetchAssets()

	return scraper
}

func (s *BitflowAssetSource) fetchAssets() {
	s.logger.Info("scraping assets...")

	tokens, err := s.bitflowClient.GetAllTokens()
	if err != nil {
		s.logger.WithError(err).Error("error when scraping assets")
		return
	}

	for _, token := range tokens {
		if token.Symbol == "STX" && token.TokenContract == "null" {
			token.TokenContract = "0x0000000000000000000000000000000000000000"
		}
		s.assetChannel <- dia.Asset{
			Address:    token.TokenContract,
			Name:       token.Name,
			Symbol:     token.Symbol,
			Decimals:   uint8(token.TokenDecimals),
			Blockchain: s.blockchain,
		}
		for _, wrappedToken := range token.WrapTokens {
			s.assetChannel <- dia.Asset{
				Address:    wrappedToken.TokenContract,
				Name:       wrappedToken.Name,
				Decimals:   uint8(wrappedToken.TokenDecimals),
				Blockchain: s.blockchain,
			}
		}
	}

	s.doneChannel <- true
}

func (s *BitflowAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *BitflowAssetSource) Done() chan bool {
	return s.doneChannel
}
