package substratehelper

import (
	"errors"
	"fmt"

	gsrpc "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/retriever"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/state"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"

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

func (s *SubstrateEventHelper) ListenForSpecificBlock(blockNumber uint64, callback func([]*parser.Event, uint64)) error {
	blockHash, err := s.API.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		message := fmt.Sprintf("Failed to fetch block hash: %v", err)
		s.logger.Errorf(message, err)
		return errors.New(message)
	}

	events, err := s.DecodeEvents(blockHash)
	if err != nil {
		message := fmt.Sprintf("Failed to decode events: %v", err)
		s.logger.Errorf(message, err)
		return err
	}

	callback(events, blockNumber)

	return nil
}

func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) ([]*parser.Event, error) {
	r, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(s.API.RPC.State), s.API.RPC.State)

	if err != nil {
		return nil, fmt.Errorf("couldn't create event retriever: %s", err)
	}

	events, err := r.GetEvents(blockHash)

	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve events for block hash %d: %s", blockHash, err)
	}

	s.logger.Infof("Found %d events\n", len(events))

	return events, nil
}

// ListenForNewBlocks listens for new blocks and continuously decodes events.
func (s *SubstrateEventHelper) ListenForNewBlocks(callback func([]*parser.Event, uint64)) error {
	sub, err := s.API.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %v", err)
	}
	s.logger.Info("Listening for new blocks...")

	for {
		head := <-sub.Chan()
		blockNumber := uint64(head.Number)
		blockHash, err := s.API.RPC.Chain.GetBlockHash(blockNumber)
		if err != nil {
			s.logger.Errorf("Failed to fetch block hash: %v\n", err)
			continue
		}
		s.logger.Infof("\nNew block detected! Block number: %v, Block hash: %v\n", head.Number, blockHash.Hex())

		events, err := s.DecodeEvents(blockHash)
		if err != nil {
			s.logger.Errorf("Failed to decode events: %v\n", err)
			continue
		}

		callback(events, blockNumber)
	}
}
