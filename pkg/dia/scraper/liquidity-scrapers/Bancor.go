package liquidityscrapers

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/token"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/bancor/ConverterTypeFour"
	ConvertertypeOne "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/bancor/ConverterTypeOne"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/bancor/ConverterTypeThree"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/bancor/ConverterTypeZero"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BancorPoolScraper struct {
	RestClient   *ethclient.Client
	datastore    *models.DB
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	exchangeName string
	pool         string
}

type BancorPool struct {
	Reserves []struct {
		DltID   string `json:"dlt_id"`
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Balance struct {
			Usd string `json:"usd"`
		} `json:"balance"`
		Weight int `json:"weight"`
		Price  struct {
			Usd string `json:"usd"`
		} `json:"price"`
		Price24HAgo struct {
			Usd string `json:"usd"`
		} `json:"price_24h_ago"`
		Volume24H struct {
			Usd  string `json:"usd"`
			Base string `json:"base"`
		} `json:"volume_24h"`
	} `json:"reserves"`
	DltType        string `json:"dlt_type"`
	DltID          string `json:"dlt_id"`
	Type           int    `json:"type"`
	Version        int    `json:"version"`
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Supply         string `json:"supply"`
	ConverterDltID string `json:"converter_dlt_id"`
	ConversionFee  string `json:"conversion_fee"`
	Liquidity      struct {
		Usd string `json:"usd"`
	} `json:"liquidity"`
	Volume24H struct {
		Usd string `json:"usd"`
	} `json:"volume_24h"`
	Fees24H struct {
		Usd string `json:"usd"`
	} `json:"fees_24h"`
}

type BancorPools struct {
	Data      []BancorPool `json:"data"`
	Timestamp struct {
		Ethereum struct {
			Block     int   `json:"block"`
			Timestamp int64 `json:"timestamp"`
		} `json:"ethereum"`
	} `json:"timestamp"`
}

func NewBancorPoolScraper(exchange dia.Exchange, datastore *models.DB) *BancorPoolScraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *BancorPoolScraper
		pool        = "0xC0205e203F423Bcd8B2a4d6f8C8A154b0Aa60F19"
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", curveRestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	scraper = &BancorPoolScraper{
		RestClient:   restClient,
		datastore:    datastore,
		poolChannel:  poolChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
		pool:         pool,
	}

	go func() {
		scraper.fetchPools()
		scraper.doneChannel <- true
	}()

	return scraper
}

func (scraper *BancorPoolScraper) readPools() ([]BancorPool, error) {
	var bpools BancorPools
	pairs, _, err := utils.GetRequest("https://api-v2.bancor.network/pools")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(pairs, &bpools)
	if err != nil {
		log.Error("Error reading json", err)

	}
	return bpools.Data, nil

}

func (scraper *BancorPoolScraper) fetchpooltokenaddress(pool BancorPool) (address []common.Address, balances []*big.Int, err error) {

	switch pool.Type {
	case 0:
		{
			address, balances, err = scraper.ConverterTypeZero(common.HexToAddress(pool.ConverterDltID))
			if err != nil {
				log.Errorln("Error getting Address", err)
			}

		}
	case 1:
		{
			address, balances, err = scraper.ConverterTypeOne(common.HexToAddress(pool.ConverterDltID))
			if err != nil {
				log.Errorln("Error getting Address", err)
			}
		}
	case 3:
		{
			address, balances, err = scraper.ConverterTypeThree(common.HexToAddress(pool.ConverterDltID))
			if err != nil {
				log.Errorln("Error getting Address", err)
			}
		}
	case 4:
		{
			address, balances, err = scraper.ConverterTypeFour(common.HexToAddress(pool.ConverterDltID))
			if err != nil {
				log.Errorln("Error getting Address", err)
			}
		}
	default:
		err = errors.New("type not available")
	}

	return
}

func (scraper *BancorPoolScraper) loadPoolData(poolCoinAddresses []common.Address, poolBalances []*big.Int, poolAddress common.Address) (pool dia.Pool) {

	var (
		poolAssets []dia.Asset
	)

	var err error
	for _, c := range poolCoinAddresses {
		var (
			coinCaller  *token.TokenCaller
			symbol      string
			decimals    uint8
			decimalsBig *big.Int
			name        string
		)
		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else if c == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			symbol = "ETH"
			decimals = uint8(18)
			name = "Ether"
			c = common.HexToAddress("0x0000000000000000000000000000000000000000")
		} else {
			log.Infoln("Getting token details for : ", c)

			coinCaller, err = token.NewTokenCaller(c, scraper.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}
			symbol, err = coinCaller.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}
			decimalsBig, err = coinCaller.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}
			decimals = uint8(decimalsBig.Uint64())
			name, err = coinCaller.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}
		}
		log.Info(symbol, " ", decimals, " ", "'", name, "'", " ", c)

		poolAssets = append(poolAssets, dia.Asset{
			Address:    c.Hex(),
			Blockchain: scraper.blockchain,
			Decimals:   decimals,
			Symbol:     symbol,
			Name:       name,
		})

	}

	for i := range poolAssets {
		var volume float64
		if poolBalances[i] != nil {
			volume, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(poolBalances[i]), new(big.Float).SetFloat64(math.Pow10(int(poolAssets[i].Decimals)))).Float64()

		}
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
			Asset:  poolAssets[i],
			Volume: volume,
		})
	}
	pool.Exchange = dia.Exchange{Name: scraper.exchangeName}
	pool.Blockchain = dia.BlockChain{Name: scraper.blockchain}
	pool.Address = poolAddress.Hex()

	// Determine USD liquidity.
	if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
		scraper.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}

	pool.Time = time.Now()

	return pool
}

func (scraper *BancorPoolScraper) ConverterTypeZero(address common.Address) (tokenAddress []common.Address, poolBalances []*big.Int, err error) {

	contract, err := ConverterTypeZero.NewConverterTypeZeroCaller(address, scraper.RestClient)
	if err != nil {
		return
	}

	tokenCount, err := contract.ConnectorTokenCount(&bind.CallOpts{})
	if err != nil {
		return
	}

	var i uint16

	for i = 0; i < tokenCount; i++ {
		token1, err := contract.ConnectorTokens(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Errorln("Error", err)
		}
		tokenAddress = append(tokenAddress, token1)

		reserve, err := contract.GetReserveBalance(&bind.CallOpts{}, tokenAddress[int64(i)])
		if err != nil {
			log.Errorln("Error", err)
		}
		poolBalances = append(poolBalances, reserve)
	}

	return

}

func (scraper *BancorPoolScraper) ConverterTypeOne(address common.Address) (tokenAddress []common.Address, poolBalances []*big.Int, err error) {

	contract, err := ConvertertypeOne.NewConvertertypeOne(address, scraper.RestClient)
	if err != nil {
		return
	}

	tokenCount, err := contract.ConnectorTokenCount(&bind.CallOpts{})
	if err != nil {
		return
	}

	var i uint16

	for i = 0; i < tokenCount; i++ {
		token1, err := contract.ConnectorTokens(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Errorln("Error", err)
		}
		tokenAddress = append(tokenAddress, token1)

		reserve, err := contract.ReserveBalance(&bind.CallOpts{}, tokenAddress[int64(i)])
		if err != nil {
			log.Errorln("Error", err)
		}
		poolBalances = append(poolBalances, reserve)
	}

	return

}

func (scraper *BancorPoolScraper) ConverterTypeThree(address common.Address) (tokenAddress []common.Address, poolBalances []*big.Int, err error) {

	contract, err := ConverterTypeThree.NewConverterTypeThree(address, scraper.RestClient)
	if err != nil {
		return
	}

	tokenCount, err := contract.ConnectorTokenCount(&bind.CallOpts{})
	if err != nil {
		return
	}

	var i uint16

	for i = 0; i < tokenCount; i++ {
		token1, err := contract.ConnectorTokens(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Errorln("Error", err)
		}
		tokenAddress = append(tokenAddress, token1)

		reserve, err := contract.ReserveBalance(&bind.CallOpts{}, tokenAddress[int64(i)])
		if err != nil {
			log.Errorln("Error", err)
		}
		poolBalances = append(poolBalances, reserve)
	}

	return

}

func (scraper *BancorPoolScraper) ConverterTypeFour(address common.Address) (tokenAddress []common.Address, poolBalances []*big.Int, err error) {

	contract, err := ConverterTypeFour.NewConverterTypeFour(address, scraper.RestClient)
	if err != nil {
		return
	}

	tokenCount, err := contract.ConnectorTokenCount(&bind.CallOpts{})
	if err != nil {
		return
	}

	var i uint16

	for i = 0; i < tokenCount; i++ {
		token1, err := contract.ConnectorTokens(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Errorln("Error", err)
		}
		tokenAddress = append(tokenAddress, token1)

		reserve, err := contract.GetReserveBalance(&bind.CallOpts{}, tokenAddress[int64(i)])
		if err != nil {
			log.Errorln("Error", err)
		}
		poolBalances = append(poolBalances, reserve)
	}

	return

}

// fetchPools collects all available pools and sends them into the pool channel.
func (scraper *BancorPoolScraper) fetchPools() {

	bpools, err := scraper.readPools()
	if err != nil {
		log.Error("Couldn't obtain Bancor product ids:", err)
	}

	for _, bpool := range bpools {

		poolassetaddress, balances, err := scraper.fetchpooltokenaddress(bpool)
		if err != nil {
			log.Errorln("error getting pool detal", err)
			continue
		}

		pool := scraper.loadPoolData(poolassetaddress, balances, common.HexToAddress(bpool.ConverterDltID))
		scraper.poolChannel <- pool

	}

}

func (scraper *BancorPoolScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *BancorPoolScraper) Done() chan bool {
	return scraper.doneChannel
}
