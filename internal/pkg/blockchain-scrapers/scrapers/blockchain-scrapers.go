package blockchainscrapers

import (
	"log"
	"time"

	"github.com/blockstatecom/go-bitcoind"
	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	USESSL = false
)

type BlockchainScraper struct {
	bitcoind    *bitcoind.Bitcoind
	client      *dia.Client
	symbol      string
	elapsedTime time.Duration
}

func numberOfCoinsFor(blockNumber float64, subsidy float64, totalCoins float64, rewardModulo int64) float64 {
	var i int64 = 1
	for i < int64(blockNumber) {
		if i%rewardModulo == 0 {
			subsidy = subsidy / 2
		}
		totalCoins += subsidy
		i++
	}
	return totalCoins
}

func (s *BlockchainScraper) numberOfCoinsFor(blockNumber float64) float64 {
	switch s.symbol {
	case "LTC":
		return numberOfCoinsFor(blockNumber, 50.0, 50.0, 840000)
	default:
		return numberOfCoinsFor(blockNumber, 50.0, 50.0, 210000)
	}
}

func NewScraper(client *dia.Client, symbol string, serverHost string, serverPort int, user string, passwd string, elapsedTime int) *BlockchainScraper {
	bc, err := bitcoind.New(serverHost, serverPort, user, passwd, USESSL)
	if err != nil {
		log.Println(err)
		return nil
	}
	t := &BlockchainScraper{
		bitcoind:    bc,
		client:      client,
		symbol:      symbol,
		elapsedTime: time.Duration(elapsedTime) * time.Second,
	}
	return t
}

func (s *BlockchainScraper) Run() {
	for {
		rinfo, err := s.bitcoind.GetBlockchainInfo()
		if err == nil {
			log.Println("GetBlockchainInfo:", rinfo)
			m := time.Unix(rinfo.Mediantime, 0)
			l := time.Now().Sub(m)
			circulatingSupply := s.numberOfCoinsFor(rinfo.Blocks)
			log.Println("ElapsedTime block:", l, circulatingSupply)
			if l < s.elapsedTime {
				err = s.client.SendSupply(&dia.Supply{
					Symbol:            s.symbol,
					CirculatingSupply: circulatingSupply,
				})
				if err != nil {
					log.Println("Err communicating with api:", err)
				}
			}
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(10 * time.Second)
	}
}
