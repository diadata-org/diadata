package db

import (
	"bufio"
	"context"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	postgresKey = "postgres_credentials.txt"
)

func GetPostgresClient() *pgx.Conn {
	var postgresClient *pgx.Conn
	var err error

	log.Info("connect to postgres server...")
	postgresClient, err = pgx.Connect(context.Background(), GetPostgresURL())
	if err != nil {
		log.Error(err)
		return nil
	}
	log.Info("...connection to postgres server established.")

	return postgresClient
}

func GetPostgresURL() (url string) {
	if os.Getenv("USE_ENV") == "true" {
		return "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/"+ os.Getenv("POSTGRES_DATABASE")
	}
	if utils.Getenv("EXEC_MODE","") == "production" {
		return "postgresql://postgres/postgres?user=postgres&password=" + getPostgresKeyFromSecrets()
	}
	return "postgresql://localhost/postgres?user=postgres&password=" + getPostgresKeyFromSecrets()
}

func getPostgresKeyFromSecrets() string {
	var lines []string
	var file *os.File
	var err error
	if  utils.Getenv("EXEC_MODE","") == "production" {
		pwd, _ := os.Getwd()
		log.Info("current directory: ", pwd)
		file, err = os.Open("/run/secrets/" + "postgres_credentials")
		if err != nil {
			log.Fatal(err)
		}
	} else {
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

