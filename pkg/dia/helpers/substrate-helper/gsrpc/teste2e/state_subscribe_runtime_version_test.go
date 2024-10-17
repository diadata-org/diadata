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

	"github.com/stretchr/testify/assert"
)

func TestState_SubscribeRuntimeVersion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode.")
	}
	api := subscriptionsAPI

	sub, err := api.RPC.State.SubscribeRuntimeVersion()
	assert.NoError(t, err)
	defer sub.Unsubscribe()

	timeout := time.After(10 * time.Second)

	for {
		select {
		case rv := <-sub.Chan():
			fmt.Printf("%#v\n", rv)
			return
		case <-timeout:
			assert.FailNow(t, "timeout reached without getting a notification from subscription")
			return
		}
	}
}
