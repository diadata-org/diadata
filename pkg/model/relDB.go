package models

import (
	"bufio"
	"context"
	"os"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// RelDatastore is a (persistent) relational database with an additional redis caching layer
type RelDatastore interface {
	GetAvailablePairs(exchange string) (pairs []dia.ExchangePair, err error)
	SetExchangePair(exchange string, pair dia.ExchangePair)

	// Assets methods
	// Persistent
	SetAsset(asset dia.Asset) error
	GetAsset(symbol, name string) (asset dia.Asset, err error)
	GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error)
	Count() (count uint32, err error)
	// Caching
	SetAssetCache(asset dia.Asset) error
	GetAssetCache(symbol, name string) (dia.Asset, error)
	CountCache() (uint32, error)
	SetAvailablePairsCache(exchange string, pairs []dia.ExchangePair) error
}

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

func NewCachingLayer() (*RelDB, error) {
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

// GetAvailablePairs returns all trading pairs on @exchange from exchangepair table
func (rdb *RelDB) GetAvailablePairs(exchange string) (pairs []dia.ExchangePair, err error) {

	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,foreignname from exchangepair where exchange=$1", exchange)
	for rows.Next() {
		pair := dia.ExchangePair{}
		rows.Scan(&pair.Symbol, &pair.ForeignName)
		pairs = append(pairs, pair)
	}

	return []dia.ExchangePair{}, nil
}

// SetExchangePair adds @pair to exchangepair table
func (rdb *RelDB) SetExchangePair(exchange string, pair dia.ExchangePair) error {
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
