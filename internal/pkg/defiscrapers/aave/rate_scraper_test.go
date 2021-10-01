package aave

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mgurevin/ethlogscanner"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var protocol = dia.DefiProtocol{
	Name:                 "aave",
	Address:              "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9",
	UnderlyingBlockchain: "Ethereum",
	Token:                "AAVE",
}

type mockCursorStore struct {
	v ethlogscanner.Cursor
}

func (s *mockCursorStore) Store(ctx context.Context, cursor ethlogscanner.Cursor) error {
	s.v = cursor
	return nil
}

func (s *mockCursorStore) Load(context.Context) (ethlogscanner.Cursor, error) {
	return s.v, nil
}

func TestRateScraper(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	ethC, err := ethclient.DialContext(ctx, "ETH_MAINNET_WEB3_ADDR")
	assert.Nil(t, err)

	logger := logrus.StandardLogger()
	logger.Level = logrus.InfoLevel

	chMsg := make(chan *dia.DefiRate)

	cursorStore := &mockCursorStore{v: ethlogscanner.MakeCursor(11420207, 0, 0)} // scan from block 11420207

	s, err := newRateScraper(chMsg, cursorStore, &scraperDeps{
		EthClient: ethC,
		Protocol:  protocol,
		ERC20MD:   NewERC20MetadataCache(ethC),
		Logger:    logger.WithContext(ctx),
	})

	assert.Nil(t, err)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-s.chMsg:
				log.Printf("%s | %s  Lending: %.2f  Borrowing: %.2f", msg.Timestamp.Format(time.RFC3339), msg.Asset, msg.LendingRate, msg.LendingRate)
			}
		}
	}()

	err = s.scrapRates(ctx, 15*time.Second)
	assert.Nil(t, err)
}
