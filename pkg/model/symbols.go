package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (db *DB) GetAllSymbols() []string {
	r := make(map[string]string)

	// TODO: search in redis instead
	for _, e := range dia.Exchanges() {
		p, err := db.GetAvailablePairsForExchange(e)
		if err == nil {
			for _, v := range p {
				r[v.Symbol] = v.Symbol
			}
		} else {
			log.Error("GetAllSymbols", err)
		}
	}
	s := []string{}
	for _, value := range r {
		s = append(s, value)
	}
	return s
}

func (db *DB) GetSymbols(exchange string) ([]string, error) {
	var result []string
	var cursor uint64
	key := "dia_" + dia.FilterKing + "_"
	for {
		var keys []string
		var err error
		keys, cursor, err = db.redisClient.Scan(cursor, key+"*", 10).Result()
		if err != nil {
			log.Error("GetPairs err", err)
			return result, err
		}
		for _, value := range keys {

			filteredKey := strings.Replace(strings.Replace(value, key, "", 1), "_ZSET", "", 1)
			s := strings.Split(strings.Replace(filteredKey, key, "", 1), "_")
			if exchange == "" {
				if len(s) == 1 {
					result = append(result, s[0])
				}
			} else {
				if s[1] == exchange {
					result = append(result, s[0])
				}
			}
		}
		if cursor == 0 {
			log.Debug("GetSymbols %v returns %v", key, result)
			return result, nil
		}
	}
}

func (db *DB) GetSymbolExchangeDetails(symbol string, exchange string) (*SymbolExchangeDetails, error) {
	result := &SymbolExchangeDetails{
		Name: exchange,
	}
	v, err := db.GetPrice(symbol, exchange)
	if err == nil {
		result.Price = v
	}

	py, err2 := db.GetPriceYesterday(symbol, exchange)
	if err2 == nil {
		result.PriceYesterday = &py
	}

	v2, _ := db.GetVolumeExchange(symbol, exchange)
	result.VolumeYesterdayUSD = v2
	d, _ := db.GetLastTradeTimeForExchange(symbol, exchange)
	result.Time = d

	t, _ := db.GetLastTrades(symbol, exchange, 10)
	result.LastTrades = t

	return result, err
}

func (db *DB) GetSymbolDetails(symbol string) (*SymbolDetails, error) {
	q, err := db.GetQuotation(symbol)
	if err != nil {
		return nil, err
	} else {
		r := &SymbolDetails{
			Coin: Coin{
				Symbol:             q.Symbol,
				Name:               q.Name,
				Price:              q.Price,
				VolumeYesterdayUSD: q.VolumeYesterdayUSD,
				Time:               q.Time,
				PriceYesterday:     q.PriceYesterday,
			},
			Exchanges: []SymbolExchangeDetails{},
		}
		r.Change, _ = db.GetCurrencyChange()
		s, err := db.GetSupply(symbol)
		if err == nil {
			r.Coin.CirculatingSupply = &s.CirculatingSupply
		}
		exs, err := db.GetExchangesForSymbol(symbol)
		if err == nil {
			for _, e := range exs {
				s, err2 := db.GetSymbolExchangeDetails(symbol, e)
				if err2 == nil {
					r.Exchanges = append(r.Exchanges, *s)
				}
			}
		}
		return r, err
	}
}
