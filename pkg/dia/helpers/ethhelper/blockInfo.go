package ethhelper

import (
	"context"
	"math/big"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/structs"
)

// GetBlockData returns relevant block data from block with @blockNumber.
func GetBlockData(blockNumber int64, relDB *models.RelDB, client *ethclient.Client) (blockdata dia.BlockData, err error) {
	blockdata, err = relDB.GetBlockData(dia.ETHEREUM, blockNumber)
	if err != nil {
		if err.Error() == "no rows in result set" {
			blockdata, err = GetBlockDataOnChain(blockNumber, client)
			if err != nil {
				return
			}
		}
	}
	return blockdata, err
}

func GetBlockDataOnChain(blockNumber int64, client *ethclient.Client) (dia.BlockData, error) {

	var ethblockdata dia.EthereumBlockData
	var blockdata dia.BlockData
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return dia.BlockData{}, err
	}

	ethblockdata.Coinbase = block.Coinbase()
	ethblockdata.Difficulty = block.Difficulty()
	ethblockdata.Extra = block.Extra()
	ethblockdata.GasLimit = block.GasLimit()
	ethblockdata.GasUsed = block.GasUsed()
	ethblockdata.MixDigest = block.MixDigest()
	ethblockdata.Nonce = block.Nonce()
	ethblockdata.Number = block.NumberU64()
	ethblockdata.ParentHash = block.ParentHash()
	ethblockdata.ReceiptHash = block.ReceiptHash()
	ethblockdata.Root = block.Root()
	ethblockdata.Size = block.Size()
	ethblockdata.Time = block.Time()
	ethblockdata.TxHash = block.TxHash()
	ethblockdata.UncleHash = block.UncleHash()

	blockdata.BlockchainName = dia.ETHEREUM
	blockdata.BlockNumber = int64(ethblockdata.Number)
	blockdata.Data = structs.Map(ethblockdata)

	return blockdata, nil

}

// GetBlockTimeEth returns the block time of @blockNumber on Ethereum mainnet.
func GetBlockTimeEth(blockNumber int64, relDB *models.RelDB, client *ethclient.Client) (blockTime time.Time, err error) {
	blockData, err := GetBlockData(blockNumber, relDB, client)
	if err != nil {
		return
	}
	switch blockData.Data["Time"].(type) {
	case float64:
		blockTime = time.Unix(int64(blockData.Data["Time"].(float64)), 0)
	case uint64:
		blockTime = time.Unix(int64(blockData.Data["Time"].(uint64)), 0)
	}
	return
}
