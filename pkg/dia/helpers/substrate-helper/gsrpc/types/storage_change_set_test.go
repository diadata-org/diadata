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

package types_test

import (
	"encoding/json"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestStorageChangeSet_UnmarshalMarshalJSON(t *testing.T) {
	s := []byte("{\"block\":\"0xa230d0b6dc75868237b08d71618f3d19526b8aa346d94c792a4fce0a945b1e3f\",\"changes\":[[\"0xcc956bdb7605e3547539f321ac2bc95c\",\"0x0800000000000000000001000000000000\"],[\"0xcc956bdb7605e3547539f321ac2bc95d\"]]}") //nolint:lll

	var cs StorageChangeSet

	err := json.Unmarshal(s, &cs)
	assert.NoError(t, err)

	assert.Equal(t, StorageChangeSet{
		Block: NewHash([]byte{0xa2, 0x30, 0xd0, 0xb6, 0xdc, 0x75, 0x86, 0x82, 0x37, 0xb0, 0x8d, 0x71, 0x61, 0x8f, 0x3d, 0x19, 0x52, 0x6b, 0x8a, 0xa3, 0x46, 0xd9, 0x4c, 0x79, 0x2a, 0x4f, 0xce, 0xa, 0x94, 0x5b, 0x1e, 0x3f}), //nolint:lll
		Changes: []KeyValueOption{
			{
				StorageKey:     MustHexDecodeString("0xcc956bdb7605e3547539f321ac2bc95c"),
				HasStorageData: true,
				StorageData:    MustHexDecodeString("0x0800000000000000000001000000000000"),
			},
			{
				StorageKey:     MustHexDecodeString("0xcc956bdb7605e3547539f321ac2bc95d"),
				HasStorageData: false,
			},
		},
	}, cs)

	b, err := json.Marshal(cs)
	assert.NoError(t, err)
	assert.Equal(t, s, b)
}

func TestKeyValueOption_UnmarshalJSON(t *testing.T) {
	s := []byte("[\"0xcc956bdb7605e3547539f321ac2bc95c\",\"0x0800000000000000000001000000000000\"]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.NoError(t, err)

	assert.Equal(t, KeyValueOption{
		StorageKey:     MustHexDecodeString("0xcc956bdb7605e3547539f321ac2bc95c"),
		HasStorageData: true,
		StorageData:    MustHexDecodeString("0x0800000000000000000001000000000000"),
	}, kv)
}

func TestKeyValueOption_UnmarshalJSON2(t *testing.T) {
	s := []byte("[\"0xcc956bdb7605e3547539f321ac2bc95c\"]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.NoError(t, err)

	assert.Equal(t, KeyValueOption{
		StorageKey:     MustHexDecodeString("0xcc956bdb7605e3547539f321ac2bc95c"),
		HasStorageData: false,
	}, kv)
}

func TestKeyValueOption_UnmarshalMarshalJSON(t *testing.T) {
	s := []byte("[\"0xcc956bdb7605e3547539f321ac2bc95c\",\"0x0800000000000000000001000000000000\"]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.NoError(t, err)

	b, err := json.Marshal(kv)
	assert.NoError(t, err)
	assert.Equal(t, s, b)
}

func TestKeyValueOption_UnmarshalMarshalJSON2(t *testing.T) {
	s := []byte("[\"0xcc956bdb7605e3547539f321ac2bc95c\"]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.NoError(t, err)

	b, err := json.Marshal(kv)
	assert.NoError(t, err)
	assert.Equal(t, s, b)
}

func TestKeyValueOption_UnmarshalErrorEmpty(t *testing.T) {
	s := []byte("[]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.Errorf(t, err, "expected at least one entry for KeyValueOption")
}

func TestKeyValueOption_UnmarshalErrorTooMany(t *testing.T) {
	s := []byte("[\"0x01\", \"0x12\", \"0x23\"]")

	var kv KeyValueOption

	err := json.Unmarshal(s, &kv)
	assert.Errorf(t, err, "expected 1 or 2 entries for KeyValueOption, got 3")
}
