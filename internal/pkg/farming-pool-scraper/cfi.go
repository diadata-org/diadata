package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	platform "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/platform"
	standard "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/curveficontracts/standard"

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
	var coins []string
	//var underlyingCoins []string
	var balances float64
	var rates float64
	normalPoolAndSwap := []string{"gusd", "husd", "usdk", "usdn", "3", "sbtc", "hbtc", "ren"}

	for _, poolDetail := range cv.getCFIPools() {

		platform, err := platform.NewPlatformCaller(common.HexToAddress(poolDetail.PoolAddress), cv.RestClient)
		if err != nil {
			return err
		}

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		for _, name := range normalPoolAndSwap {
			if name == poolDetail.PoolName {

				coinsAddress0, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(0))
				balance0, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(0))
				balance0ToFLoat64, _ := new(big.Float).SetInt(balance0).Float64()
				balances = balance0ToFLoat64
				coins = append(coins, coinsAddress0.Hex())

				strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

				if err != nil {
					return err
				}

				decimals, err := strategy0.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}
				precision := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
				oneFloat := big.NewFloat(1)
				rate0, _ := oneFloat.Quo(oneFloat, precision).Float64()
				rates = rate0

				if poolDetail.PoolName == "3" || poolDetail.PoolName == "sbtc" || poolDetail.PoolName == "hbtc" || poolDetail.PoolName == "ren" {
					coinsAddress1, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(1))
					balance1, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(1))
					balance1ToFLoat64, _ := new(big.Float).SetInt(balance1).Float64()
					coins = append(coins, coinsAddress1.Hex())

					strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

					if err != nil {
						return err
					}

					decimals1, err := strategy1.Decimals(&bind.CallOpts{})
					if err != nil {
						return err
					}
					precision1 := new(big.Float).SetFloat64(math.Pow10(int(decimals1)))
					rate1, _ := oneFloat.Quo(oneFloat, precision1).Float64()

					rates = rate0 + rate1
					balances = balance0ToFLoat64 + balance1ToFLoat64

					if poolDetail.PoolName == "3" || poolDetail.PoolName == "sbtc" {
						coinsAddress2, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(2))
						balance2, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(2))
						balance2ToFLoat64, _ := new(big.Float).SetInt(balance2).Float64()
						coins = append(coins, coinsAddress2.Hex())

						strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

						if err != nil {
							return err
						}

						decimals2, err := strategy2.Decimals(&bind.CallOpts{})
						if err != nil {
							return err
						}
						precision2 := new(big.Float).SetFloat64(math.Pow10(int(decimals2)))
						rate2, _ := oneFloat.Quo(oneFloat, precision2).Float64()

						balances = balances + balance2ToFLoat64
						rates = rates + rate2
					}
				}
			}
		}

		if poolDetail.PoolName == "Compound" || poolDetail.PoolName == "USDT" {
			coinsAddress0, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(0))
			balance0, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(0))

			coins = append(coins, coinsAddress0.Hex())

			comp0, err := standard.NewStandardCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)
			if err != nil {
				return err
			}

			exchangeRateStored, err := comp0.ExchangeRateStored(&bind.CallOpts{})
			if err != nil {
				return err
			}

			supplyRatePerBlock, err := comp0.SupplyRatePerBlock(&bind.CallOpts{})
			if err != nil {
				return err
			}

			accrualBlockNumber, _ := comp0.AccrualBlockNumber(&bind.CallOpts{})

			//c-rate = exchangeRateStored() * (1 + supplyRatePerBlock() * (currentBlock - accrualBlockNumber()) / 1e18)

			blockDifference := header.Number.Sub(header.Number, accrualBlockNumber)
			incrementedSupplyRatePerBlock := supplyRatePerBlock.Mul(blockDifference, supplyRatePerBlock)

			addOneToIncrementedSupplyRatePerBlock := incrementedSupplyRatePerBlock.Add(incrementedSupplyRatePerBlock, big.NewInt(1))
			updatedExchangeRateStored := exchangeRateStored.Mul(exchangeRateStored, addOneToIncrementedSupplyRatePerBlock)
			updatedExchangeRate := new(big.Float).SetInt(updatedExchangeRateStored)
			rate0 := updatedExchangeRate.Quo(updatedExchangeRate, new(big.Float).SetFloat64(1e18))

			coinsAddress1, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(1))
			balance1, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())

			comp1, err := standard.NewStandardCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)
			if err != nil {
				return err
			}

			exchangeRateStored1, err := comp1.ExchangeRateStored(&bind.CallOpts{})
			if err != nil {
				return err
			}

			supplyRatePerBlock1, err := comp1.SupplyRatePerBlock(&bind.CallOpts{})
			if err != nil {
				return err
			}

			accrualBlockNumber1, _ := comp1.AccrualBlockNumber(&bind.CallOpts{})

			//c-rate = exchangeRateStored() * (1 + supplyRatePerBlock() * (currentBlock - accrualBlockNumber()) / 1e18)

			incrementedSupplyRatePerBlock1 := supplyRatePerBlock1.Mul(header.Number.Sub(header.Number, accrualBlockNumber1), supplyRatePerBlock)

			addOneToIncrementedSupplyRatePerBlock1 := incrementedSupplyRatePerBlock1.Add(incrementedSupplyRatePerBlock1, big.NewInt(1))
			updatedExchangeRateStored1 := exchangeRateStored.Mul(exchangeRateStored1, addOneToIncrementedSupplyRatePerBlock1)
			updatedExchangeRate1 := new(big.Float).SetInt(updatedExchangeRateStored1)
			rate1 := updatedExchangeRate1.Quo(updatedExchangeRate1, new(big.Float).SetFloat64(1e18))

			balance0ToFLoat64, _ := new(big.Float).SetInt(balance0).Float64()
			balance1ToFLoat64, _ := new(big.Float).SetInt(balance1).Float64()
			balances = balance0ToFLoat64 + balance1ToFLoat64

			rate0ToFLoat64, _ := rate0.Float64()
			rate1ToFLoat64, _ := rate1.Float64()

			rates = rate0ToFLoat64 + rate1ToFLoat64

			if poolDetail.PoolName == "USDT" {
				coinsAddress2, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(2))
				balance2, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(2))
				balance2ToFLoat64, _ := new(big.Float).SetInt(balance2).Float64()
				balances = balances + balance2ToFLoat64
				coins = append(coins, coinsAddress2.Hex())

				strategy, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)

				if err != nil {
					return err
				}

				decimals, err := strategy.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}
				precision := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
				oneFloat := big.NewFloat(1)
				rate2, _ := oneFloat.Quo(oneFloat, precision).Float64()
				rates = rate2 + rates
			}
		}

		if poolDetail.PoolName == "PAX" || poolDetail.PoolName == "BUSD" || poolDetail.PoolName == "y" {

			coinsAddress0, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(0))
			coins = append(coins, coinsAddress0.Hex())
			balance0, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(0))
			balance0ToFLoat64, _ := new(big.Float).SetInt(balance0).Float64()
			strategy0, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress0.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			pricePerFullShareFromContract, err := strategy0.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
			if err != nil {
				return err

			}
			pricePerFullShare := new(big.Float).SetInt(pricePerFullShareFromContract)
			decimals, err := strategy0.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			precision := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
			rate0, _ := big.NewFloat(1).Quo(pricePerFullShare.Quo(pricePerFullShare, new(big.Float).SetFloat64(1e18)), precision).Float64()

			coinsAddress1, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(1))
			coins = append(coins, coinsAddress1.Hex())
			balance1, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(1))
			balance1ToFLoat64, _ := new(big.Float).SetInt(balance1).Float64()
			strategy1, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress1.Hex()), cv.RestClient)

			if err != nil {
				return err
			}

			pricePerFullShareFromContract1, err := strategy1.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
			if err != nil {
				return err

			}
			pricePerFullShare1 := new(big.Float).SetInt(pricePerFullShareFromContract1)
			decimals1, err := strategy1.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			precision1 := new(big.Float).SetFloat64(math.Pow10(int(decimals1)))
			rate1, _ := big.NewFloat(1).Quo(pricePerFullShare.Quo(pricePerFullShare1, new(big.Float).SetFloat64(1e18)), precision1).Float64()

			coinsAddress2, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(2))
			coins = append(coins, coinsAddress2.Hex())
			balance2, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(2))
			balance2ToFLoat64, _ := new(big.Float).SetInt(balance2).Float64()
			strategy2, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress2.Hex()), cv.RestClient)
			if err != nil {
				return err
			}

			pricePerFullShareFromContract2, err := strategy2.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
			if err != nil {
				return err

			}
			pricePerFullShare2 := new(big.Float).SetInt(pricePerFullShareFromContract2)
			decimals2, err := strategy2.Decimals(&bind.CallOpts{})
			if err != nil {
				return err
			}
			precision2 := new(big.Float).SetFloat64(math.Pow10(int(decimals2)))
			rate2, _ := big.NewFloat(1).Quo(pricePerFullShare.Quo(pricePerFullShare2, new(big.Float).SetFloat64(1e18)), precision2).Float64()

			if poolDetail.PoolName == "PAX" {
				coinsAddress3, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(3))
				balance3, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(3))
				balance3ToFLoat64, _ := new(big.Float).SetInt(balance3).Float64()
				balances = balance3ToFLoat64 + balance2ToFLoat64 + balance0ToFLoat64 + balance1ToFLoat64
				coins = append(coins, coinsAddress3.Hex())

				strategy, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress3.Hex()), cv.RestClient)

				if err != nil {
					return err
				}

				decimals3, err := strategy.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}
				precision3 := new(big.Float).SetFloat64(math.Pow10(int(decimals3)))
				oneFloat := big.NewFloat(1)
				rate3, _ := oneFloat.Quo(oneFloat, precision3).Float64()

				rates = rate2 + rate3 + rate1 + rate0

			} else {
				coinsAddress3, _ := platform.Coins(&bind.CallOpts{}, big.NewInt(3))
				coins = append(coins, coinsAddress3.Hex())
				balance3, _ := platform.Balances(&bind.CallOpts{}, big.NewInt(3))
				balance3ToFLoat64, _ := new(big.Float).SetInt(balance3).Float64()
				balances = balance3ToFLoat64 + balance2ToFLoat64 + balance0ToFLoat64 + balance1ToFLoat64
				strategy3, err := strategy.NewStrategyCaller(common.HexToAddress(coinsAddress3.Hex()), cv.RestClient)
				if err != nil {
					return err
				}

				pricePerFullShareFromContract3, err := strategy3.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
				if err != nil {
					return err

				}
				pricePerFullShare3 := new(big.Float).SetInt(pricePerFullShareFromContract3)
				decimals3, err := strategy3.Decimals(&bind.CallOpts{})
				if err != nil {
					return err
				}
				precision3 := new(big.Float).SetFloat64(math.Pow10(int(decimals3)))
				rate3, _ := big.NewFloat(1).Quo(pricePerFullShare.Quo(pricePerFullShare3, new(big.Float).SetFloat64(1e18)), precision3).Float64()

				rates = rate2 + rate3 + rate1 + rate0

			}

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

	pools = append(pools, &CFIPoolDetail{PoolName: "Compound", PoolAddress: "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56", PoolID: "Compound Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "USDT", PoolAddress: "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C", PoolID: "USDT Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "PAX", PoolAddress: "0x06364f10B501e868329afBc005b3492902d6C763", PoolID: "PAX Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "y", PoolAddress: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51", PoolID: "y Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "BUSD", PoolAddress: "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27", PoolID: "BUSD Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "sUSD", PoolAddress: "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD", PoolID: "sUSD v2 Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "Ren", PoolAddress: "0x93054188d876f558f4a66B2EF1d97d16eDf0895B", PoolID: "REN Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "sbtc", PoolAddress: "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714", PoolID: "sBTC Swap"})
	pools = append(pools, &CFIPoolDetail{PoolName: "hbtc", PoolAddress: "0x4CA9b3063Ec5866A4B82E437059D2C43d1be596F", PoolID: "HBTC Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "3", PoolAddress: "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7", PoolID: "DAI/USDC/USDT Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "gusd", PoolAddress: "0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956", PoolID: "GUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "husd", PoolAddress: "0x3eF6A01A0f81D6046290f3e2A8c5b843e738E604", PoolID: "HUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "usdk", PoolAddress: "0x3E01dD8a5E1fb3481F0F589056b428Fc308AF0Fb", PoolID: "USDK Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "usdn", PoolAddress: "0x0f9cb53Ebe405d49A0bbdBD291A65Ff571bC83e1", PoolID: "USDN Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "musd", PoolAddress: "0x8474DdbE98F5aA3179B3B3F5942D724aFcdec9f6", PoolID: "mUSD Pool"})
	pools = append(pools, &CFIPoolDetail{PoolName: "tbtc", PoolAddress: "0xC25099792E9349C7DD09759744ea681C7de2cb66", PoolID: "TBTC Pool"})

	return

}

type CFIPoolDetail struct {
	PoolName    string
	PoolAddress string
	PoolID      string
}
