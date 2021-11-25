package source

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

type SerumPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     string
}

const (
	// Public Solana clients.
	wsDialSolana   = "http://192.168.1.88:8900"
	restDialSolana = "http://192.168.1.88:8899"
)

type SerumAssetSource struct {
	assetChannel chan dia.Asset
	blockchain   string
}

func NewSerumAssetSource(exchange dia.Exchange) *SerumAssetSource {

	var assetChannel = make(chan dia.Asset)
	var sas *SerumAssetSource

	exchangeFactoryContractAddress = ""

	sas = &SerumAssetSource{
		assetChannel: assetChannel,
		blockchain:   dia.SOLANA,
	}

	go func() {
		sas.fetchAssets()
	}()
	return sas

}

func (sas *SerumAssetSource) Asset() chan dia.Asset {
	return sas.assetChannel
}

func (sas *SerumAssetSource) fetchAssets() {
	// TO DO: Implement fetch pairs/assets logic here.
}
