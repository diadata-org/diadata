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

var childStorageKey = types.NewStorageKey(codec.MustHexDecodeString(mockSrv.childStorageKeyHex))
var key = types.NewStorageKey(codec.MustHexDecodeString(mockSrv.childStorageTrieKeyHex))

func TestState_GetChildStorageLatest(t *testing.T) {
	var decoded ChildStorageTrieTestVal
	ok, err := testState.GetChildStorageLatest(childStorageKey, key, &decoded)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.childStorageTrieValue, decoded)
}

func TestState_GetChildStorage(t *testing.T) {
	var decoded ChildStorageTrieTestVal
	ok, err := testState.GetChildStorageLatest(childStorageKey, key, &decoded)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.childStorageTrieValue, decoded, mockSrv.blockHashLatest)
}

func TestState_GetChildStorageRawLatest(t *testing.T) {
	data, err := testState.GetChildStorageRawLatest(childStorageKey, key)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.childStorageTrieValueHex, data.Hex())
}

func TestState_GetChildStorageRaw(t *testing.T) {
	data, err := testState.GetChildStorageRaw(childStorageKey, key, mockSrv.blockHashLatest)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.childStorageTrieValueHex, data.Hex())
}
