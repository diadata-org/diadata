package liquidityscrapers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	orcaWhirlpoolIdlBind "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/orca/whirlpool"
	models "github.com/diadata-org/diadata/pkg/model"
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

type OrcaScraper struct {
	blockchain   string
	exchangeName string
	RestClient   *rpc.Client
	datastore    *models.DB
	poolChannel  chan dia.Pool
	doneChannel  chan bool
}

func NewOrcaScraper(exchange dia.Exchange, datastore *models.DB) *OrcaScraper {

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	restClient := rpc.New(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", orcaSolanaHttpEndpoint))

	scraper := &OrcaScraper{
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
		RestClient:   restClient,
		datastore:    datastore,
		poolChannel:  make(chan dia.Pool),
		doneChannel:  make(chan bool),
	}

	go func() {
		err := scraper.loadMarketsMetadata()
		if err != nil {
			log.Error(err)
		}
		scraper.doneChannel <- true
	}()

	return scraper
}

// Load markets and tokens metadata
func (s *OrcaScraper) loadMarketsMetadata() (err error) {
	log.Infof("loading initial data from pools ...")
	start := time.Now()
	err = s.loadMarketPools()
	if err != nil {
		return
	}
	log.Infof("loaded legacy-pools data in %.1fs", time.Since(start).Seconds())
	start = time.Now()
	err = s.loadMarketWhirlpools()
	if err != nil {
		return
	}
	log.Infof("loaded whirlpools data in %.1fs", time.Since(start).Seconds())
	return
}

// Get Orca market legacy pools
func (s *OrcaScraper) loadMarketPools() (err error) {
	return
}

// Get Orca market whirlpools
func (s *OrcaScraper) loadMarketWhirlpools() (err error) {
	hardcodedTokenMeta := scrapers.GetOrcaTokensMetadata()
	resp, err := s.RestClient.GetProgramAccountsWithOpts(
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
				var tokenA, tokenB dia.Asset

				// Get token A mint data and metadata
				if mintData, err := s.getTokenMintData(w.TokenMintA.String()); err == nil {
					if mintData.IsInitialized {
						tokenA.Decimals = mintData.Decimals
					}
				} else {
					return err
				}
				if metadata, err := s.getTokenMetadata(w.TokenMintA.String()); err != nil {
					if v, ok := hardcodedTokenMeta[w.TokenMintA.String()]; ok {
						tokenA.Symbol = v.(scrapers.OrcaTokenMetadata).GetSymbol()
						tokenA.Name = v.(scrapers.OrcaTokenMetadata).GetName()
					} else {
						log.Warnf("cannot found token metadata for %s: %s", w.TokenMintA.String(), err)
						continue
					}
				} else {
					tokenA.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
					tokenA.Name = strings.TrimRight(metadata.Data.Name, "\x00")
				}
				tokenA.Address = w.TokenMintA.String()
				tokenA.Blockchain = "Solana"

				// Get token B mint data and metadata
				if mintData, err := s.getTokenMintData(w.TokenMintB.String()); err == nil {
					tokenB.Decimals = mintData.Decimals
				} else {
					return err
				}
				if metadata, err := s.getTokenMetadata(w.TokenMintB.String()); err != nil {
					if v, ok := hardcodedTokenMeta[w.TokenMintB.String()]; ok {
						tokenB.Symbol = v.(scrapers.OrcaTokenMetadata).GetSymbol()
						tokenB.Name = v.(scrapers.OrcaTokenMetadata).GetName()
					} else {
						log.Warnf("cannot found token metadata for %s: %s", w.TokenMintB.String(), err)
						continue
					}
				} else {
					tokenB.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
					tokenB.Name = strings.TrimRight(metadata.Data.Name, "\x00")
				}
				tokenB.Address = w.TokenMintB.String()
				tokenB.Blockchain = "Solana"

				tokenABalance, err := s.RestClient.GetTokenAccountBalance(context.TODO(), w.TokenVaultA, rpc.CommitmentFinalized)
				if err != nil {
					return fmt.Errorf("GetTokenAccountBalance: %s", err.Error())
				}
				tokenBBalance, err := s.RestClient.GetTokenAccountBalance(context.TODO(), w.TokenVaultB, rpc.CommitmentFinalized)
				if err != nil {
					return fmt.Errorf("GetTokenAccountBalance: %s", err.Error())
				}

				var pool dia.Pool
				log.Info("pool asset: ", tokenA)
				pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
					Asset:  tokenA,
					Volume: *tokenABalance.Value.UiAmount,
				})
				log.Info("pool asset: ", tokenB)
				pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
					Asset:  tokenB,
					Volume: *tokenBBalance.Value.UiAmount,
				})

				// Determine USD liquidity.
				if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
					s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
				}

				pool.Exchange = dia.Exchange{Name: s.exchangeName}
				pool.Blockchain = dia.BlockChain{Name: s.blockchain}
				pool.Address = pubKey
				pool.Time = time.Now()
				s.Pool() <- pool
			}
		}
	}
	return
}

// Get Solana token mint data
func (s *OrcaScraper) getTokenMintData(account string) (mint token.Mint, err error) {
	resp, err := s.RestClient.GetAccountInfoWithOpts(
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
func (s *OrcaScraper) getTokenMetadata(account string) (metadata tokenmetadata.Metadata, err error) {
	accMint := solana.MustPublicKeyFromBase58(account)
	tMeta, err := tokenregistry.GetTokenRegistryEntry(context.TODO(), s.RestClient, accMint)
	if err != nil {
		metaAddress, _, err := solana.FindTokenMetadataAddress(accMint)
		if err != nil {
			return metadata, err
		}
		resp, err := s.RestClient.GetAccountInfo(
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

func (scraper *OrcaScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *OrcaScraper) Done() chan bool {
	return scraper.doneChannel
}
