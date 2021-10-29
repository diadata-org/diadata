package aave

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chRate := make(chan *dia.DefiRate)
	chState := make(chan *dia.DefiProtocolState)

	defer close(chRate)
	defer close(chState)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg, ok := <-chRate:
				if ok {
					fmt.Printf("%s | %s: Lending Rate: %.2f, Borrowing Rate: %.2f\n", msg.Timestamp.Format(time.RFC3339), msg.Asset, msg.LendingRate, msg.BorrowingRate)
				}

			case msg, ok := <-chState:
				if ok {
					fmt.Printf("%s | Total value locked: %.2f USD, %.2f ETH\n", msg.Timestamp.Format(time.RFC3339), msg.TotalUSD, msg.TotalETH)
				}
			}
		}
	}()

	ethC, err := ethhelper.NewETHClient()
	assert.Nil(t, err)

	defer ethC.Close()

	assert.Nil(t, Read(ctx, ethC, dia.DefiProtocol{
		Name:                 "AAVE",
		Address:              "0x3dfd23A6c5E8BbcFc9581d2E864a68feb6a076d3",
		UnderlyingBlockchain: "Ethereum",
		Token:                "",
	}, chRate, chState))
}
