package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/diaoracle"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*

1) Get Wallet
2) establish connection to oracle
3) Deploy Oracle

*/

var log = logrus.New()

func main() {

	flag.String("config", "config", "Name for config file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	setConfigDefault()

	secretsFile := viper.GetString("secretsFile")
	blockchainNode := viper.GetString("blockchainNode")
	deployedContract := viper.GetString("deployedContract")
	diaServer := viper.GetString("diaServer")
	frequencySeconds := viper.GetInt64("frequencySeconds")
	symbols := viper.GetStringSlice("symbols")
	deviationPermille := viper.GetInt("deviationPermille")
	updateOnDeviation := viper.GetBool("updateOnDeviation")

	chainId := viper.GetInt64("chainId")
	oracleType := viper.GetString("oracleType")
	defirate := viper.GetStringMapString("defiRate")
	defiState := viper.GetStringSlice("defiState")
	dexChartPoints := viper.GetStringMapString("dexChartPoints")
	farmingPools := viper.GetStringMapString("farmingPools")
	indexSymbols := viper.GetStringSlice("indexSymbols")

	wallet, err := diaoracle.NewWallet(secretsFile)
	if err != nil {
		log.Fatalf("Failed to Get Wallet: %v", err)
	}
	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the EVM client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(wallet.Key), wallet.PassPhrase, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	log.Infoln("Running Oracle Type", oracleType)
	log.Infoln("Running Oracle Type", oracleType)

	log.Infoln("Running Oracle deviationPermille", deviationPermille)

	var oracle diaoracle.Oracalize

	switch oracleType {
	case "legacy":
		oracle = diaoracle.NewLegacyContract(conn, auth, deployedContract, symbols, defirate, defiState, dexChartPoints, farmingPools, diaServer)
	case "signed":
		oracle = diaoracle.NewSignedOracleContract(conn, auth, deployedContract, deviationPermille, updateOnDeviation, symbols)
	case "unsigned":
		oracle = diaoracle.NewUnSignedOracleContract(conn, auth, deployedContract, deviationPermille, updateOnDeviation, symbols, indexSymbols)
	}

	log.Infoln("time to update Oracle", frequencySeconds)

	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for range ticker.C {
			log.Infoln("time to update Oracle")
			oracle.Update()

		}
	}()
	select {}

}

func setConfigDefault() {
	config := viper.GetString("config")
	if config != "" {
		viper.SetConfigName(config)

	} else {
		viper.SetConfigName("config")

	}
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error config file: %w ", err))
	}

	viper.SetDefault("deployedContract", "")
	viper.SetDefault("deviationPermille", 30)
	viper.SetDefault("updateOnDeviation", false)

	viper.SetDefault("secretsFile", "")
	viper.SetDefault("blockchainNode", "")
	viper.SetDefault("sleepSeconds", "")
	viper.SetDefault("frequencySeconds", 1)
	viper.SetDefault("chainId", "")

	viper.SetEnvPrefix("oracle")
	viper.BindEnv("deployedContract")

}
