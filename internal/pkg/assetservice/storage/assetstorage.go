package assetstorage

import "github.com/diadata-org/diadata/pkg/dia"

type StorageHelper interface {
	Save(asset dia.Asset)
	Get() dia.Asset
}





