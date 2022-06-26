package resolver

import (
	"context"
)

type FilterPointMetaResolver struct {
	fpr *[]*FilterPointResolver
	max float64
	min float64
}

func (qr *FilterPointMetaResolver) Max(ctx context.Context) (*float64, error) {
	return &qr.max, nil
}

func (qr *FilterPointMetaResolver) Min(ctx context.Context) (*float64, error) {
	return &qr.min, nil
}

func (qr *FilterPointMetaResolver) Points(ctx context.Context) (*[]*FilterPointResolver, error) {
	return qr.fpr, nil
}
