package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/pkg/model"
	"time"
)

type Filter interface {
	compute(trade dia.Trade)
	finalComputeEndOfBlock(t time.Time)
	filterPointForBlock() *dia.FilterPoint
	save(ds models.Datastore) error
}
