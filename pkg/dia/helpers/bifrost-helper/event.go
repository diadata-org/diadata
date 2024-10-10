package bifrosthelper

import (
	"bytes"
	"fmt"
	"io"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	// "github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	// "github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/bifrost-helper/decoder"

	"github.com/sirupsen/logrus"
)

type SubstrateEventHelper struct {
	logger *logrus.Entry
	API    *gsrpc.SubstrateAPI
}

func NewSubstrateEventHelper(nodeURL string, logger *logrus.Entry) (*SubstrateEventHelper, error) {
	api, err := gsrpc.NewSubstrateAPI(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Substrate node: %v", err)
	}
	return &SubstrateEventHelper{API: api, logger: logger}, nil
}

type StableSwapEvent struct {
	TxID            string   `json:"txId"`
	Timestamp       int64    `json:"timestamp"`
	BlockHash       string   `json:"blockHash"`
	Swapper         string   `json:"swapper"`
	PoolID          string   `json:"poolId"`
	A               string   `json:"a"`
	InputAsset      string   `json:"inputAsset"`
	OutputAsset     string   `json:"outputAsset"`
	InputAmount     string   `json:"inputAmount"`
	MinOutputAmount string   `json:"minOutputAmount"`
	Balances        []string `json:"balances"`
	TotalSupply     string   `json:"totalSupply"`
	OutputAmount    string   `json:"outputAmount"`
}

// type CustomEventRecords struct {
// 	StableAsset_TokenSwaped []EventSellExecuted
// }

type EventSellExecuted struct {
	InputAsset   types.U32  `json:"input_asset"`
	OutputAsset  types.U32  `json:"output_asset"`
	InputAmount  types.U128 `json:"input_amount"`
	OutputAmount types.U128 `json:"output_amount"`
}

// // Define the custom CurrencyId struct (assuming bifrost_primitives:currency:CurrencyId)
// type CurrencyId struct {
// 	VToken2 uint32 // Placeholder type for the currency (adjust according to your runtime)
// 	Token2  uint32 // Placeholder type for another asset (adjust accordingly)
// }

// // Define the struct for the TokenSwapped event with all 10 fields
// type EventSellExecuted struct {
// 	Swapper         types.AccountID `json:"swapper"`
// 	PoolId          types.U32       `json:"pool_id"`
// 	A               types.U128      `json:"a"`
// 	InputAsset      CurrencyId      `json:"input_asset"`
// 	OutputAsset     CurrencyId      `json:"output_asset"`
// 	InputAmount     types.U128      `json:"input_amount"`
// 	MinOutputAmount types.U128      `json:"min_output_amount"`
// 	Balances        []types.U128    `json:"balances"`
// 	TotalSupply     types.U128      `json:"total_supply"`
// 	OutputAmount    types.U128      `json:"output_amount"`
// }

func (s *SubstrateEventHelper) ListenForSpecificBlock(blockNumber uint64, callback func(*[]EventSellExecuted)) error {
	blockHash, err := s.API.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		message := fmt.Sprintf("Failed to fetch block hash: %v", err)
		s.logger.Errorf(message, err)
		return fmt.Errorf(message)
	}

	// block, err := s.API.RPC.Chain.GetBlock(blockHash)
	// if err != nil {
	// 	message := fmt.Sprintf("Failed to fetch block details: %v", err)
	// 	s.logger.Errorf(message, err)
	// 	return err
	// }

	// s.logger.Info("ExtrinsicRoot: ", block.Block.Header.ExtrinsicsRoot.Hex())

	//event, err := s.DecodeEvents(block.Block.Header.ExtrinsicsRoot)
	events, err := s.DecodeEvents(blockHash)
	if err != nil {
		message := fmt.Sprintf("Failed to decode events: %v", err)
		s.logger.Errorf(message, err)
		return err
	}

	for _, event := range *events {
		s.logger.Info("Event details: ", event)
		s.logger.Info("Event inputamount: ", event.InputAmount.Int64())
		s.logger.Info("Event outputamount: ", event.OutputAmount.Int64())
	}

	callback(events)

	return nil
}

type StableSwapEvent2 struct {
}

// DecodeEvents fetches and decodes events for a specific block hash using CustomEventRecords
// func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*[]EventSellExecuted, error) {
// 	retriever, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(s.API.RPC.State), s.API.RPC.State)
// 	if err != nil {
// 		log.Printf("Couldn't create event retriever: %s", err)
// 		return nil, fmt.Errorf("failed to create event retriever: %v", err)
// 	}

// 	events, err := retriever.GetEvents(blockHash)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve events: %v", err)
// 	}

// 	log.Printf("Found %d events'%s'", len(events))

// 	// Example of the events returned structure
// 	for _, event := range events {
// 		log.Printf("Event ID: %x \n", event.EventID)
// 		log.Printf("Event Name: %s \n", event.Name)
// 		log.Printf("Event Fields Count: %d \n", len(event.Fields))
// 		for k, v := range event.Fields {
// 			log.Printf("Field Name: %s \n", k)
// 			log.Printf("Field Type: %v \n", reflect.TypeOf(v))
// 			log.Printf("Field Value: %v \n", v)
// 		}
// 	}

// 	return nil, nil
// }

// TODO: MOVE THIS TO COMMON PARACHAIN HELPER
// DecodeEvents fetches and decodes events for a specific block hash using CustomEventRecords
// func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*CustomEventRecords, error) {
func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*[]EventSellExecuted, error) {
	meta, err := s.API.RPC.State.GetMetadataLatest()
	if err != nil {

		return nil, fmt.Errorf("failed to get latest metadata: %v", err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage key for events: %v", err)
	}

	s.logger.Info("BlockHash: ", blockHash.Hex())

	events := []EventSellExecuted{}
	// ok, err := s.API.RPC.State.GetStorage(key, &events, blockHash)
	// if err != nil || !ok {
	// 	s.logger.WithError(err).Error("failed to get events from block")
	// 	return nil, fmt.Errorf("failed to get events from block: %v", err)
	// }

	rawEvent, err := s.API.RPC.State.GetStorageRaw(key, blockHash)
	if err != nil {
		s.logger.WithError(err).Error("failed to get events from block")
		return nil, fmt.Errorf("failed to get events from block: %v", err)
	}

	runtimeVersion, err := s.API.RPC.State.GetRuntimeVersionLatest() // rpc chain_getRuntimeVersion
	if err != nil {
		return nil, fmt.Errorf("failed to get runtime version: %v", err)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)

	encoder := scale.NewEncoder(writer)
	err = meta.Encode(*encoder)

	if err != nil {
		return nil, fmt.Errorf("failed to encode metadata: %v", err)
	}

	spec := int(runtimeVersion.SpecVersion)
	runtimeRaw := decoder.RuntimeRaw{
		Spec: spec,
		Raw:  decoder.BytesToHex(buf.Bytes()),
	}
	metadataInstant := decoder.Latest(&runtimeRaw)
	event, err := decoder.DecodeEvent(rawEvent.Hex(), metadataInstant, spec)
	if err != nil {
		return nil, fmt.Errorf("failed to decode event: %v", err)
	}

	s.logger.Info("Event: ", event)

	if events, ok := event.([]interface{}); ok {
		for i, e := range events {
			s.logger.Info("Event ", i)
			if eventMap, ok := e.(map[string]interface{}); ok {
				s.logger.Info("Event details:")
				for key, value := range eventMap {
					s.logger.Info("key: ", key)
					s.logger.Info("value: ", value)
				}
			} else {
				s.logger.Info("Event is not a map[string]interface{}")
			}
		}
	} else {
		s.logger.Info("Event is not a []interface{}")
	}

	for _, event := range events {
		s.logger.Info("Event details: ")
		s.logger.Info("InputAsset: ", event.InputAsset)
		s.logger.Info("OutputAsset: ", event.OutputAsset)
		s.logger.Info("InputAmount: ", event.InputAmount)
		s.logger.Info("OutputAmount: ", event.OutputAmount)
	}

	return &events, nil
}

// ListenForNewBlocks listens for new blocks and continuously decodes events.
func (s *SubstrateEventHelper) ListenForNewBlocks(callback func([]StableSwapEvent)) error {
	sub, err := s.API.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %v", err)
	}
	fmt.Println("Listening for new blocks...")

	for {
		head := <-sub.Chan()

		blockHash, err := s.API.RPC.Chain.GetBlockHash(uint64(head.Number))
		if err != nil {
			s.logger.Errorf("Failed to fetch block hash: %v\n", err)
			continue
		}
		fmt.Printf("\nNew block detected! Block number: %v, Block hash: %v\n", head.Number, blockHash.Hex())

		block, err := s.API.RPC.Chain.GetBlock(blockHash)
		if err != nil {
			s.logger.Errorf("Failed to fetch block details: %v\n", err)
			continue
		}

		s.logger.Info("Block details: ", block.Block.Header)
		events, err := s.DecodeEvents(block.Block.Header.ExtrinsicsRoot)
		if err != nil {
			s.logger.Errorf("Failed to decode events: %v\n", err)
			continue
		}

		s.logger.Info(events)

		//newEvents := []StableSwapEvent{}
		//s.FilterSellExecuted(events, newEvents)

		//callback(newEvents)
	}
}
