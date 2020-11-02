package defiscrapers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	compoundcontract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/compound"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
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
	assets["UNI"] = "0xd88b94128ff2b8cf2d7886cd1c1e46757418ca2a"




	decimals := make(map[string]int)
	decimals["ETH"] = 18
	decimals["BAT"] = 18
	decimals["DAI"] = 18
	decimals["REP"] = 18
	decimals["USDC"] = 6
	decimals["USDT"] = 6
	decimals["WBTC"] = 8
	decimals["ZRX"] = 18
	decimals["UNI"] = 18


	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	return &CompoundProtocol{scraper: scraper, protocol: protocol, assets: assets, decimals: decimals, connection: connection}
}

func NewCreamFinance(scraper *DefiScraper, protocol dia.DefiProtocol) *CompoundProtocol {
	assets := make(map[string]string)
	//https://compound.finance/docs#networks
	assets["ETH"] = "0xD06527D5e56A3495252A528C4987003b712860eE"   //done
	assets["USDC"] = "0x44fbeBd2F576670a6C33f6Fc0B00aA8c5753b322"  //done
	assets["USDT"] = "0x797AAB1ce7c01eB727ab980762bA88e7133d2157"  //done
	assets["MTA"] = "0x3623387773010d9214B10C551d6e7fc375D31F58"   //done
	assets["COMP"] = "0x19D1666f543D42ef17F66E376944A22aEa1a8E46"  //done
	assets["BAL"] = "0xcE4Fe9b4b8Ff61949DCfeB7e03bc9FAca59D2Eb3"   //done
	assets["YFI"] = "0xCbaE0A83f4f9926997c8339545fb8eE32eDc6b76"   //done
	assets["yCRV"] = "0x9baF8a5236d44AC410c0186Fe39178d5AAD0Bb87"  //done
	assets["LINK"] = "0x697256CAA3cCaFD62BB6d3Aa1C7C5671786A5fD9"  //done
	assets["CREAM"] = "0x892B14321a4FCba80669aE30Bd0cd99a7ECF6aC0" //done
	assets["LEND"] = "0x8B86e0598616a8d4F1fdAE8b59E55FB5Bc33D0d6"  //done
	assets["CRV"] = "0xc7Fd8Dcee4697ceef5a2fd4608a7BD6A94C77480"   //done
	assets["BUSD"] = "0x1FF8CDB51219a8838b52E9cAc09b71e591BC998e"  //done
	assets["yUSD"] = "0x4EE15f44c6F0d8d1136c83EfD2e8E4AC768954c6"  //done
	assets["SUSHI"] = "0x338286C0BC081891A4Bda39C7667ae150bf5D206" //done
	assets["FTT"] = "0x10FDBD1e48eE2fD9336a482D746138AE19e649Db"   //done
	assets["yETH"] = "0x01da76DEa59703578040012357b81ffE62015C2d"  //done
	assets["SRM"] = "0xef58b2d5A1b8D3cDE67b8aB054dC5C831E9Bc025"   //done
	assets["UNI"] = "0xe89a6D0509faF730BD707bf868d9A2A744a363C7"   //done

	assets["renBTC"] = "0x17107f40d70f4470d20CB3f138a052cAE8EbD4bE" //done

	decimals := make(map[string]int)
	decimals["ETH"] = 18
	decimals["USDT"] = 6
	decimals["USDC"] = 6
	decimals["COMP"] = 18
	decimals["BAL"] = 18
	decimals["YFI"] = 18
	decimals["yCRV"] = 18
	decimals["LINK"] = 18
	decimals["CREAM"] = 18
	decimals["LEND"] = 18
	decimals["CRV"] = 18
	decimals["renBTC"] = 8
	decimals["BUSD"] = 18
	decimals["MTA"] = 18
	decimals["yUSD"] = 18
	decimals["SUSHI"] = 18
	decimals["FTT"] = 18
	decimals["yETH"] = 18
	decimals["SRM"] = 6
	decimals["UNI"] = 18

	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
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
