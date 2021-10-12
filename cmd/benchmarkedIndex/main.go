package main

import (
	"encoding/csv"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
	"indexEngine/db"
	"indexEngine/sftp"
	"os"
	"strings"
)

func main() {
	sftpClient := sftp.GetConnection()

	files := strings.Split(utils.Getenv("SFTP_FILELIST", ""), ";")
	for _, file := range files {
		pathParts := strings.Split(file, "/")
		filename := pathParts[len(pathParts)-1]
		if _, err := os.Stat("./" + filename); err == nil {
			errRemove := os.Remove("./" + filename)
			if errRemove != nil {
				log.Error(filename, ":", errRemove)
			}

		}
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
		indexed := indexedTimestamp[len(indexedTimestamp)-1]

		// save indexed into influx
		symbol := strings.Split(file, "_")[1]
		db.SaveIndexEngineIntoInflux(indexed, symbol)
		errFileHandle := fileHandle.Close()
		if errFileHandle != nil {
			log.Error(errFileHandle)
		}
	}
}
