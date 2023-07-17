package snapshotServer

import (
	"archive/zip"
	"database/sql"
	"fmt"
	db2 "github.com/diadata-org/diadata/pkg/dia/helpers/db"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	zipPath = "latest.zip"
)

func main() {
	r := gin.Default()

	r.GET("/latest", func(c *gin.Context) {
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
				log.Printf("Failed to create zip file: %v", err)
			} else {
				log.Println("Zip file created successfully.")
			}

			time.Sleep(24 * time.Hour)
		}
	}()

	log.Fatal(r.Run(":8080"))
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

	queries := []string{
		"COPY (SELECT * FROM asset) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM blockchain WHERE \"name\" IN (SELECT DISTINCT blockchain FROM asset WHERE asset_id IN (SELECT DISTINCT asset_id FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days'))) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM exchange WHERE centralized = true) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM exchange WHERE \"name\" IN (SELECT distinct exchange FROM pool WHERE pool_id IN (SELECT DISTINCT pool_id FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days'))) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM exchangepair) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM pool WHERE pool_id IN (SELECT DISTINCT pool_id FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days')) TO STDOUT WITH (format csv, delimiter ';')",
		"COPY (SELECT * FROM poolasset WHERE \"time_stamp\" >= now() - interval '7 days') TO STDOUT WITH (format csv, delimiter ';')",
	}

	for i, query := range queries {
		filePath := fmt.Sprintf("output%d.csv", i+1)

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer file.Close()

		cmd := exec.Command("psql", getPostgresUrl(), "-c", query)
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
