package main

import (
	"archive/zip"
	"database/sql"
	"fmt"
	db2 "github.com/diadata-org/diadata/pkg/dia/helpers/db"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	zipPath = "latest.zip"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	urlFolderPrefix := utils.Getenv("URL_FOLDER_PREFIX", "/")
	if !strings.HasPrefix(urlFolderPrefix, "/") {
		urlFolderPrefix = "/" + urlFolderPrefix
	}

	r.GET(urlFolderPrefix+"/latest", func(c *gin.Context) {
		if _, err := os.Stat(zipPath); os.IsNotExist(err) {
			c.String(http.StatusNotFound, "No latest data available.")
			return
		}

		c.FileAttachment(zipPath, "latest.zip")
	})

	go func() {
		for {
			err := runPostgresScripts()
			if err != nil {
				log.Printf("Failed to run Postgres scripts: %v", err)
			} else {
				log.Println("Postgres scripts executed successfully.")
			}

			err = createZipFile()
			if err != nil {
				log.Fatal("Failed to create zip file: %v", err)
			} else {
				log.Println("Zip file created successfully.")
			}

			time.Sleep(24 * time.Hour)
		}
	}()

	log.Fatal(r.Run(utils.Getenv("LISTEN_PORT", ":8080")))
}

var postgresUrl string

func getPostgresUrl() string {
	if postgresUrl != "" {
		return postgresUrl
	}
	postgresUrl := db2.GetPostgresURL()
	return postgresUrl
}

func runPostgresScripts() error {
	db, err := sql.Open("postgres", getPostgresUrl())

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("error connecting to database: ", err)
		}
	}(db)

	queries := []struct {
		Query     string
		OutputCSV string
	}{
		{
			Query:     "COPY (SELECT * FROM asset) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_assets.csv",
		},
		{
			Query:     "COPY (SELECT * FROM blockchain ) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_blockchain.csv",
		},
		{
			Query:     "COPY (SELECT * FROM exchange) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_exchange.csv",
		},
		{
			Query:     "COPY (SELECT * FROM exchange WHERE centralized = true) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_cex.csv",
		},
		{
			Query:     "COPY (SELECT * FROM exchange WHERE  centralized = false) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_dex.csv",
		},
		{
			Query:     "COPY (SELECT * FROM exchangepair) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_exchangepair.csv",
		},
		{
			Query:     "COPY (SELECT * FROM pool WHERE pool_id IN (SELECT DISTINCT pool_id FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days')) TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_pool.csv",
		},
		{
			Query:     "COPY (SELECT * FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days') TO STDOUT WITH (format csv, delimiter ';')",
			OutputCSV: "output_poolasset.csv",
		},
	}

	for _, query := range queries {
		filePath := query.OutputCSV

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("Zip file save error.")
			}
		}(file)

		cmd := exec.Command("psql", getPostgresUrl(), "-c", query.Query)
		cmd.Stdout = file
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run Postgres script: %v", err)
		}
	}

	return nil
}

func createZipFile() error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	files, err := filepath.Glob("output*.csv")
	if err != nil {
		return fmt.Errorf("failed to find CSV files: %v", err)
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer f.Close()

		zipEntry, err := zipWriter.Create(filepath.Base(file))
		if err != nil {
			return fmt.Errorf("failed to create zip entry: %v", err)
		}

		_, err = io.Copy(zipEntry, f)
		if err != nil {
			return fmt.Errorf("failed to copy file to zip: %v", err)
		}
	}

	return nil
}
