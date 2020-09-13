package defiscrapers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/fortubev2"
	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/fortubev2/token"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ForTubeRate struct {
	SupplyRate float64
	BorrowRate float64
	Cash       *big.Int
	Decimal    uint8
	Symbol     string
}

type ForTubeProtocol struct {
	scraper    *DefiScraper
	protocol   dia.DefiProtocol
	connection *ethclient.Client
}

func NewForTube(scraper *DefiScraper, protocol dia.DefiProtocol) *ForTubeProtocol {

	connection, err := ethhelper.NewETHClient()

	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	return &ForTubeProtocol{scraper: scraper, protocol: protocol, connection: connection}
}

func (proto *ForTubeProtocol) fetchALL() (rates []ForTubeRate, err error) {
	fortubeV2Address := common.HexToAddress(proto.protocol.Address)

	fortubeV2, err := fortubev2.NewFortubev2Caller(fortubeV2Address, proto.connection)
	if err != nil {
		log.Error(err)
	}
	markets, err := fortubeV2.GetAllMarkets(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	for _, tkn := range markets {

		tokenV2caller, err := token.NewTokenCaller(tkn, proto.connection)
		if err != nil {
			log.Error(err)
		}

		apr, err := tokenV2caller.APR(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		apy, err := tokenV2caller.APY(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		totalCash, err := tokenV2caller.TotalCash(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}

		underlying, err := tokenV2caller.Underlying(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		uTokenCaller, err := token.NewTokenCaller(underlying, proto.connection)
		if err != nil {
			log.Error(err)
		}

		uAsset, err := uTokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		uDecimals, err := uTokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}

		if underlying.Hex() == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" {
			// fasset == "fETH" {
			uDecimals = 18
			uAsset = "ETH"
		}

		rate := ForTubeRate{
			Symbol:     uAsset,
			Decimal:    uDecimals,
			BorrowRate: proto.calculateAPY(apr),
			SupplyRate: proto.calculateAPY(apy),
			Cash:       totalCash,
		}
		// fmt.Println(rate)
		rates = append(rates, rate)

	}

	return
}

func (proto *ForTubeProtocol) calculateAPY(rate *big.Int) float64 {

	//Calculate APY
	rateInt, err := strconv.ParseFloat(rate.String(), 64)
	if err != nil {
		return 0
	}
	rates := (rateInt / 1e18)
	return rates * 100
}

func (proto *ForTubeProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
	markets, err := proto.fetchALL()
	if err != nil {
		log.Error("error fetching rates %+v\n ", err)
		return err
	}

	for _, market := range markets {
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Symbol,
			Protocol:      proto.protocol.Name,
			LendingRate:   market.SupplyRate,
			BorrowingRate: market.BorrowRate,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.connection.Close()
		proto.scraper.RateChannel() <- asset
	}
	log.Info("Update complete")
	return nil
}

func (proto *ForTubeProtocol) getTotalValueLocked() (float64, error) {
	markets, err := proto.fetchALL()
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for i := 0; i < len(markets); i++ {
		supply := float64(0)
		if markets[i].Cash != nil {
			supply, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(markets[i].Cash), new(big.Float).SetFloat64(math.Pow10(int(markets[i].Decimal)))).Float64()
		}
		price, err := utils.GetCoinPrice(markets[i].Symbol)
		if err != nil {
			fmt.Println("error getting prices: ", err)
		}
		sum += supply * price

	}
	return sum, nil
}

func (proto *ForTubeProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	totalValueLocked, err := proto.getTotalValueLocked()
	if err != nil {
		log.Error(err)
	}
	ETHPrice, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return err
	}
	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalValueLocked,
		TotalETH:  totalValueLocked / ETHPrice,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())
	log.Info("Update State complete")

	return nil

}
