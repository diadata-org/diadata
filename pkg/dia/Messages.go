package dia

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"
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
	BIFROST_POLKADOT                        = "Bifrost-polkadot"
	UNREAL_TESTNET                          = "Unreal-Testnet"
	UNREAL                                  = "Unreal"
	LINEA                                   = "Linea"
	OPTIMISM                                = "Optimism"
	ALEPHIUM                                = "Alephium"
	BASE                                    = "Base"
	FILECOIN                                = "Filecoin"
	HYDRATION                               = "Hydration"
	STACKS                                  = "Stacks"
	SWELLCHAIN                              = "Swellchain"
	SONIC                                   = "Sonic"
)

var CRYPTO_ZERO_UNIX_TIME = time.Unix(1230768000, 0)

type VerificationMechanism string

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
	Asset     Asset   `json:"Asset"`
	Volume    float64 `json:"Volume"`
	VolumeUSD float64 `json:"VolumeUSD"`
	Index     uint8   `json:"Index"`
}

type AssetLiquidity struct {
	Asset     Asset   `json:"Asset"`
	Volume    float64 `json:"Liquidity"`
	VolumeUSD float64 `json:"LiquidityUSD"`
	Index     uint8   `json:"Index"`
}

type TopAsset struct {
	Asset          Asset               `json:"Asset"`
	Volume         float64             `json:"Volume"`
	Price          float64             `json:"Price"`
	PriceYesterday float64             `json:"PriceYesterday"`
	Source         map[string][]string `json:"Source"`
}

type PairVolume struct {
	Pair        Pair    `json:"Pair"`
	PoolAddress string  `json:"Pooladdress"`
	Volume      float64 `json:"Volume"`
	TradesCount int64   `json:"TradesCount"`
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

type AssetList struct {
	AssetName   string
	CustomName  string
	Symbol      string
	Methodology string
	ListName    string

	Exchanges []ExchangeList
}

func (a AssetList) String() string {
	var exchanges []string
	for _, exchange := range a.Exchanges {
		exchanges = append(exchanges, exchange.String())
	}

	return fmt.Sprintf(
		" Exchanges: [%s]",
		strings.Join(exchanges, "; "),
	)
}

type ExchangeList struct {
	Name  string
	Pairs []string
}

func (e ExchangeList) String() string {
	return fmt.Sprintf("%s: [%s]", e.Name, strings.Join(e.Pairs, ", "))
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

// SufficientNativeBalance returns true if all pool assets have at least @threshold liquidity.
func (p *Pool) SufficientNativeBalance(threshold float64) bool {
	sufficientNativeBalance := true
	for _, av := range p.Assetvolumes {
		if av.Volume < threshold {
			sufficientNativeBalance = false
		}
	}
	return sufficientNativeBalance
}

// GetPoolLiquidityUSD returns the total USD liquidity if available.
// @lowerBound is true in case USD liquidity is not available for all pool assets.
func (p *Pool) GetPoolLiquidityUSD() (totalLiquidity float64, lowerBound bool) {
	for _, av := range p.Assetvolumes {
		// For some pools, for instance on BalancerV2 type contracts, the pool contains itself as an asset.
		if av.Asset.Address == p.Address {
			continue
		}
		if av.VolumeUSD == 0 {
			lowerBound = true
		}
		totalLiquidity += av.VolumeUSD
	}
	return
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

type FeedSelection struct {
	Asset              Asset
	Exchangepairs      []ExchangepairSelection
	LiquidityThreshold float64
}

type FeedSelectionAggregated struct {
	Exchange         string
	Quotetoken       Asset
	Basetoken        Asset
	Pooladdress      string
	PoolLiquidityUSD float64
	TradesCount      int32
	Volume           float64
	LastPrice        float64
	Starttime        time.Time
	Endtime          time.Time
	StatusMessage    string
	StatusCode       int32
}

type ExchangepairSelection struct {
	Exchange           Exchange
	Pairs              []Pair
	Pools              []Pool
	LiquidityThreshold float64
}

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
	PoolAddress       string    `json:"PoolAddress"`
	ForeignTradeID    string    `json:"ForeignTradeID"`
	EstimatedUSDPrice float64   `json:"EstimatedUSDPrice"` // will be filled by the TradesBlockService
	Source            string    `json:"Source"`
	VerifiedPair      bool      `json:"VerifiedPair"` // will be filled by the pairDiscoveryService
}

func (t *Trade) VolumeUSD() float64 {
	return t.EstimatedUSDPrice * math.Abs(t.Volume)
}

// NormalizeSymbols normalizes @t.Symbol and @t.Pair in a trade struct to
// upper case letters like so A@pairSplitterB. For instance, btcusdt will be BTC-USDT.
func (t *Trade) NormalizeSymbols(upperCase bool, pairSplitter string) error {
	symbols, err := GetPairSymbols(ExchangePair{Exchange: t.Source, ForeignName: t.Pair, Symbol: t.Symbol})
	if err != nil {
		return err
	}
	if len(symbols) == 2 {
		if upperCase {
			t.Pair = strings.ToUpper(symbols[0] + pairSplitter + symbols[1])
			t.Symbol = strings.ToUpper(symbols[0])
		} else {
			t.Pair = strings.ToLower(symbols[0] + pairSplitter + symbols[1])
			t.Symbol = strings.ToLower(symbols[0])
		}
	}
	return nil
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

type FilterPointExtended struct {
	FilterPoint FilterPoint
	// Pools and pairs of the filter point's underlying trades.
	Pools         []Pool
	Pairs         []ExchangePair
	TradesCount   int32
	StatusMessage string
	StatusCode    int32
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
	Deleted            bool
	Draft              bool
	FeederSelection    string
	Expired            bool
	ExpiredDate        time.Time
	ExpiringDate       time.Time
	LastOracleUpdate   time.Time
	CustomerID         string
	Billable           bool
	Name               string
	Ecosystem          bool
}

func (e *OracleConfig) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *OracleConfig) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

type OracleUpdate struct {
	OracleAddress     string
	TransactionHash   string
	TransactionCost   string
	AssetKey          string
	AssetPrice        string
	UpdateBlock       uint64
	UpdateFrom        string
	FromBalance       string
	GasCost           string
	GasUsed           float64
	ChainID           string
	UpdateTime        time.Time
	CreationBlock     uint64
	CreationBlockTime time.Time
}

type FeedUpdates struct {
	Day         time.Time `json:"Day"`
	UpdateCount int32     `json:"UpdateCount"`
	GasUsed     float64   `json:"GasUsed"`
}
