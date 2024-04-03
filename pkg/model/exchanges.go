package models

import (
	// "encoding/json"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// GetActiveExchangesAndPairs returns all exchanges the asset with @address and @blockchain was
// traded on in the given time-range as keys of a map. The map's values are the underlying pairs.
// Additionally, a map is returned where keys are exchange/pair identifier and values are the number
// of trades in the respective time-range.
// The pair has to have at least @numTrades trades in the given time-range in order to get returned.
func (datastore *DB) GetActiveExchangesAndPairs(
	address string,
	blockchain string,
	numTradesThreshold int64,
	starttime time.Time,
	endtime time.Time,
) (map[string][]dia.Pair, map[string]int64, error) {
	exchangepairmap := make(map[string][]dia.Pair)
	pairCountTradesMap := make(map[string]int64)

	query := `
	SELECT count(*)
	FROM %s 
	WHERE time>%d AND time<=%d 
	AND quotetokenaddress='%s' AND quotetokenblockchain='%s'
	AND verified='true'
	GROUP BY "exchange","pair","basetokenaddress","basetokenblockchain"
	`

	q := fmt.Sprintf(query, influxDbTradesTable, starttime.UnixNano(), endtime.UnixNano(), address, blockchain)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return exchangepairmap, pairCountTradesMap, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series {

			if len(row.Values[0]) > 1 {
				numTrades, err := row.Values[0][1].(json.Number).Int64()
				if err != nil {
					log.Warn("parse number of trades: ", err)
				}

				exchange := row.Tags["exchange"]
				// Only include pair if it has more than @numTradesThreshold trades.
				if numTrades >= numTradesThreshold {
					pair := dia.Pair{
						QuoteToken: dia.Asset{Blockchain: blockchain, Address: address},
						BaseToken:  dia.Asset{Blockchain: row.Tags["basetokenblockchain"], Address: row.Tags["basetokenaddress"]},
					}
					exchangepairmap[exchange] = append(exchangepairmap[exchange], pair)
					pairCountTradesMap[pair.PairExchangeIdentifier(exchange)] = numTrades
				}
			}
		}
	}

	return exchangepairmap, pairCountTradesMap, nil

}

func (rdb *RelDB) GetExchangesForSymbol(symbol string) (exchanges []string, err error) {

	query := fmt.Sprintf("select distinct(exchange) from %s where symbol=$1", exchangesymbolTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query, symbol)
	if err != nil {
		return
	}
	for rows.Next() {
		exchange := ""
		err = rows.Scan(&exchange)
		if err != nil {
			return []string{}, err
		}
		exchanges = append(exchanges, exchange)
	}
	return
}

// SetAvailablePairs stores @pairs in redis
// TO DO: Setter and getter should act on RelDB
func (datastore *DB) SetAvailablePairs(exchange string, pairs []dia.ExchangePair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return datastore.redisClient.Set(context.Background(), key, &p, 0).Err()
}

// GetAvailablePairs a slice of all pairs available in the exchange in the internal redis db
func (datastore *DB) GetAvailablePairs(exchange string) ([]dia.ExchangePair, error) {
	key := "dia_available_pairs_" + exchange
	p := dia.Pairs{}
	err := datastore.redisClient.Get(context.Background(), key).Scan(&p)
	if err != nil {
		log.Errorf("Error: %v on GetAvailablePairs %v\n", err, exchange)
		return nil, err
	}
	return p, nil
}

func (rdb *RelDB) SetExchange(exchange dia.Exchange) (err error) {
	fields := fmt.Sprintf("INSERT INTO %s (name,centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay,scraper_active) VALUES ", exchangeTable)
	values := "($1,$2,$3,NULLIF($4,''),$5,NULLIF($6,''),NULLIF($7,''),NULLIF($8,''),$9,$10)"
	conflict := " ON CONFLICT (name) DO UPDATE SET contract=NULLIF($4,''),rest_api=$6,ws_api=$7,pairs_api=$8,watchdog_delay=$9,scraper_active=$10"

	query := fields + values + conflict
	_, err = rdb.postgresClient.Exec(context.Background(), query,
		exchange.Name,
		exchange.Centralized,
		exchange.Bridge,
		exchange.Contract,
		exchange.BlockChain.Name,
		exchange.RestAPI,
		exchange.WsAPI,
		exchange.PairsAPI,
		exchange.WatchdogDelay,
		exchange.ScraperActive,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetExchange(name string) (exchange dia.Exchange, err error) {
	query := fmt.Sprintf(`
	SELECT centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay,scraper_active 
	FROM %s 
	WHERE name=$1`,
		exchangeTable,
	)
	var (
		contract       sql.NullString
		blockchainName sql.NullString
		restAPI        sql.NullString
		wsAPI          sql.NullString
		pairsAPI       sql.NullString
	)
	err = rdb.postgresClient.QueryRow(context.Background(), query, name).Scan(
		&exchange.Centralized,
		&exchange.Bridge,
		&contract,
		&blockchainName,
		&restAPI,
		&wsAPI,
		&pairsAPI,
		&exchange.WatchdogDelay,
		&exchange.ScraperActive,
	)
	if err != nil {
		return
	}
	if contract.Valid {
		exchange.Contract = contract.String
	}
	if blockchainName.Valid {
		exchange.BlockChain.Name = blockchainName.String
	}
	if restAPI.Valid {
		exchange.RestAPI = restAPI.String
	}
	if wsAPI.Valid {
		exchange.WsAPI = wsAPI.String
	}
	if pairsAPI.Valid {
		exchange.PairsAPI = pairsAPI.String
	}
	exchange.Name = name
	return
}

// GetAllExchanges returns all exchanges existent in the exchange table.
func (rdb *RelDB) GetAllExchanges() (exchanges []dia.Exchange, err error) {
	query := fmt.Sprintf(`
	SELECT name,centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay,scraper_active 
	FROM %s`,
		exchangeTable,
	)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []dia.Exchange{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var exchange dia.Exchange
		var contract sql.NullString
		var blockchainName sql.NullString
		var restAPI sql.NullString
		var wsAPI sql.NullString
		var pairsAPI sql.NullString
		err := rows.Scan(
			&exchange.Name,
			&exchange.Centralized,
			&exchange.Bridge,
			&contract,
			&blockchainName,
			&restAPI,
			&wsAPI,
			&pairsAPI,
			&exchange.WatchdogDelay,
			&exchange.ScraperActive,
		)
		if err != nil {
			return []dia.Exchange{}, err
		}
		if contract.Valid {
			exchange.Contract = contract.String
		}
		if blockchainName.Valid {
			exchange.BlockChain.Name = blockchainName.String
		}
		if restAPI.Valid {
			exchange.RestAPI = restAPI.String
		}
		if wsAPI.Valid {
			exchange.WsAPI = wsAPI.String
		}
		if pairsAPI.Valid {
			exchange.PairsAPI = pairsAPI.String
		}
		exchanges = append(exchanges, exchange)
	}

	return exchanges, nil
}

// GetExchangeNames returns the names of all available exchanges.
func (rdb *RelDB) GetExchangeNames() (allExchanges []string, err error) {
	exchanges, err := rdb.GetAllExchanges()
	if err != nil {
		return
	}
	for _, exchange := range exchanges {
		allExchanges = append(allExchanges, exchange.Name)
	}
	sort.Strings(allExchanges)
	return
}

func GetExchangeType(exchange dia.Exchange) string {
	if exchange.Centralized {
		return "CEX"
	} else if exchange.Bridge {
		return "Bridge"
	} else {
		return "DEX"
	}
}
