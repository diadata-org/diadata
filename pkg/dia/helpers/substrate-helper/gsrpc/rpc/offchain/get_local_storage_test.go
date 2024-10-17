// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2022 Snowfork
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

package offchain

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffchain_LocalStorageGetSet(t *testing.T) {
	key := make([]byte, 20)
	n, err := rand.Read(key)
	assert.NoError(t, err)
	assert.Equal(t, 20, n)

	value := []byte{0, 1, 2}

	data, err := testOffchain.LocalStorageGet(Persistent, key)
	assert.NoError(t, err)
	assert.Empty(t, data)

	err = testOffchain.LocalStorageSet(Persistent, key, value)
	assert.NoError(t, err)

	data, err = testOffchain.LocalStorageGet(Persistent, key)
	assert.NoError(t, err)

	assert.Equal(t, value, []byte(*data))
}
