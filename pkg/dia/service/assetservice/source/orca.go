package source

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	orcaWhirlpoolIdlBind "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/orca/whirlpool"
	"github.com/diadata-org/diadata/pkg/utils"
	bin "github.com/gagliardetto/binary"
	tokenmetadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/programs/tokenregistry"
	"github.com/gagliardetto/solana-go/rpc"
)

const (
	orcaSolanaHttpEndpoint = "https://rpc.ankr.com/solana"
)

// The scraper object for Orca.
type OrcaAssetSource struct {
	RestClient   *rpc.Client
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
}

// Returns a new Orca asset scraper.
func NewOrcaScraper(exchange dia.Exchange) *OrcaAssetSource {

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		pas          *OrcaAssetSource
	)

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	restClient := rpc.New(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", orcaSolanaHttpEndpoint))

	pas = &OrcaAssetSource{
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		RestClient:   restClient,
	}
	go func() {
		pas.loadMarketsMetadata()
	}()
	return pas
}

// Load markets and tokens metadata
func (pas *OrcaAssetSource) loadMarketsMetadata() {
	log.Infof("loading initial data from pools ...")
	start := time.Now()
	err := pas.loadMarketPools()
	if err != nil {
		log.Error(err)
	}
	log.Infof("loaded legacy-pools data in %.1fs", time.Since(start).Seconds())
	start = time.Now()
	err = pas.loadMarketWhirlpools()
	if err != nil {
		log.Error(err)
	}
	log.Infof("loaded whirlpools data in %.1fs", time.Since(start).Seconds())
	pas.doneChannel <- true
}

// Get Orca market whirlpools
func (pas *OrcaAssetSource) loadMarketWhirlpools() (err error) {
	resp, err := pas.RestClient.GetProgramAccountsWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(scrapers.OrcaProgWhirlpoolAddr),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					DataSize: scrapers.OrcaProgWhirlpoolAccountDataSize,
				},
			},
		},
	)
	if err != nil {
		return
	}
	if resp == nil {
		return fmt.Errorf("program account not found")
	}
	log.Infof("discovered %d accounts in whirlpool program, retrieving metadata ...", len(resp))
	cachedAssets := make(map[string]bool)
	for _, progAcc := range resp {
		acct := progAcc.Account
		pubKey := progAcc.Pubkey.String()
		if acct.Owner.String() == scrapers.OrcaProgWhirlpoolAddr {
			d := bin.NewBorshDecoder(acct.Data.GetBinary())
			var w orcaWhirlpoolIdlBind.Whirlpool
			err = d.Decode(&w)
			if err != nil {
				return err
			}
			// Blacklist XXX/USDC, ATLAS/USDC, SHIB/USDC
			if pubKey == "FfBeru58Q7hjqHq9T2Trw1BeyjE1YwHsx9MivKUwoTLQ" || pubKey == "9vqFu6v9CcVDaSx2oRD3jo8H5gqkE2urYQgpT16V1BTa" || pubKey == "DahhciLA89UkZoqrqVWL2nojwPLmSVkXQGTiEhAtkaFa" {
				continue
			}
			if w.WhirlpoolsConfig.String() == scrapers.OrcaProgWhirlpoolConfigAddr {
				// Get token A mint data and metadata
				if _, ok := cachedAssets[w.TokenMintA.String()]; !ok {
					token, err := pas.getAssetData(w.TokenMintA.String())
					if err != nil {
						return err
					}

					pas.Asset() <- token
					cachedAssets[token.Address] = true
				}

				// Get token B mint data and metadata
				if _, ok := cachedAssets[w.TokenMintB.String()]; !ok {
					token, err := pas.getAssetData(w.TokenMintB.String())
					if err != nil {
						return err
					}

					pas.Asset() <- token
					cachedAssets[token.Address] = true
				}
			}
		}
	}
	return
}

func (pas *OrcaAssetSource) getAssetData(mintAddress string) (asset dia.Asset, err error) {
	if mintData, err := pas.getTokenMintData(mintAddress); err == nil {
		asset.Decimals = mintData.Decimals
	} else {
		return asset, err
	}
	if metadata, err := pas.getTokenMetadata(mintAddress); err != nil {
		hardcodedTokenMeta := scrapers.GetOrcaTokensMetadata()
		if v, ok := hardcodedTokenMeta[mintAddress]; ok {
			asset.Symbol = v.(scrapers.OrcaTokenMetadata).GetSymbol()
			asset.Name = v.(scrapers.OrcaTokenMetadata).GetName()
		} else {
			log.Warnf("cannot found token metadata for %s: %s", mintAddress, err)
			return asset, nil
		}
	} else {
		asset.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
		asset.Name = strings.TrimRight(metadata.Data.Name, "\x00")
	}
	asset.Address = mintAddress
	asset.Blockchain = "Solana"
	return asset, nil
}

// Get Solana token mint data
func (pas *OrcaAssetSource) getTokenMintData(account string) (mint token.Mint, err error) {
	resp, err := pas.RestClient.GetAccountInfoWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(account),
		&rpc.GetAccountInfoOpts{},
	)
	if err != nil {
		return
	}
	d := bin.NewBorshDecoder(resp.Value.Data.GetBinary())
	err = d.Decode(&mint)
	if err != nil {
		return
	}
	return
}

// Get Solana token metadata
func (pas *OrcaAssetSource) getTokenMetadata(account string) (metadata tokenmetadata.Metadata, err error) {
	accMint := solana.MustPublicKeyFromBase58(account)
	tMeta, err := tokenregistry.GetTokenRegistryEntry(context.TODO(), pas.RestClient, accMint)
	if err != nil {
		metaAddress, _, err := solana.FindTokenMetadataAddress(accMint)
		if err != nil {
			return metadata, err
		}
		resp, err := pas.RestClient.GetAccountInfo(
			context.TODO(),
			metaAddress,
		)
		if err != nil {
			return metadata, err
		}
		d := bin.NewBorshDecoder(resp.Value.Data.GetBinary())
		err = d.Decode(&metadata)
		if err != nil {
			return metadata, err
		}
		return metadata, nil
	}
	return tokenmetadata.Metadata{Data: tokenmetadata.Data{Symbol: tMeta.Symbol.String()}}, nil
}

// Get Orca market legacy pools
func (pas *OrcaAssetSource) loadMarketPools() (err error) {
	return
}

func (pas *OrcaAssetSource) Asset() chan dia.Asset {
	return pas.assetChannel
}

func (pas *OrcaAssetSource) Done() chan bool {
	return pas.doneChannel
}
