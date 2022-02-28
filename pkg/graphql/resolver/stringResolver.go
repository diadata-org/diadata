package resolver

import (
	"context"
)

type StringResolver struct {
	q string
}

func (qr *QuotationResolver) Data(ctx context.Context) (*string, error) {
	return &qr.q.Symbol, nil
}
