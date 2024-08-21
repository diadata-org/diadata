package hydrationhelper

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

const (
	AssetAddressURI = "AssetRegistry:Assets"
	Blockchain      = "Polkadot"
	AssetPallet     = "AssetRegistry"
	AssetStorage    = "Assets"
	DexUrl          = "wss://rpc.hydradx.cloud"
)

type HydrationClient struct {
	client            *gsrpc.SubstrateAPI
	logger            *logrus.Entry
	sleepBetweenCalls time.Duration
	metadata          types.Metadata
	debug             bool
	data              *AssetData
}

func NewHydrationClient(logger *logrus.Entry, sleepBetweenCalls time.Duration, isDebug bool) *HydrationClient {

	client, err := gsrpc.NewSubstrateAPI(DexUrl)
	if err != nil {
		logger.Fatalf("Failed to connect to Hydration RPC: %v", err)
	}

	metadata, err := client.RPC.State.GetMetadataLatest()
	if err != nil {
		logger.WithError(err).Error("Failed to get metadata")
	}

	return &HydrationClient{
		client:            client,
		logger:            logger,
		sleepBetweenCalls: sleepBetweenCalls,
		debug:             isDebug,
		metadata:          *metadata,
		data:              &AssetData{},
	}
}

func (c *HydrationClient) waiting() {
	time.Sleep(c.sleepBetweenCalls)
}

func (c *HydrationClient) GetCurrentHeight() (uint32, error) {
	header, err := c.client.RPC.Chain.GetHeaderLatest()
	if err != nil {
		c.logger.WithError(err).Error("Failed to get the latest block header")
		return 0, err
	}
	return uint32(header.Number), nil
}

func (c *HydrationClient) GetStorageKeys(prefix string) {
	fmt.Printf("Fetching all storage keys for prefix: %sn", prefix)

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

func (c *HydrationClient) GetAssetById(assetId uint32) (*dia.Asset, error) {
	pallet := "AssetRegistry"
	storage := "Assets"

	buf := bytes.NewBuffer(nil)
	enc := scale.NewEncoder(buf)
	if err := enc.Encode(types.NewU32(assetId)); err != nil {
		return nil, err
	}
	args := buf.Bytes()

	keyPrefix, err := types.CreateStorageKey(&c.metadata, pallet, storage, args)
	if err != nil {
		return nil, err
	}

	var storageResp PalletAssetRegistryAssetDetails
	ok, err := c.client.RPC.State.GetStorageLatest(keyPrefix, &storageResp)
	if err != nil {
		c.logger.WithError(err).Error("Failed to get storage")
		return nil, err
	}

	if !ok {
		c.logger.Infof("No data found for AssetId: %d", assetId)
		return nil, nil
	}

	return c.MapFromPalletAsset(assetId, &storageResp)
}

func (c *HydrationClient) ScrapAssets() ([]*dia.Asset, error) {
	assetIds := c.GetAssetIds()
	assets := []*dia.Asset{}

	for _, assetId := range assetIds {
		asset, err := c.GetAssetById(assetId)
		if err != nil {
			c.logger.WithError(err).Errorf("Failed to get asset by id: %d", assetId)
			continue
		}
		if asset != nil {
			assets = append(assets, asset)
			c.ToString(*asset)
		}
	}

	c.logger.Infof("Scraped %d assets", len(assets))

	return assets, nil
}

func (c *HydrationClient) MapFromPalletAsset(assetId uint32, palletAsset *PalletAssetRegistryAssetDetails) (*dia.Asset, error) {
	asset := &dia.Asset{
		Address:    fmt.Sprintf("%s:%s:%d", AssetPallet, AssetStorage, assetId),
		Blockchain: Blockchain,
	}

	ok, nameBytes := palletAsset.Name.Unwrap()
	if ok {
		if utf8.Valid(nameBytes) {
			asset.Name = string(nameBytes)
		} else {
			fmt.Println("Invalid UTF-8 sequence in Name, replacing invalid bytes")
			asset.Name = c.sanitizeToUTF8(nameBytes)
		}
	}
	asset.Name = strings.ReplaceAll(asset.Name, "\u0000", "")

	ok, symbolBytes := palletAsset.Symbol.Unwrap()
	if ok {
		if utf8.Valid(symbolBytes) {
			asset.Symbol = string(symbolBytes)
		} else {
			fmt.Println("Invalid UTF-8 sequence in Symbol, replacing invalid bytes")
			asset.Symbol = c.sanitizeToUTF8(symbolBytes)
		}
	}
	asset.Symbol = strings.ReplaceAll(asset.Symbol, "\u0000", "")

	ok, decimals := palletAsset.Decimals.Unwrap()
	if ok {
		asset.Decimals = uint8(decimals)
	} else {
		asset.Decimals = 0
	}
	return asset, nil
}

func (c *HydrationClient) GetAssetIds() []uint32 {
	assetIds := []uint32{}

	dataAssetIds := c.data.GetAssetIds()
	for _, dataId := range dataAssetIds {
		assetIds = append(assetIds, dataId.Value)
	}

	return assetIds
}

func (c *HydrationClient) ToString(a dia.Asset) {
	fmt.Printf("{\n")
	fmt.Printf("\t\"name\": \"%s\",\n", a.Name)
	fmt.Printf("\t\"blockchain\": \"%s\",\n", a.Blockchain)
	fmt.Printf("\t\"address\": \"%v\",\n", a.Address)
	fmt.Printf("\t\"symbol\": \"%s\",\n", a.Symbol)
	fmt.Printf("\t\"decimals\": \"%v\"\n", a.Decimals)
	fmt.Printf("}\n")
}

func (c *HydrationClient) GetBlockEvents(blockHash types.Hash) ([]types.EventRecords, error) {
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

func (c *HydrationClient) sanitizeToUTF8(data []byte) string {
	var buf bytes.Buffer
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			data = data[size:]
			continue
		}
		if r == 0 {
			data = data[size:]
			continue
		}
		buf.WriteRune(r)
		data = data[size:]
	}
	return buf.String()
}
