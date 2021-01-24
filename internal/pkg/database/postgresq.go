package database

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/jackc/pgx/v4"
)

type PostgresDatabase struct {
	URI  string
	conn *pgx.Conn
}

//Save(asset dia.Asset) error
//GetByName(name string)   (dia.Asset,error)
//GetPage(pageNumber int64) (assets []dia.Asset, err error)
//Count() (int64, error)

func NewPostgres(url string) (*PostgresDatabase, error) {
	pg := &PostgresDatabase{}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	pg.conn = conn

	return pg, nil
}

func (pg *PostgresDatabase) Save(asset dia.Asset) error {
	_, err := pg.conn.Exec(context.Background(), "insert into asset(symbol,name,address,decimals) values ($1, $2,$3,$4)", asset.Symbol, asset.Name, asset.Address, asset.Decimals)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PostgresDatabase) GetByName(name string) (asset dia.Asset, err error) {
	err = pg.conn.QueryRow(context.Background(), "select symbol,name,address,decimals from asset where name=$1", name).Scan(&asset.Symbol,&asset.Name,&asset.Address,&asset.Decimals)
	if err != nil {
		return
	}
	return
}
func (pg *PostgresDatabase) Count() (count int64, err error) {
	err = pg.conn.QueryRow(context.Background(), "select count(*) from asset").Scan(&count)
	if err != nil {
		return
	}
	return
}

func (pg *PostgresDatabase) GetPage(pageNumber int64) (assets []dia.Asset, err error){

	var limit int64
	limit = 100
	skip := limit * pageNumber
	rows,err := pg.conn.Query(context.Background(), "select symbol,name,address,decimals from asset LIMIT $1 OFFSET $2 ", limit,skip)

	if err!=nil{
		return
	}
	for rows.Next(){
		fmt.Println("---")
		 var asset dia.Asset
		 rows.Scan(&asset.Symbol,&asset.Name,&asset.Address,&asset.Decimals)
		assets = append(assets,asset)
	}


	return

}

