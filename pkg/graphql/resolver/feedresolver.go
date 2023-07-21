package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/graph-gophers/graphql-go"
)

type FilterPointExtendedResolver struct {
	q dia.FilterPointExtended
}

func (qr *FilterPointExtendedResolver) Name(ctx context.Context) (*string, error) {
	return &qr.q.FilterPoint.Name, nil
}

func (qr *FilterPointExtendedResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.FilterPoint.Asset.Symbol, nil
}

func (qr *FilterPointExtendedResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.FilterPoint.Time}, nil
}

func (qr *FilterPointExtendedResolver) Value(ctx context.Context) (*float64, error) {
	return &qr.q.FilterPoint.Value, nil
}

func (qr *FilterPointExtendedResolver) Address(ctx context.Context) (*string, error) {
	return &qr.q.FilterPoint.Asset.Address, nil
}

func (qr *FilterPointExtendedResolver) Blockchain(ctx context.Context) (*string, error) {
	return &qr.q.FilterPoint.Asset.Blockchain, nil
}

func (qr *FilterPointExtendedResolver) Pools(ctx context.Context) (*[]string, error) {
	pools := []string{}
	for _, p := range qr.q.Pools {
		pools = append(pools, p.Exchange.Name+":"+p.Address)
	}
	return &pools, nil
}

func (qr *FilterPointExtendedResolver) Pairs(ctx context.Context) (*[]string, error) {
	pairs := []string{}
	for _, p := range qr.q.Pairs {
		pairs = append(pairs, p.Exchange+":"+p.ForeignName)
	}
	return &pairs, nil
}

func (qr *FilterPointExtendedResolver) TradesCount(ctx context.Context) (*int32, error) {
	return &qr.q.TradesCount, nil
}
