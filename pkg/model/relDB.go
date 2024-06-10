package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"

	"github.com/go-redis/redis"
)

// RelDatastore is a (persistent) relational database with an additional redis caching layer
type RelDatastore interface {

	// --- Assets methods ---
	// --------- Persistent ---------
	SetAsset(asset dia.Asset) error
	GetAsset(address, blockchain string) (dia.Asset, error)
	GetAssetByID(ID string) (dia.Asset, error)
	GetAssetsBySymbolName(symbol, name string) ([]dia.Asset, error)
	GetAllAssets(blockchain string) ([]dia.Asset, error)
	GetFiatAssetBySymbol(symbol string) (asset dia.Asset, err error)
	GetAssetID(asset dia.Asset) (string, error)
	GetPage(pageNumber uint32) ([]dia.Asset, bool, error)
	Count() (uint32, error)
	SetAssetVolume24H(asset dia.Asset, volume float64, timestamp time.Time) error
	GetLastAssetVolume24H(asset dia.Asset) (float64, error)
	GetAssetsWithVOL(starttime time.Time, numAssets int64, skip int64, onlycex bool, substring string) ([]dia.AssetVolume, error)
	GetAssetSource(asset dia.Asset, onlycex bool) ([]string, error)
	GetAssetsWithVolByBlockchain(starttime time.Time, endtime time.Time, blockchain string) ([]dia.AssetVolume, error)

	// --------------- asset methods for exchanges ---------------
	SetExchangePair(exchange string, pair dia.ExchangePair, cache bool) error
	GetExchangePair(exchange string, foreignname string, caseSensitive bool) (exchangepair dia.ExchangePair, err error)
	GetExchangePairSeparator(exchange string) (string, error)
	GetPairsForExchange(exchange dia.Exchange, filterVerified bool, verified bool) ([]dia.ExchangePair, error)
	GetPairsForAsset(asset dia.Asset, filterVerified bool, verified bool) ([]dia.ExchangePair, error)
	GetExchangepairsByAsset(asset dia.Asset, exchange string, basetoken bool) ([]dia.ExchangePair, error)
	GetExchangePairSymbols(exchange string) ([]dia.ExchangePair, error)
	GetNumPairs(exchange dia.Exchange) (int, error)
	SetExchangeSymbol(exchange string, symbol string) error
	GetExchangeSymbol(exchange string, symbol string) (dia.Asset, error)
	GetExchangeSymbols(exchange string, substring string) ([]string, error)
	GetUnverifiedExchangeSymbols(exchange string) ([]string, error)
	VerifyExchangeSymbol(exchange string, symbol string, assetID string) (bool, error)
	GetExchangeSymbolAssetID(exchange string, symbol string) (string, bool, error)
	GetAllExchangeAssets(verified bool) ([]dia.Asset, error)

	// ----------------- Historical quotations methods -------------------
	SetHistoricalQuotation(quotation AssetQuotation) error
	GetHistoricalQuotations(asset dia.Asset, starttime time.Time, endtime time.Time) ([]AssetQuotation, error)
	GetLastHistoricalQuotationTimestamp(asset dia.Asset) (time.Time, error)

	// ----------------- exchange methods -------------------
	SetExchange(exchange dia.Exchange) error
	GetExchange(name string) (dia.Exchange, error)
	GetAllExchanges() ([]dia.Exchange, error)
	GetExchangeNames() ([]string, error)

	// ----------------- pool methods -------------------
	SetPool(pool dia.Pool) error
	SetScraperIndex(exchange string, scraperType string, indexType string, index int64) error
	GetScraperIndex(exchange string, scraperType string) (string, int64, error)
	GetAllDEXPoolsCount() (map[string]int, error)
	GetPoolByAddress(blockchain string, address string) (pool dia.Pool, err error)
	GetAllPoolAddrsExchange(exchange string, liquiThreshold float64) ([]string, error)
	GetAllPoolsExchange(exchange string, liquiThreshold float64) ([]dia.Pool, error)
	GetPoolsByAsset(asset dia.Asset, liquidityThreshold float64, liquidityThresholdUSD float64) ([]dia.Pool, error)

	// ----------------- blockchain methods -------------------
	SetBlockchain(blockchain dia.BlockChain) error
	GetBlockchain(name string) (dia.BlockChain, error)
	GetAllAssetsBlockchains() ([]string, error)
	GetAllBlockchains(fullAsset bool) ([]dia.BlockChain, error)

	// ------ Caching ------
	SetAssetCache(asset dia.Asset) error
	GetAssetCache(blockchain string, address string) (dia.Asset, error)
	SetExchangePairCache(exchange string, pair dia.ExchangePair) error
	GetExchangePairCache(exchange string, foreignName string) (dia.ExchangePair, error)
	CountCache() (uint32, error)

	// ---------------- NFT methods -------------------
	// NFT class methods
	SetNFTClass(nftClass dia.NFTClass) error
	GetAllNFTClasses(blockchain string) ([]dia.NFTClass, error)
	GetTradedNFTClasses(starttime time.Time) ([]dia.NFTClass, error)
	GetNFTClasses(limit, offset uint64) ([]dia.NFTClass, error)
	GetNFTClass(address string, blockchain string) (dia.NFTClass, error)
	GetNFTClassID(address string, blockchain string) (string, error)
	GetNFTClassByID(id string) (dia.NFTClass, error)
	GetNFTClassesByNameSymbol(searchstring string) ([]dia.NFTClass, error)
	UpdateNFTClassCategory(nftclassID string, category string) (bool, error)
	GetNFTCategories() ([]string, error)

	// NFT methods
	SetNFT(nft dia.NFT) error
	GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error)
	GetNFTID(address string, blockchain string, tokenID string) (string, error)

	// NFT trading and bidding methods
	SetNFTTrade(trade dia.NFTTrade) error
	SetNFTTradeToTable(trade dia.NFTTrade, table string) error
	GetNFTTrades(address string, blockchain string, tokenID string, starttime time.Time, endtime time.Time) ([]dia.NFTTrade, error)
	GetNFTTradesCollection(address string, blockchain string, starttime time.Time, endtime time.Time) ([]dia.NFTTrade, error)
	GetAllLastTrades(nftclass dia.NFTClass) ([]dia.NFTTrade, error)
	GetNFTOffers(address string, blockchain string, tokenID string) ([]dia.NFTOffer, error)
	GetNFTBids(address string, blockchain string, tokenID string) ([]dia.NFTBid, error)
	GetNFTFloor(nftclass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, noBundles bool, exchange string) (float64, error)
	GetNFTFloorLevel(nftclass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, currencies []dia.Asset, level float64, noBundles bool, exchange string) (float64, error)
	GetNFTFloorRecursive(nftClass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, stepBackLimit int, noBundles bool, exchange string) (float64, error)
	GetNFTFloorRange(nftClass dia.NFTClass, starttime time.Time, endtime time.Time, floorWindowSeconds time.Duration, stepBackLimit int, noBundles bool, exchange string) ([]float64, error)
	GetLastBlockheightTopshot(upperBound time.Time) (uint64, error)
	SetNFTBid(bid dia.NFTBid) error
	GetLastNFTBid(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (dia.NFTBid, error)
	GetLastBlockNFTBid(nftclass dia.NFTClass) (uint64, error)
	GetLastBlockNFTOffer(nftclass dia.NFTClass) (uint64, error)
	GetLastBlockNFTTrade(nftclass dia.NFTClass) (uint64, error)
	SetNFTOffer(offer dia.NFTOffer) error
	GetLastNFTOffer(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (offer dia.NFTOffer, err error)

	// NFT stats
	GetTopNFTsEth(numCollections int, offset int64, exchanges []string, starttime time.Time, endtime time.Time) ([]struct {
		Name       string
		Address    string
		Blockchain string
		Volume     float64
	}, error)
	GetNumNFTTrades(address string, blockchain string, exchange string, starttime time.Time, endtime time.Time) (int, error)
	GetNFTVolume(address string, blockchain string, exchange string, starttime time.Time, endtime time.Time) (float64, error)

	// General methods
	GetKeys(table string) ([]string, error)

	// Scraper config and state
	GetScraperState(ctx context.Context, scraperName string, state ScraperState) error
	SetScraperState(ctx context.Context, scraperName string, state ScraperState) error
	GetScraperConfig(ctx context.Context, scraperName string, config ScraperConfig) error
	SetScraperConfig(ctx context.Context, scraperName string, config ScraperConfig) error

	// Blockchain data
	SetBlockData(dia.BlockData) error
	GetBlockData(blockchain string, blocknumber int64) (dia.BlockData, error)
	GetLastBlockBlockscraper(blockchain string) (int64, error)

	//NFT exchange methods

	GetAllNFTExchanges() (exchanges []dia.NFTExchange, err error)
	GetNFTExchange(name string) (exchange dia.Exchange, err error)
	SetNFTExchange(exchange dia.NFTExchange) (err error)
	GetCollectionCountByExchange(exchange string) (int64, error)
	Get24HoursNFTExchangeVolume(exchange dia.NFTExchange) (float64, error)
	Get24HoursNFTExchangeTrades(exchange dia.NFTExchange) (int64, error)

	//Oracle builder
	SetKeyPair(publickey string, privatekey string) error
	GetKeyPairID(publickey string) string
	GetFeederAccessByID(id string) (owner string)
	GetFeederByID(id string) (owner string)
	SetOracleConfig(address, feederID, owner, feederAddress, symbols, feedSelection, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, mandatoryFrequency string) error
	SetFeederConfig(feederid, oracleconfigid string) error
	GetFeederID(address string) (feederId string)
	GetFeederLimit(owner string) (limit int)
	GetTotalFeeder(owner string) (total int)
	GetOracleConfig(address, chainid string) (oracleconfig dia.OracleConfig, err error)
	ChangeOracleState(feederID string, active bool) (err error)
	DeleteOracle(feederID string) (err error)
	GetOraclesByOwner(owner string) (oracleconfigs []dia.OracleConfig, err error)
	GetAllFeeders(bool, bool) (oracleconfigs []dia.OracleConfig, err error)
	GetFeederResources() (addresses []string, err error)
	GetOracleUpdates(address string, chainid string, offset int) ([]dia.OracleUpdate, error)
	GetOracleUpdateCount(address string, chainid string, symbol string) (int64, error)
	UpdateFeederAddressCheckSum(oracleaddress string) error
	GetExpiredFeeders() (oracleconfigs []dia.OracleConfig, err error)
	GetFeeder(feederID string) (oracleconfig dia.OracleConfig, err error)

	GetDayWiseUpdates(address string, chainid string) ([]dia.FeedUpdates, float64, float64, error)
	GetOracleLastUpdate(address, chainid, symbol string) (time.Time, string, error)
	GetOracleUpdatesByTimeRange(address, chainid, symbol string, offset int, startTime, endTime time.Time) ([]dia.OracleUpdate, error)

	// Asset List methods
	SetAssetList(asset dia.AssetList) error
	GetAssetListBySymbol(symbol string, listname string) ([]dia.AssetList, error)
	DeleteAssetList(sheetName string) error

	CreateCustomer(email string, customerPlan int, paymentStatus string, paymentSource string, numberOfDataFeeds int, walletPublicKeys []string) error
}

const (

	// postgres tables
	assetTable               = "asset"
	assetIdent               = "assetIdent"
	exchangepairTable        = "exchangepair"
	exchangesymbolTable      = "exchangesymbol"
	poolTable                = "pool"
	poolassetTable           = "poolasset"
	scraperCronjobStateTable = "scraper_cronjob_state"
	exchangeTable            = "exchange"
	nftExchangeTable         = "nftexchange"
	chainconfigTable         = "chainconfig"
	blockchainTable          = "blockchain"
	assetVolumeTable         = "assetvolume"
	historicalQuotationTable = "historicalquotation"

	// cache keys
	keyAssetCache        = "dia_asset_"
	keyExchangePairCache = "dia_exchangepair_"

	blockdataTable       = "blockdata"
	nftcategoryTable     = "nftcategory"
	nftclassTable        = "nftclass"
	nftTable             = "nft"
	NfttradeCurrTable    = "nfttradecurrent"
	NfttradeSumeriaTable = "nfttradesumeria"
	nftbidTable          = "nftbid"
	nftofferTable        = "nftoffer"
	scrapersTable        = "scrapers"
	keypairTable         = "keypair"
	oracleconfigTable    = "oracleconfig"
	feederconfigTable    = "feederconfig"
	feederaccessTable    = "feederaccess"
	feederResourceTable  = "feederresource"
	feederupdatesTable   = "feederupdates"

	// time format for blockchain genesis dates
	// timeFormatBlockchain = "2006-01-02"
)

// RelDB is a relative database with redis caching layer.
type RelDB struct {
	URI            string
	postgresClient *pgxpool.Pool
	redisClient    *redis.Client
	redisPipe      redis.Pipeliner
	pagesize       uint32
}

// NewRelDataStore returns a datastore with postgres client and redis cache.
func NewRelDataStore() (*RelDB, error) {
	log.Info("NewRelDataStore: Initialised")
	return NewRelDataStoreWithOptions(true, true)
}

// NewPostgresDataStore returns a datastore with postgres client and without redis caching layer.
func NewPostgresDataStore() (*RelDB, error) {
	return NewRelDataStoreWithOptions(true, false)
}

// NewCachingLayer returns a datastore with redis caching layer and without postgres client.
func NewCachingLayer() (*RelDB, error) {
	return NewRelDataStoreWithOptions(false, true)
}

// NewRelDataStoreWithOptions returns a postgres datastore and/or redis caching layer.
func NewRelDataStoreWithOptions(withPostgres bool, withRedis bool) (*RelDB, error) {
	var (
		postgresClient *pgxpool.Pool
		redisClient    *redis.Client
		redisPipe      redis.Pipeliner
		url            string
	)

	if withPostgres {
		url = db.GetPostgresURL()
		postgresClient = db.PostgresDatabase()
	}
	if withRedis {
		redisClient = db.GetRedisClient()
		redisPipe = redisClient.TxPipeline()
	}
	return &RelDB{url, postgresClient, redisClient, redisPipe, 32}, nil
}

// GetKeys returns a slice of strings holding the names of the keys of @table in postgres
func (rdb *RelDB) GetKeys(table string) (keys []string, err error) {
	query := fmt.Sprintf("SELECT column_name from information_schema.columns WHERE table_name='%s'", table)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			return keys, err
		}
		keys = append(keys, val[0].(string))
	}
	return
}
