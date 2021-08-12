package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Asset struct {
	bun.BaseModel `bun:"asset"`

	Id         uuid.UUID `bun:"type:uuid,column_name:asset_id, default:gen_random_uuid(), pk"`
	Symbol     string    `bun:"type:text,column_name:symbol, notnull,unique:exchangesymbol_symbol_exchange_key"`
	Name       string    `bun:"type:text,column_name:name, notnull"`
	Blockchain string    `bun:"type:text,column_name:blockchain"`
	address    string    `bun:"type:text,column_name:address, notnull"`
}
