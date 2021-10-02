package defiscrapers

import (
	"context"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type AAVEProtocol struct {
	ethC     *ethclient.Client
	protocol dia.DefiProtocol
	scraper  *DefiScraper
}

func NewAAVE(scraper *DefiScraper, protocol dia.DefiProtocol) *AAVEProtocol {
	ethC, err := ethhelper.NewETHClient()
	if err != nil {
		panic(err) // TODO: add an error to the function signature
	}

	return &AAVEProtocol{
		ethC:     ethC,
		protocol: protocol,
		scraper:  scraper,
	}
}

func (proto *AAVEProtocol) UpdateRate() error {
	return errors.WithStack(aave.Read(context.Background(), proto.ethC, proto.protocol, proto.scraper.chanDefiRate, nil))
}

func (proto *AAVEProtocol) UpdateState() error {
	return errors.WithStack(aave.Read(context.Background(), proto.ethC, proto.protocol, nil, proto.scraper.chanDefiState))
}
