package liquidityscrapers

import (
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
)

const stellarDecimals = 7
const stellarDefaultTradeCursor = ""

type StellarScraper struct {
	horizonClient *horizonclient.Client
	relDB         *models.RelDB
	datastore     *models.DB
	poolChannel   chan dia.Pool
	doneChannel   chan bool
	blockchain    string
	exchangeName  string
	cursor        *string
	cachedAssets  map[string]dia.Asset
}

func NewStellarScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *StellarScraper {
	var (
		horizonClient *horizonclient.Client
		poolChannel   = make(chan dia.Pool)
		doneChannel   = make(chan bool)
		scraper       *StellarScraper
	)

	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_CURSOR", "")
	if cursor == "" {
		cursor = stellarDefaultTradeCursor
	}

	horizonClient = horizonclient.DefaultPublicNetClient

	scraper = &StellarScraper{
		horizonClient: horizonClient,
		relDB:         relDB,
		datastore:     datastore,
		poolChannel:   poolChannel,
		doneChannel:   doneChannel,
		exchangeName:  exchange.Name,
		cursor:        &cursor,
	}

	go func() {
		scraper.fetchPools()
	}()

	return scraper
}
func (scraper *StellarScraper) fetchPools() {
	page, err := scraper.horizonClient.LiquidityPools(horizonclient.LiquidityPoolsRequest{
		Cursor: *scraper.cursor,
	})
	if err != nil {
		log.Error(err)
		return
	}

	recordsFound := len(page.Embedded.Records) > 0
	for recordsFound {
		for _, stellarPool := range page.Embedded.Records {
			log.Infof("pool: %s", stellarPool.ID)

			pool := scraper.getDIAPool(stellarPool)

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

func (scraper *StellarScraper) getDIAPool(stellarPool hProtocol.LiquidityPool) dia.Pool {
	assetvolumes := make([]dia.AssetVolume, 0)
	for i, asset := range stellarPool.Reserves {
		amount, _ := strconv.ParseFloat(asset.Amount, 64)
		s := strings.Split(asset.Asset, ":")
		assetVolume := dia.AssetVolume{
			Asset: dia.Asset{
				Address:    s[1],
				Symbol:     s[0],
				Decimals:   stellarDecimals,
				Name:       asset.Asset,
				Blockchain: scraper.blockchain,
			},
			Volume: amount,
			Index:  uint8(i),
		}
		assetvolumes = append(assetvolumes, assetVolume)
	}
	pool := dia.Pool{
		Exchange:     dia.Exchange{Name: scraper.exchangeName},
		Blockchain:   dia.BlockChain{Name: scraper.blockchain},
		Address:      stellarPool.ID,
		Assetvolumes: assetvolumes,
		Time:         time.Now(),
	}
	return pool
}

func (scraper *StellarScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *StellarScraper) Done() chan bool {
	return scraper.doneChannel
}
