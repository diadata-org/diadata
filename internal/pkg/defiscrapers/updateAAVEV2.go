package defiscrapers

import (
	"context"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type AAVEv2Protocol struct {
	ethC     *ethclient.Client
	protocol dia.DefiProtocol
	scraper  *DefiScraper
}

func NewAAVEv2(scraper *DefiScraper, protocol dia.DefiProtocol) *AAVEv2Protocol {
	ethC, err := ethhelper.NewETHClient()
	if err != nil {
		panic(err) // TODO: add an error to the function signature
	}

	return &AAVEv2Protocol{
		ethC:     ethC,
		protocol: protocol,
		scraper:  scraper,
	}
}

func (proto *AAVEv2Protocol) UpdateRate() error {
	return errors.WithStack(aave.Read(context.Background(), proto.ethC, proto.protocol, proto.scraper.chanDefiRate, nil))
}

func (proto *AAVEv2Protocol) UpdateState() error {
	return errors.WithStack(aave.Read(context.Background(), proto.ethC, proto.protocol, nil, proto.scraper.chanDefiState))
}
