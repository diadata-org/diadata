package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgconn"

	nftsource "github.com/diadata-org/diadata/internal/pkg/nftService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

var (
	log       = logrus.New()
	nftSource *string
	secret    *string
)

const (
	fetchPeriodMinutes = 24 * 60
)

func init() {
	nftSource = flag.String("source", "Opensea", "Data source for nft data collection")
	secret = flag.String("secret", "", "secret for asset source")
	flag.Parse()
}

// NewAssetScraper returns a scraper for assets on @exchange.
// For NewJSONReader @exchange is the folder in the config folder and @secret the filename.
func NewNFTClassCollector(nftSource string, secret string) nftsource.NFTClassSource {
	switch nftSource {
	case "Opensea":
		return nftsource.NewOpenseaNFTSource(secret)
	// case "nftlists":
	// 	return nftsource.NewJSONReader(nftSource)
	default:
		return nil
	}
}

func main() {

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}

	// Initial run.
	err = runNFTSource(relDB, *nftSource, *secret)
	if err != nil {
		log.Errorln("initial run: ", err)
	}

	// load gitcoin files for categorization of NFT classes.
	gitcoinfiles := iterateDirectory("nftClassesGitcoin")
	for _, file := range gitcoinfiles {
		log.Info("write file: ", file)
		path := "nftClassesGitcoin/" + file
		gitcoinCategories, err := readFile(path)
		if err != nil {
			log.Errorln("Error while reading  file", path, err)
			continue
		}
		setGitcoinCategories(gitcoinCategories, relDB)
	}

	// Afterwards, run every @fetchPeriodMinutes.
	ticker := time.NewTicker(fetchPeriodMinutes * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				runNFTSource(relDB, *nftSource, *secret)
			}
		}
	}()
	select {}
}

func runNFTSource(relDB *models.RelDB, source string, secret string) error {

	log.Println("Fetching asset from ", source)
	nftCollector := NewNFTClassCollector(source, secret)
	for {
		select {
		case receivedClass := <-nftCollector.NFTClass():
			// Set to persistent DB
			err := relDB.SetNFTClass(receivedClass)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code == "23505" {
						log.Infof("nft %v already in db. continue.", receivedClass)
						continue
					} else {
						log.Errorf("postgres error saving nft %v: %v", receivedClass, err)
					}
				} else {
					log.Errorf("Error saving nft %v: %v", receivedClass, err)
				}
			} else {
				log.Info("successfully set nft ", receivedClass)
			}

		case <-nftCollector.Close():
			return nil
		}
	}
}

// ----------------------------------------------------------------------------
// Gitcoin submission functionality
// ----------------------------------------------------------------------------

type GitcoinSubmission struct {
	AllItems []SubmissionItem `json:"NFTClasses"`
}

type SubmissionItem struct {
	Address      string `json:"Address"`
	Symbol       string `json:"Symbol"`
	Name         string `json:"Name"`
	Blockchain   string `json:"Blockchain"`
	ContractType string `json:"ContractType"`
	Category     string `json:"Category"`
}

// readFile reads a gitcoin submission json file and returns a slice of NFTClasse items.
func readFile(filename string) (items GitcoinSubmission, err error) {
	var (
		jsonFile  *os.File
		filebytes []byte
	)
	path := configCollectors.ConfigFileConnectors(filename, "")
	jsonFile, err = os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	filebytes, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	json.Unmarshal(filebytes, &items)
	return
}

// setGitcoinCategories writes all mappings from submissions into exchangesymbol table.
func setGitcoinCategories(submissions GitcoinSubmission, relDB *models.RelDB) {
	for _, nftClass := range submissions.AllItems {
		// Get ID of underlying NFTClass.
		nftClassID, err := relDB.GetNFTClassID(common.HexToAddress(nftClass.Address), nftClass.Blockchain)
		if err != nil {
			log.Error("class not in db yet. set: ", nftClass)
			err = relDB.SetNFTClass(dia.NFTClass{
				Address:      common.HexToAddress(nftClass.Address),
				Symbol:       nftClass.Symbol,
				Name:         nftClass.Name,
				Blockchain:   nftClass.Blockchain,
				ContractType: nftClass.ContractType,
				Category:     nftClass.Category,
			})
			if err != nil {
				log.Error("setting nft class: ", err)
			}
			continue
		}

		// Write into nftclass table and update Category.
		success, err := relDB.UpdateNFTClassCategory(nftClassID, nftClass.Category)
		if err != nil || !success {
			log.Error(err)
		}
	}
}

func iterateDirectory(foldername string) (files []string) {
	path := configCollectors.ConfigFileConnectors(foldername, "")
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fileExtension := filepath.Ext(path)

		if fileExtension == ".json" {
			files = append(files, info.Name())
		}
		return nil
	})
	return
}
