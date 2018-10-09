package filters

import (
	"github.com/diadata-org/api-golang/internal/pkg/model"
	"github.com/diadata-org/api-golang/pkg/dia"
	"time"
)

type Filter interface {
	compute(trade dia.Trade)
	finalComputeEndOfBlock(t time.Time)
	filterPointForBlock() *dia.FilterPoint
	save(ds models.Datastore) error
}
