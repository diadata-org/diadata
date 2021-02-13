package assetstore

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const (
	postgresKey = "postgres_key.txt"
)

// RelDB is a relative database with redis caching layer
type RelDB struct {
	URI            string
	postgresClient *pgx.Conn
	redisClient    *redis.Client
	pagesize       uint32
}

func NewRelDataStore() (*RelDB, error) {
	return NewRelDataStoreWithOptions(true, true)
}

func NewPostgresDataStore() (*RelDB, error) {
	return NewRelDataStoreWithOptions(true, false)
}

func NewRedisDataStore() (*RelDB, error) {
	return NewRelDataStoreWithOptions(false, true)
}

func NewRelDataStoreWithOptions(withPostgres bool, withRedis bool) (*RelDB, error) {
	var postgresClient *pgx.Conn
	var redisClient *redis.Client
	var err error
	// This environment variable is either set in docker-compose or empty
	executionMode := os.Getenv("EXEC_MODE")
	address := ""
	url := "postgresql://localhost/postgres?user=postgres&password=" + getPostgresKeyFromSecrets()
	if withPostgres {

		postgresClient, err = pgx.Connect(context.Background(), url)
		if err != nil {
			return nil, err
		}
	}
	if withRedis {
		// Run localhost for testing and server for production
		if executionMode == "production" {
			address = "redis:6379"
		} else {
			address = "localhost:6379"
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong2, err := redisClient.Ping().Result()
		if err != nil {
			log.Error("NewDataStore redis", err)
		}
		log.Debug("NewDB", pong2)
	}
	return &RelDB{url, postgresClient, redisClient, 32}, nil
}

// SetAsset stores an asset into postgres
func (rdb *RelDB) SetAsset(asset dia.Asset) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetAsset returns a dia.Asset by its symbol and name from postgres
func (rdb *RelDB) GetAsset(symbol, name string) (asset dia.Asset, err error) {
	var decimals string
	err = rdb.postgresClient.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1 and name=$2", symbol, name).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	// TO DO: Get Blockchain by name from postgres and add to asset
	return
}

// GetPage returns assets per page number. @hasNext is true iff there is a non-empty next page.
func (rdb *RelDB) GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error) {

	pagesize := rdb.pagesize
	skip := pagesize * pageNumber
	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip)
	if err != nil {
		return
	}
	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals)
		assets = append(assets, asset)
	}
	// Last page (or empty page)
	if len(rows.RawValues()) < int(pagesize) {
		hasNextPage = false
		return
	}
	// No next page
	nextPageRows, err := rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip+1)
	if len(nextPageRows.RawValues()) == 0 {
		hasNextPage = false
		return
	}
	hasNextPage = true
	return
}

// Count returns the number of assets stored in postgres
func (rdb *RelDB) Count() (count uint32, err error) {
	err = rdb.postgresClient.QueryRow(context.Background(), "select count(*) from asset").Scan(&count)
	if err != nil {
		return
	}
	return
}

// GetAvailablePairs returns all trading pairs on @exchange from exchangepair table
func (rdb *RelDB) GetAvailablePairs(exchange string) (pairs []dia.Pair, err error) {

	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,foreignname from exchangepair where exchange=$1", exchange)
	for rows.Next() {
		pair := dia.Pair{}
		rows.Scan(&pair.Symbol, &pair.ForeignName)
		pairs = append(pairs, pair)
	}

	return []dia.Pair{}, nil
}

// SetExchangePair adds @pair to exchangepair table
func (rdb *RelDB) SetExchangePair(exchange string, pair dia.Pair) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into exchangepair(symbol,foreignname,exchange) values ($1,$2,$3)", pair.Symbol, pair.ForeignName, exchange)
	if err != nil {
		return err
	}
	return nil
}

func getPostgresKeyFromSecrets() string {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../../secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file should have exactly one line")
	}
	return lines[0]
}
