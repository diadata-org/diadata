package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ExchangeSymbol struct {
	bun.BaseModel `bun:"exchangesymbol"`

	Id       uuid.UUID `bun:"type:uuid,column_name:exchangesymbol_id, default:gen_random_uuid(), pk"`
	Symbol   string    `bun:"type:text,column_name:symbol, notnull,unique:exchangesymbol_symbol_exchange_key"`
	Exchange string    `bun:"type:text,column_name:exchange, notnull,unique:exchangesymbol_symbol_exchange_key"`
	Verified bool      `bun:"type:boolean,column_name:verified, default:false"`
	//AssetId		uuid.UUID	`bun:"type:uuid"`
	Assets []*Asset `bun:"rel:has-many,join:asset_id=asset_id"`
}
