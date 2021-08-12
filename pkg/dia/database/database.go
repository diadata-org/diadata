package database

import (
	postgres "github.com/diadata-org/diadata/pkg/dia/database/dialects"
	"github.com/uptrace/bun"
)

func GetReadDb() *bun.DB {
	return postgres.GetConnection()
}
