package liquidityscrapers

import (
	"context"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/velarhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type VelarLiquidityScraper struct {
	logger       *logrus.Entry
	api          *velarhelper.VelarClient
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	exchangeName string
	relDB        *models.RelDB
	datastore    *models.DB
	handlerType  string
}

// NewVelarLiquidityScraper returns a new VelarLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	VELAR_DEBUG - (optional, bool), make stdout output with velar client http call, default = false
func NewVelarLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *VelarLiquidityScraper {
	envPrefix := strings.ToUpper(exchange.Name)
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	velarClient := velarhelper.NewVelarClient(
		log.WithContext(context.Background()).WithField("context", "VelarClient"),
		isDebug,
	)

	s := &VelarLiquidityScraper{
		poolChannel:  make(chan dia.Pool),
		doneChannel:  make(chan bool),
		exchangeName: exchange.Name,
		blockchain:   exchange.BlockChain.Name,
		api:          velarClient,
		relDB:        relDB,
		datastore:    datastore,
		handlerType:  "liquidity",
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("handlerType", s.handlerType).
		WithField("context", "VelarLiquidityScraper")

	go s.fetchPools()

	return s
}

func (s *VelarLiquidityScraper) fetchPools() {
	tickers, err := s.api.GetAllTickers()
	if err != nil {
		s.logger.WithError(err).Error("failed to fetch velar tickers")
		return
	}

	for _, t := range tickers {
		tokens := [...]string{t.BaseCurrency, t.TargetCurrency}
		dbAssets := make([]dia.Asset, 0, len(tokens))

		for _, address := range tokens {
			assset, err := s.relDB.GetAsset(address, s.blockchain)
			if err != nil {
				s.logger.WithError(err).Errorf("failed to GetAsset with key: %s", address)
				continue
			}
			dbAssets = append(dbAssets, assset)
		}

		if len(dbAssets) != len(tokens) {
			s.logger.Error("found less than 2 assets for the pool pair")
			continue
		}

		balances := [...]float64{t.BaseVolume, t.TargetVolume}
		assetVolumes := make([]dia.AssetVolume, len(balances))

		for i, a := range dbAssets {
			assetVolumes[i] = dia.AssetVolume{
				Index:  uint8(i),
				Asset:  a,
				Volume: balances[i],
			}
		}

		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      t.ID,
			Time:         time.Now(),
			Assetvolumes: assetVolumes,
		}

		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		s.logger.WithField("pool", pool).Info("sending pool to poolChannel")
		s.poolChannel <- pool
	}

	s.doneChannel <- true
}

func (s *VelarLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *VelarLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
