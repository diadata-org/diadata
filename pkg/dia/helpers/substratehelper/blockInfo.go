package subhelper

import (
	"bytes"
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/fatih/structs"
	"github.com/vedhavyas/go-subkey/scale"
)

// GetBlockData returns relevant block data from block with @blockNumber.
func GetBlockData(substrateBlockchainName string, blockNumber int64, relDB *models.RelDB, client *gsrpc.SubstrateAPI, key types.StorageKey, meta *types.Metadata) (blockdata dia.BlockData, err error) {
	blockdata, err = relDB.GetBlockData(substrateBlockchainName, blockNumber)
	if err != nil {
		if err.Error() == "no rows in result set" {
			blockdata, err = GetBlockDataOnChain(substrateBlockchainName, uint64(blockNumber), client, key, meta)
			if err != nil {
				return
			}
		}
	}

	return blockdata, err
}

func GetBlockDataOnChain(substrateBlockchainName string, blockNumber uint64, client *gsrpc.SubstrateAPI, key types.StorageKey, meta *types.Metadata) (dia.BlockData, error) {

	var subblockdata dia.SubstrateBlockData
	var blockdata dia.BlockData
	blockHash, err := client.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		return dia.BlockData{}, err
	}
	block, err := client.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return dia.BlockData{}, err
	}

	set, err := client.RPC.State.QueryStorageAt([]types.StorageKey{key}, blockHash)
	if err != nil {
		panic(err)
	}

	subblockdata.Number = uint64(block.Block.Header.Number)
	subblockdata.ParentHash = block.Block.Header.ParentHash
	subblockdata.Root = block.Block.Header.StateRoot
	subblockdata.Events = make([]dia.SubstrateChangeSet, len(set))
	eventNamesMap := make(map[string]bool)
	for k, changes := range set {
		subblockdata.Events[k] = dia.SubstrateChangeSet{}
		for _, change := range changes.Changes {
			if !types.Eq(change.StorageKey, key) || !change.HasStorageData {
				// skip, we are only interested in events with content
				continue
			}
			// Decode the event records
			decoder := scale.NewDecoder(bytes.NewReader(types.EventRecordsRaw(change.StorageData)))
			// determigo get github.com/vedhavyas/go-subkey/scale@v1.0.3ne number of events
			n, err := decoder.DecodeUintCompact()
			if err != nil {
				log.Error(err)
				continue
			}

			// Get event names
			for i := uint64(0); i < n.Uint64(); i++ {
				log.Debug(fmt.Sprintf("decoding event #%v", i))

				// decode Phase
				phase := types.Phase{}
				err := decoder.Decode(&phase)
				if err != nil {
					log.Errorf("skipping... unable to decode Phase for event #%v: %v", i, err)
					continue
				}

				// decode EventID
				id := types.EventID{}
				err = decoder.Decode(&id)
				if err != nil {
					log.Errorf("unable to decode EventID for event #%v: %v", i, err)
					continue
				}

				log.Debug(fmt.Sprintf("event #%v has EventID %v", i, id))

				// ask metadata for method & event name for event
				moduleName, eventName, err := meta.FindEventNamesForEventID(id)
				// moduleName, eventName, err := "System", "ExtrinsicSuccess", nil
				if err != nil {
					log.Errorf("unable to find event with EventID %v in metadata for event #%v: %s", id, i, err)
					continue
				}

				log.Debug(fmt.Sprintf("event #%v is in module %v with event name %v", i, moduleName, eventName))
				eventNamesMap[fmt.Sprintf("%s(%s)", moduleName, eventName)] = true
			}
			subblockdata.Events[k].Data = change.StorageData
		}

	}
	eventNames := make([]string, 0)
	for k := range eventNamesMap {
		eventNames = append(eventNames, k)
	}
	subblockdata.EventNames = eventNames

	blockdata.BlockchainName = substrateBlockchainName
	blockdata.BlockNumber = int64(subblockdata.Number)
	blockdata.Data = structs.Map(subblockdata)

	return blockdata, nil

}
