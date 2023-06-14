package source

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type GRPCClient struct {
	ctx      context.Context
	grpcConn *grpc.ClientConn
	bank     banktypes.QueryClient
	ibc      ibctransfertypes.QueryClient
}

// The scraper object for Osmosis.
type OsmosisAssetSource struct {
	GRPCClient   *GRPCClient
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	assetsMap    map[string]*OsmosisAsset
}

type OsmosisAssetResponse struct {
	Assets *[]OsmosisAsset
}

type OsmosisAsset struct {
	Base        string // base_denom
	Name        string // name of denom
	Display     string // human name
	Symbol      string // symbol of the denom
	Decimals    uint8  // to be loaded from Denom_units
	Denom_units *[]banktypes.DenomUnit
}

// Returns a new Osmosis asset scraper.
func NewOsmosisScraper(exchange dia.Exchange) *OsmosisAssetSource {
	assetChannel := make(chan dia.Asset)
	doneChannel := make(chan bool)

	assetsMap, err := GetAssetsJson()
	if err != nil {
		log.Fatal("failed to download assets json file: ", err)
	}

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
		log.Fatal("failed to create GRPC client: ", err)
	}
	oas := &OsmosisAssetSource{
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		GRPCClient:   grpcClient,
		assetsMap:    assetsMap,
	}

	go oas.loadMarketsMetadata()
	return oas
}

func (oas *OsmosisAssetSource) loadMarketsMetadata() {
	oas.getAssetsTotalSupply(nil)
}

func (oas *OsmosisAssetSource) getAssetsTotalSupply(page *[]byte) {
	reqPage := &banktypes.QueryTotalSupplyRequest{Pagination: &query.PageRequest{}}
	if page != nil {
		reqPage.Pagination.Key = *page
	}
	res, err := oas.GRPCClient.bank.TotalSupply(oas.GRPCClient.ctx, reqPage)
	if err != nil {
		log.Warn("failed to get data ", err)
		return
	}

	go func() {
		if res.Pagination.NextKey != nil {
			oas.getAssetsTotalSupply(&res.Pagination.NextKey)
		}
	}()

	for _, coin := range res.Supply {
		oas.getAssetInfo(coin)
	}

	if res.Pagination.NextKey == nil {
		defer func() {
			oas.GRPCClient.grpcConn.Close()
			oas.Done() <- true
		}()
	}
}

// Get osmosis/ibc asset metadata
func (oas *OsmosisAssetSource) getAssetInfo(coin sdk.Coin) {
	a := &dia.Asset{
		Blockchain: "Osmosis",
	}
	// ibc denoms start with `ibc/`
	splittedDenom := strings.Split(coin.Denom, "/")
	switch splittedDenom[0] {
	case "ibc":
		{
			// TODO: once https://github.com/cosmos/ibc-go/pull/3104 gets merged
			// we can use the following codes to get the metadata:
			//
			// 		res, err := oas.GRPCClient.ibc.DenomTrace(
			// 			oas.GRPCClient.ctx,
			// 			&ibctransfertypes.QueryDenomTraceRequest{Hash: splittedDenom[1]},
			// 		)
			// 		if err != nil {
			// 			log.Warn("can't get denom metadata for ", coin.Denom)
			// 			return
			// 		}

			osmosisAsset := oas.assetsMap[coin.Denom]
			if osmosisAsset == nil {
				log.Warn("Couldn't find any info about ", coin.Denom)
				return
			}
			a.Name = osmosisAsset.Display
			a.Symbol = osmosisAsset.Symbol
			a.Address = coin.Denom // keeping base denom as address
			a.Decimals = osmosisAsset.Decimals
			oas.Asset() <- *a
			return
		}
	case "factory":
	case "gamm":
		{
			// we don't want factory/gamm assets in the db as they're not swappable.
			// https://docs.osmosis.zone/osmosis-core/modules/tokenfactory/
			// https://docs.osmosis.zone/osmosis-core/modules/gamm/
			return
		}
	default:
		{
			res, err := oas.GRPCClient.bank.DenomMetadata(
				oas.GRPCClient.ctx,
				&banktypes.QueryDenomMetadataRequest{Denom: coin.Denom},
			)
			if err != nil {
				log.Warn("can't get denom metadata for ", coin.Denom, err)
				return
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
			oas.Asset() <- *a
		}
	}
}

func (oas *OsmosisAssetSource) Asset() chan dia.Asset {
	return oas.assetChannel
}

func (oas *OsmosisAssetSource) Done() chan bool {
	return oas.doneChannel
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
	bank := banktypes.NewQueryClient(grpcConn)
	ibc := ibctransfertypes.NewQueryClient(grpcConn)

	c := &GRPCClient{
		ctx:      ctx,
		grpcConn: grpcConn,
		bank:     bank,
		ibc:      ibc,
	}

	return c, nil
}

func GetAssetsJson() (map[string]*OsmosisAsset, error) {
	url := "https://raw.githubusercontent.com/osmosis-labs/assetlists/main/osmosis-1/osmosis-1.assetlist.json"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var assetsResponse OsmosisAssetResponse
	err = json.Unmarshal([]byte(data), &assetsResponse)
	if err != nil {
		return nil, err
	}
	assetsMap := make(map[string]*OsmosisAsset)
	for _, asset := range *assetsResponse.Assets {
		denomUnitsMap := make(map[string]uint32)
		for _, unit := range *asset.Denom_units {
			denomUnitsMap[unit.Denom] = unit.Exponent
		}
		assetCopy := asset
		assetCopy.Decimals = uint8(denomUnitsMap[asset.Display])
		assetsMap[asset.Base] = &assetCopy
	}
	return assetsMap, nil
}
