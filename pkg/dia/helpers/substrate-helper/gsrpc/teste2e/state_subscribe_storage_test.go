// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package teste2e

import (
	"fmt"
	"testing"
	"time"

	gsrpc "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/config"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestState_SubscribeStorage_EventsRaw(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode.")
	}

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	assert.NoError(t, err)

	meta, err := api.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	key, err := types.CreateStorageKey(meta, "System", "Events", nil)
	assert.NoError(t, err)

	sub, err := api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	assert.NoError(t, err)
	defer sub.Unsubscribe()

	timeout := time.After(10 * time.Second)
	received := 0

	for {
		select {
		case set := <-sub.Chan():
			fmt.Printf("%#v\n", set)
			received++

			if received >= 2 {
				return
			}
		case <-timeout:
			assert.FailNow(t, "timeout reached without getting 2 notifications from subscription")
			return
		}
	}
}

func TestState_SubscribeStorage_Events(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode.")
	}

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	assert.NoError(t, err)

	meta, err := api.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	key, err := types.CreateStorageKey(meta, "System", "Events", nil)
	assert.NoError(t, err)

	sub, err := api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	assert.NoError(t, err)
	defer sub.Unsubscribe()

	timeout := time.After(10 * time.Second)
	received := 0

	for {
		select {
		case set := <-sub.Chan():
			fmt.Printf("%#v\n", set)
			for _, chng := range set.Changes {
				if !codec.Eq(chng.StorageKey, key) || !chng.HasStorageData {
					// skip, we are only interested in events with content
					continue
				}

				fmt.Printf("%s\n", chng.StorageKey.Hex())
				fmt.Printf("%#x\n", chng.StorageData)

				events := types.EventRecords{}
				err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(meta, &events)
				assert.NoError(t, err)

				fmt.Printf("%#v\n", events)
			}

			received++

			if received >= 2 {
				return
			}
		case <-timeout:
			assert.FailNow(t, "timeout reached without getting 2 notifications from subscription")
			return
		}
	}
}
