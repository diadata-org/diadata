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
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestStorageDataRaw_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewStorageDataRaw([]byte{12, 251, 42}), 3},
		{NewStorageDataRaw([]byte{}), 0},
	})
}

func TestStorageDataRaw_Encode(t *testing.T) {
	bz := []byte{12, 251, 42}
	dataRaw := NewStorageDataRaw(bz)
	encoded, err := Encode(dataRaw)
	assert.NoError(t, err)
	assert.Equal(t, bz, encoded)
}

func TestStorageDataRaw_Decode(t *testing.T) {
	bz := []byte{12, 251, 42}
	decoded := make(StorageDataRaw, len(bz))
	err := Decode(bz, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, StorageDataRaw(bz), decoded)
}

func TestStorageDataRaw_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewStorageDataRaw([]byte{0, 42, 254}), MustHexDecodeString(
			"0x537db36f5b5970b679a28a3df8d219317d658014fb9c3d409c0c799d8ecf149d")},
		{NewStorageDataRaw([]byte{}), MustHexDecodeString(
			"0x0e5751c026e543b2e8ab2eb06099daa1d1e5df47778f7787faab45cdf12fe3a8")},
	})
}

func TestStorageDataRaw_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewStorageDataRaw([]byte{0, 0, 0}), "0x000000"},
		{NewStorageDataRaw([]byte{171, 18, 52}), "0xab1234"},
		{NewStorageDataRaw([]byte{0, 1}), "0x0001"},
		{NewStorageDataRaw([]byte{18, 52, 86}), "0x123456"},
	})
}

func TestStorageDataRaw_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewStorageDataRaw([]byte{0, 0, 0}), "[0 0 0]"},
		{NewStorageDataRaw([]byte{171, 18, 52}), "[171 18 52]"},
		{NewStorageDataRaw([]byte{0, 1}), "[0 1]"},
		{NewStorageDataRaw([]byte{18, 52, 86}), "[18 52 86]"},
	})
}

func TestStorageDataRaw_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewStorageDataRaw([]byte{1, 0, 0}), NewStorageDataRaw([]byte{1, 0}), false},
		{NewStorageDataRaw([]byte{0, 0, 1}), NewStorageDataRaw([]byte{0, 1}), false},
		{NewStorageDataRaw([]byte{0, 0, 0}), NewStorageDataRaw([]byte{0, 0}), false},
		{NewStorageDataRaw([]byte{12, 48, 255}), NewStorageDataRaw([]byte{12, 48, 255}), true},
		{NewStorageDataRaw([]byte{0}), NewStorageDataRaw([]byte{0}), true},
		{NewStorageDataRaw([]byte{1}), NewBool(true), false},
		{NewStorageDataRaw([]byte{0}), NewBool(false), false},
	})
}
