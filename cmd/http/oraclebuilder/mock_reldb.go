package main

import (
	"context"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/stretchr/testify/mock"
)

// RelDatastoreMock is a mock type for the RelDatastore interface
type RelDatastoreMock struct {
	mock.Mock
}

func (m *RelDatastoreMock) InsertLoopPaymentTransferProcessed(record models.LoopPaymentTransferProcessed) error {
	args := m.Called(record)
	return args.Error(0)
}
func (m *RelDatastoreMock) GetLastPaymentByEndUser(endUser string) (models.LoopPaymentTransferProcessed, error) {
	args := m.Called(endUser)
	return args.Get(0).(models.LoopPaymentTransferProcessed), args.Error(1)
}

func (m *RelDatastoreMock) UpdateFeederAddressCheckSum(oracleaddress string) error {
	args := m.Called(oracleaddress)
	return args.Error(0)
}

func (m *RelDatastoreMock) UpdateAccessLevel(accessLevel, publicKey string) error {
	args := m.Called(accessLevel, publicKey)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetScraperState(ctx context.Context, scraperName string, state models.ScraperState) error {
	args := m.Called(ctx, scraperName, state)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetScraperConfig(ctx context.Context, scraperName string, config models.ScraperConfig) error {
	args := m.Called(ctx, scraperName, config)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetOracleConfig(oracleaddress, feederID, creator, address, symbols, feedSelection, chainID, frequency, sleepSeconds, deviationPermille, blockchainnode, mandatoryFrequency string) error {
	args := m.Called(oracleaddress, feederID, creator, address, symbols, feedSelection, chainID, frequency, sleepSeconds, deviationPermille, blockchainnode, mandatoryFrequency)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetNFTOffer(offer dia.NFTOffer) error {
	args := m.Called(offer)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetNFTExchange(exchange dia.NFTExchange) (err error) {
	args := m.Called(exchange)
	return args.Error(0)
}
func (m *RelDatastoreMock) SetNFTBid(bid dia.NFTBid) error {
	args := m.Called(bid)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetFeederConfig(feederid, oracleconfigid string) error {
	args := m.Called(feederid, oracleconfigid)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetKeyPair(publickey string, privatekey string) error {
	args := m.Called(publickey, privatekey)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetBlockData(blockData dia.BlockData) error {
	args := m.Called(blockData)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetAssetList(asset dia.AssetList) error {
	args := m.Called(asset)
	return args.Error(0)
}

func (m *RelDatastoreMock) RemoveWalletKeys(publicKey []string) error {
	args := m.Called(publicKey)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetTotalGasSpend(address string, chainid string) (float64, error) {
	args := m.Called(address, chainid)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetTotalOracles(owner string) (total int) {
	args := m.Called(owner)
	return args.Int(0)
}

func (m *RelDatastoreMock) GetTopNFTsEth(numCollections int, offset int64, exchanges []string, starttime time.Time, endtime time.Time) ([]struct {
	Name       string
	Address    string
	Blockchain string
	Volume     float64
}, error) {
	args := m.Called(numCollections, offset, exchanges, starttime, endtime)
	return args.Get(0).([]struct {
		Name       string
		Address    string
		Blockchain string
		Volume     float64
	}), args.Error(1)
}

func (m *RelDatastoreMock) GetScraperState(ctx context.Context, scraperName string, state models.ScraperState) error {
	args := m.Called(ctx, scraperName, state)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetScraperConfig(ctx context.Context, scraperName string, config models.ScraperConfig) error {
	args := m.Called(ctx, scraperName, config)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetOraclesByOwner(owner string) (oracleconfigs []dia.OracleConfig, err error) {
	args := m.Called(owner)
	return args.Get(0).([]dia.OracleConfig), args.Error(1)
}

func (m *RelDatastoreMock) GetOracleUpdatesByTimeRange(address, chainid, symbol string, offset int, startTime, endTime time.Time) ([]dia.OracleUpdate, error) {
	args := m.Called(address, chainid, symbol, offset, startTime, endTime)
	return args.Get(0).([]dia.OracleUpdate), args.Error(1)
}

func (m *RelDatastoreMock) GetOracleUpdates(address string, chainid string, offset int) ([]dia.OracleUpdate, error) {
	args := m.Called(address, chainid, offset)
	return args.Get(0).([]dia.OracleUpdate), args.Error(1)
}

func (m *RelDatastoreMock) GetOracleUpdateCount(address string, chainid string, symbol string) (int64, error) {
	args := m.Called(address, chainid, symbol)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RelDatastoreMock) GetOracleConfig(oracleaddress, chainID string) (dia.OracleConfig, error) {
	args := m.Called(oracleaddress, chainID)
	return args.Get(0).(dia.OracleConfig), args.Error(1)
}

func (m *RelDatastoreMock) GetNumNFTTrades(address string, blockchain string, exchange string, starttime time.Time, endtime time.Time) (int, error) {
	args := m.Called(address, blockchain, exchange, starttime, endtime)
	return args.Int(0), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTVolume(address string, blockchain string, exchange string, starttime time.Time, endtime time.Time) (float64, error) {
	args := m.Called(address, blockchain, exchange, starttime, endtime)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTExchange(name string) (exchange dia.Exchange, err error) {
	args := m.Called(name)
	return args.Get(0).(dia.Exchange), args.Error(1)
}

func (m *RelDatastoreMock) GetLastOracleUpdate(address string, chainid string) ([]dia.OracleUpdate, error) {
	args := m.Called(address, chainid)
	return args.Get(0).([]dia.OracleUpdate), args.Error(1)
}

func (m *RelDatastoreMock) GetLastNFTOffer(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (dia.NFTOffer, error) {
	args := m.Called(address, blockchain, tokenID, blockNumber, blockPosition)
	return args.Get(0).(dia.NFTOffer), args.Error(1)
}

func (m *RelDatastoreMock) GetLastNFTBid(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (dia.NFTBid, error) {
	args := m.Called(address, blockchain, tokenID, blockNumber, blockPosition)
	return args.Get(0).(dia.NFTBid), args.Error(1)
}

func (m *RelDatastoreMock) GetLastBlockheightTopshot(upperBound time.Time) (uint64, error) {
	args := m.Called(upperBound)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *RelDatastoreMock) GetLastBlockNFTTrade(nftclass dia.NFTClass) (uint64, error) {
	args := m.Called(nftclass)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *RelDatastoreMock) GetLastBlockNFTOffer(nftclass dia.NFTClass) (uint64, error) {
	args := m.Called(nftclass)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *RelDatastoreMock) GetLastBlockNFTBid(nftclass dia.NFTClass) (uint64, error) {
	args := m.Called(nftclass)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *RelDatastoreMock) GetLastBlockBlockscraper(blockchain string) (int64, error) {
	args := m.Called(blockchain)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RelDatastoreMock) GetKeys(table string) ([]string, error) {
	args := m.Called(table)
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetFeederResources() (addresses []string, err error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetKeyPairID(publickey string) string {
	args := m.Called(publickey)
	return args.String(0)
}

func (m *RelDatastoreMock) GetFeederLimit(owner string) (limit int) {
	args := m.Called(owner)
	return args.Int(0)
}

func (m *RelDatastoreMock) GetFeederID(address string) (feederId string) {
	args := m.Called(address)
	return args.String(0)
}

func (m *RelDatastoreMock) GetFeederByID(id string) (owner string) {
	args := m.Called(id)
	return args.String(0)
}

func (m *RelDatastoreMock) GetAllFeeders(bool, bool) (oracleconfigs []dia.OracleConfig, err error) {
	args := m.Called()
	return args.Get(0).([]dia.OracleConfig), args.Error(1)
}
func (m *RelDatastoreMock) SetAsset(asset dia.Asset) error {
	args := m.Called(asset)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetAsset(address, blockchain string) (dia.Asset, error) {
	args := m.Called(address, blockchain)
	return args.Get(0).(dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetByID(ID string) (dia.Asset, error) {
	args := m.Called(ID)
	return args.Get(0).(dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetsBySymbolName(symbol, name string) ([]dia.Asset, error) {
	args := m.Called(symbol, name)
	return args.Get(0).([]dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetAllAssets(blockchain string) ([]dia.Asset, error) {
	args := m.Called(blockchain)
	return args.Get(0).([]dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetFiatAssetBySymbol(symbol string) (dia.Asset, error) {
	args := m.Called(symbol)
	return args.Get(0).(dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetID(asset dia.Asset) (string, error) {
	args := m.Called(asset)
	return args.String(0), args.Error(1)
}

func (m *RelDatastoreMock) GetPage(pageNumber uint32) ([]dia.Asset, bool, error) {
	args := m.Called(pageNumber)
	return args.Get(0).([]dia.Asset), args.Bool(1), args.Error(2)
}

func (m *RelDatastoreMock) Count() (uint32, error) {
	args := m.Called()
	return args.Get(0).(uint32), args.Error(1)
}

func (m *RelDatastoreMock) SetAssetVolume24H(asset dia.Asset, volume float64, timestamp time.Time) error {
	args := m.Called(asset, volume, timestamp)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetLastAssetVolume24H(asset dia.Asset) (float64, error) {
	args := m.Called(asset)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetsWithVOL(starttime time.Time, numAssets int64, skip int64, onlycex bool, substring string) ([]dia.AssetVolume, error) {
	args := m.Called(starttime, numAssets, skip, onlycex, substring)
	return args.Get(0).([]dia.AssetVolume), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetSource(asset dia.Asset, onlycex bool) ([]string, error) {
	args := m.Called(asset, onlycex)
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetsWithVolByBlockchain(starttime time.Time, endtime time.Time, blockchain string) ([]dia.AssetVolume, error) {
	args := m.Called(starttime, endtime, blockchain)
	return args.Get(0).([]dia.AssetVolume), args.Error(1)
}

func (m *RelDatastoreMock) SetExchangePair(exchange string, pair dia.ExchangePair, cache bool) error {
	args := m.Called(exchange, pair, cache)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetExchangePair(exchange string, foreignname string, caseSensitive bool) (dia.ExchangePair, error) {
	args := m.Called(exchange, foreignname, caseSensitive)
	return args.Get(0).(dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangePairSeparator(exchange string) (string, error) {
	args := m.Called(exchange)
	return args.String(0), args.Error(1)
}

func (m *RelDatastoreMock) GetPairsForExchange(exchange dia.Exchange, filterVerified bool, verified bool) ([]dia.ExchangePair, error) {
	args := m.Called(exchange, filterVerified, verified)
	return args.Get(0).([]dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) GetPairsForAsset(asset dia.Asset, filterVerified bool, verified bool) ([]dia.ExchangePair, error) {
	args := m.Called(asset, filterVerified, verified)
	return args.Get(0).([]dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangepairsByAsset(asset dia.Asset, exchange string, basetoken bool) ([]dia.ExchangePair, error) {
	args := m.Called(asset, exchange, basetoken)
	return args.Get(0).([]dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangePairSymbols(exchange string) ([]dia.ExchangePair, error) {
	args := m.Called(exchange)
	return args.Get(0).([]dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) GetNumPairs(exchange dia.Exchange) (int, error) {
	args := m.Called(exchange)
	return args.Int(0), args.Error(1)
}

func (m *RelDatastoreMock) SetExchangeSymbol(exchange string, symbol string) error {
	args := m.Called(exchange, symbol)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetExchangeSymbol(exchange string, symbol string) (dia.Asset, error) {
	args := m.Called(exchange, symbol)
	return args.Get(0).(dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangeSymbols(exchange string, substring string) ([]string, error) {
	args := m.Called(exchange, substring)
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetUnverifiedExchangeSymbols(exchange string) ([]string, error) {
	args := m.Called(exchange)
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) VerifyExchangeSymbol(exchange string, symbol string, assetID string) (bool, error) {
	args := m.Called(exchange, symbol, assetID)
	return args.Bool(0), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangeSymbolAssetID(exchange string, symbol string) (string, bool, error) {
	args := m.Called(exchange, symbol)
	return args.String(0), args.Bool(1), args.Error(2)
}

func (m *RelDatastoreMock) GetAllExchangeAssets(verified bool) ([]dia.Asset, error) {
	args := m.Called(verified)
	return args.Get(0).([]dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) SetHistoricalQuotation(quotation models.AssetQuotation) error {
	args := m.Called(quotation)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetHistoricalQuotations(asset dia.Asset, starttime time.Time, endtime time.Time) ([]models.AssetQuotation, error) {
	args := m.Called(asset, starttime, endtime)
	return args.Get(0).([]models.AssetQuotation), args.Error(1)
}

func (m *RelDatastoreMock) GetLastHistoricalQuotationTimestamp(asset dia.Asset) (time.Time, error) {
	args := m.Called(asset)
	return args.Get(0).(time.Time), args.Error(1)
}

func (m *RelDatastoreMock) SetExchange(exchange dia.Exchange) error {
	args := m.Called(exchange)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetExchange(name string) (dia.Exchange, error) {
	args := m.Called(name)
	return args.Get(0).(dia.Exchange), args.Error(1)
}

func (m *RelDatastoreMock) GetAllExchanges() ([]dia.Exchange, error) {
	args := m.Called()
	return args.Get(0).([]dia.Exchange), args.Error(1)
}

func (m *RelDatastoreMock) GetExchangeNames() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) SetPool(pool dia.Pool) error {
	args := m.Called(pool)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetScraperIndex(exchange string, scraperType string, indexType string, index int64) error {
	args := m.Called(exchange, scraperType, indexType, index)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetScraperIndex(exchange string, scraperType string) (string, int64, error) {
	args := m.Called(exchange, scraperType)
	return args.String(0), args.Get(1).(int64), args.Error(2)
}

func (m *RelDatastoreMock) GetAllDEXPoolsCount() (map[string]int, error) {
	args := m.Called()
	return args.Get(0).(map[string]int), args.Error(1)
}

func (m *RelDatastoreMock) GetPoolByAddress(blockchain string, address string) (dia.Pool, error) {
	args := m.Called(blockchain, address)
	return args.Get(0).(dia.Pool), args.Error(1)
}

func (m *RelDatastoreMock) GetAllPoolAddrsExchange(exchange string, liquiThreshold float64) ([]string, error) {
	args := m.Called(exchange, liquiThreshold)
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetAllPoolsExchange(exchange string, liquiThreshold float64) ([]dia.Pool, error) {
	args := m.Called(exchange, liquiThreshold)
	return args.Get(0).([]dia.Pool), args.Error(1)
}

func (m *RelDatastoreMock) GetPoolsByAsset(asset dia.Asset, liquidityThreshold float64, liquidityThresholdUSD float64) ([]dia.Pool, error) {
	args := m.Called(asset, liquidityThreshold, liquidityThresholdUSD)
	return args.Get(0).([]dia.Pool), args.Error(1)
}

func (m *RelDatastoreMock) SetBlockchain(blockchain dia.BlockChain) error {
	args := m.Called(blockchain)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetBlockchain(name string) (dia.BlockChain, error) {
	args := m.Called(name)
	return args.Get(0).(dia.BlockChain), args.Error(1)
}

func (m *RelDatastoreMock) GetAllAssetsBlockchains() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) GetAllBlockchains(fullAsset bool) ([]dia.BlockChain, error) {
	args := m.Called(fullAsset)
	return args.Get(0).([]dia.BlockChain), args.Error(1)
}

func (m *RelDatastoreMock) SetAssetCache(asset dia.Asset) error {
	args := m.Called(asset)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetAssetCache(blockchain string, address string) (dia.Asset, error) {
	args := m.Called(blockchain, address)
	return args.Get(0).(dia.Asset), args.Error(1)
}

func (m *RelDatastoreMock) SetExchangePairCache(exchange string, pair dia.ExchangePair) error {
	args := m.Called(exchange, pair)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetExchangePairCache(exchange string, foreignName string) (dia.ExchangePair, error) {
	args := m.Called(exchange, foreignName)
	return args.Get(0).(dia.ExchangePair), args.Error(1)
}

func (m *RelDatastoreMock) CountCache() (uint32, error) {
	args := m.Called()
	return args.Get(0).(uint32), args.Error(1)
}

func (m *RelDatastoreMock) SetNFTClass(nftClass dia.NFTClass) error {
	args := m.Called(nftClass)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetAllNFTClasses(blockchain string) ([]dia.NFTClass, error) {
	args := m.Called(blockchain)
	return args.Get(0).([]dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) GetTradedNFTClasses(starttime time.Time) ([]dia.NFTClass, error) {
	args := m.Called(starttime)
	return args.Get(0).([]dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTClasses(limit, offset uint64) ([]dia.NFTClass, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTClass(address string, blockchain string) (dia.NFTClass, error) {
	args := m.Called(address, blockchain)
	return args.Get(0).(dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTClassID(address string, blockchain string) (string, error) {
	args := m.Called(address, blockchain)
	return args.String(0), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTClassByID(id string) (dia.NFTClass, error) {
	args := m.Called(id)
	return args.Get(0).(dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTClassesByNameSymbol(searchstring string) ([]dia.NFTClass, error) {
	args := m.Called(searchstring)
	return args.Get(0).([]dia.NFTClass), args.Error(1)
}

func (m *RelDatastoreMock) UpdateNFTClassCategory(nftclassID string, category string) (bool, error) {
	args := m.Called(nftclassID, category)
	return args.Bool(0), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTCategories() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *RelDatastoreMock) SetNFT(nft dia.NFT) error {
	args := m.Called(nft)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error) {
	args := m.Called(address, blockchain, tokenID)
	return args.Get(0).(dia.NFT), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTID(address string, blockchain string, tokenID string) (string, error) {
	args := m.Called(address, blockchain, tokenID)
	return args.String(0), args.Error(1)
}

func (m *RelDatastoreMock) SetNFTTrade(trade dia.NFTTrade) error {
	args := m.Called(trade)
	return args.Error(0)
}

func (m *RelDatastoreMock) SetNFTTradeToTable(trade dia.NFTTrade, table string) error {
	args := m.Called(trade, table)
	return args.Error(0)
}

func (m *RelDatastoreMock) GetNFTTrades(address string, blockchain string, tokenID string, starttime time.Time, endtime time.Time) ([]dia.NFTTrade, error) {
	args := m.Called(address, blockchain, tokenID, starttime, endtime)
	return args.Get(0).([]dia.NFTTrade), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTTradesCollection(address string, blockchain string, starttime time.Time, endtime time.Time) ([]dia.NFTTrade, error) {
	args := m.Called(address, blockchain, starttime, endtime)
	return args.Get(0).([]dia.NFTTrade), args.Error(1)
}

func (m *RelDatastoreMock) GetAllLastTrades(nftclass dia.NFTClass) ([]dia.NFTTrade, error) {
	args := m.Called(nftclass)
	return args.Get(0).([]dia.NFTTrade), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTOffers(address string, blockchain string, tokenID string) ([]dia.NFTOffer, error) {
	args := m.Called(address, blockchain, tokenID)
	return args.Get(0).([]dia.NFTOffer), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTBids(address string, blockchain string, tokenID string) ([]dia.NFTBid, error) {
	args := m.Called(address, blockchain, tokenID)
	return args.Get(0).([]dia.NFTBid), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTFloor(nftclass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, noBundles bool, exchange string) (float64, error) {
	args := m.Called(nftclass, timestamp, floorWindowSeconds, noBundles, exchange)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTFloorLevel(nftclass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, currencies []dia.Asset, level float64, noBundles bool, exchange string) (float64, error) {
	args := m.Called(nftclass, timestamp, floorWindowSeconds, currencies, level, noBundles, exchange)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTFloorRecursive(nftClass dia.NFTClass, timestamp time.Time, floorWindowSeconds time.Duration, stepBackLimit int, noBundles bool, exchange string) (float64, error) {
	args := m.Called(nftClass, timestamp, floorWindowSeconds, stepBackLimit, noBundles, exchange)
	return args.Get(0).(float64), args.Error(1)
}

func (m *RelDatastoreMock) GetNFTFloorRange(nftClass dia.NFTClass, starttime time.Time, endtime time.Time, floorWindowSeconds time.Duration, stepBackLimit int, noBundles bool, exchange string) ([]float64, error) {
	args := m.Called(nftClass, starttime, endtime, floorWindowSeconds, stepBackLimit, noBundles, exchange)
	return args.Get(0).([]float64), args.Error(1)
}

// func (m *RelDatastoreMock) SetNFTTransfer(transfer dia.NFTTransfer) error {
// 	args := m.Called(transfer)
// 	return args.Error(0)
// }

func (m *RelDatastoreMock) Get24HoursNFTExchangeVolume(exchange dia.NFTExchange) (float64, error) {
	args := m.Called(exchange)
	return args.Get(0).(float64), args.Error(1)
}

// func (m *RelDatastoreMock) GetNFTTransfers(address string, blockchain string, tokenID string, starttime time.Time, endtime time.Time) ([]dia.NFTTransfer, error) {
// 	args := m.Called(address, blockchain, tokenID, starttime, endtime)
// 	return args.Get(0).([]dia.NFTTransfer), args.Error(1)
// }

// func (m *RelDatastoreMock) GetAllNFTTransfers(address string, blockchain string, starttime time.Time, endtime time.Time) ([]dia.NFTTransfer, error) {
// 	args := m.Called(address, blockchain, starttime, endtime)
// 	return args.Get(0).([]dia.NFTTransfer), args.Error(1)
// }

func (m *RelDatastoreMock) GetOracleLastUpdate(address, chainid, symbol string) (time.Time, string, error) {
	args := m.Called(address, chainid, symbol)
	return args.Get(0).(time.Time), args.Get(1).(string), args.Error(2)
}

func (m *RelDatastoreMock) AddWalletKeys(owner string, publicKey []string) error {
	args := m.Called(owner, publicKey)
	return args.Error(0)
}

func (m *RelDatastoreMock) ChangeOracleState(feederID string, active bool) error {
	args := m.Called(feederID, active)
	return args.Error(0)
}

func (m *RelDatastoreMock) CreateCustomer(email string, customerPlan int, paymentStatus string, paymentSource string, numberOfDataFeeds int, walletPublicKeys []string) error {
	args := m.Called(email, customerPlan, paymentStatus, paymentSource, numberOfDataFeeds, walletPublicKeys)
	return args.Error(0)
}
func (m *RelDatastoreMock) DeleteAssetList(sheetName string) error {
	args := m.Called(sheetName)
	return args.Error(0)
}
func (m *RelDatastoreMock) DeleteOracle(feederID string) error {
	args := m.Called(feederID)
	return args.Error(0)
}

func (m *RelDatastoreMock) Get24HoursNFTExchangeTrades(exchange dia.NFTExchange) (int64, error) {
	args := m.Called(exchange)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RelDatastoreMock) GetAccessLevel(publicKey string) (string, error) {
	args := m.Called(publicKey)
	return args.String(0), args.Error(1)
}

func (m *RelDatastoreMock) GetAllChains() (chainconfigs []dia.ChainConfig, err error) {
	args := m.Called()
	return args.Get(0).([]dia.ChainConfig), args.Error(1)
}

func (m *RelDatastoreMock) GetAllNFTExchanges() (exchanges []dia.NFTExchange, err error) {
	args := m.Called()
	return args.Get(0).([]dia.NFTExchange), args.Error(1)
}

func (m *RelDatastoreMock) GetAssetListBySymbol(symbol string, listname string) ([]dia.AssetList, error) {
	args := m.Called(symbol, listname)
	return args.Get(0).([]dia.AssetList), args.Error(1)
}

func (m *RelDatastoreMock) GetBalanceRemaining(address string, chainid string) (float64, error) {
	args := m.Called(address, chainid)
	return args.Get(0).(float64), args.Error(1)
}
func (m *RelDatastoreMock) GetBlockData(blockchain string, blocknumber int64) (dia.BlockData, error) {
	args := m.Called(blockchain, blocknumber)
	return args.Get(0).(dia.BlockData), args.Error(1)
}
func (m *RelDatastoreMock) GetCollectionCountByExchange(exchange string) (int64, error) {
	args := m.Called(exchange)
	return args.Get(0).(int64), args.Error(1)
}
func (m *RelDatastoreMock) GetCustomerByPublicKey(publicKey string) (*models.Customer, error) {
	args := m.Called(publicKey)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *RelDatastoreMock) GetCustomerIDByWalletPublicKey(publicKey string) (int, error) {
	args := m.Called(publicKey)
	return args.Int(0), args.Error(1)
}
func (m *RelDatastoreMock) GetDayWiseUpdates(address string, chainid string) ([]dia.FeedUpdates, float64, float64, error) {
	args := m.Called(address, chainid)
	return args.Get(0).([]dia.FeedUpdates), args.Get(1).(float64), args.Get(2).(float64), args.Error(3)
}

func (m *RelDatastoreMock) GetExpiredFeeders() ([]dia.OracleConfig, error) {
	args := m.Called()
	return args.Get(0).([]dia.OracleConfig), args.Error(1)
}
func (m *RelDatastoreMock) GetFeeder(feederID string) (dia.OracleConfig, error) {
	args := m.Called(feederID)
	return args.Get(0).(dia.OracleConfig), args.Error(1)
}

func (m *RelDatastoreMock) GetFeederAccessByID(id string) (owner string) {
	args := m.Called(id)
	return args.String(0)
}
