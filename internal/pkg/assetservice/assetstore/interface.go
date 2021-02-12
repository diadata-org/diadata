package assetstore

import "github.com/diadata-org/diadata/pkg/dia"

// AssetDB is the interface for our asset database(s).
// It can be implemented by methods using underlying DBs such as postgres, redis,...
type AssetDB interface {
	SetAsset(asset dia.Asset) error
	GetAsset(symbol, name string) (dia.Asset, error)
	GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error)
	Count() (uint32, error)
}
