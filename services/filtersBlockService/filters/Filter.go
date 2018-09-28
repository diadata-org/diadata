package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/services/model"
	"time"
)

type Filter interface {
	compute(trade *dia.Trade)
	filterPoint(time time.Time) dia.FilterPoint
	save(ds models.Datastore) error
	copyToFilterBlock() bool
}
