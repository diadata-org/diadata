package resolver

import (
	"context"
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/graphql/utils"
	"github.com/graph-gophers/graphql-go"
)

type FilterPointResolver struct {
	q dia.FilterPoint
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
func (qr *FilterPointResolver) Sign(ctx context.Context) (*string, error) {
	su, err := utils.NewSignatureUtil()
	if err != nil {
		return nil, err
	}
	messageTosign, err := json.Marshal(qr.q)
	if err != nil {
		return nil, err
	}

	log.Println("messageTosign", messageTosign)
	log.Println("messageTosign", qr.q)

	signedMessage := su.Sign(messageTosign)
	return &signedMessage, nil
}
