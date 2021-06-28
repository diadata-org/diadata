package utils

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthTxFilterCriteria used for filtering transaction records
type EthTxFilterCriteria struct {
	StartBlockNum      uint64           // inclusive. filter transactions from the specific block
	StartTxIndex       uint             // inclusive. filter transactions from specific index in the given StartBlockNum
	LimitBlocks        int              // filter transactions in the specific number of blocks, zero means up to highest possible block
	BehindHighestBlock int              // stay behind the highest synched block if StartBlockNum+LimitBlocks in excess of the head
	EvAddrs            []common.Address // list of addresses from which transaction events should originate
	Events             []common.Hash    // list of events from which transactions should contain
}

// EthTxFilterResult describes filter results and holds found transactions
type EthTxFilterResult struct {
	Synced       bool             // it means the most recent inspected block matches the highest known block (see BehindHighestBlock)
	NumBlocks    int              // number of blocks found
	NumTXs       int              // number of transactions found
	NumLogs      int              // number of log records found
	TXs          []*EthFilteredTx // list of found transactions
	LastBlockNum uint64           // block number of most recent transaction inspected
}

// EthFilteredTx holds limited info for found transactions
type EthFilteredTx struct {
	BlockNum  uint64
	BlockHash common.Hash
	TXIndex   uint
	TXHash    common.Hash
	Logs      []types.Log // list of matched log records if Events or EvAddrs were used, otherwise all logs of the transaction
}

// EthFilterTXs returns transactions filtered by log records
func EthFilterTXs(ctx context.Context, ethClient *ethclient.Client, filter EthTxFilterCriteria) (*EthTxFilterResult, error) {
	endBlockNum, synced, err := ethFilterTXsCalcEndBlockNum(ctx, ethClient, filter.StartBlockNum, uint64(filter.BehindHighestBlock), uint64(filter.LimitBlocks))
	if err != nil {
		return nil, err
	}

	query := make([][]interface{}, 1)
	query[0] = make([]interface{}, len(filter.Events))

	for i, ev := range filter.Events {
		query[0][i] = ev
	}

	topics, err := abi.MakeTopics(query...)
	if err != nil {
		return nil, err
	}

	logs, err := ethClient.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: filter.EvAddrs,
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(filter.StartBlockNum),
		ToBlock:   new(big.Int).SetUint64(endBlockNum),
	})

	if err != nil {
		return nil, err
	}

	txMap := make(map[common.Hash]*EthFilteredTx)

	for _, log := range logs {
		// skip if the log removed due to chain re-organization
		if log.Removed {
			continue
		}

		// skip if the event has already been passed
		if log.BlockNumber == filter.StartBlockNum && log.TxIndex < filter.StartTxIndex {
			continue
		}

		tx, ok := txMap[log.TxHash]
		if !ok {
			tx = &EthFilteredTx{
				BlockNum:  log.BlockNumber,
				BlockHash: log.BlockHash,
				TXIndex:   log.TxIndex,
				TXHash:    log.TxHash,
				Logs:      make([]types.Log, 0, 10),
			}

			txMap[log.TxHash] = tx
		}

		tx.Logs = append(tx.Logs, log)
	}

	result := &EthTxFilterResult{
		Synced:       synced,
		TXs:          make([]*EthFilteredTx, 0, len(txMap)),
		LastBlockNum: endBlockNum,
	}

	lastBlockNum := uint64(0)
	for _, tx := range txMap {
		if lastBlockNum != tx.BlockNum {
			lastBlockNum = tx.BlockNum
			result.NumBlocks++
		}

		sort.Slice(tx.Logs, func(i, j int) bool { return tx.Logs[i].Index < tx.Logs[j].Index })

		result.NumTXs++
		result.NumLogs += len(tx.Logs)

		result.TXs = append(result.TXs, tx)
	}

	sort.Slice(result.TXs, func(i, j int) bool {
		if result.TXs[i].BlockNum < result.TXs[j].BlockNum {
			return true
		}

		if result.TXs[i].BlockNum > result.TXs[j].BlockNum {
			return false
		}

		return result.TXs[i].TXIndex < result.TXs[j].TXIndex
	})

	return result, nil
}

func ethFilterTXsCalcEndBlockNum(ctx context.Context, ethClient *ethclient.Client, start, stayBehind, limit uint64) (uint64, bool, error) {
	end, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return 0, false, err
	}

	syncProgress, err := ethClient.SyncProgress(ctx)
	if err != nil {
		return 0, false, err
	}

	if syncProgress == nil { // means the connected node is synced
		end -= stayBehind
	} else {
		max := syncProgress.HighestBlock - stayBehind
		if syncProgress.CurrentBlock < max {
			end = syncProgress.CurrentBlock
		}
	}

	synced := true

	if to := start + limit; limit != 0 && to < end {
		synced = false
		end = to
	}

	if start > end {
		end = start
	}

	return end, synced, nil
}
