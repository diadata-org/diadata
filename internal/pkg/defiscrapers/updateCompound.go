package defiscrapers

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"

	compoundcontract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/compound"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const (
	cTokenDecimals = 8
)

type CompoundRate struct {
	SupplyRate    float64
	BorrowRate    float64
	TotalSupply   *big.Int
	TotalBorrow   *big.Int
	TotalReserves *big.Int
	Cash          *big.Int
	Decimal       int
	Symbol        string
}

type CompoundProtocol struct {
	scraper    *DefiScraper
	protocol   dia.DefiProtocol
	connection *ethclient.Client
	assets     map[string]string // assetname and address
	decimals   map[string]int
}

func NewCompound(scraper *DefiScraper, protocol dia.DefiProtocol) *CompoundProtocol {
	assets := make(map[string]string)
	//https://compound.finance/docs#networks
	assets["ETH"] = "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
	assets["BAT"] = "0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"
	assets["DAI"] = "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"
	assets["REP"] = "0x158079ee67fce2f58472a96584a73c7ab9ac95c1"
	assets["USDC"] = "0x39aa39c021dfbae8fac545936693ac917d5e7563"
	assets["USDT"] = "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"
	assets["WBTC"] = "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"
	assets["ZRX"] = "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"

	decimals := make(map[string]int)
	decimals["ETH"] = 18
	decimals["BAT"] = 18
	decimals["DAI"] = 18
	decimals["REP"] = 18
	decimals["USDC"] = 6
	decimals["USDT"] = 6
	decimals["WBTC"] = 8
	decimals["ZRX"] = 18

	executionMode := os.Getenv("EXEC_MODE")
	var connection *ethclient.Client
	err := errors.New("")
	if executionMode == "production" {
		connection, err = ethclient.Dial("http://159.69.120.42:8545/")
		if err != nil {
			log.Error("Error connecting Eth Client")
		}
	} else {
		connection, err = ethclient.Dial("https://mainnet.infura.io/v3/806b0419b2d041869fc83727e0043236")
		if err != nil {
			log.Error("Error connecting Eth Client")
		}
	}

	return &CompoundProtocol{scraper: scraper, protocol: protocol, assets: assets, decimals: decimals, connection: connection}
}

func (proto *CompoundProtocol) fetch(asset string) (rate CompoundRate, err error) {
	var contract *compoundcontract.CTokenCaller
	contract, err = compoundcontract.NewCTokenCaller(common.HexToAddress(proto.assets[asset]), proto.connection)

	// var underlyingContract *compoundcontract.CErc20Caller
	// underlyingContract, err = compoundcontract.NewCErc20Caller(common.HexToAddress(proto.assets[asset]), proto.connection)
	// address, _ := underlyingContract.Underlying(&bind.CallOpts{})

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
	totalBorrow, err := contract.TotalBorrows(&bind.CallOpts{})
	if err != nil {
		return
	}
	totalReserves, err := contract.TotalReserves(&bind.CallOpts{})
	if err != nil {
		return
	}
	cash, err := contract.GetCash(&bind.CallOpts{})
	if err != nil {
		return
	}
	rate = CompoundRate{
		Symbol:        asset,
		Decimal:       proto.decimals[asset],
		BorrowRate:    proto.calculateAPY(borrowInterestRate),
		SupplyRate:    proto.calculateAPY(supplyInterestRate),
		TotalSupply:   totalSupply,
		TotalBorrow:   totalBorrow,
		TotalReserves: totalReserves,
		Cash:          cash,
	}
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
	for asset := range proto.assets {
		Compoundrate, err := proto.fetch(asset)
		if err != nil {
			log.Error("error fetching asset: ", err)
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

func (proto *CompoundProtocol) getTotalSupply() (float64, error) {
	// total supply for a market is explained here:
	// https://compound.finance/docs/ctokens#get-cash
	markets, err := proto.fetchALL()
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for i := 0; i < len(markets); i++ {
		supply, err := strconv.ParseFloat(markets[i].Cash.String(), 64)
		if err != nil {
			return 0, err
		}
		normalizedSupply := supply / math.Pow10(int(markets[i].Decimal))
		price, err := utils.GetCoinPrice(markets[i].Symbol)
		if err != nil {
			fmt.Println("error getting prices: ", err)
		}
		sum += normalizedSupply * price
	}
	return sum, nil
}

func (proto *CompoundProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	totalValueLocked, err := proto.getTotalSupply()
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
