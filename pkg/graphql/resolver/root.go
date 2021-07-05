package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
)

// Resolver is the root resolver
type DiaResolver struct {
	DS models.DB
}

// GetQuotation Get quotation
func (r *DiaResolver) GetQuotation(ctx context.Context, args struct{ Symbol graphql.NullString }) (*QuotationResolver, error) {
	q, err := r.DS.GetQuotation(*args.Symbol.Value)
	if err != nil {
		return nil, err
	}
	return &QuotationResolver{q: *q}, nil
}

func (r *DiaResolver) GetSupply(ctx context.Context, args struct{ Symbol graphql.NullString }) (*SupplyResolver, error) {
	q, err := r.DS.GetLatestSupply(*args.Symbol.Value)
	if err != nil {
		return nil, err
	}
	return &SupplyResolver{q: q}, nil
}

func (r *DiaResolver) GetSupplies(ctx context.Context, args struct{ Symbol graphql.NullString }) (*[]*SupplyResolver, error) {
	// starttime := time.Unix(1, 0)
	// endtime := time.Now()
	// q, err := r.DS.GetSupply(*args.Symbol.Value, starttime, endtime)
	// if err != nil {
	// 	return nil, err
	// }

	var sr []*SupplyResolver

	sr = append(sr, &SupplyResolver{q: &dia.Supply{Symbol: "ss"}})
	sr = append(sr, &SupplyResolver{q: &dia.Supply{Symbol: "ss"}})

	return &sr, nil
}
