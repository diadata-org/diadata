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
	testMultiLocationV0n1 = MultiLocationV0{
		IsNull: true,
	}
	testMultiLocationV0n2 = MultiLocationV0{
		IsX1: true,
		X1:   testJunction1,
	}
	testMultiLocationV0n3 = MultiLocationV0{
		IsX2: true,
		X2: [2]JunctionV0{
			testJunction1,
			testJunction2,
		},
	}
	testMultiLocationV0n4 = MultiLocationV0{
		IsX3: true,
		X3: [3]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
		},
	}
	testMultiLocationV0n5 = MultiLocationV0{
		IsX4: true,
		X4: [4]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
			testJunction4,
		},
	}
	testMultiLocationV0n6 = MultiLocationV0{
		IsX5: true,
		X5: [5]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
			testJunction4,
			testJunction5,
		},
	}
	testMultiLocationV0n7 = MultiLocationV0{
		IsX6: true,
		X6: [6]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
			testJunction4,
			testJunction5,
			testJunction6,
		},
	}
	testMultiLocationV0n8 = MultiLocationV0{
		IsX7: true,
		X7: [7]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
			testJunction4,
			testJunction5,
			testJunction6,
			testJunction7,
		},
	}
	testMultiLocationV0n9 = MultiLocationV0{
		IsX8: true,
		X8: [8]JunctionV0{
			testJunction1,
			testJunction2,
			testJunction3,
			testJunction4,
			testJunction5,
			testJunction6,
			testJunction7,
			testJunction8,
		},
	}
	multiLocationV0FuzzOpts = CombineFuzzOpts(
		junctionV0FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(m *MultiLocationV0, c fuzz.Continue) {
				switch c.Intn(9) {
				case 0:
					m.IsNull = true
				case 1:
					m.IsX1 = true

					c.Fuzz(&m.X1)
				case 2:
					m.IsX2 = true

					c.Fuzz(&m.X2)
				case 3:
					m.IsX3 = true

					c.Fuzz(&m.X3)
				case 4:
					m.IsX4 = true

					c.Fuzz(&m.X4)
				case 5:
					m.IsX5 = true

					c.Fuzz(&m.X5)
				case 6:
					m.IsX6 = true

					c.Fuzz(&m.X6)
				case 7:
					m.IsX7 = true

					c.Fuzz(&m.X7)
				case 8:
					m.IsX8 = true

					c.Fuzz(&m.X8)
				}
			}),
		},
	)
)

func TestMultiLocationV0_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiLocationV0](t, 1000, multiLocationV0FuzzOpts...)
	AssertDecodeNilData[MultiLocationV0](t)
	AssertEncodeEmptyObj[MultiLocationV0](t, 0)
}

func TestMultiLocationV0_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testMultiLocationV0n1,
			MustHexDecodeString("0x00"),
		},
		{
			testMultiLocationV0n2,
			MustHexDecodeString("0x0100"),
		},
		{
			testMultiLocationV0n3,
			MustHexDecodeString("0x0200010b000000"),
		},
		{
			testMultiLocationV0n4,
			MustHexDecodeString("0x0300010b00000002000c010203"),
		},
		{
			testMultiLocationV0n5,
			MustHexDecodeString("0x0400010b00000002000c01020303001000000000000000"),
		},
		{
			testMultiLocationV0n6,
			MustHexDecodeString("0x0500010b00000002000c01020303001000000000000000040300"),
		},
		{
			testMultiLocationV0n7,
			MustHexDecodeString("0x0600010b00000002000c010203030010000000000000000403000504"),
		},
		{
			testMultiLocationV0n8,
			MustHexDecodeString("0x0700010b00000002000c010203030010000000000000000403000504062a000000000000000000000000000000"),
		},
		{
			testMultiLocationV0n9,
			MustHexDecodeString("0x0800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608"),
		},
	})
}

func TestMultiLocationV0_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x00"),
			testMultiLocationV0n1,
		},
		{
			MustHexDecodeString("0x0100"),
			testMultiLocationV0n2,
		},
		{
			MustHexDecodeString("0x0200010b000000"),
			testMultiLocationV0n3,
		},
		{
			MustHexDecodeString("0x0300010b00000002000c010203"),
			testMultiLocationV0n4,
		},
		{
			MustHexDecodeString("0x0400010b00000002000c01020303001000000000000000"),
			testMultiLocationV0n5,
		},
		{
			MustHexDecodeString("0x0500010b00000002000c01020303001000000000000000040300"),
			testMultiLocationV0n6,
		},
		{
			MustHexDecodeString("0x0600010b00000002000c010203030010000000000000000403000504"),
			testMultiLocationV0n7,
		},
		{
			MustHexDecodeString("0x0700010b00000002000c010203030010000000000000000403000504062a000000000000000000000000000000"),
			testMultiLocationV0n8,
		},
		{
			MustHexDecodeString("0x0800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608"),
			testMultiLocationV0n9,
		},
	})
}
