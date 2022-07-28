package models

import (
	"context"
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
