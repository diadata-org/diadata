package pool

import (
	"context"
	"math/big"
	"time"

	gauge "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/gauge"
	platform "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/platform"
	"github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/special"
	ethhelper "github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CFIScraper struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

type CFIPool struct {
	PoolName     string
	PoolAddress  string
	GaugeAddress string
}

func NewCFIScraper(scraper *PoolScraper) *CFIScraper {
	log.Info("Curvefi scrapers is built and triggered")
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	cv := &CFIScraper{scraper: scraper, RestClient: restClient, WsClient: wsClient}

	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *CFIScraper) mainLoop() {

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case <-cv.scraper.tickerRate.C:
				err := cv.scrapePools()
				if err != nil {
					log.Errorln("Error while Scraping", err)
				}
			}

		}
	}()

}

func (cv *CFIScraper) scrapePools() (err error) {

	for _, poolDetail := range cv.getCFIPools() {

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err
		}
		platform, err := platform.NewPlatformCaller(common.HexToAddress(poolDetail.PoolAddress), cv.RestClient)
		if err != nil {
			return err
		}
		gauge, err := gauge.NewGauge(common.HexToAddress(poolDetail.GaugeAddress), cv.RestClient)
		if err != nil {
			return err
		}

		// Get supply of lp tokens in gauge
		totalSupplyBig, err := gauge.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return err
		}
		tmpSupply := new(big.Float).SetInt(totalSupplyBig)
		totalSupplyGauge, _ := tmpSupply.Quo(tmpSupply, new(big.Float).SetFloat64(1e18)).Float64()

		// Get virtual price
		virtualPriceBig, _ := platform.GetVirtualPrice(&bind.CallOpts{})
		tmpVirtualPrice := new(big.Float).SetInt(virtualPriceBig)
		virtualPrice, _ := tmpVirtualPrice.Quo(tmpVirtualPrice, new(big.Float).SetFloat64(1e18)).Float64()

		// Get constituent tokens
		coins, err := cv.getCoins(common.HexToAddress(poolDetail.PoolAddress))
		if err != nil {
			return err
		}

		// Get pool's lp token symbol
		lpToken, err := gauge.LpToken(&bind.CallOpts{})
		if err != nil {
			log.Errorf("error getting lp token for pool %s \n ", poolDetail.PoolAddress)
		}
		tokenCaller, _ := ethhelper.NewTokenCaller(lpToken, cv.RestClient)
		lpTokenSymbol := new([]interface{})
		tokenCaller.Contract.Call(&bind.CallOpts{}, lpTokenSymbol, "symbol")

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = virtualPrice
		pr.Balance = totalSupplyGauge
		pr.ProtocolName = "CURVEFI"
		pr.PoolID = poolDetail.PoolName
		pr.OutputAsset = []string{(*lpTokenSymbol)[0].(string)}
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = coins
		cv.scraper.chanPoolInfo <- &pr
	}
	return
}

// getCoins returns the symbols for the constituent coins of the gauge with address @poolAddress
func (cv *CFIScraper) getCoins(poolAddress common.Address) ([]string, error) {
	var symbols []string
	var coin common.Address
	platform, err := platform.NewPlatformCaller(poolAddress, cv.RestClient)
	if err != nil {
		return nil, err
	}
	special, err := special.NewSpecial(poolAddress, cv.RestClient)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		// If present, get underlying coin...
		coin, err = platform.UnderlyingCoins(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			// ... if not, get coin itself from platform if present...
			coin, err = platform.Coins(&bind.CallOpts{}, big.NewInt(int64(i)))
		}
		if err != nil {
			// ... if not get from special
			coin, err = special.Coins(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				continue
			}
		}
		tokenCaller, _ := ethhelper.NewTokenCaller(coin, cv.RestClient)
		symbol := new([]interface{})
		err = tokenCaller.Contract.Call(&bind.CallOpts{}, symbol, "symbol")
		symbols = append(symbols, (*symbol)[0].(string))
	}

	return symbols, nil
}

func (cv *CFIScraper) getCFIPools() (pools []*CFIPool) {

	pools = append(pools, &CFIPool{PoolName: "Compound", PoolAddress: "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56", GaugeAddress: "0x7ca5b0a2910B33e9759DC7dDB0413949071D7575"})
	pools = append(pools, &CFIPool{PoolName: "USDT", PoolAddress: "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C", GaugeAddress: "0xBC89cd85491d81C6AD2954E6d0362Ee29fCa8F53"})
	pools = append(pools, &CFIPool{PoolName: "PAX", PoolAddress: "0x06364f10B501e868329afBc005b3492902d6C763", GaugeAddress: "0x64E3C23bfc40722d3B649844055F1D51c1ac041d"})
	pools = append(pools, &CFIPool{PoolName: "y", PoolAddress: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51", GaugeAddress: "0xFA712EE4788C042e2B7BB55E6cb8ec569C4530c1"})
	pools = append(pools, &CFIPool{PoolName: "BUSD", PoolAddress: "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27", GaugeAddress: "0x69Fb7c45726cfE2baDeE8317005d3F94bE838840"})
	pools = append(pools, &CFIPool{PoolName: "sUSD", PoolAddress: "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD", GaugeAddress: "0xA90996896660DEcC6E997655E065b23788857849"})
	pools = append(pools, &CFIPool{PoolName: "ren", PoolAddress: "0x93054188d876f558f4a66B2EF1d97d16eDf0895B", GaugeAddress: "0xB1F2cdeC61db658F091671F5f199635aEF202CAC"})
	pools = append(pools, &CFIPool{PoolName: "sbtc", PoolAddress: "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714", GaugeAddress: "0x705350c4BcD35c9441419DdD5d2f097d7a55410F"})
	pools = append(pools, &CFIPool{PoolName: "hbtc", PoolAddress: "0x4CA9b3063Ec5866A4B82E437059D2C43d1be596F", GaugeAddress: "0x4c18E409Dc8619bFb6a1cB56D114C3f592E0aE79"})
	pools = append(pools, &CFIPool{PoolName: "3", PoolAddress: "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7", GaugeAddress: "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"})
	pools = append(pools, &CFIPool{PoolName: "gusd", PoolAddress: "0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956", GaugeAddress: "0xC5cfaDA84E902aD92DD40194f0883ad49639b023"})
	pools = append(pools, &CFIPool{PoolName: "husd", PoolAddress: "0x3eF6A01A0f81D6046290f3e2A8c5b843e738E604", GaugeAddress: "0x2db0E83599a91b508Ac268a6197b8B14F5e72840"})
	pools = append(pools, &CFIPool{PoolName: "usdk", PoolAddress: "0x3E01dD8a5E1fb3481F0F589056b428Fc308AF0Fb", GaugeAddress: "0xC2b1DF84112619D190193E48148000e3990Bf627"})
	pools = append(pools, &CFIPool{PoolName: "usdn", PoolAddress: "0x0f9cb53Ebe405d49A0bbdBD291A65Ff571bC83e1", GaugeAddress: "0xF98450B5602fa59CC66e1379DFfB6FDDc724CfC4"})
	pools = append(pools, &CFIPool{PoolName: "musd", PoolAddress: "0x8474DdbE98F5aA3179B3B3F5942D724aFcdec9f6", GaugeAddress: "0x5f626c30EC1215f4EdCc9982265E8b1F411D1352"})
	pools = append(pools, &CFIPool{PoolName: "tbtc", PoolAddress: "0xC25099792E9349C7DD09759744ea681C7de2cb66", GaugeAddress: "0x6828bcF74279eE32f2723eC536c22c51Eed383C6"})
	pools = append(pools, &CFIPool{PoolName: "rsv", PoolAddress: "0xC18cC39da8b11dA8c3541C598eE022258F9744da", GaugeAddress: "0x4dC4A289a8E33600D8bD4cf5F6313E43a37adec7"})

	return

}
