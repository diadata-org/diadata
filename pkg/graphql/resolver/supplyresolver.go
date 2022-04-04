package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/graphql/utils"
	"github.com/graph-gophers/graphql-go"
)

type SupplyResolver struct {
	q *dia.Supply
}

func (qr *SupplyResolver) Symbol(ctx context.Context) (*string, error) {
	fmt.Println("Symbol")

	return &qr.q.Asset.Symbol, nil
}

func (qr *SupplyResolver) Name(ctx context.Context) (*string, error) {
	fmt.Println("Name")

	return &qr.q.Asset.Name, nil
}

func (qr *SupplyResolver) Source(ctx context.Context) (*string, error) {
	fmt.Println("Source")

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

func (qr *SupplyResolver) Sign(ctx context.Context) (*string, error) {
	su, err := utils.NewSignatureUtil()
	if err != nil {
		return nil, err
	}
	messageTosign, err := json.Marshal(qr.q)
	if err != nil {
		return nil, err
	}
	signedMessage := su.Sign(messageTosign)
	return &signedMessage, nil
}
