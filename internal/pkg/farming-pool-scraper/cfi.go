package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	platform "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/platform"
	ren "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/ren"
	special "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/special"
	suds "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/susd"

	strategy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/yficontracts/strategy"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CFIPool struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewCFIPool(scraper *PoolScraper) *CFIPool {
	log.Info("Curvefi pool is built and triggered")
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	cv := &CFIPool{scraper: scraper, RestClient: restClient, WsClient: wsClient}

	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *CFIPool) mainLoop() {

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case <-cv.scraper.tickerRate.C:
				err := cv.scrapePools()
				if err != nil {
					log.Errorln("Error while Scrapping", err)
				}
			}

		}
	}()

}

func (cv *CFIPool) scrapePools() (err error) {

	//var underlyingCoins []string

	metaPool := []string{"gusd", "husd", "usdk", "usdn", "musd", "rsv", "tbtc", "3", "hbtc"}

	for _, poolDetail := range cv.getCFIPools() {
		var coins []string

		var balances float64

		platform, err := platform.NewPlatformCaller(common.HexToAddress(poolDetail.PoolAddress[0]), cv.RestClient)
		if err != nil {
			return err
		}

		special, err := special.NewSpecial(common.HexToAddress(poolDetail.PoolAddress[0]), cv.RestClient)
		if err != nil {
			return err
		}

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		virtualPrice, _ := platform.GetVirtualPrice(&bind.CallOpts{})
		virtualPriceRate := new(big.Float).SetInt(virtualPrice)
		rates, _ := virtualPriceRate.Quo(virtualPriceRate, new(big.Float).SetFloat64(1e18)).Float64()

		if poolDetail.PoolName == "sUSD" {
			susd, err := suds.NewSudsCaller(common.HexToAddress(poolDetail.PoolAddress[0]), cv.RestClient)
			if err != nil {
				return err
			}

			coinsAddress0, _ := susd.Coins(&bind.CallOpts{}, big.NewInt(0))
			coins = append(coins, coinsAddress0.Hex())

			strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals0, err := strategy0.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance0, _ := susd.Balances(&bind.CallOpts{}, big.NewInt(0))
			balance0ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0), new(big.Float).SetFloat64(math.Pow10(int(decimals0)))).Float64()

			coinsAddress1, _ := susd.Coins(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())

			strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals1, err := strategy1.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			balance1, _ := susd.Balances(&bind.CallOpts{}, big.NewInt(1))
			balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

			coinsAddress2, _ := susd.Coins(&bind.CallOpts{}, big.NewInt(2))
			coins = append(coins, coinsAddress2.Hex())

			strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals2, err := strategy2.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			balance2, _ := susd.Balances(&bind.CallOpts{}, big.NewInt(2))
			balance2ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance2), new(big.Float).SetFloat64(math.Pow10(int(decimals2)))).Float64()

			coinsAddress3, _ := susd.Coins(&bind.CallOpts{}, big.NewInt(3))
			coins = append(coins, coinsAddress3.Hex())

			strategy3, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress3.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals3, err := strategy3.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			balance3, _ := susd.Balances(&bind.CallOpts{}, big.NewInt(3))
			balance3ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance3), new(big.Float).SetFloat64(math.Pow10(int(decimals3)))).Float64()

			balances = balance0ToFLoat64 + balance1ToFLoat64 + balance2ToFLoat64 + balance3ToFLoat64
		}

		if "ren" == poolDetail.PoolName {
			ren, err := ren.NewRenCaller(common.HexToAddress(poolDetail.PoolAddress[0]), cv.RestClient)
			if err != nil {
				return err
			}

			coinsAddress0, _ := ren.Coins(&bind.CallOpts{}, big.NewInt(0))
			coins = append(coins, coinsAddress0.Hex())

			strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals0, err := strategy0.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance0, _ := ren.Balances(&bind.CallOpts{}, big.NewInt(0))
			balance0ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0), new(big.Float).SetFloat64(math.Pow10(int(decimals0)))).Float64()

			coinsAddress1, _ := ren.Coins(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())

			strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals1, err := strategy1.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance1, _ := ren.Balances(&bind.CallOpts{}, big.NewInt(1))
			balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

			balances = balance1ToFLoat64 + balance0ToFLoat64
		}

		for _, name := range metaPool {
			if name == poolDetail.PoolName {
				coinsAddress0, _ := special.Coins(&bind.CallOpts{}, big.NewInt(0))
				coins = append(coins, coinsAddress0.Hex())

				strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

				if err != nil {
					return err
				}

				decimals0, err := strategy0.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}

				balance0, _ := special.Balances(&bind.CallOpts{}, big.NewInt(0))
				balance0ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0), new(big.Float).SetFloat64(math.Pow10(int(decimals0)))).Float64()

				balances = balance0ToFLoat64

				if poolDetail.PoolName == "3" {
					coinsAddress1, _ := special.Coins(&bind.CallOpts{}, big.NewInt(1))
					coins = append(coins, coinsAddress1.Hex())

					strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

					if err != nil {
						return err
					}

					decimals1, err := strategy1.Decimals(&bind.CallOpts{})
					if err != nil {
						return err
					}

					balance1, _ := special.Balances(&bind.CallOpts{}, big.NewInt(1))
					balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

					coinsAddress2, _ := special.Coins(&bind.CallOpts{}, big.NewInt(2))
					coins = append(coins, coinsAddress2.Hex())

					strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

					if err != nil {
						return err
					}

					decimals2, err := strategy2.Decimals(&bind.CallOpts{})
					if err != nil {
						return err
					}

					balance2, _ := special.Balances(&bind.CallOpts{}, big.NewInt(2))
					balance2ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance2), new(big.Float).SetFloat64(math.Pow10(int(decimals2)))).Float64()

					balances = balances + balance1ToFLoat64 + balance2ToFLoat64

				} else {
					coinsAddress1, _ := special.Coins(&bind.CallOpts{}, big.NewInt(1))
					coins = append(coins, coinsAddress1.Hex())

					strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

					if err != nil {
						return err
					}

					decimals1, err := strategy1.Decimals(&bind.CallOpts{})
					if err != nil {
						return err
					}

					balance1, _ := special.Balances(&bind.CallOpts{}, big.NewInt(1))
					balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

					balances = balances + balance1ToFLoat64
				}
			}
		}

		if poolDetail.PoolName == "Compound" || poolDetail.PoolName == "USDT" || poolDetail.PoolName == "sbtc" {
			coinsAddress0, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(0))
			coins = append(coins, coinsAddress0.Hex())

			strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals0, err := strategy0.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			balance0, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(0))
			balance0ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0), new(big.Float).SetFloat64(math.Pow10(int(decimals0)))).Float64()

			coinsAddress1, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())

			strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals1, err := strategy1.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance1, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(1))
			balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

			balances = balance0ToFLoat64 + balance1ToFLoat64
			if poolDetail.PoolName == "USDT" || poolDetail.PoolName == "sbtc" {
				coinsAddress2, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(2))
				coins = append(coins, coinsAddress2.Hex())

				strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

				if err != nil {
					return err
				}

				decimals2, err := strategy2.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}

				balance2, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(2))
				balance2ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance2), new(big.Float).SetFloat64(math.Pow10(int(decimals2)))).Float64()

				balances = balances + balance2ToFLoat64
			}
		}

		if poolDetail.PoolName == "PAX" || poolDetail.PoolName == "BUSD" || poolDetail.PoolName == "y" {
			coinsAddress0, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(0))
			coins = append(coins, coinsAddress0.Hex())

			strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals0, err := strategy0.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance0, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(0))
			balance0ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0), new(big.Float).SetFloat64(math.Pow10(int(decimals0)))).Float64()

			coinsAddress1, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())

			strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals1, err := strategy1.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance1, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(1))
			balance1ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1), new(big.Float).SetFloat64(math.Pow10(int(decimals1)))).Float64()

			coinsAddress2, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(2))
			coins = append(coins, coinsAddress2.Hex())

			strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals2, err := strategy2.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance2, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(2))
			balance2ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance2), new(big.Float).SetFloat64(math.Pow10(int(decimals2)))).Float64()

			coinsAddress3, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(3))
			coins = append(coins, coinsAddress3.Hex())

			strategy3, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress3.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			decimals3, err := strategy3.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}

			balance3, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(3))
			balance3ToFLoat64, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance3), new(big.Float).SetFloat64(math.Pow10(int(decimals3)))).Float64()

			balances = balance0ToFLoat64 + balance1ToFLoat64 + balance2ToFLoat64 + balance3ToFLoat64
		}

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = rates
		pr.Balance = balances
		pr.ProtocolName = cv.scraper.poolName
		pr.PoolID = poolDetail.PoolID
		pr.OutputAsset = coins
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = coins
		cv.scraper.chanPoolInfo <- &pr
	}
	return

}

func (cv *CFIPool) getCFIPools() (pools []*CFIPoolDetail) {

	pools = append(pools, &CFIPoolDetail{PoolName: "Compound", PoolAddress: []string{"0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56", "0x7ca5b0a2910B33e9759DC7dDB0413949071D7575"}, PoolID: "Compound Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "USDT", PoolAddress: []string{"0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C", "0xBC89cd85491d81C6AD2954E6d0362Ee29fCa8F53"}, PoolID: "USDT Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "PAX", PoolAddress: []string{"0x06364f10B501e868329afBc005b3492902d6C763", "0x64E3C23bfc40722d3B649844055F1D51c1ac041d"}, PoolID: "PAX Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "y", PoolAddress: []string{"0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51", "0xFA712EE4788C042e2B7BB55E6cb8ec569C4530c1"}, PoolID: "y Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "BUSD", PoolAddress: []string{"0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27", "0x69Fb7c45726cfE2baDeE8317005d3F94bE838840"}, PoolID: "BUSD Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "sUSD", PoolAddress: []string{"0xA5407eAE9Ba41422680e2e00537571bcC53efBfD", "0xA90996896660DEcC6E997655E065b23788857849"}, PoolID: "sUSD v2 Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "ren", PoolAddress: []string{"0x93054188d876f558f4a66B2EF1d97d16eDf0895B", "0xB1F2cdeC61db658F091671F5f199635aEF202CAC"}, PoolID: "REN Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "sbtc", PoolAddress: []string{"0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714", "0x705350c4BcD35c9441419DdD5d2f097d7a55410F"}, PoolID: "sBTC Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "hbtc", PoolAddress: []string{"0x4CA9b3063Ec5866A4B82E437059D2C43d1be596F", "0x4c18E409Dc8619bFb6a1cB56D114C3f592E0aE79"}, PoolID: "HBTC Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "3", PoolAddress: []string{"0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7", "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"}, PoolID: "DAI/USDC/USDT Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "gusd", PoolAddress: []string{"0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956", "0xC5cfaDA84E902aD92DD40194f0883ad49639b023"}, PoolID: "GUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "husd", PoolAddress: []string{"0x3eF6A01A0f81D6046290f3e2A8c5b843e738E604", "0x2db0E83599a91b508Ac268a6197b8B14F5e72840"}, PoolID: "HUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "usdk", PoolAddress: []string{"0x3E01dD8a5E1fb3481F0F589056b428Fc308AF0Fb", "0xC2b1DF84112619D190193E48148000e3990Bf627"}, PoolID: "USDK Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "usdn", PoolAddress: []string{"0x0f9cb53Ebe405d49A0bbdBD291A65Ff571bC83e1", "0xF98450B5602fa59CC66e1379DFfB6FDDc724CfC4"}, PoolID: "USDN Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "musd", PoolAddress: []string{"0x8474DdbE98F5aA3179B3B3F5942D724aFcdec9f6", "0x5f626c30EC1215f4EdCc9982265E8b1F411D1352"}, PoolID: "mUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "tbtc", PoolAddress: []string{"0xC25099792E9349C7DD09759744ea681C7de2cb66", "0x6828bcF74279eE32f2723eC536c22c51Eed383C6"}, PoolID: "TBTC Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "rsv", PoolAddress: []string{"0xC18cC39da8b11dA8c3541C598eE022258F9744da", "0x4dC4A289a8E33600D8bD4cf5F6313E43a37adec7"}, PoolID: "RSV Pool"})

	return

}

type CFIPoolDetail struct {
	PoolName    string
	PoolAddress []string
	PoolID      string
}
