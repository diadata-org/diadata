package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	"time"
)

type Filter interface {
	compute(trade dia.Trade)
	finalComputeEndOfBlock(t time.Time)
	filterPointForBlock() *dia.FilterPoint
	save(ds models.Datastore) error
}
