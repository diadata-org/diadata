package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ScraperConfig is a JSON compatible struct to keep the configuration of a scraper
type ScraperConfig interface{}

// ScraperState is a JSON compatible struct to keep the state of a scraper
type ScraperState interface{}

func (rdb *RelDB) GetScraperState(ctx context.Context, scraperName string, state ScraperState) error {
	return rdb.postgresClient.QueryRow(ctx, fmt.Sprintf("select state from %s where name=$1", scrapersTable), scraperName).Scan(state)
}

func (rdb *RelDB) SetScraperState(ctx context.Context, scraperName string, state ScraperState) error {
	_, err := rdb.postgresClient.Exec(ctx, fmt.Sprintf("insert into %s(name, state) values($1, $2) on conflict(name) do update set state=excluded.state", scrapersTable),
		scraperName,
		state,
	)

	return err
}

func (rdb *RelDB) GetScraperConfig(ctx context.Context, scraperName string, config ScraperConfig) error {
	return rdb.postgresClient.QueryRow(ctx, fmt.Sprintf("select conf from %s where name=$1", scrapersTable), scraperName).Scan(config)
}

func (rdb *RelDB) SetScraperConfig(ctx context.Context, scraperName string, config ScraperConfig) error {
	_, err := rdb.postgresClient.Exec(ctx, fmt.Sprintf("insert into %s(name, conf) values($1, $2) on conflict(name) do update set conf=excluded.conf", scrapersTable),
		scraperName,
		config,
	)

	return err
}

func (rdb *RelDB) SetScraperIndex(exchange string, scraperType string, indexType string, index int64) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (scraper,index_type,index_value) VALUES ($1,$2,$3) 
		ON CONFLICT (scraper,index_type)
		DO UPDATE SET index_value=EXCLUDED.index_value`,
		scraperCronjobStateTable,
	)
	_, err := rdb.postgresClient.Exec(
		context.Background(),
		query,
		exchange+"_"+scraperType,
		indexType,
		index,
	)
	return err
}

func (rdb *RelDB) GetScraperIndex(exchange string, scraperType string) (string, int64, error) {
	query := fmt.Sprintf(
		"SELECT index_type,index_value FROM %s WHERE scraper=$1",
		scraperCronjobStateTable,
	)

	var indexType sql.NullString
	var index sql.NullFloat64
	err := rdb.postgresClient.QueryRow(context.Background(), query, exchange+"_"+scraperType).Scan(&indexType, &index)
	if err != nil {
		return "", 0, err
	}
	if index.Valid {
		return indexType.String, int64(index.Float64), nil
	}
	return "", 0, errors.New("blocknumber not valid")
}
