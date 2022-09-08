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
	FIAT                                    = "Fiat"
)

type VerificationMechanism string

// NFTClass is the container for an nft class defined by
// a contract (address) on a blockchain.
type NFTClass struct {
	Address      string
	Symbol       string
	Name         string
	Blockchain   string
	ContractType string
	Category     string
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
	NFTClass       NFTClass
	TokenID        string
	CreationTime   time.Time
	CreatorAddress string
	URI            string
	// @Attributes is a collection of attributes from on- and off-chain
	// TO DO: Should we split up into two fields?
	Attributes NFTAttributes
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
	NFT         NFT
	Price       *big.Int
	PriceUSD    float64
	FromAddress string
	ToAddress   string
	Currency    Asset
	BundleSale  bool
	BlockNumber uint64
	Timestamp   time.Time
	TxHash      string
	Exchange    string
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
	NFT         NFT
	Value       *big.Int
	FromAddress string

	CurrencySymbol   string
	CurrencyAddress  string
	CurrencyDecimals int32

	BlockNumber   uint64
	BlockPosition uint64
	Timestamp     time.Time
	TxHash        string
	Exchange      string
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

type NFTOffer struct {
	NFT NFT
	// Start and EndValue are for auction types. Otherwise, use StartValue
	// and leave EndValue blank.
	StartValue *big.Int
	EndValue   *big.Int
	// Duration of the offer/auction measured in seconds
	Duration    time.Duration
	FromAddress string
	// Type of offer can be auction, simple offer,...
	AuctionType string

	CurrencySymbol   string
	CurrencyAddress  string
	CurrencyDecimals int32

	BlockNumber   uint64
	BlockPosition uint64
	Timestamp     time.Time
	TxHash        string
	Exchange      string
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

type AggregatedVolume struct {
	Pair             Pair      `json:"Pair"`
	Volume           float64   `json:"Volume"`
	Exchange         string    `json:"ExchangeVolumes"`
	TimeRangeSeconds int64     `json:"TimeRangeSeconds"`
	Timestamp        time.Time `json:"Timestamp"`
}

type TradesDistribution struct {
	Asset            Asset     `json:"Asset"`
	NumTradesTotal   int       `json:"NumTradesTotal"`
	NumLowBins       int       `json:"NumberLowBins"`
	Threshold        int       `json:"Threshold"`
	SizeBinSeconds   int64     `json:"SizeBin"`
	AvgNumPerBin     float64   `json:"AverageNumberPerBin"`
	StdDeviation     float64   `json:"StandardDeviation"`
	TimeRangeSeconds int64     `json:"TimeRangeSeconds"`
	Timestamp        time.Time `json:"Timestamp"`
}

type EthereumBlockData struct {
	GasLimit    uint64             `json:"gas_limit"`
	GasUsed     uint64             `json:"gas_used"`
	Difficulty  *big.Int           `json:"difficulty"`
	Time        uint64             `json:"time"`
	Size        common.StorageSize `json:"size"`
	Number      uint64             `json:"number"`
	MixDigest   common.Hash        `json:"mix_digest"`
	Nonce       uint64             `json:"nonce"`
	Coinbase    common.Address     `json:"coinbase"`
	Root        common.Hash        `json:"root"`
	ParentHash  common.Hash        `json:"parent_hash"`
	TxHash      common.Hash        `json:"tx_hash"`
	ReceiptHash common.Hash        `json:"receipt_hash"`
	UncleHash   common.Hash        `json:"uncle_hash"`
	Extra       []byte             `json:"extra"`
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
}

type Supply struct {
	Asset             Asset
	Supply            float64
	CirculatingSupply float64
	Source            string
	Time              time.Time
}

// Asset is the data type for all assets, ranging from fiat to crypto.
type Asset struct {
	Symbol     string
	Name       string
	Address    string
	Decimals   uint8
	Blockchain string
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
	RestURL string `json:"restURL"`
	WSURL   string `json:"wsURL"`
	ChainID string `json:"ChainID"`
}

// Pair substitues the old dia.Pair. It includes the new asset type.
type Pair struct {
	QuoteToken Asset
	BaseToken  Asset
}

// ForeignName returns the foreign name of the pair @p, i.e. the string Quotetoken-Basetoken
func (p *Pair) ForeignName() string {
	return p.QuoteToken.Symbol + "-" + p.BaseToken.Symbol
}

// Pool is the container for liquidity pools on DEXes.
type Pool struct {
	Exchange   Exchange
	Blockchain BlockChain
	Address    string
	// Assetvolumes map[Asset]float64
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
	Symbol         string
	ForeignName    string
	Exchange       string
	Verified       bool
	UnderlyingPair Pair
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
	Symbol string
	Pair   string
	// Final fields for trade
	QuoteToken        Asset
	BaseToken         Asset
	Price             float64
	Volume            float64 // Quantity of bought/sold units of Quote token. Negative if result of Market order Sell
	Time              time.Time
	ForeignTradeID    string
	EstimatedUSDPrice float64 // will be filled by the TradesBlockService
	Source            string
	VerifiedPair      bool // will be filled by the pairDiscoveryService
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
	BlockHash      string
	IndexBlockData IndexBlockData
}

type IndexBlockData struct {
	FiltersBlockHash    string
	SuppliesBlockHash   string
	VolatilityBlockHash string
	IndexElements       []IndexElement
	IndexElementsNumber int
	Time                time.Time
	IndexValue          float64
	ValueTokenette      float64
	ValueToken          float64
	USDPerPointsOfIndex float64
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
	BlockHash string
	BlockData SuppliesBlockData
}

type SuppliesBlockData struct {
	Time     time.Time
	Supplies []Supply
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
