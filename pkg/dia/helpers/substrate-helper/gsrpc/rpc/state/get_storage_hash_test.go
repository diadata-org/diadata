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

package state

import (
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestState_GetStorageHashLatest(t *testing.T) {
	key := types.NewStorageKey(codec.MustHexDecodeString("0x3a636f6465"))
	hash, err := testState.GetStorageHashLatest(key)
	assert.NoError(t, err)
	var expected types.Hash
	copy(expected[:], codec.MustHexDecodeString(mockSrv.storageHashHex))
	assert.Equal(t, expected, hash)
}

func TestState_GetStorageHash(t *testing.T) {
	key := types.NewStorageKey(codec.MustHexDecodeString("0x3a636f6465"))
	hash, err := testState.GetStorageHash(key, mockSrv.blockHashLatest)
	assert.NoError(t, err)
	var expected types.Hash
	copy(expected[:], codec.MustHexDecodeString(mockSrv.storageHashHex))
	assert.Equal(t, expected, hash)
}
