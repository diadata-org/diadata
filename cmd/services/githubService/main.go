package main

import (
	"bufio"
	"flag"
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

	// Set basic structure
	var nameUser = flag.String("username", "diadata-org", "github username")
	var nameRepository = flag.String("repository", "diadata", "name of the repository")
	flag.Parse()
	apiKey := getAPIKeyFromSecrets()
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	// Get latest commit from database, if existent
	latestCommit, err := ds.GetLatestCommit(*nameUser, *nameRepository)
	if err != nil {
		log.Fatal("error getting latest commit: ", err)
	}

	// If no commit is in the DB, fetch all commits up until now
	if (latestCommit) == (models.GithubCommit{}) {
		commits, err := githubservice.FetchAllCommits(*nameUser, *nameRepository, 100, apiKey)
		if err != nil {
			log.Fatal("error fetching all commits: ", err)
		}
		for _, commit := range commits {
			err = ds.SetCommit(&commit)
			if err != nil {
				log.Fatal("error setting commit: ", err)
			}
			log.Info("set commit: ", commit)
		}
	}

	// Continuously update commits once every 6h
	// ticker := time.NewTicker(6 * time.Hour)
	ticker := time.NewTicker(20 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				latestCommit, err := ds.GetLatestCommit(*nameUser, *nameRepository)
				if err != nil {
					log.Fatal("error getting latest commit: ", err)
				}
				lastTimestamp := latestCommit.Timestamp
				commits, err := githubservice.FetchCommitsByDate("diadata-org", "diadata", apiKey, lastTimestamp, time.Now())
				if err != nil {
					log.Fatal(err)
				}
				for _, commit := range commits {
					err = ds.SetCommit(&commit)
					if err != nil {
						log.Fatal(err)
					}
					log.Info("set commit: ", commit)
				}

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
