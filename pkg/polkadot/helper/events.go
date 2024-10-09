package helper

import (
	"fmt"
	"log"
	"reflect"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/sirupsen/logrus"
)

type EventSellExecuted struct {
	AssetIn   types.U32  `json:"asset_in"`
	AssetOut  types.U32  `json:"asset_out"`
	AmountIn  types.U128 `json:"amount_in"`
	AmountOut types.U128 `json:"amount_out"`
}

type SubstrateEventHelper struct {
	logger *logrus.Entry
	API    *gsrpc.SubstrateAPI
}

// NewSubstrateEventHelper creates a new SubstrateEventHelper and connects to the node.
func NewSubstrateEventHelper(logger *logrus.Entry, nodeURL string) (*SubstrateEventHelper, error) {
	api, err := gsrpc.NewSubstrateAPI(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Substrate node: %v", err)
	}
	return &SubstrateEventHelper{API: api, logger: logger}, nil
}

type Event struct {
	Name    string
	EventID types.EventID
	Phase   *types.Phase
	Topics  []types.Hash
}

// DecodeEvents fetches and decodes events for a specific block hash using CustomEventRecords
func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*[]EventSellExecuted, error) {
	retriever, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(s.API.RPC.State), s.API.RPC.State)
	if err != nil {
		log.Printf("Couldn't create event retriever: %s", err)
		return nil, fmt.Errorf("failed to create event retriever: %v", err)
	}

	events, err := retriever.GetEvents(blockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve events: %v", err)
	}

	log.Printf("Found %d events'%s'", len(events))

	// Example of the events returned structure
	for _, event := range events {
		log.Printf("Event ID: %x \n", event.EventID)
		log.Printf("Event Name: %s \n", event.Name)
		log.Printf("Event Fields Count: %d \n", len(event.Fields))
		for k, v := range event.Fields {
			log.Printf("Field Name: %s \n", k)
			log.Printf("Field Type: %v \n", reflect.TypeOf(v))
			log.Printf("Field Value: %v \n", v)
		}
	}

	return nil, nil
}

func (s *SubstrateEventHelper) ListenForSpecificBlock(blockNumber uint64, callback func(*[]EventSellExecuted)) error {
	blockHash, err := s.API.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		message := fmt.Sprintf("Failed to fetch block hash: %v", err)
		s.logger.Errorf(message, err)
		return fmt.Errorf(message)
	}

	events, err := s.DecodeEvents(blockHash)
	if err != nil {
		message := fmt.Sprintf("Failed to decode events: %v", err)
		s.logger.Errorf(message, err)
		return err
	}

	callback(events)

	return nil
}

// ListenForNewBlocks listens for new blocks and continuously decodes events.
func (s *SubstrateEventHelper) ListenForNewBlocks(callback func([]EventSellExecuted)) error {
	sub, err := s.API.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %v", err)
	}

	for {
		head := <-sub.Chan()

		blockHash, err := s.API.RPC.Chain.GetBlockHash(uint64(head.Number))
		if err != nil {
			s.logger.Errorf("Failed to fetch block hash: %v\n", err)
			continue
		}
		fmt.Printf("\nNew block detected! Block number: %v, Block hash: %v\n", head.Number, blockHash)

		block, err := s.API.RPC.Chain.GetBlock(blockHash)
		if err != nil {
			s.logger.Errorf("Failed to fetch block details: %v\n", err)
			continue
		}

		events, err := s.DecodeEvents(block.Block.Header.ExtrinsicsRoot)
		if err != nil {
			s.logger.Errorf("Failed to decode events: %v\n", err)
			continue
		}

		callback(*events)
	}
}
