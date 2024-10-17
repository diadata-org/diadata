//go:build live

package retriever

import (
	"log"
	"os"
	"strconv"
	"sync"
	"testing"

	gsrpc "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/state"
)

var (
	eventTestURLs = []string{
		"wss://fullnode.parachain.centrifuge.io",
		"wss://rpc.polkadot.io",
		"wss://acala-rpc-0.aca-api.network",
		"wss://wss.api.moonbeam.network",
	}
)

func TestLive_EventRetriever_GetEvents(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup

	eventsThreshold, err := strconv.Atoi(os.Getenv("GSRPC_LIVE_TEST_EVENTS_THRESHOLD"))
	if err != nil {
		eventsThreshold = 50
		log.Printf("Env Var GSRPC_LIVE_TEST_EVENTS_THRESHOLD not set, defaulting to %d", eventsThreshold)
	}

	for _, testURL := range eventTestURLs {
		testURL := testURL

		wg.Add(1)

		go func() {
			defer wg.Done()

			api, err := gsrpc.NewSubstrateAPI(testURL)

			if err != nil {
				log.Printf("Couldn't connect to '%s': %s\n", testURL, err)
				return
			}

			retriever, err := NewDefaultEventRetriever(state.NewEventProvider(api.RPC.State), api.RPC.State)

			if err != nil {
				log.Printf("Couldn't create event retriever: %s", err)
				return
			}

			header, err := api.RPC.Chain.GetHeaderLatest()

			if err != nil {
				log.Printf("Couldn't get latest header for '%s': %s\n", testURL, err)
				return
			}

			eventsCount := 0

			for {
				blockHash, err := api.RPC.Chain.GetBlockHash(uint64(header.Number))

				if err != nil {
					log.Printf("Couldn't retrieve blockHash for '%s', block number %d: %s\n", testURL, header.Number, err)
					return
				}

				events, err := retriever.GetEvents(blockHash)

				if err != nil {
					log.Printf("Couldn't retrieve events for '%s', block number %d: %s\n", testURL, header.Number, err)
					return
				}

				log.Printf("Found %d events for '%s', at block number %d.\n", len(events), testURL, header.Number)

				eventsCount += len(events)

				if eventsCount > eventsThreshold {
					log.Printf("Retrieved a total of %d events for '%s', last block number %d. Stopping now.\n", eventsCount, testURL, header.Number)

					return
				}

				header, err = api.RPC.Chain.GetHeader(header.ParentHash)

				if err != nil {
					log.Printf("Couldn't retrieve header for block number '%d' for '%s': %s\n", header.Number, testURL, err)

					return
				}
			}
		}()
	}

	wg.Wait()
}
