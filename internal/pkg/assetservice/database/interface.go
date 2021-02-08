package database

import "github.com/diadata-org/diadata/pkg/dia"

// AssetStore is the interface for our asset database.
// It can be implemented by methods using underlying DBs such as postgres, redis,...
type AssetStore interface {
	Save(asset dia.Asset) error
	GetByName(name string) (dia.Asset, error)
	GetPage(pageNumber int64) (assets []dia.Asset, err error)
	Count() (int64, error)
}
