package main

import (
	"benchmarkedIndex/db"
	"benchmarkedIndex/sftp"
	"encoding/csv"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {

	files := strings.Split(utils.Getenv("SFTP_FILELIST", ""), ";")
	for _, file := range files {
		sftpClient := sftp.GetConnection()
		filename := getFilename(file)
		symbol := getSymbol(file)
		errDownload := sftp.DownloadFile(*sftpClient, file, "./"+filename)
		if errDownload != nil {
			log.Error(errDownload)
			return
		}
		fileHandle, err := os.Open("./" + filename)
		if err != nil {
			log.Error(err)
		}
		csvReader := csv.NewReader(fileHandle)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Error(err)
		}

		indexedTimestamp := db.CreateIndexTimestampList(data)
		if utils.Getenv("INDEX_HISTORICAL_DATA", "false") == "true" {
			for _, indexed := range indexedTimestamp {
				db.SaveIndexEngineIntoInflux(indexed, symbol)
			}
		} else {
			indexed := indexedTimestamp[len(indexedTimestamp)-1]
			db.SaveIndexEngineIntoInflux(indexed, symbol)
		}

		errFileHandle := fileHandle.Close()
		if errFileHandle != nil {
			log.Error(errFileHandle)
		}
	}
}

func getFilename(file string) string {
	pathParts := strings.Split(file, "/")
	filename := pathParts[len(pathParts)-1]
	if _, err := os.Stat("./" + filename); err == nil {
		errRemove := os.Remove("./" + filename)
		if errRemove != nil {
			log.Error(filename, ":", errRemove)
		}
	}
	return filename
}

func getSymbol(file string) string {
	symbol := strings.Split(file, "_")[1]
	return symbol
}
