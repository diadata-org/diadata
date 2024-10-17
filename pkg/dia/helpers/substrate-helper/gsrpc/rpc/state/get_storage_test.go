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

func TestState_GetStorageLatest(t *testing.T) {
	var decoded types.U64
	ok, err := testState.GetStorageLatest(codec.MustHexDecodeString(mockSrv.storageKeyHex), &decoded)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, types.U64(0x5d892db8), decoded)
}

func TestState_GetStorage(t *testing.T) {
	var decoded types.U64
	ok, err := testState.GetStorage(codec.MustHexDecodeString(mockSrv.storageKeyHex), &decoded, mockSrv.blockHashLatest)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, types.U64(0x5d892db8), decoded)
}

func TestState_GetStorageEmpty(t *testing.T) {
	var decoded types.U64
	ok, err := testState.GetStorage([]byte{0xab}, &decoded, mockSrv.blockHashLatest)
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestState_GetStorageRawLatest(t *testing.T) {
	data, err := testState.GetStorageRawLatest(codec.MustHexDecodeString(mockSrv.storageKeyHex))
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.storageDataHex, data.Hex())
}

func TestState_GetStorageRaw(t *testing.T) {
	data, err := testState.GetStorageRaw(codec.MustHexDecodeString(mockSrv.storageKeyHex), mockSrv.blockHashLatest)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.storageDataHex, data.Hex())
}
