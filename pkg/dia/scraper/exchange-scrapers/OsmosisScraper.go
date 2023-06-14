package scrapers

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	stdtypes "github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v3/modules/core/types"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-resty/resty/v2"
	gammtypes "github.com/osmosis-labs/osmosis/v6/x/gamm/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v6/x/lockup/types"
	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tendermint "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	osmosisRefreshDelay = time.Second * 30 * 1
)

type OsmosisEncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

type OsmosisConfig struct {
	Bech32AddrPrefix  string
	Bech32ValPrefix   string
	Bech32PkPrefix    string
	Bech32PkValPrefix string
	RpcURL            string
	Encoding          *OsmosisEncodingConfig
}

// Contains info about a transaction log event key/val attribute
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Contains info about an attribute value keyed by attribute type
type ValueByAttribute map[string]string

// Contains info about transaction events keyed by message index
type EventsByMsgIndex map[string]AttributesByEvent

// Contains info about a transaction log event
type Event struct {
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

// Contains info about event attributes keyed by event type
type AttributesByEvent map[string]ValueByAttribute

type OsmosisScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock            sync.RWMutex
	error                error
	closed               bool
	pairScrapers         map[string]*OsmosisPairScraper // pc.ExchangePair -> pairScraperSet
	wsClient             *tendermint.WSClient
	rpcClient            *resty.Client
	encoding             *OsmosisEncodingConfig
	ticker               *time.Ticker
	exchangeName         string
	chanTrades           chan *dia.Trade
	db                   *models.RelDB
	blockTimestampsCache map[int64]*time.Time
}

// NewOsmosisScraper returns a new OsmosisScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewOsmosisScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *OsmosisScraper {
	encoding := NewOsmosisEncoding()

	cfg := &OsmosisConfig{
		Bech32AddrPrefix:  "osmo",
		Bech32PkPrefix:    "osmopub",
		Bech32ValPrefix:   "osmovaloper",
		Bech32PkValPrefix: "osmovalpub",
		Encoding:          encoding,
		RpcURL:            utils.Getenv("OSMOSIS_RPC_URL", ""),
	}
	wsClient, err := NewWsClient(*cfg)
	if err != nil {
		panic("failed to create ws client")
	}
	rpcClient, err := NewRpcClient(*cfg)
	if err != nil {
		panic("failed to create rpc client")
	}
	s := &OsmosisScraper{
		shutdown:             make(chan nothing),
		shutdownDone:         make(chan nothing),
		pairScrapers:         make(map[string]*OsmosisPairScraper),
		blockTimestampsCache: make(map[int64]*time.Time),
		wsClient:             wsClient,
		rpcClient:            rpcClient,
		ticker:               time.NewTicker(osmosisRefreshDelay),
		exchangeName:         exchange.Name,
		encoding:             encoding,
		error:                nil,
		chanTrades:           make(chan *dia.Trade),
		db:                   relDB,
	}
	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *OsmosisScraper) mainLoop() {
	isWsRunning := s.wsClient.IsRunning()
	if !isWsRunning {
		s.Start()
	}
	for {
		select {
		case <-s.shutdown: // user requested shutdown
			log.Printf("OsmosisScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *OsmosisScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	s.wsClient.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *OsmosisScraper) Close() error {
	if s.closed {
		return errors.New("OsmosisScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// OsmosisPairScraper implements PairScraper for Osmosis
type OsmosisPairScraper struct {
	parent     *OsmosisScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *OsmosisScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("OsmosisScraper: Call ScrapePair on closed scraper")
	}
	ps := &OsmosisPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0, //TODO FIX to figure out the last we got...
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *OsmosisScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *OsmosisScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return []dia.ExchangePair{}, errors.New("FetchAvailablePairs() not implemented")
}

// NormalizePair accounts for the par
func (ps *OsmosisScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *OsmosisScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *OsmosisPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *OsmosisPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *OsmosisPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func NewWsClient(conf OsmosisConfig) (*tendermint.WSClient, error) {
	sdk.GetConfig().SetBech32PrefixForAccount(conf.Bech32AddrPrefix, conf.Bech32PkPrefix)
	sdk.GetConfig().SetBech32PrefixForValidator(conf.Bech32ValPrefix, conf.Bech32PkValPrefix)

	client, err := tendermint.NewWS(conf.RpcURL, "/websocket")
	if err != nil {
		log.Fatal("failed to create websocket client: ", err)
	}
	client.Dialer = net.Dial
	return client, nil
}

func NewRpcClient(conf OsmosisConfig) (*resty.Client, error) {
	sdk.GetConfig().SetBech32PrefixForAccount(conf.Bech32AddrPrefix, conf.Bech32PkPrefix)
	sdk.GetConfig().SetBech32PrefixForValidator(conf.Bech32ValPrefix, conf.Bech32PkValPrefix)

	headers := map[string]string{"Accept": "application/json"}
	client := resty.New().SetBaseURL(conf.RpcURL).SetHeaders(headers)
	return client, nil
}

func NewOsmosisEncoding() *OsmosisEncodingConfig {
	registry := codectypes.NewInterfaceRegistry()

	ibctransfertypes.RegisterInterfaces(registry)
	gammtypes.RegisterInterfaces(registry)
	lockuptypes.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)
	authztypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	distributiontypes.RegisterInterfaces(registry)
	govtypes.RegisterInterfaces(registry)
	ibccoretypes.RegisterInterfaces(registry)
	liquiditytypes.RegisterInterfaces(registry)
	stakingtypes.RegisterInterfaces(registry)
	stdtypes.RegisterInterfaces(registry)

	marshaler := codec.NewProtoCodec(registry)

	return &OsmosisEncodingConfig{
		InterfaceRegistry: registry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}

func (s *OsmosisScraper) Start() error {
	err := s.wsClient.Start()
	if err != nil {
		log.Warn("failed to start websocket client: ", err)
		return err
	}

	err = s.wsClient.Subscribe(context.Background(), tmtypes.EventQueryTx.String())
	if err != nil {
		log.Warn(err, "failed to subscribe to txs")
		return err
	}

	go s.listen()

	return nil
}

func (s *OsmosisScraper) listen() {
	for r := range s.wsClient.ResponsesCh {
		if r.Error != nil {
			// resubscribe if subscription is cancelled by the server for reason:
			// client is not pulling messages fast enough
			// experimental rpc config available to help mitigate this issue:
			// https://github.com/tendermint/tendermint/blob/main/config/config.go#L373
			if r.Error.Code == -32000 {
				err := s.wsClient.UnsubscribeAll(context.Background())
				if err != nil {
					log.Fatal(err, "failed to unsubscribe from all subscriptions")
				}

				err = s.wsClient.Subscribe(context.Background(), tmtypes.EventQueryTx.String())
				if err != nil {
					log.Fatal(err, "failed to subscribe to txs")
				}

				continue
			}

			log.Error(r.Error.Error())
			continue
		}

		result := &coretypes.ResultEvent{}
		if err := tmjson.Unmarshal(r.Result, result); err != nil {
			log.Errorf("failed to unmarshal tx message: %v", err)
			continue
		}

		if result.Data != nil {
			switch result.Data.(type) {
			case tmtypes.EventDataTx:
				go s.handleTx(result.Data.(tmtypes.EventDataTx))
			default:
				fmt.Printf("unsupported result type: %T", result.Data)
			}
		}
	}

	// if reconnect fails, ResponsesCh is closed
	log.Fatal("websocket client connection closed by server")
}

func (s *OsmosisScraper) handleTx(tx tmtypes.EventDataTx) {
	decodedTx, err := DecodeTx(*s.encoding, tx.Tx)
	if err != nil {
		// unsupported tx
		return
	}

	txid := fmt.Sprintf("%X", sha256.Sum256(tx.Tx))
	events := ParseEvents(tx.Result.Log)
	messages := ParseMessages(decodedTx.GetMsgs(), events)
	// messages var is empty for any types other than `*gammtypes.MsgSwapExactAmountIn`
	if len(messages) > 0 {
		quoteToken, err := s.db.GetAsset(messages[0].Token, "Osmosis")
		if err != nil {
			log.Error(err, ", failed to get asset: ", messages[0].Token)
			return
		}
		baseToken, err := s.db.GetAsset(messages[1].Token, "Osmosis")
		if err != nil {
			log.Error(err, ", failed to get asset: ", messages[1].Token)
			return
		}
		volumeOut, err := strconv.ParseFloat(messages[0].Value, 64)
		if err != nil {
			log.Error(err, ", failed to parse volume of: ", txid)
			return
		}
		volumeIn, err := strconv.ParseFloat(messages[1].Value, 64)
		if err != nil {
			log.Error(err, ", failed to parse volume of: ", txid)
			return
		}

		if volumeOut == 0 {
			return
		}
		price := (volumeIn / math.Pow(10, float64(baseToken.Decimals))) / (volumeOut / math.Pow(10, float64(quoteToken.Decimals)))
		timestamp := s.blockTimestampsCache[tx.Height]
		if timestamp == nil {
			// get timestamp from rpc
			rpcTimestamp, err := s.GetBlock(int(tx.Height))
			if err != nil {
				log.Error(err, ", failed to get block timestampfor: ", txid)
				return
			}
			s.blockTimestampsCache[tx.Height] = rpcTimestamp
			timestamp = rpcTimestamp
		}
		t := &dia.Trade{
			Symbol:         quoteToken.Symbol,
			Pair:           quoteToken.Symbol + "-" + baseToken.Symbol,
			Volume:         volumeOut / math.Pow(10, float64(quoteToken.Decimals)),
			Price:          price,
			Time:           *timestamp,
			ForeignTradeID: txid,
			Source:         dia.OsmosisExchange,
			BaseToken:      baseToken,
			QuoteToken:     quoteToken,
			VerifiedPair:   true,
		}
		log.Info("New Trade: ", t)
		s.chanTrades <- t
	}
}

// DecodeTx will attempt to decode a raw transaction in the form of
// a base64 encoded string or a protobuf encoded byte slice
func DecodeTx(encoding OsmosisEncodingConfig, rawTx interface{}) (sdk.Tx, error) {
	var txBytes []byte

	switch rawTx := rawTx.(type) {
	case string:
		var err error

		txBytes, err = base64.StdEncoding.DecodeString(rawTx)
		if err != nil {
			return nil, fmt.Errorf("error decoding transaction from base64: %s", err)
		}
	case []byte:
		txBytes = rawTx
	case tmtypes.Tx:
		txBytes = rawTx
	default:
		return nil, fmt.Errorf("rawTx must be string or []byte")
	}

	tx, err := encoding.TxConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding transaction from protobuf: %s", err)
	}

	return tx, nil
}

func ParseEvents(log string) EventsByMsgIndex {
	events := make(EventsByMsgIndex)

	logs, err := sdk.ParseABCILogs(log)
	if err != nil {
		// transaction error logs are not in json format and will fail to parse
		// return error event with the log message
		events["0"] = AttributesByEvent{"error": ValueByAttribute{"message": log}}
		return events
	}

	for _, l := range logs {
		msgIndex := strconv.Itoa(int(l.GetMsgIndex()))
		events[msgIndex] = make(AttributesByEvent)

		for _, e := range l.GetEvents() {
			attributes := make(ValueByAttribute)
			for _, a := range e.Attributes {
				attributes[a.Key] = a.Value
			}

			events[msgIndex][e.Type] = attributes
		}
	}

	return events
}

// Contains info about a transaction message
type Message struct {
	Value string
	Token string
}

// ParseMessages will parse any osmosis or cosmos-sdk message types
func ParseMessages(msgs []sdk.Msg, events EventsByMsgIndex) []Message {
	messages := []Message{}

	if _, ok := events["0"]["error"]; ok {
		return messages
	}

	for i, msg := range msgs {
		switch v := msg.(type) {
		case *gammtypes.MsgSwapExactAmountIn:
			swappedTokensOut := events[strconv.Itoa(i)]["token_swapped"]["tokens_out"]

			tokenOut, err := sdk.ParseCoinNormalized(swappedTokensOut)
			if err != nil && swappedTokensOut != "" {
				log.Error(err)
			}

			msgs := []Message{
				// token in (sell)
				{
					Token: (&v.TokenIn).Denom,
					Value: (&v.TokenIn).Amount.String(),
				},
				// token out (buy)
				{
					Token: (&tokenOut).Denom,
					Value: (&tokenOut).Amount.String(),
				},
			}
			messages = append(messages, msgs...)
		default:
		}
	}

	return messages
}

func (s *OsmosisScraper) GetBlock(height int) (*time.Time, error) {
	res := &rpctypes.RPCResponse{}
	_, err := s.rpcClient.R().SetResult(res).SetError(res).SetQueryParam("height", strconv.Itoa(height)).Get("/block")
	if err != nil {
		return nil, err
	}
	result := &coretypes.ResultBlock{}
	if err := tmjson.Unmarshal(res.Result, result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal block result: %v: %s", res.Result, res.Error.Error())
	}

	return &result.Block.Time, nil
}
