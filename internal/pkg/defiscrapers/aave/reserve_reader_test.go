package aave

import (
	"context"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestReserveReader(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	ethC, err := ethclient.DialContext(ctx, "ETH_MAINNET_WEB3_ADDR")
	assert.Nil(t, err)

	logger := logrus.StandardLogger()
	logger.Level = logrus.InfoLevel

	chMsg := make(chan *dia.DefiProtocolState)

	r, err := newReserveReader(chMsg, &scraperDeps{
		EthClient: ethC,
		Protocol:  protocol,
		ERC20MD:   NewERC20MetadataCache(ethC),
		Logger:    logger.WithContext(ctx),
	})

	assert.Nil(t, err)

	err = r.read(ctx)

	assert.Nil(t, err)
}
