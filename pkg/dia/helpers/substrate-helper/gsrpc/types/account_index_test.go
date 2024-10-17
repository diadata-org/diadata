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
)

func TestAccountIndex_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AccountIndex](t, 100)
	AssertDecodeNilData[AccountID](t)
	AssertEncodeEmptyObj[AccountID](t, 32)
}

func TestAccountIndex_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewAccountIndex(336794129), 4},
	})
}

func TestAccountIndex_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewAccountIndex(336794129), MustHexDecodeString("0x11121314")},
	})
}

func TestAccountIndex_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewAccountIndex(336794129), MustHexDecodeString(
			"0xa6730c0d3a95e0ff2068fa9a6ecf82c42c494c8c2cdd65379c898a4b88dd7138")},
	})
}

func TestAccountIndex_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewAccountIndex(336794129), "0x11121314"},
	})
}

func TestAccountIndex_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewAccountIndex(336794129), "336794129"},
	})
}

func TestAccountIndex_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewAccountIndex(336794129), NewAccountIndex(336794129), true},
		{NewAccountIndex(336794129), NewAccountIndex(12), false},
		{NewAccountIndex(336794129), NewBool(false), false},
	})
}
