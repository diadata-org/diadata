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

func TestHealth_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, Health{3, false, true})
	AssertRoundtrip(t, Health{1, true, true})
	AssertRoundtrip(t, Health{0, true, false})
	AssertRoundTripFuzz[Health](t, 100)
	AssertDecodeNilData[Health](t)
	AssertEncodeEmptyObj[Health](t, 10)
}

func TestHealth_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{Health{3, false, true}, MustHexDecodeString("0x03000000000000000001")},
		{Health{1, true, true}, MustHexDecodeString("0x01000000000000000101")},
		{Health{0, true, false}, MustHexDecodeString("0x00000000000000000100")},
	})
}

func TestHealth_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x03000000000000000001"), Health{3, false, true}},
		{MustHexDecodeString("0x01000000000000000101"), Health{1, true, true}},
		{MustHexDecodeString("0x00000000000000000100"), Health{0, true, false}},
	})
}
