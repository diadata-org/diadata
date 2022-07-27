package resolver

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type vwalp struct {
	Symbol string
	Value  float64
	Time   time.Time
}

type VWALPResolver struct {
	q vwalp
}

func (vr *VWALPResolver) Symbol(ctx context.Context) (*string, error) {
	return &vr.q.Symbol, nil
}

func (vr *VWALPResolver) Value(ctx context.Context) (*float64, error) {
	return &vr.q.Value, nil
}

func (vr *VWALPResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: vr.q.Time}, nil
}
