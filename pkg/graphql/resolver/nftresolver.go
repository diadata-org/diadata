package resolver

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/graph-gophers/graphql-go"
)

type NFTResolver struct {
	n dia.NFT
}

func (nr *NFTResolver) Address(ctx context.Context) (*string, error) {
	return &nr.n.NFTClass.Address, nil
}

func (nr *NFTResolver) Blockchain(ctx context.Context) (*string, error) {
	return &nr.n.NFTClass.Blockchain, nil
}

func (nr *NFTResolver) TokenID(ctx context.Context) (*string, error) {
	return &nr.n.TokenID, nil
}

func (nr *NFTResolver) CreationTime(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: nr.n.CreationTime}, nil
}

func (nr *NFTResolver) CreatorAddress(ctx context.Context) (*string, error) {
	return &nr.n.CreatorAddress, nil
}

func (nr *NFTResolver) URI(ctx context.Context) (*string, error) {
	return &nr.n.URI, nil
}

// func (nr *NFTResolver) Attributes(ctx context.Context) (*dia.NFTAttributes, error) {
// 	return &nr.n.Attributes, nil
// }

// ----------------------------------------------------------------------------

type NFTTradeResolver struct {
	trade dia.NFTTrade
}

func (tr *NFTTradeResolver) Address(ctx context.Context) (*string, error) {
	return &tr.trade.NFT.NFTClass.Address, nil
}

func (tr *NFTTradeResolver) Blockchain(ctx context.Context) (*string, error) {
	return &tr.trade.NFT.NFTClass.Blockchain, nil
}

func (tr *NFTTradeResolver) TokenID(ctx context.Context) (*string, error) {
	return &tr.trade.NFT.TokenID, nil
}

func (tr *NFTTradeResolver) Price(ctx context.Context) (*string, error) {
	price := tr.trade.Price.String()
	return &price, nil
}

// func (tr *NFTTradeResolver) CurrencyAddress(ctx context.Context) (*string, error) {
// 	return &tr.trade.CurrencyAddress, nil
// }

// func (tr *NFTTradeResolver) CurrencySymbol(ctx context.Context) (*string, error) {
// 	return &tr.trade.CurrencySymbol, nil
// }

// func (tr *NFTTradeResolver) CurrencyDecimals(ctx context.Context) (*int32, error) {
// 	decimals := tr.trade.CurrencyDecimals
// 	return &decimals, nil
// }

func (tr *NFTTradeResolver) FromAddress(ctx context.Context) (*string, error) {
	return &tr.trade.FromAddress, nil
}

func (tr *NFTTradeResolver) ToAddress(ctx context.Context) (*string, error) {
	return &tr.trade.FromAddress, nil
}

func (tr *NFTTradeResolver) Timestamp(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: tr.trade.Timestamp}, nil
}

func (tr *NFTTradeResolver) Blocknumber(ctx context.Context) (*int32, error) {
	blocknumber := int32(tr.trade.BlockNumber)
	return &blocknumber, nil
}

func (tr *NFTTradeResolver) TxHash(ctx context.Context) (*string, error) {
	return &tr.trade.TxHash, nil
}

func (tr *NFTTradeResolver) Exchange(ctx context.Context) (*string, error) {
	return &tr.trade.Exchange, nil
}

// -----------------------------------------------------------------------------

type NFTOfferResolver struct {
	offer dia.NFTOffer
}

func (or *NFTOfferResolver) Address(ctx context.Context) (*string, error) {
	return &or.offer.NFT.NFTClass.Address, nil
}

func (or *NFTOfferResolver) Blockchain(ctx context.Context) (*string, error) {
	return &or.offer.NFT.NFTClass.Blockchain, nil
}

func (or *NFTOfferResolver) TokenID(ctx context.Context) (*string, error) {
	return &or.offer.NFT.TokenID, nil
}

func (or *NFTOfferResolver) StartValue(ctx context.Context) (*string, error) {
	var startvalue string
	if or.offer.StartValue != nil {
		startvalue = or.offer.StartValue.String()
	}
	return &startvalue, nil
}

func (or *NFTOfferResolver) EndValue(ctx context.Context) (*string, error) {
	var endvalue string
	if or.offer.EndValue != nil {
		endvalue = or.offer.EndValue.String()
	}
	return &endvalue, nil
}

func (or *NFTOfferResolver) Duration(ctx context.Context) (*int32, error) {
	duration := int32(0)
	return &duration, nil
	// return &or.offer.CurrencyAddress, nil
}

func (or *NFTOfferResolver) AuctionType(ctx context.Context) (*string, error) {
	return &or.offer.AuctionType, nil
}

func (or *NFTOfferResolver) CurrencyAddress(ctx context.Context) (*string, error) {
	return &or.offer.CurrencyAddress, nil
}

func (or *NFTOfferResolver) CurrencySymbol(ctx context.Context) (*string, error) {
	return &or.offer.CurrencySymbol, nil
}

func (or *NFTOfferResolver) CurrencyDecimals(ctx context.Context) (*int32, error) {
	decimals := or.offer.CurrencyDecimals
	return &decimals, nil
}

func (or *NFTOfferResolver) FromAddress(ctx context.Context) (*string, error) {
	return &or.offer.FromAddress, nil
}

func (or *NFTOfferResolver) Timestamp(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: or.offer.Timestamp}, nil
}

func (or *NFTOfferResolver) Blocknumber(ctx context.Context) (*int32, error) {
	blocknumber := int32(or.offer.BlockNumber)
	return &blocknumber, nil
}

func (or *NFTOfferResolver) TxHash(ctx context.Context) (*string, error) {
	return &or.offer.TxHash, nil
}

func (or *NFTOfferResolver) Exchange(ctx context.Context) (*string, error) {
	return &or.offer.Exchange, nil
}

// -----------------------------------------------------------------------------

type NFTBidResolver struct {
	bid dia.NFTBid
}

func (br *NFTBidResolver) Address(ctx context.Context) (*string, error) {
	return &br.bid.NFT.NFTClass.Address, nil
}

func (br *NFTBidResolver) Blockchain(ctx context.Context) (*string, error) {
	return &br.bid.NFT.NFTClass.Blockchain, nil
}

func (br *NFTBidResolver) TokenID(ctx context.Context) (*string, error) {
	return &br.bid.NFT.TokenID, nil
}

func (br *NFTBidResolver) BidValue(ctx context.Context) (*string, error) {
	var bidvalue string
	if br.bid.Value != nil {
		bidvalue = br.bid.Value.String()
	}
	return &bidvalue, nil
}

func (br *NFTBidResolver) CurrencyAddress(ctx context.Context) (*string, error) {
	return &br.bid.CurrencyAddress, nil
}

func (br *NFTBidResolver) CurrencySymbol(ctx context.Context) (*string, error) {
	return &br.bid.CurrencySymbol, nil
}

func (br *NFTBidResolver) CurrencyDecimals(ctx context.Context) (*int32, error) {
	decimals := br.bid.CurrencyDecimals
	return &decimals, nil
}

func (br *NFTBidResolver) FromAddress(ctx context.Context) (*string, error) {
	return &br.bid.FromAddress, nil
}

func (br *NFTBidResolver) Timestamp(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: br.bid.Timestamp}, nil
}

func (br *NFTBidResolver) Blocknumber(ctx context.Context) (*int32, error) {
	blocknumber := int32(br.bid.BlockNumber)
	return &blocknumber, nil
}

func (br *NFTBidResolver) TxHash(ctx context.Context) (*string, error) {
	return &br.bid.TxHash, nil
}

func (br *NFTBidResolver) Exchange(ctx context.Context) (*string, error) {
	return &br.bid.Exchange, nil
}
