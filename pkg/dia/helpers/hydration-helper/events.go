package hydrationhelper

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
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

// DecodeEvents fetches and decodes events for a specific block hash using CustomEventRecords
func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*[]EventSellExecuted, error) {
	meta, err := s.API.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest metadata: %v", err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage key for events: %v", err)
	}

	events := []EventSellExecuted{}
	ok, err := s.API.RPC.State.GetStorage(key, &events, blockHash)
	if err != nil || !ok {
		s.logger.Info(err)
		return nil, fmt.Errorf("failed to get events from block: %v", err)
	}

	if err != nil || !ok {
		s.logger.WithError(err).Info("Error:")
		return nil, fmt.Errorf("failed to get events from block: %v", err)
	}

	s.logger.Info("Events: ", events)

	return &events, nil
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
