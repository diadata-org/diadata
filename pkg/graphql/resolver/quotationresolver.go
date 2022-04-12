package resolver

import (
	"context"
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/graphql/utils"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/graph-gophers/graphql-go"
)

type QuotationResolver struct {
	q models.Quotation
}

func (qr *QuotationResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.Symbol, nil
}

func (qr *QuotationResolver) Name(ctx context.Context) (*string, error) {
	return &qr.q.Name, nil
}

func (qr *QuotationResolver) Price(ctx context.Context) (*float64, error) {
	return &qr.q.Price, nil
}

func (qr *QuotationResolver) Source(ctx context.Context) (*string, error) {
	return &qr.q.Source, nil
}

func (qr *QuotationResolver) ITIN(ctx context.Context) (*string, error) {
	return &qr.q.ITIN, nil
}

func (qr *QuotationResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Time}, nil
}

func (qr *QuotationResolver) MAIR(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Time}, nil
}

func (qr *QuotationResolver) Sign(ctx context.Context) (*string, error) {
	su, err := utils.NewSignatureUtil()
	if err != nil {
		return nil, err
	}
	messageTosign, err := json.Marshal(&qr.q)
	if err != nil {
		return nil, err
	}
	signedMessage := su.Sign(messageTosign)
	return &signedMessage, nil
}
