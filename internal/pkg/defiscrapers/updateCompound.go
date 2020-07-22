package defiscrapers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	compoundcontract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/compound"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type CompoundRate struct {
	SupplyRate  float64
	BorrowRate  float64
	TotalSupply *big.Int
	Decimal     *big.Int
	Symbol      string
}

type CompoundProtocol struct {
	scraper    *DefiScraper
	protocol   dia.DefiProtocol
	connection *ethclient.Client
	assets     map[string]string // assetname and address
}

func NewCompound(scraper *DefiScraper, protocol dia.DefiProtocol) *CompoundProtocol {
	assets := make(map[string]string)
	//https://compound.finance/docs#networks
	assets["BAT"] = "0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"
	assets["DAI"] = "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"
	assets["ETH"] = "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
	assets["REP"] = "0x158079ee67fce2f58472a96584a73c7ab9ac95c1"
	assets["SAI"] = "0xf5dce57282a584d2746faf1593d3121fcac444dc"
	assets["USDC"] = "0x39aa39c021dfbae8fac545936693ac917d5e7563"
	assets["USDT"] = "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"
	assets["WBTC"] = "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"
	assets["ZRX"] = "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"

	// connection, err := ethclient.Dial("https://mainnet.infura.io/v3/806b0419b2d041869fc83727e0043236")
	connection, err := ethclient.Dial("https://mainnet.infura.io/v3/f619e28e13f0428cba6f9243b09d4af0")
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	return &CompoundProtocol{scraper: scraper, protocol: protocol, assets: assets, connection: connection}
}

func (proto *CompoundProtocol) fetch(asset string) (rate CompoundRate, err error) {
	var contract *compoundcontract.CTokenCaller
	contract, err = compoundcontract.NewCTokenCaller(common.HexToAddress(proto.assets[asset]), proto.connection)
	supplyInterestRate, err := contract.SupplyRatePerBlock(&bind.CallOpts{})
	if err != nil {
		return
	}
	borrowInterestRate, err := contract.BorrowRatePerBlock(&bind.CallOpts{})
	if err != nil {
		return
	}
	totalSupply, err := contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return
	}
	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return
	}
	rate = CompoundRate{Symbol: asset, Decimal: decimals, BorrowRate: proto.calculateAPY(borrowInterestRate), SupplyRate: proto.calculateAPY(supplyInterestRate), TotalSupply: totalSupply}
	return
}

func (proto *CompoundProtocol) calculateAPY(rate *big.Int) float64 {
	// https://compound.finance/docs#protocol-math
	//Calculate APY
	var blocksPerDay float64
	var daysPerYear float64
	var ethMantissa float64
	ethMantissa = 1e18
	blocksPerDay = 4 * 60 * 24
	daysPerYear = 365

	rateInt, err := strconv.ParseFloat(rate.String(), 64)
	if err != nil {
		return 0
	}
	rates := math.Pow((rateInt/ethMantissa*blocksPerDay)+1, daysPerYear-1) - 1
	return rates * 100
}

func (proto *CompoundProtocol) fetchALL() (rates []CompoundRate, err error) {
	for asset, _ := range proto.assets {
		Compoundrate, err := proto.fetch(asset)
		if err != nil {
			log.Error(err)
		}
		rates = append(rates, Compoundrate)
	}
	return
}

func (proto *CompoundProtocol) UpdateRate() error {
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

func (proto *CompoundProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	usdcMarket, err := proto.fetch("USDC")
	if err != nil {
		return err
	}
	ethMarket, err := proto.fetch("ETH")
	if err != nil {
		return err
	}
	totalSupplyUSDC, err := strconv.ParseFloat(usdcMarket.TotalSupply.String(), 64)
	fmt.Println("usdcMarket: ", usdcMarket)
	if err != nil {
		return err
	}
	totalSupplyDecimal, err := strconv.ParseFloat(usdcMarket.Decimal.String(), 64)
	if err != nil {
		return err
	}

	totalBorrowETH, err := strconv.ParseFloat(ethMarket.TotalSupply.String(), 64)
	if err != nil {
		return err
	}

	totalBorrowETHDecimal, err := strconv.ParseFloat(ethMarket.Decimal.String(), 64)
	if err != nil {
		return err
	}

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSDC / math.Pow(10, totalSupplyDecimal),
		TotalETH:  totalBorrowETH / math.Pow(10, totalBorrowETHDecimal),
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
