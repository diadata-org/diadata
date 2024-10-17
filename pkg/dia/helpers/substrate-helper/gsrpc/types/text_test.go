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

func TestString_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewText(""))
	AssertRoundtrip(t, NewText("My nice string"))
	AssertRoundTripFuzz[Text](t, 100)
	AssertEncodeEmptyObj[Text](t, 1)
}

func TestString_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewText(""), 1},
		{NewText("My nice string"), 15},
		{NewText("ښ"), 3},
	})
}

func TestString_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewText(""), MustHexDecodeString("0x00")},
		{NewText("My nice string"), MustHexDecodeString("0x384d79206e69636520737472696e67")},
		{NewText("ښ"), MustHexDecodeString("0x08da9a")},
	})
}

func TestString_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewText(""), MustHexDecodeString("0x03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c111314")},
		{NewText("My nice string"), MustHexDecodeString(
			"0x30f31ff5f82596c990e5afd2b656db0ca30e1a5a9cc7eb5f83ee18731eecd453")},
	})
}

func TestString_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewText(""), "0x00"},
		{NewText("My nice string"), "0x384d79206e69636520737472696e67"},
		{NewText("ښ"), "0x08da9a"},
	})
}

func TestString_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewText(""), ""},
		{NewText("My nice string"), "My nice string"},
		{NewText("ښ"), "ښ"},
	})
}

func TestString_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewText("My nice string"), NewText("My nice string"), true},
		{NewText(""), NewText("23"), false},
		{NewText("My nice string"), NewU8(23), false},
		{NewText("My nice string"), NewBool(false), false},
	})
}
