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

var testPeerInfo = PeerInfo{
	PeerID:          "abc12345",
	Roles:           "some roles",
	ProtocolVersion: 123,
	BestHash:        NewHash([]byte{0xab, 0xcd}),
	BestNumber:      1337,
}

func TestPeerInfo_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testPeerInfo)
	AssertRoundTripFuzz[PeerInfo](t, 100)
	AssertEncodeEmptyObj[PeerInfo](t, 42)
}

func TestPeerInfo_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testPeerInfo,
			MustHexDecodeString("0x20616263313233343528736f6d6520726f6c65737b000000abcd00000000000000000000000000000000000000000000000000000000000039050000"),
		},
	})
}

func TestPeerInfo_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x20616263313233343528736f6d6520726f6c65737b000000abcd00000000000000000000000000000000000000000000000000000000000039050000"),
			testPeerInfo,
		},
	})
}
