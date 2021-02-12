package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
)

// AssetDB is our single source of truth for assets
type AssetDB struct {
	URI      string
	conn     *pgx.Conn
	pagesize uint32
}

// NewAssetDB returns a postres database connection
func NewAssetDB(url string) (*AssetDB, error) {
	pg := &AssetDB{}
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	pg.conn = conn
	return pg, nil
}

// SetAsset stores an asset into postgres
func (pg *AssetDB) SetAsset(asset dia.Asset) error {
	_, err := pg.conn.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetAsset returns a dia.Asset by its symbol and name from postgres
func (pg *AssetDB) GetAsset(symbol, name string) (asset dia.Asset, err error) {
	var decimals string
	err = pg.conn.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1 and name=$2", symbol, name).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	// TO DO: Get Blockchain by name from postgres and add to asset
	return
}

// GetPage returns assets per page number. @hasNext is true iff there is a non-empty next page.
func (pg *AssetDB) GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error) {

	pagesize := pg.pagesize
	skip := pagesize * pageNumber
	rows, err := pg.conn.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip)
	if err != nil {
		return
	}
	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals)
		assets = append(assets, asset)
	}
	// Last page (or empty page)
	if len(rows.RawValues()) < int(pagesize) {
		hasNextPage = false
		return
	}
	// No next page
	nextPageRows, err := pg.conn.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip+1)
	if len(nextPageRows.RawValues()) == 0 {
		hasNextPage = false
		return
	}
	hasNextPage = true
	return
}

// Count returns the number of assets stored in postgres
func (pg *AssetDB) Count() (count uint32, err error) {
	err = pg.conn.QueryRow(context.Background(), "select count(*) from asset").Scan(&count)
	if err != nil {
		return
	}
	return
}
