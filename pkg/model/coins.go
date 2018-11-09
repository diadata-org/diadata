package models

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"sort"
)

const (
	coinsPerPage = 50
)

func (db *DB) GetCoins() (*Coins, error) {
	symbols := db.GetAllSymbols()

	var coins Coins
	key := "dia_coins"
	err := db.redisClient.Get(key).Scan(&coins)
	if err != nil {
		if err != redis.Nil {
			return &coins, err
		}

		coins.Coins = []Coin{}
		coins.CompleteCoinList = []CoinSymbolAndName{}
		coins.Change, _ = db.GetCurrencyChange()

		for _, symbol := range symbols {

			var c1 Coin
			log.Debug("Adding symbol", symbol)
			supply, _ := db.GetSupply(symbol)
			price, _ := db.GetQuotation(symbol)
			if price != nil {
				volume, _ := db.GetVolume(symbol)
				if volume != nil {
					if *volume < 1.0 {
						log.Warning("GetCoins: skipping ", symbol, "because <1.0 volume")
						continue
					}
					c1.Price = price.Price
					c1.Name = price.Name
					c1.Symbol = price.Symbol
					if price.PriceYesterday != nil {
						c1.PriceYesterday = price.PriceYesterday
					}
					c1.Time = price.Time
					c1.VolumeYesterdayUSD = volume
					if supply != nil {
						c1.CirculatingSupply = &supply.CirculatingSupply
					}
					coins.Coins = append(coins.Coins, c1)
				}
			}
		}
		sort.Slice(coins.Coins, func(i, j int) bool {

			if coins.Coins[i].CirculatingSupply == nil {
				return false
			}
			if coins.Coins[j].CirculatingSupply == nil {
				return true
			}
			return (*coins.Coins[i].CirculatingSupply * coins.Coins[i].Price) > (*coins.Coins[j].CirculatingSupply * coins.Coins[j].Price)
		})
		for _, coin := range coins.Coins {
			coins.CompleteCoinList = append(coins.CompleteCoinList, CoinSymbolAndName{coin.Symbol, coin.Name})
		}
		if len(coins.Coins) > coinsPerPage {
			coins.Coins = coins.Coins[:coinsPerPage]
		}
		err = db.redisClient.Set(key, &coins, timeOutRedisOneBlock).Err()
		if err != nil {
			log.Error("Error: on GetCoin setting cache\n", err)
		}
	}
	return &coins, nil
}
