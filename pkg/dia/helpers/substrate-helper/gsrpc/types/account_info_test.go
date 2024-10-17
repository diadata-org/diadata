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

func TestAccountInfoV4_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AccountInfoV4](t, 100)
	AssertEncodeEmptyObj[AccountInfoV4](t, 9)
}

func TestAccountInfoV4_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), 12},
	})
}

func TestAccountInfoV4_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), MustHexDecodeString("0x0c0102030d00000000000000")},
	})
}

func TestAccountInfoV4_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), MustHexDecodeString(
			"0x4fac0dfeb9b4efd2518c762e7d097fafaffaf8d56a2e784f9fc9919c22277804")},
	})
}

func TestAccountInfoV4_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), "0x0c0102030d00000000000000"},
	})
}

func TestAccountInfoV4_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), "{[1 2 3] 13}"},
	})
}

func TestAccountInfoV4_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), NewAccountInfoV4([]byte{1, 2, 3}, 13), true},
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), NewAccountInfoV4([]byte{1, 2, 2}, 13), false},
		{NewAccountInfoV4([]byte{1, 2, 3}, 13), NewBool(false), false},
	})
}
