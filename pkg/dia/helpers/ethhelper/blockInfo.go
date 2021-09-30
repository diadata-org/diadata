package ethhelper

import (
	"context"
	"math/big"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/structs"
)

// GetBlockData returns
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
