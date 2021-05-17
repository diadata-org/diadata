package dia

import (
	"encoding/json"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	log "github.com/sirupsen/logrus"
)

const (
	Diadata                                 = "diadata.org"
	PROOF_OF_STAKE    VerificationMechanism = "pos"
	PROOF_OF_WORK     VerificationMechanism = "pow"
	BITCOIN                                 = "Bitcoin"
	ETHEREUM                                = "Ethereum"
	BINANCESMARTCHAIN                       = "BinanceSmartChain"
)

type VerificationMechanism string

// NFTClass is the container for an nft class defined by
// a contract (address) on a blockchain.
type NFTClass struct {
	Address      common.Address
	Symbol       string
	Name         string
	Blockchain   string
	ContractType string
	Category     string
}

// MarshalBinary for DefiProtocolState
func (nc *NFTClass) MarshalBinary() ([]byte, error) {
	return json.Marshal(nc)
}

// UnmarshalBinary for DefiProtocolState
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
	TokenID        uint64
	CreationTime   time.Time
	CreatorAddress common.Address
	URI            string
	Attributes     []byte
}

// MarshalBinary for DefiProtocolState
func (n *NFT) MarshalBinary() ([]byte, error) {
	return json.Marshal(n)
}

// UnmarshalBinary for DefiProtocolState
func (n *NFT) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	return nil
}

type NFTTrade struct {
	NFT         NFT
	BlockNumber *big.Int
	PriceUSD    float64
	FromAddress common.Address
	ToAddress   common.Address
	Exchange    string
}

// MarshalBinary for DefiProtocolState
func (ns *NFTTrade) MarshalBinary() ([]byte, error) {
	return json.Marshal(ns)
}

// UnmarshalBinary for DefiProtocolState
func (ns *NFTTrade) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &ns); err != nil {
		return err
	}
	return nil
}

type Exchange struct {
	Name          string         `json:"Name"`
	Centralized   bool           `json:"Centralized"`
	Contract      common.Address `json:"Contract"`
	BlockChain    BlockChain     `json:"BlockChain"`
	WatchdogDelay int            `json:"WatchdogDelay"`
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
	Name                  string                `json:"Name"`
	GenesisDate           time.Time             `json:"GenesisDate"`
	NativeToken           string                `json:"NativeToken"`
	VerificationMechanism VerificationMechanism `json:"VerificationMechanism"`
}

// Pair substitues the old dia.Pair. It includes the new asset type.
type Pair struct {
	BaseToken  Asset
	QuoteToken Asset
}

// ForeignName returns the foreign name of the pair @p, i.e. the string Quotetoken-Basetoken
func (p *Pair) ForeignName() string {
	return p.QuoteToken.Symbol + "-" + p.BaseToken.Symbol
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

type ItinToken struct {
	Itin               string
	Symbol             string
	Label              string
	Url_website        string
	Coinmarketcap_url  string
	Coinmarketcap_slug string
}

type OptionType int

// signals if the option is call or a put
const (
	CallOption OptionType = iota + 1
	PutOption
)

type OptionOrderbookDatum struct {
	InstrumentName  string
	ObservationTime time.Time
	AskPrice        float64
	BidPrice        float64
	AskSize         float64
	BidSize         float64
}

type OptionMeta struct {
	InstrumentName string
	BaseCurrency   string
	ExpirationTime time.Time
	StrikePrice    float64
	OptionType     OptionType
}

type OptionMetaIndex struct {
	OptionMeta
	OptionOrderbookDatum
}

type OptionMetaForward struct {
	GeneralizedInstrumentName string
	StrikePrice               float64
	CallPrice                 float64
	PutPrice                  float64 // this, as well as the above is defined as the bid price at a given strike price
	ExpirationTime            time.Time
}

type CviDataPoint struct {
	Timestamp time.Time
	Value     float64
}

type DefiProtocol struct {
	Name                 string
	Address              string
	UnderlyingBlockchain string
	Token                string
}

type DefiProtocolState struct {
	TotalUSD  float64
	TotalETH  float64
	Timestamp time.Time
	Protocol  DefiProtocol
}

type DefiRate struct {
	Timestamp     time.Time
	LendingRate   float64
	BorrowingRate float64
	Asset         string
	Protocol      string
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
	Asset Asset
	Value float64
	Name  string
	Time  time.Time
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

// MarshalBinary for DefiProtocolState
func (e *DefiProtocolState) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiProtocolState
func (e *DefiProtocolState) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary for DefiRate
func (e *DefiRate) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiRate
func (e *DefiRate) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// MarshalBinary for DefiProtocol
func (e *DefiProtocol) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for DefiProtocol
func (e *DefiProtocol) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
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

// MarshalBinary -
func (e *ItinToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *ItinToken) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

func (e *OptionMeta) MarshalBinary() ([]byte, error) {
	basicOptionMeta := struct {
		InstrumentName string     `json:"instrumentname"`
		BaseCurrency   string     `json:"basecurrency"`
		ExpirationTime string     `json:"expirationtime"`
		StrikePrice    float64    `json:"strikeprice"`
		OptionType     OptionType `json:"optiontype"`
	}{
		InstrumentName: e.InstrumentName,
		BaseCurrency:   e.BaseCurrency,
		ExpirationTime: e.ExpirationTime.Format(time.RFC3339),
		StrikePrice:    e.StrikePrice,
		OptionType:     e.OptionType,
	}

	return json.Marshal(basicOptionMeta)
}

func (e *OptionMeta) UnmarshalBinary(data []byte) error {
	var rawStrings map[string]interface{}

	err := json.Unmarshal(data, &rawStrings)
	if err != nil {
		log.Error(err)
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "strikeprice" {
			e.StrikePrice = v.(float64)
		}
		if strings.ToLower(k) == "instrumentname" {
			e.InstrumentName = v.(string)
		}
		if strings.ToLower(k) == "basecurrency" {
			e.BaseCurrency = v.(string)
		}
		if strings.ToLower(k) == "optiontype" {
			if int(v.(float64)) == 2 {
				e.OptionType = PutOption
			} else {
				e.OptionType = CallOption
			}
		}
		if strings.ToLower(k) == "expirationtime" {
			t, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return err
			}
			e.ExpirationTime = t
		}
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
