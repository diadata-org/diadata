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

func TestUSize_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, USize(0))
	AssertRoundtrip(t, USize(12))
}

func TestUsize_JSONMarshalUnmarshal(t *testing.T) {
	u := USize(11)

	AssertJSONRoundTrip(t, &u)
}

func TestUSize_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{USize(13), 4}})
}

func TestUSize_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{USize(29), MustHexDecodeString("0x1d000000")},
	})
}

func TestUSize_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{USize(29), MustHexDecodeString("0x60ebb66f09bc7fdd21772ab1ed0efb1fd1208e3f5cd20d2d9a29a2a79b6f953f")},
	})
}

func TestUSize_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{USize(29), "0x1d000000"},
	})
}

func TestUSize_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{USize(29), "29"},
	})
}

func TestUSize_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{USize(23), USize(23), true},
		{USize(23), NewBool(false), false},
	})
}
