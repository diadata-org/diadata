package main

import (
	"fmt"
	"os"

	builderUtils "github.com/diadata-org/diadata/cmd/oraclebuildertools/utils"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var log = logrus.New()

/*

1. Command to restart all oracle feeder
2. Command to restart specific oracle feeder
3. Command to delete specific feeder
*/

func main() {

	relStore, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}
	oraclebaseimage := utils.Getenv("ORACLE_BASE_IMAGE", "us.icr.io/dia-registry/oracles/oracle-baseimage:latest")
	oraclenamespace := utils.Getenv("ORACLE_NAMESPACE", "dia-oracle-feeder")

	var restartFeeder = &cobra.Command{
		Use:   "restart",
		Short: "restart all feeders",
		Run: func(cmd *cobra.Command, args []string) {
			oracleconfigs, err := relStore.GetAllFeeders()
			if err != nil {
				log.Errorln("error getting feeders", err)
				return
			}
			ph := builderUtils.NewPodHelper(oraclebaseimage, oraclenamespace)

			for _, oracleconfig := range oracleconfigs {
				err = ph.RestartOracleFeeder(cmd.Context(), oracleconfig.FeederID, oracleconfig)
				if err != nil {
					log.Errorln("error RestartOracleFeeder ", err)
				}

			}
		},
	}

	var deleteFeeder = &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var updateAddressChecksum = &cobra.Command{
		Use: "updateaddresschecksum",
		Run: func(cmd *cobra.Command, args []string) {
			oracleconfigs, err := relStore.GetAllFeeders()
			if err != nil {
				log.Errorln("error getting feeders", err)
				return
			}
			for _, oracleconfig := range oracleconfigs {
				log.Infoln("updating oracle config for ", oracleconfig.Address)
				err = relStore.UpdateFeederAddressCheckSum(oracleconfig.Address)
				if err != nil {
					log.Errorln("error UpdateFeederAddressCheckSum ", err, oracleconfig.Address)
				}
			}
		},
	}

	var rootCmd = &cobra.Command{
		Use: "oraclebuildertool",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	rootCmd.AddCommand(restartFeeder, deleteFeeder, updateAddressChecksum)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
