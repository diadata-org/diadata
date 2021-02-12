package database

import "github.com/diadata-org/diadata/pkg/dia"

// AssetStore is the interface for our asset database.
// It can be implemented by methods using underlying DBs such as postgres, redis,...
type AssetStore interface {
	SetAsset(asset dia.Asset) error
	GetAsset(symbol, name string) (dia.Asset, error)
	GetPage(pageNumber uint32) (assets []dia.Asset, hasNext bool, err error)
	Count() (uint32, error)
}
