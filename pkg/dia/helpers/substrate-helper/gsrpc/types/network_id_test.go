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
	fuzz "github.com/google/gofuzz"
)

var (
	testNetworkID1 = NetworkID{
		IsAny: true,
	}
	testNetworkID2 = NetworkID{
		IsNamed:      true,
		NamedNetwork: []U8{1, 2, 3, 4},
	}
	testNetworkID3 = NetworkID{
		IsPolkadot: true,
	}
	testNetworkID4 = NetworkID{
		IsKusama: true,
	}

	networkIDFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(n *NetworkID, c fuzz.Continue) {
			switch c.Intn(4) {
			case 0:
				n.IsAny = true
			case 1:
				n.IsNamed = true

				c.Fuzz(&n.NamedNetwork)
			case 2:
				n.IsPolkadot = true
			case 3:
				n.IsKusama = true
			}
		}),
	}
)

func TestNetworkID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[NetworkID](t, 100, networkIDFuzzOpts...)
	AssertDecodeNilData[NetworkID](t)
	AssertEncodeEmptyObj[NetworkID](t, 0)
}

func TestNetworkID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testNetworkID1, MustHexDecodeString("0x00")},
		{testNetworkID2, MustHexDecodeString("0x011001020304")},
		{testNetworkID3, MustHexDecodeString("0x02")},
		{testNetworkID4, MustHexDecodeString("0x03")},
	})
}

func TestNetworkID_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testNetworkID1},
		{MustHexDecodeString("0x011001020304"), testNetworkID2},
		{MustHexDecodeString("0x02"), testNetworkID3},
		{MustHexDecodeString("0x03"), testNetworkID4},
	})
}
