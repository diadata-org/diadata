package liquidityscrapers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	gammtypes "github.com/osmosis-labs/osmosis/v6/x/gamm/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type GRPCClient struct {
	ctx      context.Context
	grpcConn *grpc.ClientConn
	gamm     gammtypes.QueryClient
	bank     banktypes.QueryClient
}

type OsmosisScraper struct {
	blockchain   string
	exchangeName string
	grpcClient   *GRPCClient
	relDb        *models.RelDB
	poolChannel  chan dia.Pool
	doneChannel  chan bool
}

func NewOsmosisScraper(exchange dia.Exchange, relDb *models.RelDB) *OsmosisScraper {
	log.Infof("init rpc client for %s", exchange.BlockChain.Name)
	// encoding := scrapers.NewOsmosisEncoding()

	cfg := &scrapers.OsmosisConfig{
		Bech32AddrPrefix:  "osmo",
		Bech32PkPrefix:    "osmopub",
		Bech32ValPrefix:   "osmovaloper",
		Bech32PkValPrefix: "osmovalpub",
		Encoding:          nil,
		RpcURL:            utils.Getenv("OSMOSIS_RPC_URL", ""),
	}

	grpcClient, err := NewGRPCClient(cfg)
	if err != nil {
		panic("failed to create rpc client")
	}

	scraper := &OsmosisScraper{
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
		grpcClient:   grpcClient,
		relDb:        relDb,
		poolChannel:  make(chan dia.Pool),
		doneChannel:  make(chan bool),
	}

	go func() {
		err := scraper.loadMarketsMetadata()
		if err != nil {
			log.Error(err)
		}
	}()

	return scraper
}

// Load markets and tokens metadata
func (s *OsmosisScraper) loadMarketsMetadata() (err error) {
	log.Infof("loading initial data from pools ...")
	start := time.Now()
	s.loadMarketPools()
	log.Infof("loaded pools data in %.1fs", time.Since(start).Seconds())
	return
}

// Get Osmosis market pools
func (s *OsmosisScraper) loadMarketPools() {
	s.loadMarketPoolsPage(nil)
}

func (s *OsmosisScraper) loadMarketPoolsPage(page *[]byte) {
	reqPage := &gammtypes.QueryPoolsRequest{Pagination: &query.PageRequest{}}
	if page != nil {
		reqPage.Pagination.Key = *page
	}
	res, err := s.grpcClient.gamm.Pools(s.grpcClient.ctx, reqPage)
	if err != nil {
		panic(err)
	}

	go func() {
		if res.Pagination.NextKey != nil {
			s.loadMarketPoolsPage(&res.Pagination.NextKey)
		}
	}()

	for _, pool := range res.Pools {
		p := &gammtypes.Pool{}
		err = p.Unmarshal(pool.Value)
		if err != nil {
			log.Warn("error parsing pool protobuf msg with type: ", pool.TypeUrl)
			continue
		}
		s.HandlePool(p)
	}
	if res.Pagination.NextKey == nil {
		defer func() {
			s.grpcClient.grpcConn.Close()
			s.Done() <- true
		}()
	}

}

func (s *OsmosisScraper) HandlePool(pool *gammtypes.Pool) {
	diaPool := &dia.Pool{}
	for idx, asset := range pool.PoolAssets {
		diaAsset, err := s.relDb.GetAsset(asset.Token.Denom, "Osmosis")
		if err != nil {
			// db didn't have the asset, searching bank sdk
			diaAsset, err = s.GetAssetMetadata(asset)
			if err != nil {
				// we don't have any info on this asset
				log.Warn("Unknown asset: ", asset.Token.Denom)
				return
			}
		}

		poolAsset := dia.Asset{
			Address:    asset.Token.Denom,
			Blockchain: "Osmosis",
			Decimals:   diaAsset.Decimals,
			Symbol:     diaAsset.Symbol,
			Name:       diaAsset.Name,
		}
		log.Info("pool asset: ", poolAsset)

		diaPool.Assetvolumes = append(diaPool.Assetvolumes, dia.AssetVolume{
			Asset:  poolAsset,
			Volume: asset.Token.Amount.ToDec().MustFloat64(),
			Index:  uint8(idx),
		})
	}

	diaPool.Exchange = dia.Exchange{Name: s.exchangeName}
	diaPool.Blockchain = dia.BlockChain{Name: s.blockchain}
	diaPool.Address = pool.Address

	diaPool.Time = time.Now()

	s.Pool() <- *diaPool
}

func (s *OsmosisScraper) GetAssetMetadata(asset gammtypes.PoolAsset) (dia.Asset, error) {
	a := &dia.Asset{}
	res, err := s.grpcClient.bank.DenomMetadata(
		s.grpcClient.ctx,
		&banktypes.QueryDenomMetadataRequest{Denom: asset.Token.Denom},
	)
	if err != nil {
		log.Warn("can't get denom metadata for ", asset.Token.Denom, err)
		return *a, err
	}
	a.Name = res.Metadata.Display
	a.Symbol = res.Metadata.Symbol
	// fallback to Display if Metadata.Symbol is empty
	if a.Symbol == "" {
		a.Symbol = res.Metadata.Display
	}
	denomUnitsMap := make(map[string]uint8)
	for _, du := range res.Metadata.DenomUnits {
		denomUnitsMap[du.Denom] = uint8(du.Exponent)
	}
	a.Decimals = denomUnitsMap[res.Metadata.Display]
	a.Address = res.Metadata.Base
	return *a, nil
}

func (scraper *OsmosisScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *OsmosisScraper) Done() chan bool {
	return scraper.doneChannel
}

func NewGRPCClient(conf *scrapers.OsmosisConfig) (*GRPCClient, error) {
	sdk.GetConfig().SetBech32PrefixForAccount(conf.Bech32AddrPrefix, conf.Bech32PkPrefix)
	md := metadata.Pairs()
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcURL, err := url.Parse(conf.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GRPCURL: %s", conf.RpcURL)
	}
	grpcConn, err := grpc.DialContext(
		context.Background(),
		grpcURL.String(),
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
	)
	if err != nil {
		if grpcConn != nil {
			grpcConn.Close()
		}
		return nil, fmt.Errorf("unable to connect to grpcConn: %s", err)
	}
	gamm := gammtypes.NewQueryClient(grpcConn)
	bank := banktypes.NewQueryClient(grpcConn)

	c := &GRPCClient{
		ctx:      ctx,
		grpcConn: grpcConn,
		gamm:     gamm,
		bank:     bank,
	}
	return c, nil
}
