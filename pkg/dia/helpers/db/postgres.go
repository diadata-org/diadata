package db

import (
	"bufio"
	"context"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const (
	postgresKey = "postgres_credentials.txt"

//	reconnectWaitSeconds = 5
//	maxRetry             = 120
)

func PostgresDatabase() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), GetPostgresURL())
	if err != nil {
		log.Error(err)
	}
	return pool

}

/*
var postgresClient *pgxpool.Pool
func GetPostgresClient() (*pgx.Conn, error) {
	var err error
	log.Info("connect to postgres server...")
	postgresClient, err = pgxpool.Connect(context.Background(), GetPostgresURL())
	if err != nil {
		log.Error(err)
		return &pgx.Conn{}, err
	}
	log.Info("...connection to postgres server established.")

	return postgresClient.Conn(), err
}

func PostgresDatabase() *pgx.Conn {
	connected := false
	var err error
	if postgresClient == nil {
		// during startup - if it does not exist, create it
		postgresClient, err = GetPostgresClient()
		if err == nil {
			connected = true
		}
	}

	count := 0
	for (!connected || postgresClient.IsClosed()) && count < maxRetry {
		log.Info("Connection to Postgres was lost. Waiting for 5s...")
		time.Sleep(reconnectWaitSeconds * time.Second)
		log.Info("Reconnecting to Postgres...")
		postgresClient, err = GetPostgresClient()
		if err == nil {
			connected = true
		}
		count++
	}
	return postgresClient
}
*/

func GetPostgresURL() (url string) {
	if utils.Getenv("USE_ENV", "false") == "true" {
		return "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB")
	}
	if utils.Getenv("EXEC_MODE", "local") == "production" {
		return "postgresql://postgres/postgres?user=postgres&password=" + getPostgresKeyFromSecrets()
	}
	return "postgresql://localhost/postgres?user=postgres&password=" + getPostgresKeyFromSecrets()
}

func getPostgresKeyFromSecrets() string {
	var lines []string
	var file *os.File
	var err error
	if utils.Getenv("EXEC_MODE", "") == "production" {
		pwd, _ := os.Getwd()
		log.Info("current directory: ", pwd)
		file, err = os.Open("/run/secrets/" + "postgres_credentials")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("current directory: ", os.Getenv("GOPATH"))
		file, err = os.Open(os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file should have exactly one line")
	}
	err = file.Close()
	if err != nil {
		log.Error(err)
	}
	return lines[0]
}
