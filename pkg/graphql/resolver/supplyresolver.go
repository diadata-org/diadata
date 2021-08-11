package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/graph-gophers/graphql-go"
)

type SupplyResolver struct {
	q *dia.Supply
}

func (qr *SupplyResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Symbol, nil
}

func (qr *SupplyResolver) Name(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Name, nil
}

func (qr *SupplyResolver) Source(ctx context.Context) (*string, error) {
	return &qr.q.Source, nil
}

func (qr *SupplyResolver) CirculatingSupply(ctx context.Context) (*float64, error) {
	return &qr.q.CirculatingSupply, nil
}

func (qr *SupplyResolver) Supply(ctx context.Context) (*float64, error) {
	return &qr.q.Supply, nil
}

func (qr *SupplyResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Time}, nil
}

func (qr *SupplyResolver) MA(ctx context.Context) (*graphql.Time, error) {
	// time,
	return &graphql.Time{Time: qr.q.Time}, nil
}
