package api

import (
	"context"
	"github.com/diadata-org/diadata/pkg/dia/database"
	"github.com/diadata-org/diadata/pkg/dia/database/models"
)

func GetUnverifiedExchangeSymbols(exchange string) (symbols []string, err error) {
	ctx := context.Background()
	db := database.GetReadDb()
	exchangeSymbols := new(models.ExchangeSymbol)
	if err := db.NewSelect().
		Model(exchangeSymbols).
		Where("exchange=?", exchange).
		Where("verified=false").
		Order("symbol ASC").
		Scan(ctx, &symbols); err != nil {
		return []string{}, err
	}
	return
}
