package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/graph-gophers/graphql-go"
)

type FilterPointResolver struct {
	q dia.AssetFilterPoint
}
type TradeResolver struct {
	q dia.Trade
}

func (qr *FilterPointResolver) Name(ctx context.Context) (*string, error) {
	return &qr.q.Name, nil
}

func (qr *FilterPointResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Symbol, nil
}

func (qr *FilterPointResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Time}, nil
}

func (qr *FilterPointResolver) Value(ctx context.Context) (*float64, error) {
	return &qr.q.Value, nil
}
func (qr *FilterPointResolver) Address(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Address, nil
}

func (qr *FilterPointResolver) Blockchain(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Blockchain, nil
}

func (qr *FilterPointResolver) FirstTrade(ctx context.Context) (*TradeResolver, error) {
	return &TradeResolver{q: qr.q.FirstTrade}, nil
}
func (qr *FilterPointResolver) LastTrade(ctx context.Context) (*TradeResolver, error) {
	return &TradeResolver{q: qr.q.LastTrade}, nil
}

func (qr *TradeResolver) Price(ctx context.Context) (*float64, error) {
	return &qr.q.Price, nil
}
func (qr *TradeResolver) Pair(ctx context.Context) (*string, error) {
	return &qr.q.Pair, nil
}

func (qr *TradeResolver) Volume(ctx context.Context) (*float64, error) {
	return &qr.q.Volume, nil
}
func (qr *TradeResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.Symbol, nil
}

func (qr *TradeResolver) EstimatedUSDPrice(ctx context.Context) (*float64, error) {
	return &qr.q.EstimatedUSDPrice, nil
}
