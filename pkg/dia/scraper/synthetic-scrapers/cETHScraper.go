package synthscrapers

import (
	ceth "github.com/diadata-org/diadata/config/synthContracts/cETH"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type cETHScraper struct {
	RestClient   *ethclient.Client
	synthChannel chan dia.SynthAssetSupply
	blockchain   string
}

func NewcETHScraper() *cETHScraper {
	log.Info("TO DO")

	scraper := &cETHScraper{
		blockchain: dia.ETHEREUM,
	}

	connection, err := ethclient.Dial(utils.Getenv(scraper.blockchain+"_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	scraper.RestClient = connection
	go scraper.mainLoop()

	return scraper
}

func (scraper *cETHScraper) mainLoop() {
	// TO DO
	// Periodically run
	scraper.FetchSynthSupply()
}

func (scraper *cETHScraper) GetSynthSupplyChannel() chan dia.SynthAssetSupply {
	return scraper.synthChannel
}

func (scraper *cETHScraper) FetchSynthSupply() error {
	// TO DO...
	filterer, err := ceth.NewERC20Caller(common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"), scraper.RestClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
	}
	supply, err := filterer.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Info("get supply: ", err)
	}
	log.Info("supply: ", supply.Int64())
	return nil
}
