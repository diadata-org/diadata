package defiscrapers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/makerdao/jugcontract"
	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/makerdao/potcontract"
	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/makerdao/vatcontract"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Collateral struct {
	name      string
	vaultAddr string
	tokenAddr string
}

type MakerdaoProtocol struct {
	scraper     *DefiScraper
	protocol    dia.DefiProtocol
	connection  *ethclient.Client
	collaterals []Collateral
}

var (
	collaterals = []Collateral{
		{
			name:      "ETH-A",
			vaultAddr: "0x2F0b23f53734252Bda2277357e97e1517d6B042A",
			tokenAddr: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		},
		{
			name:      "ETH-B",
			vaultAddr: "0x08638eF1A205bE6762A8b935F5da9b700Cf7322c",
			tokenAddr: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		},
		{
			name:      "BAT-A",
			vaultAddr: "0x3D0B1912B66114d4096F48A8CEe3A56C231772cA",
			tokenAddr: "0x0D8775F648430679A709E98d2b0Cb6250d2887EF",
		},
		{
			name:      "USDC-A",
			vaultAddr: "0xA191e578a6736167326d05c119CE0c90849E84B7",
			tokenAddr: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		},
		{
			name:      "USDC-B",
			vaultAddr: "0x2600004fd1585f7270756DDc88aD9cfA10dD0428",
			tokenAddr: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		},
		{
			name:      "WBTC-A",
			vaultAddr: "0xBF72Da2Bd84c5170618Fbe5914B0ECA9638d5eb5",
			tokenAddr: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
		},
		{
			name:      "TUSD-A",
			vaultAddr: "0x4454aF7C8bb9463203b66C816220D41ED7837f44",
			tokenAddr: "0x0000000000085d4780B73119b644AE5ecd22b376",
		},
		{
			name:      "KNC-A",
			vaultAddr: "0x475F1a89C1ED844A08E8f6C50A00228b5E59E4A9",
			tokenAddr: "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		},
		{
			name:      "ZRX-A",
			vaultAddr: "0xc7e8Cd72BDEe38865b4F5615956eF47ce1a7e5D0",
			tokenAddr: "0xE41d2489571d322189246DaFA5ebDe1F4699F498",
		},
		{
			name:      "MANA-A",
			vaultAddr: "0xA6EA3b9C04b8a38Ff5e224E7c3D6937ca44C0ef9",
			tokenAddr: "0x0F5D2fB29fb7d3CFeE444a200298f468908cC942",
		},
		{
			name:      "USDT-A",
			vaultAddr: "0x0Ac6A1D74E84C2dF9063bDDc31699FF2a2BB22A2",
			tokenAddr: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		},
		{
			name:      "PAXUSD-A",
			vaultAddr: "0x7e62B7E279DFC78DEB656E34D6a435cC08a44666",
			tokenAddr: "0x8E870D67F660D95d5be530380D0eC0bd388289E1",
		},
		{
			name:      "COMP-A",
			vaultAddr: "0xBEa7cDfB4b49EC154Ae1c0D731E4DC773A3265aA",
			tokenAddr: "0xc00e94Cb662C3520282E6f5717214004A7f26888",
		},
		{
			name:      "LRC-A",
			vaultAddr: "0x6C186404A7A238D3d6027C0299D1822c1cf5d8f1",
			tokenAddr: "0xBBbbCA6A901c926F240b89EacB641d8Aec7AEafD",
		},
		{
			name:      "LINK-A",
			vaultAddr: "0xdFccAf8fDbD2F4805C174f856a317765B49E4a50",
			tokenAddr: "0x514910771AF9Ca656af840dff83E8264EcF986CA",
		},
		{
			name:      "YFI-A",
			vaultAddr: "0x3ff33d9162aD47660083D7DC4bC02Fb231c81677",
			tokenAddr: "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e",
		},
		{
			name:      "BAL-A",
			vaultAddr: "0x4a03Aa7fb3973d8f0221B466EefB53D0aC195f55",
			tokenAddr: "0xba100000625a3754423978a60c9317c58a424e3D",
		},
		{
			name:      "GUSD-A",
			vaultAddr: "0xe29A14bcDeA40d83675aa43B72dF07f649738C8b",
			tokenAddr: "0x056Fd409E1d7A124BD7017459dFEa2F387b6d5Cd",
		},
		{
			name:      "UNI-A",
			vaultAddr: "0x3BC3A58b4FC1CbE7e98bB4aB7c99535e8bA9b8F1",
			tokenAddr: "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		},
		{
			name:      "RENBTC-A",
			vaultAddr: "0xFD5608515A47C37afbA68960c1916b79af9491D0",
			tokenAddr: "0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D",
		},
		{
			name:      "AAVE-A",
			vaultAddr: "0x24e459F61cEAa7b1cE70Dbaea938940A7c5aD46e",
			tokenAddr: "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
		},
		{
			name:      "UNIV2DAIETH-A",
			vaultAddr: "0x2502F65D77cA13f183850b5f9272270454094A08",
			tokenAddr: "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
		},
	}
)

func NewMakerdao(scraper *DefiScraper, protocol dia.DefiProtocol) *MakerdaoProtocol {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	return &MakerdaoProtocol{scraper: scraper, protocol: protocol, collaterals: collaterals, connection: connection}
}

func (proto *MakerdaoProtocol) UpdateRate() error {
	// DAI lending rate
	pot, err := potcontract.NewPotcontract(common.HexToAddress("0x197E90f9FAD81970bA7976f33CbD77088E5D7cf7"), proto.connection)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)

	dsr, err := pot.Dsr(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	dsrFloat := convertBigintToFloat64(dsr)
	apy := getYearlyFromInstantaneous(dsrFloat)

	asset := &dia.DefiRate{
		Timestamp:     time.Now(),
		Asset:         fmt.Sprintf("DAI"),
		Protocol:      proto.protocol.Name,
		LendingRate:   apy,
		BorrowingRate: 0,
	}
	log.Printf("writing DEFI rate for %#v in %v\n", asset, proto.scraper.RateChannel())
	proto.scraper.RateChannel() <- asset

	// Collaterals borrowing rates
	jug, err := jugcontract.NewJugcontract(common.HexToAddress("0x19c0976f590D67707E62397C87829d896Dc0f1F1"), proto.connection)
	if err != nil {
		log.Fatal(err)
	}

	base, err := jug.Base(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	for _, col := range proto.collaterals {
		collateral := convertStringToBytes32(col.name)

		ilks, err := jug.Ilks(&bind.CallOpts{}, collateral)
		if err != nil {
			log.Fatal(err)
		}

		localBase := *base

		localBase.Add(&localBase, ilks.Duty)

		rate := convertBigintToFloat64(&localBase)
		apy := getYearlyFromInstantaneous(rate)

		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         fmt.Sprintf("%s", col.name),
			Protocol:      proto.protocol.Name,
			LendingRate:   0,
			BorrowingRate: apy,
		}

		log.Printf("writing DEFI rate for %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset
	}

	proto.connection.Close()
	log.Info("Update complete")

	return nil
}

func (proto *MakerdaoProtocol) UpdateState() error {
	log.Printf("Updating DEFI state for %+v\n ", proto.protocol)

	vat, err := vatcontract.NewVatcontract(common.HexToAddress("0x35D1b3F3D7966A1DFe207aa4514C12a259A0492B"), proto.connection)
	if err != nil {
		log.Fatal(err)
	}

	totalDaiDebt := 0.

	for _, col := range collaterals {
		ilks, err := vat.Ilks(&bind.CallOpts{}, convertStringToBytes32(col.name))
		if err != nil {
			log.Fatal(err)
		}

		art := convertBigintToFloat64(ilks.Art) / math.Pow(10, 18)
		rate := convertBigintToFloat64(ilks.Rate) / math.Pow(10, 27)

		totalDaiDebt += art * rate
	}

	ETHPrice, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return err
	}

	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalDaiDebt,
		TotalETH:  totalDaiDebt / ETHPrice,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}

	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())
	log.Info("Update State complete")

	return nil
}

func convertStringToBytes32(value string) [32]byte {
	var bytes32 [32]byte
	copy(bytes32[:], string(value))

	return bytes32
}

func convertBigintToFloat64(bigint *big.Int) float64 {
	value := bigint.String()

	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}

	return float
}

func getYearlyFromInstantaneous(instantaneous float64) float64 {
	const nbOfSecInYear float64 = 365 * 24 * 3600

	rate := instantaneous / math.Pow(10, 27)
	apy := math.Pow(rate, nbOfSecInYear)
	apy = (apy - 1) * 100
	apy = math.Round(apy*100) / 100

	return apy
}
