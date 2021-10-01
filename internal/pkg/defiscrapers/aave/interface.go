package aave

import (
	"context"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mgurevin/ethlogscanner"
	"github.com/sirupsen/logrus"
)

type CursorStore interface {
	Store(ctx context.Context, cursor ethlogscanner.Cursor) error
	Load(ctx context.Context) (ethlogscanner.Cursor, error)
}

type ERC20Metadata interface {
	Name() string
	Symbol() string
	Decimals() int
}

type ERC20MetadataProvider interface {
	ERC20(ctx context.Context, assetAddr common.Address) (ERC20Metadata, error)
}

type ScraperDeps struct {
	EthClient *ethclient.Client
	Protocol  dia.DefiProtocol
	ERC20MD   ERC20MetadataProvider
	Logger    *logrus.Entry
}
