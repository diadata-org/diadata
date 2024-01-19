package liquidityscrapers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/horizonhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
)

const (
	defaultLPRequestCursor = ""
	defaultLPRequestLimit  = 200
)

type StellarScraper struct {
	logger        *logrus.Entry
	horizonClient *horizonclient.Client
	poolChannel   chan dia.Pool
	doneChannel   chan bool
	blockchain    string
	exchangeName  string
	relDB         *models.RelDB
	datastore     *models.DB
	cursor        *string
	limit         uint
}

func NewStellarScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *StellarScraper {
	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_LPS_CURSOR", defaultLPRequestCursor)
	limit := utils.GetenvUint(strings.ToUpper(exchange.Name)+"_LPS_LIMIT", defaultLPRequestLimit)

	var (
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *StellarScraper
		logger      = logrus.New().WithFields(logrus.Fields{
			"context": "StellarLiquidityScraper",
			"cursor":  cursor,
		})
	)

	horizonClient := horizonclient.DefaultPublicNetClient

	scraper = &StellarScraper{
		logger:        logger,
		horizonClient: horizonClient,
		poolChannel:   poolChannel,
		doneChannel:   doneChannel,
		exchangeName:  exchange.Name,
		blockchain:    exchange.BlockChain.Name,
		relDB:         relDB,
		datastore:     datastore,
		cursor:        &cursor,
		limit:         limit,
	}

	go func() {
		scraper.fetchPools()
	}()

	return scraper
}

func (scraper *StellarScraper) fetchPools() {
	page, err := scraper.horizonClient.LiquidityPools(horizonclient.LiquidityPoolsRequest{
		Cursor: *scraper.cursor,
		Limit:  scraper.limit,
	})
	if err != nil {
		log.Error(err)
		return
	}

	recordsFound := len(page.Embedded.Records) > 0
	for recordsFound {
		for _, stellarPool := range page.Embedded.Records {
			log.Infof("pool: %s", stellarPool.ID)

			pool, err := scraper.getDIAPool(stellarPool)
			if err != nil {
				log.Error(err)
				continue
			}

			// Determine USD liquidity.
			if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
				scraper.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
			}

			scraper.poolChannel <- pool
		}

		nextPage, err := scraper.horizonClient.NextLiquidityPoolsPage(page)
		if err != nil {
			log.Error(err)
			return
		}
		page = nextPage

		recordsFound = len(page.Embedded.Records) > 0
		log.Infof("found %d pools", len(page.Embedded.Records))
	}
	scraper.doneChannel <- true
}

func (scraper *StellarScraper) getDIAPool(stellarPool hProtocol.LiquidityPool) (dia.Pool, error) {
	scraper.logger.
		WithFields(logrus.Fields{
			"poolId": stellarPool.ID,
		}).Info("fetching pool info")

	assetvolumes := make([]dia.AssetVolume, len(stellarPool.Reserves))
	for i, stellarAsset := range stellarPool.Reserves {
		asset, err := scraper.getDIAAsset(stellarAsset.Asset)
		if err != nil {
			return dia.Pool{}, fmt.Errorf("error getting DIA asset for %s: %v", stellarAsset.Asset, err)
		}

		volume, err := strconv.ParseFloat(stellarAsset.Amount, 64)
		if err != nil {
			return dia.Pool{}, fmt.Errorf("error parsing volume: %v", err)
		}

		assetvolumes[i] = dia.AssetVolume{
			Asset:  asset,
			Volume: volume,
			Index:  uint8(i),
		}
	}
	pool := dia.Pool{
		Exchange:     dia.Exchange{Name: scraper.exchangeName},
		Blockchain:   dia.BlockChain{Name: scraper.blockchain},
		Address:      stellarPool.ID,
		Assetvolumes: assetvolumes,
		Time:         time.Now(),
	}
	return pool, nil
}

func (scraper *StellarScraper) getDIAAsset(stellarAsset string) (asset dia.Asset, err error) {
	scraper.logger.
		WithFields(logrus.Fields{
			"asset": stellarAsset,
		}).Info("fetching asset info")

	if stellarAsset == "native" {
		asset = horizonhelper.GetStellarNativeAssetInfo(scraper.blockchain)
		return
	}

	s := strings.SplitN(stellarAsset, ":", 2)
	if len(s) != 2 {
		err = fmt.Errorf("invalid asset format: %s", stellarAsset)
		return
	}

	assetCode, assetIssuer := s[0], s[1]
	assetID := horizonhelper.GetAssetAddress(assetCode, assetIssuer)
	asset, err = scraper.relDB.GetAsset(assetID, scraper.blockchain)
	if err == nil {
		return
	}
	err = nil

	asset, err = horizonhelper.GetStellarAssetInfo(scraper.horizonClient, assetCode, assetIssuer, scraper.blockchain)
	if err != nil {
		log.Error("cannot fetch asset info: ", err)
	}
	return
}

func (scraper *StellarScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *StellarScraper) Done() chan bool {
	return scraper.doneChannel
}
