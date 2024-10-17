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

func TestType_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, Type(""))
	AssertRoundtrip(t, Type("Custom Type"))
	AssertRoundTripFuzz[Type](t, 100)
	AssertEncodeEmptyObj[Type](t, 1)
}

func TestType_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{Type(""), 1},
		{Type("Custom Type"), 12},
	})
}

func TestType_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{Type(""), MustHexDecodeString("0x00")},
		{Type("Custom Type"), MustHexDecodeString("0x2c437573746f6d2054797065")},
	})
}

func TestType_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{Type(""), MustHexDecodeString("0x03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c111314")},
		{Type("Custom Type"), MustHexDecodeString(
			"0x7c2996c160a91b8479eae96ade3ad8b287edc22997bf209b7bc0ca04d3bc0725")},
	})
}

func TestType_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{Type(""), "0x00"},
		{Type("Custom Type"), "0x2c437573746f6d2054797065"},
	})
}

func TestType_Type(t *testing.T) {
	AssertString(t, []StringAssert{
		{Type(""), ""},
		{Type("Custom Type"), "Custom Type"},
	})
}

func TestType_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{Type("Custom Type"), Type("Custom Type"), true},
		{Type(""), Type("23"), false},
		{Type("Custom Type"), NewU8(23), false},
		{Type("Custom Type"), NewBool(false), false},
	})
}
