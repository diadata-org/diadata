package defiscrapers

import (
	"context"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/mgurevin/ethlogscanner"
	"github.com/pkg/errors"
)

type AAVEProtocol struct {
	ctx           context.Context
	rateScrapDur  time.Duration
	rateScraper   *aave.RateScraper
	reserveReader *aave.ReserveReader
}

type aaveScraperState struct {
	Cursor ethlogscanner.Cursor `json:"cursor"`
}

type cursorStore struct {
	relDB *models.RelDB
}

func NewAAVE(scraper *DefiScraper, protocol dia.DefiProtocol) (*AAVEProtocol, error) {
	ethC, err := ethhelper.NewETHClient()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	deps := &aave.ScraperDeps{
		Protocol:  protocol,
		EthClient: ethC,
		ERC20MD:   aave.NewERC20MetadataCache(ethC),
		Logger:    scraper.log,
	}

	cs := &cursorStore{
		relDB: scraper.relDB,
	}

	rateScraper, err := aave.NewRateScraper(scraper.chanDefiRate, cs, deps)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reserveReader, err := aave.NewReserveReader(scraper.chanDefiState, deps)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AAVEProtocol{
		ctx:           scraper.ctx,
		rateScraper:   rateScraper,
		reserveReader: reserveReader,
		rateScrapDur:  refreshRateDelay - (5 * time.Second),
	}, nil
}

func (s *AAVEProtocol) UpdateRate() error {
	return errors.WithStack(s.rateScraper.ScrapRates(s.ctx, s.rateScrapDur))
}

func (s *AAVEProtocol) UpdateState() error {
	return errors.WithStack(s.reserveReader.Read(s.ctx))
}

func (s *cursorStore) Store(ctx context.Context, cursor ethlogscanner.Cursor) error {
	return errors.WithStack(s.relDB.SetScraperState(ctx, "aave_v2", &aaveScraperState{Cursor: cursor}))
}

func (s *cursorStore) Load(ctx context.Context) (ethlogscanner.Cursor, error) {
	state := &aaveScraperState{}

	err := s.relDB.GetScraperState(ctx, "aave_v2", state)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return state.Cursor, nil
}
