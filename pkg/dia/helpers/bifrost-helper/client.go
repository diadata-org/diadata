package bifrosthelper

import (
	// "fmt"
	"fmt"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	signature "github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/sirupsen/logrus"
)

const (
	PairContractAddress = "vyrkJHG49TXss6pGAz2dVxq5o7mBXNNXAV18nAeqVT1R"
)

const (
	DefaultRefreshDelay              = 400 // millisec
	DefaultSleepBetweenContractCalls = 300 // millisec
	DefaultEventsLimit               = 100
	DefaultSwapContractsLimit        = 100
)

type BifrostClient struct {
	client            *gsrpc.SubstrateAPI
	logger            *logrus.Entry
	sleepBetweenCalls time.Duration
	metadata          types.Metadata
	debug             bool
}

func NewBifrostClient(logger *logrus.Entry, sleepBetweenCalls time.Duration, isDebug bool) *BifrostClient {
	client, err := gsrpc.NewSubstrateAPI("wss://bifrost-parachain.api.onfinality.io/public-ws")
	if err != nil {
		logger.Fatalf("Failed to connect to Bifrost RPC: %v", err)
	}

	metadata, err := client.RPC.State.GetMetadataLatest()
	if err != nil {
		logger.WithError(err).Error("Failed to get metadata")
	}

	return &BifrostClient{
		client:            client,
		logger:            logger,
		sleepBetweenCalls: sleepBetweenCalls,
		debug:             isDebug,
		metadata:          *metadata,
	}
}

func (c *BifrostClient) waiting() {
	time.Sleep(c.sleepBetweenCalls)
}

func (c *BifrostClient) GetCurrentHeight() (uint32, error) {
	header, err := c.client.RPC.Chain.GetHeaderLatest()
	if err != nil {
		c.logger.WithError(err).Error("Failed to get the latest block header")
		return 0, err
	}
	return uint32(header.Number), nil
}

func (c *BifrostClient) GetStorageKeys(prefix string) {
	fmt.Printf("Fetching all storage keys for prefix: %s\n", prefix)

	// Iterate over all pallets in the metadata
	for _, pallet := range c.metadata.AsMetadataV14.Pallets {
		if string(pallet.Name) == prefix {
			fmt.Printf("Found pallet: %s\n", pallet.Name)

			// Iterate over all storage items in the pallet
			for _, item := range pallet.Storage.Items {
				fmt.Printf("  %s\n", item.Name)
				fmt.Printf("    Type: %s\n", item.Type)
			}
		}
	}
}

func (c *BifrostClient) ScrapAssets() {
	// Retrieve the latest metadata
	metadata, err := c.client.RPC.State.GetMetadataLatest()
	if err != nil {
		c.logger.Fatalf("Failed to get latest metadata: %v", err)
	}

	// Specify the pallet and storage item
	storagePrefix := "VtokenMinting"
	storageName := "TokenPool"

	currencyId, err := types.EncodeToBytes("00801")
	if err != nil {
		c.logger.Fatalf("Failed to encode arguments: %v", err)
	}

	// Create a storage key prefix (without arguments for prefix search)
	
	key, err := types.CreateStorageKey(metadata, storagePrefix, storageName, currencyId)
	if err != nil {
		c.logger.Fatalf("Failed to create storage key prefix: %v", err)
	}

	var balance types.U128
	ok, err := c.client.RPC.State.GetStorageLatest(key, &balance)
	if err != nil {
		c.logger.Printf("Failed to get storage data for key %s: %v", key.Hex(), err)
	}

	if ok {
		fmt.Printf("TokenPool for CurrencyId %d: %s\n", currencyId, balance.String())
	} else {
		fmt.Printf("No data found for CurrencyId %d.\n", currencyId)
	}
	c.logger.Info("ScrapAssets completed")
}

func (c *BifrostClient) GetAccountBalance() {
	alice := signature.TestKeyringPairAlice.PublicKey
	key, err := types.CreateStorageKey(&c.metadata, "Balances", "Account", alice)
	if err != nil {
		panic(err)
	}

	var accountInfo types.AccountInfo
	ok, err := c.client.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		println("Error: ", err)
		panic(err)
	}

	previous := accountInfo.Data.Free
	fmt.Printf("%#x has a balance of %v\n", alice, previous)
	fmt.Printf("You may leave this example running and transfer any value to %#x\n", alice)
}

func (c *BifrostClient) ShowPalletes() {
	// interestedPallets := []string{"Tokens", "ZenlinkProtocol", "Balances", "Assets"}

	println("Start________________________________________________________________________________")
	for _, pallet := range c.metadata.AsMetadataV14.Pallets {
		// if contains(interestedPallets, pallet.Name) {
		fmt.Printf("Pallet: %s\n", pallet.Name)
		fmt.Printf("  Prefix: %s\n", pallet.Storage.Prefix)
		for _, item := range pallet.Storage.Items {
			fmt.Printf("    Storage: %s\n", item.Name)
		}
		// }
	}
	println("____________________________________________________________________________________")
}

func contains(arr []string, name types.Text) bool {
	for _, a := range arr {
		if a == string(name) {
			return true
		}
	}
	return false
}

func (c *BifrostClient) GetBlockEvents(blockHash types.Hash) ([]types.EventRecords, error) {
	key, err := types.CreateStorageKey(&c.metadata, "System", "Events", nil, nil)
	if err != nil {
		c.logger.WithError(err).Error("Failed to create storage key")
		return nil, err
	}

	raw, err := c.client.RPC.State.GetStorageRaw(key, blockHash)
	if err != nil {
		c.logger.WithError(err).Error("Failed to get raw storage")
		return nil, err
	}

	var events []types.EventRecords
	err = types.EventRecordsRaw(*raw).DecodeEventRecords(&c.metadata, &events)
	if err != nil {
		c.logger.WithError(err).Error("Failed to decode event records")
		return nil, err
	}

	return events, nil // Modify this to filter the relevant events

}

func (c *BifrostClient) GetSwapPairsContractAddresses(swapContractsLimit int) (interface{}, error) {
	// Crie uma chave de armazenamento para o prefixo ZenlinkProtocol/
	// Liste todos os pallets e itens de armazenamento disponíveis
	for _, mod := range c.metadata.AsMetadataV14.Pallets {
		if mod.Name != "ZenlinkProtocol" && mod.Name != "Tokens" {
			continue
		}

		fmt.Printf("Pallet: %s\n", mod.Name)
		if len(mod.Storage.Items) > 0 {
			items := mod.Storage.Items
			fmt.Printf("  Prefix: %s\n", mod.Storage.Prefix)
			for _, item := range items {
				fmt.Printf("    Storage: %s\n", item.Name)
				// fmt.Printf("      Type: %s\n", item.Type)
				// fmt.Printf("      Modifier: %s\n", item.Documentation)
			}
		}
	}

	prefix := types.NewStorageKey([]byte("0x3a636f6465"))
	c.logger.Infof("Fetching all storage keys for prefix: %s", prefix.Hex())

	block, _ := c.client.RPC.Chain.GetBlockHashLatest()

	// Recupere todas as chaves (TokenPairIDs) para LiquidityPairs
	keys, err := c.client.RPC.State.GetKeys(prefix, block)
	if err != nil {
		c.logger.WithError(err).Error("Failed to get keys for LiquidityPairs")
		return nil, err
	}
	c.logger.Infof("Found %d keys for LiquidityPairs", len(keys))

	// Inicialize um slice para armazenar todos os pares de swap
	var swapPairs []interface{}

	// Itere sobre cada chave e recupere os dados associados
	for _, key := range keys {
		c.logger.Infof("Processing key: %s", key.Hex())

		// Variável para armazenar os dados do par
		var pairData interface{} // Substitua pelo tipo correto, se conhecido

		// Recupere os dados do armazenamento para a chave atual
		ok, err := c.client.RPC.State.GetStorageLatest(key, &pairData)
		if err != nil {
			c.logger.WithError(err).Errorf("Failed to get storage for key %s", key.Hex())
			continue
		}

		if ok {
			c.logger.Infof("Successfully retrieved data for key %s", key.Hex())
			swapPairs = append(swapPairs, pairData)
		} else {
			c.logger.Warnf("No data found for key %s", key.Hex())
		}
	}

	// Retorne todos os pares de swap encontrados
	return swapPairs, nil
}
