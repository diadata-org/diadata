package dia

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	Diadata                                 = "diadata.org"
	PROOF_OF_STAKE    VerificationMechanism = "pos"
	PROOF_OF_WORK     VerificationMechanism = "pow"
	BITCOIN                                 = "Bitcoin"
	ETHEREUM                                = "Ethereum"
	BINANCESMARTCHAIN                       = "BinanceSmartChain"
	POLYGON                                 = "Polygon"
	CELO                                    = "Celo"
	FANTOM                                  = "Fantom"
	NEAR                                    = "NEAR"
	AURORA                                  = "Aurora"
	SOLANA                                  = "Solana"
	FLOW                                    = "Flow"
	MOONRIVER                               = "Moonriver"
	MOONBEAM                                = "Moonbeam"
	AVALANCHE                               = "Avalanche"
	ARBITRUM                                = "Arbitrum"
	ASTAR                                   = "Astar"
	SHIDEN                                  = "Shiden"
	METIS                                   = "Metis"
	KILT                                    = "Kilt"
	FETCH                                   = "Fetch"
	FUSE                                    = "Fuse"
	TELOS                                   = "Telos"
	EVMOS                                   = "Evmos"
	KUSAMA                                  = "Kusama"
	ACALA                                   = "Acala"
	POLKADOT                                = "Polkadot"
	WANCHAIN                                = "Wanchain"
	OSMOSIS                                 = "Osmosis"
	FIAT                                    = "Fiat"
	BIFROST                                 = "Bifrost"
)

type VerificationMechanism string

// NFTClass is the container for a nft class defined by
// a contract (address) on a blockchain.
type NFTClass struct {
	Address      string `json:"Address"`
	Symbol       string `json:"Symbol"`
	Name         string `json:"Name"`
	Blockchain   string `json:"Blockchain"`
	ContractType string `json:"ContractType"`
	Category     string `json:"Category"`
}

// MarshalBinary for NFTClass
func (nc *NFTClass) MarshalBinary() ([]byte, error) {
	return json.Marshal(nc)
}

// UnmarshalBinary for NFTClass
func (nc *NFTClass) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &nc); err != nil {
		return err
	}
	return nil
}

// NFT is the container for a specific NFT defined by
// the pair (address,tokenID).
type NFT struct {
	NFTClass       NFTClass  `json:"NFTClass"`
	TokenID        string    `json:"TokenID"`
	CreationTime   time.Time `json:"CreationTme"`
	CreatorAddress string    `json:"CreatorAddress"`
	URI            string    `json:"URI"`
	// @Attributes is a collection of attributes from on- and off-chain
	// TO DO: Should we split up into two fields?
	Attributes NFTAttributes `json:"Attributes"`
}

// NFTAttributes can be stored as jasonb in postgres:
// https://www.alexedwards.net/blog/using-postgresql-jsonb
type NFTAttributes map[string]interface{}

func (a NFTAttributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *NFTAttributes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// MarshalBinary for NFT
func (n *NFT) MarshalBinary() ([]byte, error) {
	return json.Marshal(n)
}

// UnmarshalBinary for NFT
func (n *NFT) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	return nil
}

type NFTTrade struct {
	NFT         NFT       `json:"NFT"`
	Price       *big.Int  `json:"Price"`
	PriceUSD    float64   `json:"PriceUSD"`
	FromAddress string    `json:"FromAddress"`
	ToAddress   string    `json:"ToAddress"`
	Currency    Asset     `json:"Currency"`
	BundleSale  bool      `json:"BundleSale"`
	BlockNumber uint64    `json:"BlockNumber"`
	Timestamp   time.Time `json:"Timestamp"`
	TxHash      string    `json:"TxHash"`
	Exchange    string    `json:"Exchange"`
}

// MarshalBinary for NFTTrade
func (ns *NFTTrade) MarshalBinary() ([]byte, error) {
	return json.Marshal(ns)
}

// UnmarshalBinary for NFTTrade
func (ns *NFTTrade) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &ns); err != nil {
		return err
	}
	return nil
}

type NFTBid struct {
	NFT              NFT       `json:"NFT"`
	Value            *big.Int  `json:"Value"`
	FromAddress      string    `json:"FromAddress"`
	CurrencySymbol   string    `json:"CurrencySymbol"`
	CurrencyAddress  string    `json:"CurrencyAddress"`
	CurrencyDecimals int32     `json:"CurrencyDecimals"`
	BlockNumber      uint64    `json:"BlockNumber"`
	BlockPosition    uint64    `json:"BlockPosition"`
	Timestamp        time.Time `json:"Timestamp"`
	TxHash           string    `json:"TxHash"`
	Exchange         string    `json:"Exchange"`
}

// MarshalBinary for NFTBid
func (nb *NFTBid) MarshalBinary() ([]byte, error) {
	return json.Marshal(nb)
}

// UnmarshalBinary for NFTBid
func (nb *NFTBid) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &nb); err != nil {
		return err
	}
	return nil
}

type NFTExchangeStats struct {
	Exchange  string
	NumTrades uint64
	Volume    float64
}

type NFTOffer struct {
	NFT NFT `json:"NFT"`
	// Start and EndValue are for auction types. Otherwise, use StartValue
	// and leave EndValue blank.
	StartValue *big.Int `json:"StartValue"`
	EndValue   *big.Int `json:"EndValue"`
	// Duration of the offer/auction measured in seconds
	Duration    time.Duration `json:"Duration"`
	FromAddress string        `json:"FromAddress"`
	// Type of offer can be auction, simple offer,...
	AuctionType string `json:"AuctionType"`

	CurrencySymbol   string `json:"CurrencySymbol"`
	CurrencyAddress  string `json:"CurrencyAddress"`
	CurrencyDecimals int32  `json:"CurrencyDecimals"`

	BlockNumber   uint64    `json:"BlockNumber"`
	BlockPosition uint64    `json:"BlockPosition"`
	Timestamp     time.Time `json:"Timestamp"`
	TxHash        string    `json:"TxHash"`
	Exchange      string    `json:"Exchange"`
}

// MarshalBinary for NFTOffer
func (no *NFTOffer) MarshalBinary() ([]byte, error) {
	return json.Marshal(no)
}

// UnmarshalBinary for NFTOffer
func (no *NFTOffer) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &no); err != nil {
		return err
	}
	return nil
}

// BlockData stores information on a specific block in a given blockchain.
type BlockData struct {
	// Name of the blockchain, as found for instance in dia.ETHEREUM
	BlockchainName string
	// In order to keep it general, BlockNumber is a string
	BlockNumber int64
	Data        map[string]interface{}
}

type ExchangeVolume struct {
	Exchange string  `json:"Exchange"`
	Volume   float64 `json:"Volume"`
}

type ExchangeVolumesList struct {
	Volumes   []ExchangeVolume `json:"Volumes"`
	Timestamp time.Time        `json:"Timestamp"`
}

type AssetVolume struct {
	Asset  Asset   `json:"Asset"`
	Volume float64 `json:"Volume"`
	Index  uint8   `json:"Index"`
}

type AssetLiquidity struct {
	Asset  Asset   `json:"Asset"`
	Volume float64 `json:"Liquidity"`
	Index  uint8   `json:"Index"`
}

type TopAsset struct {
	Asset          Asset               `json:"Asset"`
	Volume         float64             `json:"Volume"`
	Price          float64             `json:"Price"`
	PriceYesterday float64             `json:"PriceYesterday"`
	Source         map[string][]string `json:"Source"`
}

type PairVolume struct {
	Pair   Pair    `json:"Pair"`
	Volume float64 `json:"Volume"`
}

type PairVolumesList struct {
	Volumes   []PairVolume `json:"Volumes"`
	Timestamp time.Time    `json:"Timestamp"`
}

type EthereumBlockData struct {
	GasLimit    uint64             `json:"GasLimit"`
	GasUsed     uint64             `json:"GasUsed"`
	Difficulty  *big.Int           `json:"Difficulty"`
	Time        uint64             `json:"Time"`
	Size        common.StorageSize `json:"Size"`
	Number      uint64             `json:"Number"`
	MixDigest   common.Hash        `json:"MixDigest"`
	Nonce       uint64             `json:"Nonce"`
	Coinbase    common.Address     `json:"Coinbase"`
	Root        common.Hash        `json:"Root"`
	ParentHash  common.Hash        `json:"ParentHash"`
	TxHash      common.Hash        `json:"TxHash"`
	ReceiptHash common.Hash        `json:"ReceiptHash"`
	UncleHash   common.Hash        `json:"UncleHash"`
	Extra       []byte             `json:"Extra"`
}

type Exchange struct {
	Name          string     `json:"Name"`
	Centralized   bool       `json:"Centralized"`
	Bridge        bool       `json:"Bridge"`
	Contract      string     `json:"Contract"`
	BlockChain    BlockChain `json:"BlockChain"`
	RestAPI       string     `json:"RestAPI"`
	WsAPI         string     `json:"WsAPI"`
	PairsAPI      string     `json:"PairsAPI"`
	WatchdogDelay int        `json:"WatchdogDelay"`
	ScraperActive bool       `json:"ScraperActive"`
}

type NFTExchange struct {
	Name          string     `json:"Name"`
	Centralized   bool       `json:"Centralized"`
	Contract      string     `json:"Contract"`
	BlockChain    BlockChain `json:"BlockChain"`
	RestAPI       string     `json:"RestAPI"`
	WsAPI         string     `json:"WsAPI"`
	WatchdogDelay int        `json:"WatchdogDelay"`
}

type Supply struct {
	Asset             Asset     `json:"Asset"`
	Supply            float64   `json:"Supply"`
	CirculatingSupply float64   `json:"CirculatingSupply"`
	Source            string    `json:"Source"`
	Time              time.Time `json:"Time"`
}

// Asset is the data type for all assets, ranging from fiat to crypto.
type Asset struct {
	Symbol     string `json:"Symbol"`
	Name       string `json:"Name"`
	Address    string `json:"Address"`
	Decimals   uint8  `json:"Decimals"`
	Blockchain string `json:"Blockchain"`
}

func (asset *Asset) Identifier() string {
	return asset.Blockchain + "-" + asset.Address
}

// BlockChain is the type for blockchains. Uniquely defined by its @Name.
type BlockChain struct {
	Name string `json:"Name"`
	// Genesis date is a Unix timestamp
	GenesisDate int64 `json:"GenesisDate"`
	NativeToken Asset `json:"NativeToken"`
	// Verificationmechanism is in short notation, such as pos for proof-of-stake
	VerificationMechanism VerificationMechanism `json:"VerificationMechanism"`
	// ChainID refers to EVM based chains and is thereby optional.
	ChainID string `json:"ChainID"`
}

type ChainConfig struct {
	RestURL string `json:"RestURL"`
	WSURL   string `json:"WSURL"`
	ChainID string `json:"ChainID"`
}

// Pair substitues the old dia.Pair. It includes the new asset type.
type Pair struct {
	QuoteToken Asset `json:"QuoteToken"`
	BaseToken  Asset `json:"BaseToken"`
}

func (p *Pair) Identifier() string {
	return p.QuoteToken.Blockchain + "-" + p.QuoteToken.Address + "-" + p.BaseToken.Blockchain + "-" + p.BaseToken.Address
}

func (p *Pair) PairExchangeIdentifier(exchange string) string {
	return exchange + "-" + p.Identifier()
}

// ForeignName returns the foreign name of the pair @p, i.e. the string Quotetoken-Basetoken
func (p *Pair) ForeignName() string {
	return p.QuoteToken.Symbol + "-" + p.BaseToken.Symbol
}

// Pool is the container for liquidity pools on DEXes.
type Pool struct {
	Exchange     Exchange
	Blockchain   BlockChain
	Address      string
	Assetvolumes []AssetVolume
	Time         time.Time
}

// MarshalBinary is a custom marshaller for BlockChain type
func (bc *BlockChain) MarshalBinary() ([]byte, error) {
	return json.Marshal(bc)
}

// UnmarshalBinary is a custom unmarshaller for BlockChain type
func (bc *BlockChain) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &bc); err != nil {
		return err
	}
	return nil
}

// MarshalBinary is a custom marshaller for Asset type
func (a *Asset) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

// UnmarshalBinary is a custom unmarshaller for Asset type
func (a *Asset) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return nil
}

// ExchangePair is the container for a pair as used by exchanges.
// Across exchanges, these pairs cannot be uniquely mapped on asset pairs.
type ExchangePair struct {
	Symbol         string `json:"Symbol"`
	ForeignName    string `json:"ForeignName"`
	Exchange       string `json:"EXchange"`
	Verified       bool   `json:"Verified"`
	UnderlyingPair Pair   `json:"UnderlyingPair"`
}

// MarshalBinary is a custom marshaller for ExchangePair type
func (ep *ExchangePair) MarshalBinary() ([]byte, error) {
	return json.Marshal(ep)
}

// UnmarshalBinary is a custom unmarshaller for ExchangePair type
func (ep *ExchangePair) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &ep); err != nil {
		return err
	}
	return nil
}

type Pairs []ExchangePair

// Trade remark: In a pair A-B, we call A the Quote token and B the Base token
type Trade struct {
	// TO DO: Deprecated fields. Delete as soon as token-to-type branch is deployed.
	Symbol string `json:"Symbol"`
	Pair   string `json:"Pair"`
	// Final fields for trade
	QuoteToken        Asset     `json:"QuoteToken"`
	BaseToken         Asset     `json:"BaseToken"`
	Price             float64   `json:"Price"`
	Volume            float64   `json:"Volume"` // Quantity of bought/sold units of Quote token. Negative if result of Market order Sell
	Time              time.Time `json:"Time"`
	ForeignTradeID    string    `json:"ForeignTradeID"`
	EstimatedUSDPrice float64   `json:"EstimatedUSDPrice"` // will be filled by the TradesBlockService
	Source            string    `json:"Source"`
	VerifiedPair      bool      `json:"VerifiedPair"` // will be filled by the pairDiscoveryService
}

// SynthAssetSupply is a container for data on synthetic assets such as aUSDC.
// https://etherscan.io/address/0xbcca60bb61934080951369a648fb03df4f96263c
type SynthAssetSupply struct {
	Asset            Asset   // Synthetic asset under consideration.
	AssetUnderlying  Asset   // Asset underlying the synth asset.
	Supply           float64 // Supply of the synthetic asset.
	LockedUnderlying float64 // Amount of underlying asset locked in the contract.
	NumMint          int64   // Total number of synth asset mint events (optional).
	NumRedeem        int64   // Total number of underlying asset redeem events (optional).
	BlockNumber      uint64
	Time             time.Time
	ColleteralRatio  float64
	Protocol         string
	TotalDebt        float64
}

type TradesBlockData struct {
	BeginTime    time.Time
	EndTime      time.Time
	TradesNumber int
	Trades       []Trade
}

type TradesBlock struct {
	BlockHash       string
	TradesBlockData TradesBlockData
}

type FiltersBlock struct {
	BlockHash        string
	FiltersBlockData FiltersBlockData
}

type FiltersBlockData struct {
	TradesBlockHash string
	BeginTime       time.Time
	EndTime         time.Time
	FilterPoints    []FilterPoint
	FiltersNumber   int
}

// FilterPoint contains the resulting value of a filter applied to an asset.
type FilterPoint struct {
	Asset      Asset
	Value      float64
	Name       string
	Time       time.Time
	Max        float64
	Min        float64
	FirstTrade Trade
	LastTrade  Trade
}

type IndexBlock struct {
	BlockHash      string         `json:"BlockHash"`
	IndexBlockData IndexBlockData `json:"IndexBlockData"`
}

type IndexBlockData struct {
	FiltersBlockHash    string         `json:"FiltersBlockHash"`
	SuppliesBlockHash   string         `json:"SuppliesBlockHash"`
	VolatilityBlockHash string         `json:"VolatilityBlockHash"`
	IndexElements       []IndexElement `json:"IndexElements"`
	IndexElementsNumber int            `json:"IndexElementsNumber"`
	Time                time.Time      `json:"Time"`
	IndexValue          float64        `json:"IndexValue"`
	ValueTokenette      float64        `json:"ValueTokenette"`
	ValueToken          float64        `json:"ValueToken"`
	USDPerPointsOfIndex float64        `json:"USDPerPointsOfIndex"`
}

type IndexElement struct {
	Name            string
	Symbol          string
	Percentage      float64
	FilteredPoint   FilterPoint
	Supply          Supply
	VolatilityRatio VolatilityRatio
}

type VolatilityRatio struct {
	Symbol    string
	Threehold float64
	DaysAbove int64
	DaysBelow int64
	Time      time.Time
	Selected  bool
}

type SuppliesBlock struct {
	BlockHash string            `json:"BlockHash"`
	BlockData SuppliesBlockData `json:"BlockData"`
}

type SuppliesBlockData struct {
	Time     time.Time `json:"Time"`
	Supplies []Supply  `json:"Supplies"`
}

type FilterPointMetadata struct {
	Max float64
	Min float64
}

func NewFilterPointMetadata() *FilterPointMetadata {
	return &FilterPointMetadata{Max: 0, Min: -1}
}

func (fp *FilterPointMetadata) AddPoint(value float64) {
	if fp.Max < value {
		fp.Max = value
	}
	if fp.Min > value || fp.Min == -1 {
		fp.Min = value
	}
}

// MarshalBinary -
func (e *FiltersBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *FiltersBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Trade) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Trade) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *TradesBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *TradesBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Supply) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Supply) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *Pairs) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Pairs) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

func (ib IndexBlock) Hash() string {
	return ib.BlockHash
}

// MarshalBinary -
func (e *IndexBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *IndexBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary -
func (e *SuppliesBlock) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *SuppliesBlock) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

type OracleConfig struct {
	Symbols            []string
	FeederID           string
	Address            string
	FeederAddress      string
	Owner              string
	ChainID            string
	Active             bool
	Frequency          string
	SleepSeconds       string
	DeviationPermille  string
	BlockchainNode     string
	MandatoryFrequency string
	CreatedDate        time.Time
	LastUpdate         time.Time
	Deleted            bool `json:"Deleted,omitempty"`
}
