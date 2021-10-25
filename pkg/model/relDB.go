package models

import (
	"context"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
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
	IdentifyAsset(asset dia.Asset) ([]dia.Asset, error)
	GetAssetID(asset dia.Asset) (string, error)
	GetPage(pageNumber uint32) ([]dia.Asset, bool, error)
	Count() (uint32, error)
	SetAssetVolume24H(asset dia.Asset, volume float64) error
	GetAssetVolume24H(asset dia.Asset) (float64, error)
	GetAssetsWithVOL(timeInit time.Time) ([]dia.Asset, error)

	// --------------- asset methods for exchanges ---------------
	SetExchangePair(exchange string, pair dia.ExchangePair, cache bool) error
	GetExchangePair(exchange string, foreignname string) (exchangepair dia.ExchangePair, err error)
	GetExchangePairSymbols(exchange string) ([]dia.ExchangePair, error)
	GetPairs(exchange string) ([]dia.ExchangePair, error)
	SetExchangeSymbol(exchange string, symbol string) error
	GetExchangeSymbols(exchange string) ([]string, error)
	GetUnverifiedExchangeSymbols(exchange string) ([]string, error)
	VerifyExchangeSymbol(exchange string, symbol string, assetID string) (bool, error)
	GetExchangeSymbolAssetID(exchange string, symbol string) (string, bool, error)

	// ----------------- blockchain methods -------------------
	GetBlockchain(name string) (dia.BlockChain, error)
	GetAllBlockchains() ([]string, error)

	// ------ Caching ------
	SetAssetCache(asset dia.Asset) error
	GetAssetCache(assetID string) (dia.Asset, error)
	SetExchangePairCache(exchange string, pair dia.ExchangePair) error
	GetExchangePairCache(exchange string, foreignName string) (dia.ExchangePair, error)
	CountCache() (uint32, error)

	// ---------------- NFT methods -------------------
	// NFT class methods
	SetNFTClass(nftClass dia.NFTClass) error
	GetAllNFTClasses(blockchain string) (nftClasses []dia.NFTClass, err error)
	GetNFTClasses(limit, offset uint64) (nftClasses []dia.NFTClass, err error)
	GetNFTClass(address string, blockchain string) (nftclass dia.NFTClass, err error)
	GetNFTClassID(address string, blockchain string) (ID string, err error)
	GetNFTClassByID(id string) (nftclass dia.NFTClass, err error)
	UpdateNFTClassCategory(nftclassID string, category string) (bool, error)
	GetNFTCategories() ([]string, error)

	// NFT methods
	SetNFT(nft dia.NFT) error
	GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error)
	GetNFTID(address string, blockchain string, tokenID string) (string, error)

	// NFT trading and bidding methods
	SetNFTTrade(trade dia.NFTTrade) error
	GetNFTTrades(nft dia.NFT) ([]dia.NFTTrade, error)
	GetNFTPrice30Days(nftclass dia.NFTClass) (float64, error)
	GetLastBlockheightTopshot(upperBound time.Time) (uint64, error)
	GetLastBlockNFTTradeScraper(nftclass dia.NFTClass) (uint64, error)
	SetNFTBid(bid dia.NFTBid) error
	GetLastNFTBid(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (dia.NFTBid, error)
	GetLastBlockNFTBid(nftclass dia.NFTClass) (uint64, error)
	SetNFTOffer(offer dia.NFTOffer) error
	GetLastNFTOffer(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (offer dia.NFTOffer, err error)

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
}

const (

	// postgres tables
	assetTable          = "asset"
	exchangepairTable   = "exchangepair"
	exchangesymbolTable = "exchangesymbol"
	blockchainTable     = "blockchain"
	assetVolumeTable    = "assetvolume"

	// cache keys
	keyAssetCache        = "dia_asset_"
	keyExchangePairCache = "dia_exchangepair_"

	blockdataTable   = "blockdata"
	nftcategoryTable = "nftcategory"
	nftclassTable    = "nftclass"
	nftTable         = "nft"
	nfttradeTable    = "nfttrade"
	nftbidTable      = "nftbid"
	nftofferTable    = "nftoffer"
	scrapersTable    = "scrapers"

	// time format for blockchain genesis dates
	// timeFormatBlockchain = "2006-01-02"
)

// RelDB is a relative database with redis caching layer.
type RelDB struct {
	URI            string
	postgresClient *pgx.Conn
	redisClient    *redis.Client
	pagesize       uint32
}

// NewRelDataStore returns a datastore with postgres client and redis cache.
func NewRelDataStore() (*RelDB, error) {
	log.Info("new rel datastore-------------------------")
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
	var postgresClient *pgx.Conn
	var redisClient *redis.Client

	url := db.GetPostgresURL()
	if withPostgres {
		postgresClient = db.PostgresDatabase()
	}
	if withRedis {
		redisClient = db.GetRedisClient()
	}
	return &RelDB{url, postgresClient, redisClient, 32}, nil
}

// GetKeys returns a slice of strings holding the names of the keys of @table in postgres
func (rdb *RelDB) GetKeys(table string) (keys []string, err error) {
	query := fmt.Sprintf("select column_name from information_schema.columns where table_name='%s'", table)
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
