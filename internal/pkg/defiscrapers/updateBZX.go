package defiscrapers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	bzxcontract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/bzx"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type BZXRate struct {
	SupplyRate       *big.Int
	BorrowRate       *big.Int
	TotalSupply      *big.Int
	TotalSupplyAsset *big.Int
	TotalBorrowAsset *big.Int
	Decimals         uint8
	Symbol           string
}

type BZXProtocol struct {
	scraper    *DefiScraper
	protocol   dia.DefiProtocol
	connection *ethclient.Client
	assets     map[string]string // assetname and address
}

func NewBZX(scraper *DefiScraper, protocol dia.DefiProtocol) *BZXProtocol {
	assets := make(map[string]string)
	assets["DAI"]  = "0x6b093998d36f2c7f0cc359441fbb24cc629d5ff0"
	assets["ETH"]  = "0xb983e01458529665007ff7e0cddecdb74b967eb6"
	assets["USDC"] = "0x32e4c68b3a4a813b710595aeba7f6b7604ab9c15"
	assets["WBTC"] = "0x2ffa85f655752fb2acb210287c60b9ef335f5b6e"
	assets["LEND"] = "0xab45bf58c6482b87da85d6688c4d9640e093be98"
	assets["KNC"]  = "0x687642347a9282be8fd809d8309910a3f984ac5a"
	assets["MKR"]  = "0x9189c499727f88f8ecc7dc4eea22c828e6aac015"
	assets["BZRX"] = "0x18240bd9c07fa6156ce3f3f61921cc82b2619157"
	assets["LINK"] = "0x463538705e7d22aa7f03ebf8ab09b067e1001b54"
	assets["YFI"]  = "0x7f3fe9d492a9a60aebb06d82cba23c6f32cad10b"
	assets["USDT"] = "0x7e9997a38a439b2be7ed9c9c4628391d3e055d48"

	connection, err := ethclient.Dial("http://159.69.120.42:8545/")
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	return &BZXProtocol{scraper: scraper, protocol: protocol, assets: assets, connection: connection}
}

func (proto *BZXProtocol) fetch(asset string) (bzxrate BZXRate, err error) {
	var contract *bzxcontract.LoanTokenLogicV4Caller
	contract, err = bzxcontract.NewLoanTokenLogicV4Caller(common.HexToAddress(proto.assets[asset]), proto.connection)
	if err != nil {
		log.Error(err)
		return
	}
	supplyInterestRate, err := contract.SupplyInterestRate(&bind.CallOpts{})
	if err != nil {
		return
	}
	borrowInterestRate, err := contract.BorrowInterestRate(&bind.CallOpts{})
	if err != nil {
		return
	}
	totalSupply, err := contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return
	}
	// Token supply has to be scaled by decimals
	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return
	}
	totalAssetSupply, err := contract.TotalAssetSupply(&bind.CallOpts{})
	if err != nil {
		return
	}
	totalAssetBorrow, err := contract.TotalAssetBorrow(&bind.CallOpts{})
	if err != nil {
		return
	}
	// fmt.Printf("total supp and borr for asset %s: %v , %v , %v \n", asset, totalAssetSupply, totalAssetBorrow, decimals)

	bzxrate = BZXRate{
		Symbol:           asset,
		BorrowRate:       borrowInterestRate,
		SupplyRate:       supplyInterestRate,
		TotalSupply:      totalSupply,
		TotalSupplyAsset: totalAssetSupply,
		TotalBorrowAsset: totalAssetBorrow,
		Decimals:         decimals,
	}
	return
}

func (proto *BZXProtocol) fetchALL() (bzxrates []BZXRate, err error) {
	for asset, _ := range proto.assets {
		bzxrate, err := proto.fetch(asset)
		if err != nil {
			continue
		}
		bzxrates = append(bzxrates, bzxrate)
	}
	return
}

func (proto *BZXProtocol) UpdateRate() error {
	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
	markets, err := proto.fetchALL()
	if err != nil {
		return err
	}

	for _, market := range markets {
		totalSupplyAPR := new(big.Float)
		totalSupplyAPR.SetString(market.SupplyRate.String())
		totalSupplyAPR = new(big.Float).Quo(totalSupplyAPR, big.NewFloat(math.Pow10(18)))
		totalSupplyAPRPOW18, _ := totalSupplyAPR.Float64()

		totalBorrowAPR := new(big.Float)
		totalBorrowAPR.SetString(market.BorrowRate.String())
		totalBorrowAPR = new(big.Float).Quo(totalSupplyAPR, big.NewFloat(math.Pow10(18)))
		totalBorrowAPRPOW18, _ := totalSupplyAPR.Float64()

		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         market.Symbol,
			Protocol:      proto.protocol.Name,
			LendingRate:   totalSupplyAPRPOW18,
			BorrowingRate: totalBorrowAPRPOW18,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.connection.Close()
		proto.scraper.RateChannel() <- asset

	}
	log.Info("Update complete")
	return nil
}

func (proto *BZXProtocol) getTotalSupply() (float64, error) {
	markets, err := proto.fetchALL()
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for i := 0; i < len(markets); i++ {

		TotalDiff := big.NewInt(0).Sub(markets[i].TotalSupplyAsset, markets[i].TotalBorrowAsset)
		marketLiquidityUSD, err := strconv.ParseFloat(TotalDiff.String(), 64)
		if err != nil {
			return 0, err
		}
		marketLiquidityUSD /= math.Pow10(int(markets[i].Decimals))
		sum += marketLiquidityUSD
		// fmt.Printf("market %s holds %v worth in USD \n", markets[i].Symbol, marketLiquidityUSD)
	}
	return sum, nil
}

func (proto *BZXProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)
	totalSupplyUSD, err := proto.getTotalSupply()
	if err != nil {
		return err
	}
	priceETH, err := utils.GetCoinPrice("ETH")
	if err != nil {
		fmt.Println("error getting price: ", err)
		return err
	}
	totalSupplyETH := totalSupplyUSD / priceETH
	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSD,
		TotalETH:  totalSupplyETH,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")

	return nil

}
