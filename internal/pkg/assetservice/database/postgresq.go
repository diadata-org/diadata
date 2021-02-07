package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
)

type PostgresDatabase struct {
	URI  string
	conn *pgx.Conn
}

// NewPostgres returns a postres database connection
func NewPostgres(url string) (*PostgresDatabase, error) {
	pg := &PostgresDatabase{}
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	pg.conn = conn
	return pg, nil
}

// Save stores an asset into postgres
func (pg *PostgresDatabase) Save(asset dia.Asset) error {
	_, err := pg.conn.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetByName returns a dia.Asset by its name
func (pg *PostgresDatabase) GetByName(name string) (asset dia.Asset, err error) {
	var decimals string
	err = pg.conn.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where name=$1", name).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
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

// Count returns the number of assets stored in postgres
func (pg *PostgresDatabase) Count() (count int64, err error) {
	err = pg.conn.QueryRow(context.Background(), "select count(*) from asset").Scan(&count)
	if err != nil {
		return
	}
	return
}

// GetPage returns assets per page numnber
func (pg *PostgresDatabase) GetPage(pageNumber int64) (assets []dia.Asset, err error) {

	var limit int64
	limit = 100
	skip := limit * pageNumber
	rows, err := pg.conn.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", limit, skip)
	if err != nil {
		return
	}

	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals)
		assets = append(assets, asset)
	}
	return
}
