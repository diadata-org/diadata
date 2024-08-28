package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	k8sbridge "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"

	signer "github.com/diadata-org/diadata/pkg/dia/helpers/signer/protoc"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/spf13/cobra"
)

var log = logrus.New()

/*

1. Command to restart all oracle feeder
2. Command to restart specific oracle feeder
3. Command to delete specific feeder
*/

func main() {

	relStore, err := models.NewPostgresDataStore()
	if err != nil {
		log.Errorln("NewPostgresDataStore", err)
	}

	signerservice := os.Getenv("SIGNER")
	k8bridgeClient := initializeKubernetesBridgeClient()

	log.Println("signerservice", signerservice)

	conn, err := grpc.Dial(signerservice, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorln("error while connecting to signer service", err)
		panic("signer connection error")

	}

	signerclient := signer.NewSignerClient(conn)

	// var restartFeeder = &cobra.Command{
	// 	Use:   "restart",
	// 	Short: "restart all feeders",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		oracleconfigs, err := relStore.GetAllFeeders(false)
	// 		if err != nil {
	// 			log.Errorln("error getting feeders", err)
	// 			return
	// 		}
	// 		ph := builderUtils.NewPodHelper(oraclebaseimage, oraclenamespace)

	// 		for _, oracleconfig := range oracleconfigs {
	// 			err = ph.RestartOracleFeeder(cmd.Context(), oracleconfig.FeederID, oracleconfig)
	// 			if err != nil {
	// 				log.Errorln("error RestartOracleFeeder ", err)
	// 			}

	// 		}
	// 	},
	// }

	var expireFeeder = &cobra.Command{
		Use:   "expire",
		Short: "expire stale feeder",
		Run: func(cmd *cobra.Command, args []string) {
			oracleconfigs, err := relStore.GetExpiredFeeders()
			if err != nil {
				log.Errorln("error getting feeders", err)
				return
			}

			sixtyDaysAgo := time.Now().Add(-60 * 24 * time.Hour) // 60 days = 60 * 24 hours

			for _, oracleconfig := range oracleconfigs {
				//double check if created date is greater than 60 days
				if oracleconfig.CreatedDate.Before(sixtyDaysAgo) {

					log.Infof("FeederId %s", oracleconfig.FeederID)

					fc := &k8sbridge.FeederConfig{
						FeederID: oracleconfig.FeederID,
					}

					_, err = k8bridgeClient.DeletePod(cmd.Context(), fc)
					if err != nil {
						log.Errorf("FeederId %s, DeletePod, error %v", oracleconfig.FeederID, err)
						continue
					}
					err = relStore.ChangeOracleState(oracleconfig.FeederID, false)
					if err != nil {
						log.Errorf("FeederId %s, ChangeOracleState, error %v", oracleconfig.FeederID, err)
						continue
					}
					err = relStore.ExpireOracle(oracleconfig.FeederID)
					if err != nil {
						log.Errorf("FeederId %s, ExpireOracle, error %v", oracleconfig.FeederID, err)
						continue
					}
				}

			}

			log.Println("Total oracle expired:", len(oracleconfigs))

		},
	}

	var refund = &cobra.Command{
		Use: "refund",
		Run: func(cmd *cobra.Command, args []string) {

			// Get All RPC

			chainconfig, err := relStore.GetAllChainConfig()
			if err != nil {
				log.Errorln("error  GetAllChainConfig", err)
				return
			}

			rpcClients := make(map[string]*ethclient.Client)
			rpcurls := make(map[string]string)

			for _, cc := range chainconfig {

				rpcurls[cc.ChainID] = cc.RestURL

			}
			oracleconfigs, err := relStore.GetAllFeeders(false, false)
			if err != nil {
				log.Errorln("error getting feeders", err)
				return
			}

			for chainID, rpc := range rpcurls {
				rpcClients[chainID], err = ethclient.Dial(rpc)
				if err != nil {
					log.Errorln("error getting client for chainID", chainID)
					return
				}
			}

			for _, oracleconfig := range oracleconfigs {

				from := common.HexToAddress(oracleconfig.FeederAddress)
				to := common.HexToAddress(oracleconfig.Owner)

				log.Infoln("owner", oracleconfig.Owner)
				log.Infoln("feeder address", oracleconfig.FeederAddress)

				client := rpcClients[oracleconfig.ChainID]

				balance, err := client.BalanceAt(context.Background(), from, nil)
				if err != nil {
					log.Errorln("Error BalanceAt: ", err)
					continue
				}

				nonce, err := client.PendingNonceAt(context.Background(), from)
				if err != nil {
					log.Errorln(" client.NonceAt error  tx", err)
				}

				gasPrice, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Errorln("Error gasFeeCap: ", err)
					continue
				}
				gasTipCap, err := client.SuggestGasTipCap(context.Background())
				if err != nil {
					log.Errorln("Error gasTipCap: ", err)
					continue
				}

				log.Infoln("estimated gas price", gasPrice)
				chainID, _ := strconv.ParseInt(oracleconfig.ChainID, 10, 64)

				gasRequired := new(big.Int).Mul(big.NewInt(21000), gasPrice)

				log.Infoln("estimated gas required", gasRequired)

				log.Infoln("estimated  gasTipCap", gasTipCap)

				amount := new(big.Int).Sub(balance, gasRequired)

				log.Infoln("balance", balance)

				log.Infoln("estimated  amount", amount)

				sig, tx, err := getSignedDATA(signerclient, nonce, to, amount, 21000, gasPrice, gasTipCap, nil, chainID, oracleconfig.FeederID)

				if err != nil {
					log.Errorln("getSignedDATA  tx", err)
					continue
				}

				s := types.LatestSignerForChainID(big.NewInt(chainID))

				tx, err = tx.WithSignature(s, sig)
				if err != nil {
					log.Errorln(" types.SignTx error  tx", err)
					continue

				}
				log.Infoln("tx", tx)

				msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), tx.GasPrice())
				if err != nil {
					log.Errorln("error tx.AsMessage", err)

				}
				log.Errorln("verify from address", msg.From().Hex())

				err = client.SendTransaction(context.Background(), tx)

				if err != nil {
					log.Errorln("gasFeePrice", gasPrice.String())
					log.Errorln("gasTipCap", gasTipCap.String())
					log.Errorln("error doing tx", err.Error())
					// retry if error is replacement transaction underpriced
				} else {
					log.Infoln("Transaction done", tx.Hash().String())
				}
				log.Infoln("---")

				// pay to Owner
				// pay from  FeederAddress

			}

		},
	}

	var updateAddressChecksum = &cobra.Command{
		Use: "updateaddresschecksum",
		Run: func(cmd *cobra.Command, args []string) {
			oracleconfigs, err := relStore.GetAllFeeders(false, false)
			fmt.Println(len(oracleconfigs))
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

	rootCmd.AddCommand(expireFeeder, updateAddressChecksum, refund)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func getSignedDATA(signerclient signer.SignerClient, nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice, gasTipCap *big.Int, data []byte, chainID int64, feederID string) (sig []byte, tx *types.Transaction, err error) {

	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		To:       &to,
		Value:    amount,
		Data:     data,
	})

	s := types.LatestSignerForChainID(big.NewInt(chainID))

	h := s.Hash(tx)

	sigr, err := signerclient.Sign(context.Background(), &signer.DataToSign{Data: h[:], By: feederID})

	sig = sigr.GetSigned()

	// sig, err = crypto.Sign(h[:], pk)

	return

}

func initializeKubernetesBridgeClient() k8sbridge.K8SHelperClient {
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")
	conn, err := grpc.Dial(k8bridgeurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Error while connecting to the signer service: %v", err)
		panic("signer connection error")
	}
	return k8sbridge.NewK8SHelperClient(conn)
}
