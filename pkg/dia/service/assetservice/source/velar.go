package source

import (
	"context"
	"strconv"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/velarhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// VelarAssetSource asset collector object - which serves assetCollector command
type VelarAssetSource struct {
	velarClient  *velarhelper.VelarClient
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	relDB        *models.RelDB
	logger       *logrus.Entry
}

// NewVelarAssetSource creates object to fetch velar assets
// ENV values:
//
//	VELAR_DEBUG - (optional, bool), make stdout output with velar client http call, default = false
func NewVelarAssetSource(exchange dia.Exchange, relDB *models.RelDB) *VelarAssetSource {
	envPrefix := strings.ToUpper(exchange.Name)
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	velarClient := velarhelper.NewVelarClient(
		log.WithContext(context.Background()).WithField("context", "VelarClient"),
		isDebug,
	)

	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &VelarAssetSource{
		velarClient:  velarClient,
		assetChannel: make(chan dia.Asset),
		doneChannel:  make(chan bool),
		blockchain:   exchange.BlockChain.Name,
		relDB:        relDB,
		logger:       logger,
	}

	go scraper.fetchAssets()

	return scraper
}

func (s *VelarAssetSource) fetchAssets() {
	s.logger.Info("scraping assets...")

	tokens, err := s.velarClient.GetAllTokens()
	if err != nil {
		s.logger.WithError(err).Error("error when scraping assets")
		return
	}

	for _, token := range tokens {
		decimals := len(strconv.Itoa(token.DecimalNum)) - 1

		s.assetChannel <- dia.Asset{
			Address:    token.ContractAddress,
			Name:       token.Name,
			Symbol:     token.Symbol,
			Decimals:   uint8(decimals),
			Blockchain: s.blockchain,
		}
	}

	s.doneChannel <- true
}

func (s *VelarAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *VelarAssetSource) Done() chan bool {
	return s.doneChannel
}
