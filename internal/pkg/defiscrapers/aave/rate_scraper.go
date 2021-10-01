package aave

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	contract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave/bind/aave"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mgurevin/ethlogscanner"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type RateScraper struct {
	ethC *ethclient.Client

	cursorStore CursorStore
	erc20MD     ERC20MetadataProvider

	protocol dia.DefiProtocol

	aaveContractAddr         common.Address
	aaveABI                  abi.ABI
	aaveEvReserveDataUpdated abi.Event
	aaveBoundedContract      *bind.BoundContract

	log *logrus.Entry

	chMsg chan *dia.DefiRate
}

func NewRateScraper(messageCh chan *dia.DefiRate, cursorStore CursorStore, deps *ScraperDeps) (*RateScraper, error) {
	var err error
	var ok bool

	s := &RateScraper{
		ethC:             deps.EthClient,
		aaveContractAddr: common.HexToAddress(deps.Protocol.Address),
		log:              deps.Logger.WithField("comp", "aave-rate-scraper"),
		cursorStore:      cursorStore,
		erc20MD:          deps.ERC20MD,
		chMsg:            messageCh,
	}

	s.aaveABI, err = abi.JSON(strings.NewReader(contract.ILendingPoolABI))
	if err != nil {
		return nil, err
	}

	s.aaveEvReserveDataUpdated, ok = s.aaveABI.Events["ReserveDataUpdated"]
	if !ok {
		return nil, fmt.Errorf("event not found in the ABI: ReserveDataUpdated")
	}

	s.aaveBoundedContract = bind.NewBoundContract(s.aaveContractAddr, s.aaveABI, nil, nil, nil)

	return s, nil
}

func (s *RateScraper) ScrapRates(ctx context.Context, duration time.Duration) (err error) {
	var scanner ethlogscanner.Scanner

	cursor, err := s.cursorStore.Load(ctx)
	if err != nil {
		return err
	}

	s.log.WithFields(logrus.Fields{
		"start_block_num": cursor.BlockNum(),
		"start_tx_index":  cursor.TxIndex(),
		"start_log_index": cursor.LogIndex(),
	}).Info("scanner is starting")

	defer func(t time.Time, cursor ethlogscanner.Cursor) {
		dur := time.Since(t).Truncate(time.Millisecond).String()
		blocks := scanner.Next().BlockNum() - cursor.BlockNum()

		l := s.log.WithFields(logrus.Fields{
			"next_block_num": scanner.Next().BlockNum(),
			"next_tx_index":  scanner.Next().TxIndex(),
			"next_log_index": scanner.Next().LogIndex(),
			"duration":       dur,
			"num_blocks":     blocks,
		})

		if err != nil {
			l.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("scanner is completed with an error")

			return
		}

		l.Info("scanner is completed successfully")
	}(time.Now(), cursor)

	scanner, err = ethlogscanner.Scan(ctx, s.ethC,
		ethlogscanner.WithStart(cursor),
		ethlogscanner.WithFilter(
			[]common.Address{
				s.aaveContractAddr,
			},
			[][]common.Hash{
				{
					s.aaveEvReserveDataUpdated.ID,
				},
			},
		),
	)

	if err != nil {
		return err
	}

	defer scanner.Close()

	timer := time.NewTimer(duration)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			s.log.Trace("scan time is up")

			scanner.Close()

		case err := <-scanner.Done():
			return err

		case n := <-scanner.Notify():
			if n != nil {
				switch n := n.(type) {
				case *ethlogscanner.ChunkSizeUpdated:
					s.log.WithFields(logrus.Fields{
						"previous": n.Previous,
						"updated":  n.Updated,
					}).Trace("chunk size updated")

				case *ethlogscanner.FilterStarted:
					s.log.WithFields(logrus.Fields{
						"from":  n.From,
						"to":    n.To,
						"chunk": n.ChunkSize,
					}).Trace("filtering started")

				case *ethlogscanner.FilterCompleted:
					s.log.WithFields(logrus.Fields{
						"from":    n.From,
						"to":      n.To,
						"chunk":   n.ChunkSize,
						"took":    n.Duration.Truncate(time.Millisecond).String(),
						"has_err": n.HasErr,
					}).Debug("filtering completed")

				case *ethlogscanner.CursorUpdated:
					s.log.WithFields(logrus.Fields{
						"block_num": n.Next.BlockNum(),
						"tx_index":  n.Next.TxIndex(),
						"log_index": n.Next.LogIndex(),
					}).Trace("cursor updated")

					if err := s.cursorStore.Store(ctx, n.Next); err != nil {
						return errors.Wrapf(err, "unable to store cursor value %s: %s", n.Next.String(), err.Error())
					}
				}
			}

		case err := <-scanner.Err():
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}

				s.log.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Debug("a scan error occurred")
			}

		case log := <-scanner.Log():
			if log != nil {
				l := &contract.ILendingPoolReserveDataUpdated{}

				raw := types.Log(*log)

				if err := s.aaveBoundedContract.UnpackLog(l, s.aaveEvReserveDataUpdated.RawName, raw); err != nil {
					return errors.Wrapf(err, "unable to parse the ReserveDataUpdated log in TX %s(%d): %s", log.TxHash.String(), log.Index, err.Error())
				}

				l.Raw = raw

				s.log.WithFields(logrus.Fields{
					"block_num":    log.Cursor().BlockNum(),
					"tx_hash":      log.TxHash.String(),
					"log_index":    log.Cursor().LogIndex(),
					"reserve_addr": l.Reserve.String(),
				}).Trace("found ReserveDataUpdated event")

				if err := s.handleLog(ctx, l); err != nil {
					return err
				}
			}
		}
	}
}

func (s *RateScraper) handleLog(ctx context.Context, log *contract.ILendingPoolReserveDataUpdated) error {
	liqRate, _ := new(big.Float).Quo(new(big.Float).SetInt(log.LiquidityRate), big.NewFloat(math.Pow10(25))).Float64()
	stableBorrowRate, _ := new(big.Float).Quo(new(big.Float).SetInt(log.StableBorrowRate), big.NewFloat(math.Pow10(25))).Float64()

	hdrs, err := s.ethC.HeaderByNumber(ctx, new(big.Int).SetUint64(log.Raw.BlockNumber))
	if err != nil {
		return errors.Wrapf(err, "unable to read block headers for the ReserveDataUpdate event at TX %s: %s", log.Raw.TxHash.String(), err.Error())
	}

	erc20, err := s.erc20MD.ERC20(ctx, log.Reserve)
	if err != nil {
		return errors.Wrapf(err, "unable to get ERC20 metadata of the address %s: %s", log.Reserve.String(), err.Error())
	}

	msg := &dia.DefiRate{
		Timestamp:     time.Unix(int64(hdrs.Time), 0),
		LendingRate:   liqRate,
		BorrowingRate: stableBorrowRate,
		Asset:         erc20.Symbol(),
		Protocol:      s.protocol.Name,
	}

	select {
	case <-ctx.Done():
		return errors.Wrapf(err, "context closed while processing the ReserveDataUpdate event: %s", ctx.Err())

	case s.chMsg <- msg:
		s.log.WithFields(logrus.Fields{
			"rate_timestamp": msg.Timestamp.Format(time.RFC3339),
			"lending_rate":   fmt.Sprintf("%.4f", msg.LendingRate),
			"borrowing_rate": fmt.Sprintf("%.4f", msg.BorrowingRate),
			"asset":          msg.Asset,
		}).Trace("a new DefiRate message has been sent to the update channel")

		return nil
	}
}
