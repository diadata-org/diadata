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

package state

import (
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestState_QueryStorageAtLatest(t *testing.T) {
	key := types.NewStorageKey(codec.MustHexDecodeString(mockSrv.storageKeyHex))
	data, err := testState.QueryStorageAtLatest([]types.StorageKey{key})
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.storageChangeSets, data)
}

func TestState_QueryStorageAt(t *testing.T) {
	key := types.NewStorageKey(codec.MustHexDecodeString(mockSrv.storageKeyHex))
	hash := types.NewHash(codec.MustHexDecodeString("0xdd1816b6f6889f46e23b0d6750bc441af9dad0fda8bae90677c1708d01035fbe"))
	data, err := testState.QueryStorageAt([]types.StorageKey{key}, hash)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.storageChangeSets, data)
}
