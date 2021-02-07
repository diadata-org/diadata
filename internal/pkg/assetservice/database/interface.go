package database

import "github.com/diadata-org/diadata/pkg/dia"

type AssetSaver interface {
	Save(asset dia.Asset) error
	GetByName(name string) (dia.Asset, error)
	GetPage(pageNumber int64) (assets []dia.Asset, err error)
	Count() (int64, error)
}
