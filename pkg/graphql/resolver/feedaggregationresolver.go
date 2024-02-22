package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/graph-gophers/graphql-go"
)

type FeedSelectionAggregatedResolver struct {
	q dia.FeedSelectionAggregated
}

func (qr *FeedSelectionAggregatedResolver) Exchange(ctx context.Context) (*string, error) {
	return &qr.q.Exchange, nil
}

func (qr *FeedSelectionAggregatedResolver) Quotetokenaddress(ctx context.Context) (*string, error) {
	return &qr.q.Quotetoken.Address, nil
}

func (qr *FeedSelectionAggregatedResolver) Quotetokenblockchain(ctx context.Context) (*string, error) {
	return &qr.q.Quotetoken.Blockchain, nil
}

func (qr *FeedSelectionAggregatedResolver) Quotetokensymbol(ctx context.Context) (*string, error) {
	return &qr.q.Quotetoken.Symbol, nil
}

func (qr *FeedSelectionAggregatedResolver) Basetokenaddress(ctx context.Context) (*string, error) {
	return &qr.q.Basetoken.Address, nil
}

func (qr *FeedSelectionAggregatedResolver) Basetokenblockchain(ctx context.Context) (*string, error) {
	return &qr.q.Basetoken.Blockchain, nil
}

func (qr *FeedSelectionAggregatedResolver) Basetokensymbol(ctx context.Context) (*string, error) {
	return &qr.q.Basetoken.Symbol, nil
}

func (qr *FeedSelectionAggregatedResolver) Pooladdress(ctx context.Context) (*string, error) {
	return &qr.q.Pooladdress, nil
}

func (qr *FeedSelectionAggregatedResolver) PoolLiquidityUSD(ctx context.Context) (*float64, error) {
	return &qr.q.PoolLiquidityUSD, nil
}

func (qr *FeedSelectionAggregatedResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Endtime}, nil
}

func (qr *FeedSelectionAggregatedResolver) Volume(ctx context.Context) (*float64, error) {
	return &qr.q.Volume, nil
}

func (qr *FeedSelectionAggregatedResolver) TradesCount(ctx context.Context) (*int32, error) {
	return &qr.q.TradesCount, nil
}

func (qr *FeedSelectionAggregatedResolver) LastPrice(ctx context.Context) (*float64, error) {
	return &qr.q.LastPrice, nil
}
