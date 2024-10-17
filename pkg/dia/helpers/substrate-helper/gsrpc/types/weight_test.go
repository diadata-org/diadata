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

var (
	testWeight = NewWeight(NewUCompactFromUInt(11), NewUCompactFromUInt(634))
)

func TestWeight_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Weight](t, 100)
	AssertEncodeEmptyObj[Weight](t, 2)
}

func TestWeight_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{testWeight, 3}})
}

func TestWeight_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testWeight, MustHexDecodeString("0x2ce909")},
	})
}

func TestWeight_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x2ce909"), testWeight},
	})
}

func TestWeight_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{testWeight, MustHexDecodeString("0x7daf57922bb9694b4e29da7634e1b0a6af1477a8d13b0544208cda78331ea135")},
	})
}

func TestWeight_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{testWeight, "0x2ce909"},
	})
}

func TestWeight_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{testWeight, NewWeight(NewUCompactFromUInt(11), NewUCompactFromUInt(634)), true},
		{testWeight, NewBool(false), false},
	})
}
