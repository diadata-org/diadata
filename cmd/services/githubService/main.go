package main

import (
	"bufio"
	"os"
	"time"

	githubservice "github.com/diadata-org/diadata/internal/pkg/githubService"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (
	apiKeyFilename = "github_key.txt"
)

func main() {

	// TO DO: Upon start, get last timestamp from commit table. If table is empty, run GetAllCommits and then
	// GetCommitInfo in time ranges with ticker.
	// If table is not empty, get last timestamp and run GetAllCommits from last timestamp until now.

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	// Continuously update supplies once every 6h
	// ticker := time.NewTicker(6 * time.Hour)
	ticker := time.NewTicker(20 * time.Second)
	// timeInit := time.Now().AddDate(0, -1, -10)
	// timeFinal := time.Now().AddDate(0, -1, -5)
	// fmt.Println("time: ", timeInit.Format("2006-01-02T15:04:05+00:00"))

	go func() {
		for {
			select {
			case <-ticker.C:
				apiKey := getAPIKeyFromSecrets()
				// repo, err := githubservice.GetCommitInfo("diadata-org", "diadata", apiKey, timeInit, timeFinal)
				repo, err := githubservice.GetAllCommits("diadata-org", "diadata", apiKey)
				if err != nil {
					log.Fatal(err)
				}
				var commit models.GithubCommit
				err = ds.SetCommit(&commit)
				if err != nil {
					log.Fatal(err)
				}
				log.Info("set repository: ", repo)

			}
		}
	}()
	select {}

}

// getAPIKeyFromSecrets returns a github api key
func getAPIKeyFromSecrets() string {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/" + apiKeyFilename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../../secrets/" + apiKeyFilename)
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
