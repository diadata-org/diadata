package main

import (
	"bufio"
	"flag"
	githubservice2 "github.com/diadata-org/diadata/internal/githubService"
	"os"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (
	apiKey = "api_github"
)

func main() {

	// Set basic structure
	var nameUser = flag.String("username", "diadata-org", "github username")
	var nameRepository = flag.String("repository", "diadata", "name of the repository")
	flag.Parse()
	apiSecret := getAPIKeyFromSecrets()
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	// Get latest commit from database, if existent
	latestCommit, err := ds.GetLatestCommit(*nameUser, *nameRepository)
	if err != nil {
		log.Fatal("error getting latest commit: ", err)
	}

	// If no commit is in the DB, fetch all commits until now
	if (latestCommit) == (models.GithubCommit{}) {
		log.Info("populate database...")
		commits, err := githubservice2.FetchAllCommits(*nameUser, *nameRepository, 100, apiSecret)
		if err != nil {
			log.Fatal("error fetching all commits: ", err)
		}
		for _, commit := range commits {
			err = ds.SetCommit(commit)
			if err != nil {
				log.Fatal("error setting commit: ", err)
			}
			log.Info("set commit: ", commit)
		}
		log.Info("...initial database population done.")
	}

	// Continuously update commits once every 6h
	// ticker := time.NewTicker(6 * time.Hour)
	ticker := time.NewTicker(20 * time.Second)

	go func() {
		for range ticker.C {
			log.Info("update github commits")
			latestCommit, err := ds.GetLatestCommit(*nameUser, *nameRepository)
			if err != nil {
				log.Fatal("error getting latest commit: ", err)
			}
			// Remark: if between two ticker signals no new commits are added, the latest commit is fetched again nevertheless,
			// because FetchCommitsByDate includes the borders. This does not hurt as the set of tags is the same and hence, data
			// is only stored once in influx.
			commits, err := githubservice2.FetchCommitsByDate("diadata-org", "diadata", apiKey, latestCommit.Timestamp, time.Now())
			if err != nil {
				log.Fatal(err)
			}
			for _, commit := range commits {
				err = ds.SetCommit(commit)
				if err != nil {
					log.Fatal(err)
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
		file, err = os.Open("/run/secrets/" + apiKey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../../secrets/" + apiKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

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
