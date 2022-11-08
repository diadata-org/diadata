package models

import (
	// "encoding/json"
	"context"
	"database/sql"
	"fmt"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// GetActiveExchangesAndPairs returns all exchanges the asset with @address and @blockchain was
// traded on in the given time-range as keys of a map.
// Additionaly, the map's values are the underlying pairs.
func (datastore *DB) GetActiveExchangesAndPairs(address string, blockchain string, starttime time.Time, endtime time.Time) (map[string][]dia.Pair, error) {
	exchangepairmap := make(map[string][]dia.Pair)

	query := `
	SELECT exchange,pair,quotetokenaddress,quotetokenblockchain,basetokenaddress,basetokenblockchain,LAST(estimatedUSDPrice) 
	FROM %s 
	WHERE time>%d AND time<=%d 
	AND quotetokenaddress='%s' AND quotetokenblockchain='%s'
	AND verified='true'
	GROUP BY "exchange","pair"
	`

	q := fmt.Sprintf(query, influxDbTradesTable, starttime.UnixNano(), endtime.UnixNano(), address, blockchain)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return exchangepairmap, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series {
			if len(row.Values[0]) > 1 {
				quoteToken := dia.Asset{
					Address:    row.Values[0][3].(string),
					Blockchain: row.Values[0][4].(string),
				}
				baseToken := dia.Asset{
					Address:    row.Values[0][5].(string),
					Blockchain: row.Values[0][6].(string),
				}
				pair := dia.Pair{QuoteToken: quoteToken, BaseToken: baseToken}

				exchangepairmap[row.Values[0][1].(string)] = append(exchangepairmap[row.Values[0][1].(string)], pair)
			}
		}
	}

	return exchangepairmap, nil

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
	return datastore.redisClient.Set(key, &p, 0).Err()
}

// GetAvailablePairs a slice of all pairs available in the exchange in the internal redis db
func (datastore *DB) GetAvailablePairs(exchange string) ([]dia.ExchangePair, error) {
	key := "dia_available_pairs_" + exchange
	p := dia.Pairs{}
	err := datastore.redisClient.Get(key).Scan(&p)
	if err != nil {
		log.Errorf("Error: %v on GetAvailablePairs %v\n", err, exchange)
		return nil, err
	}
	return p, nil
}

func (rdb *RelDB) SetExchange(exchange dia.Exchange) (err error) {
	fields := fmt.Sprintf("INSERT INTO %s (name,centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay) VALUES ", exchangeTable)
	values := "($1,$2,$3,NULLIF($4,''),$5,NULLIF($6,''),NULLIF($7,''),NULLIF($8,''),$9)"
	conflict := " ON CONFLICT (name) DO UPDATE SET contract=NULLIF($4,''),rest_api=$6,ws_api=$7,pairs_api=$8,watchdog_delay=$9"

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
	)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetExchange(name string) (exchange dia.Exchange, err error) {
	query := fmt.Sprintf("SELECT centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay FROM %s WHERE name=$1", exchangeTable)
	var contract sql.NullString
	var blockchainName sql.NullString
	var restAPI sql.NullString
	var wsAPI sql.NullString
	var pairsAPI sql.NullString
	err = rdb.postgresClient.QueryRow(context.Background(), query, name).Scan(
		&exchange.Centralized,
		&exchange.Bridge,
		&contract,
		&blockchainName,
		&restAPI,
		&wsAPI,
		&pairsAPI,
		&exchange.WatchdogDelay,
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
	query := fmt.Sprintf("SELECT name,centralized,bridge,contract,blockchain,rest_api,ws_api,pairs_api,watchdog_delay FROM %s", exchangeTable)
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
