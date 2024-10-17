//go:build live

package retriever

import (
	"log"
	"os"
	"strconv"
	"sync"
	"testing"

	gsrpc "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc"
)

var (
	extrinsicTestURLs = []string{
		"wss://fullnode.parachain.centrifuge.io",
		"wss://rpc.polkadot.io",
		"wss://acala-rpc-0.aca-api.network",
		"wss://wss.api.moonbeam.network",
	}
)

func TestLive_ExtrinsicRetriever_GetExtrinsics(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup

	extrinsicsThreshold, err := strconv.Atoi(os.Getenv("GSRPC_LIVE_TEST_EXTRINSICS_THRESHOLD"))
	if err != nil {
		extrinsicsThreshold = 50
		log.Printf("Env Var GSRPC_LIVE_TEST_EXTRINSICS_THRESHOLD not set, defaulting to %d", extrinsicsThreshold)
	}

	for _, testURL := range extrinsicTestURLs {
		testURL := testURL

		wg.Add(1)

		go func() {
			defer wg.Done()

			api, err := gsrpc.NewSubstrateAPI(testURL)

			if err != nil {
				log.Printf("Couldn't connect to '%s': %s\n", testURL, err)
				return
			}

			retriever, err := NewDefaultExtrinsicRetriever(api.RPC.Chain, api.RPC.State)

			if err != nil {
				log.Printf("Couldn't create extrinsic retriever: %s", err)
				return
			}

			header, err := api.RPC.Chain.GetHeaderLatest()

			if err != nil {
				log.Printf("Couldn't get latest header for '%s': %s\n", testURL, err)
				return
			}

			extrinsicCount := 0

			for {
				blockHash, err := api.RPC.Chain.GetBlockHash(uint64(header.Number))

				if err != nil {
					log.Printf("Couldn't retrieve blockHash for '%s', block number %d: %s\n", testURL, header.Number, err)
					return
				}

				extrinsics, err := retriever.GetExtrinsics(blockHash)

				if err != nil {
					log.Printf("Couldn't retrieve extrinsics for '%s', block number %d: %s\n", testURL, header.Number, err)
					return
				}

				log.Printf("Found %d extrinsics for '%s', at block number %d.\n", len(extrinsics), testURL, header.Number)

				extrinsicCount += len(extrinsics)

				if extrinsicCount > extrinsicsThreshold {
					log.Printf("Retrieved a total of %d extrinsics for '%s', last block number %d. Stopping now.\n", extrinsicCount, testURL, header.Number)

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
