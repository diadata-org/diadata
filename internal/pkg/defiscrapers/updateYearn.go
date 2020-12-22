package defiscrapers

import (
	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/yearn"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type YearnProtocol struct {
	scraper      *DefiScraper
	protocol     dia.DefiProtocol
	connection   *ethclient.Client
	yearnManager *yearn.YearnManager
}

func NewYearn(scraper *DefiScraper, protocol dia.DefiProtocol, rpcUrl, aprOracleAddress string) *YearnProtocol {
	connection, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Error("Error connecting Eth Client")
	}
	aprOracle, err := yearn.NewAproracle(common.HexToAddress(aprOracleAddress), connection)
	if err != nil {
		log.Fatalf("Failed to bind APR oracle contract: %v", err)
	}
	yearnManager := yearn.NewYearnManager(aprOracle)

	return &YearnProtocol{
		scraper:      scraper,
		protocol:     protocol,
		connection:   connection,
		yearnManager: yearnManager,
	}
}

func (proto *YearnProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
	compoundAPRResult, err := proto.yearnManager.GetAllCompoundAPR()
	if err != nil {
		return err
	}
	go proto.yearnManager.Publish(proto.protocol.Name, "CDAI", compoundAPRResult.CDAI, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CBAT", compoundAPRResult.CBAT, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CETH", compoundAPRResult.CETH, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CREP", compoundAPRResult.CREP, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CSAI", compoundAPRResult.CSAI, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CUSDC", compoundAPRResult.CUSDC, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CWBTC", compoundAPRResult.CWBTC, proto.scraper.RateChannel())
	go proto.yearnManager.Publish(proto.protocol.Name, "CZRC", compoundAPRResult.CZRC, proto.scraper.RateChannel())
	proto.yearnManager.GetAllCompoundAPR()
	proto.connection.Close()
	log.Info("Update complete")
	return nil
}

func (proto *YearnProtocol) UpdateState() error {
	//NO OP
	return nil
}
