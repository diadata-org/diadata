package models

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetSynthAssetSupply stores a synth asset along with the underlying data into postgres.
func (rdb *RelDB) SetSynthAssetSupply(synthData dia.SynthAssetSupply) error {
	// TO DO

	// query := fmt.Sprintf("INSERT INTO %s (symbol,name,address,decimals,blockchain) VALUES ($1,$2,$3,$4,$5)", assetTable)
	// _, err := rdb.postgresClient.Exec(context.Background(), query, asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (rdb *RelDB) GetSynthAssetSupply(address string, blockchain string, timestamp time.Time) (synthAssetSupply dia.SynthAssetSupply, err error) {
	// TO DO
	return
}
