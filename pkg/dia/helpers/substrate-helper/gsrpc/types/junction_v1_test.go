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
	"math/big"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
)

var (
	testJunctionV1n1 = JunctionV1{
		IsParachain: true,
		ParachainID: NewUCompactFromUInt(11),
	}
	testJunctionV1n2 = JunctionV1{
		IsAccountID32: true,
		AccountID32NetworkID: NetworkID{
			IsAny: true,
		},
		AccountID: []U8{1, 2, 3},
	}
	testJunctionV1n3 = JunctionV1{
		IsAccountIndex64: true,
		AccountIndex64NetworkID: NetworkID{
			IsAny: true,
		},
		AccountIndex: 16,
	}
	testJunctionV1n4 = JunctionV1{
		IsAccountKey20: true,
		AccountKey20NetworkID: NetworkID{
			IsKusama: true,
		},
	}
	testJunctionV1n5 = JunctionV1{
		IsPalletInstance: true,
		PalletIndex:      4,
	}
	testJunctionV1n6 = JunctionV1{
		IsGeneralIndex: true,
		GeneralIndex:   NewU128(*big.NewInt(42)),
	}
	testJunctionV1n7 = JunctionV1{
		IsGeneralKey: true,
		GeneralKey:   []U8{6, 8},
	}
	testJunctionV1n8 = JunctionV1{
		IsOnlyChild: true,
	}
	testJunctionV1n9 = JunctionV1{
		IsPlurality: true,
		BodyID: BodyID{
			IsUnit: true,
		},
		BodyPart: BodyPart{
			IsVoice: true,
		},
	}

	junctionV1FuzzOpts = CombineFuzzOpts(
		networkIDFuzzOpts,
		bodyIDFuzzOpts,
		bodyPartFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(j *JunctionV1, c fuzz.Continue) {
				switch c.Intn(9) {
				case 0:
					j.IsParachain = true

					c.Fuzz(&j.ParachainID)
				case 1:
					j.IsAccountID32 = true

					c.Fuzz(&j.AccountID32NetworkID)

					c.Fuzz(&j.AccountID)
				case 2:
					j.IsAccountIndex64 = true

					c.Fuzz(&j.AccountIndex64NetworkID)

					c.Fuzz(&j.AccountIndex)
				case 3:
					j.IsAccountKey20 = true

					c.Fuzz(&j.AccountKey20NetworkID)

					c.Fuzz(&j.AccountKey)
				case 4:
					j.IsPalletInstance = true

					c.Fuzz(&j.PalletIndex)
				case 5:
					j.IsGeneralIndex = true

					c.Fuzz(&j.GeneralIndex)
				case 6:
					j.IsGeneralKey = true

					c.Fuzz(&j.GeneralKey)
				case 7:
					j.IsOnlyChild = true
				case 8:
					j.IsPlurality = true

					c.Fuzz(&j.BodyID)

					c.Fuzz(&j.BodyPart)
				}
			}),
		},
	)
)

func TestJunctionV1_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[JunctionV1](t, 1000, junctionV1FuzzOpts...)
	AssertDecodeNilData[JunctionV1](t)
	AssertEncodeEmptyObj[JunctionV1](t, 0)
}

func TestJunctionV1_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testJunctionV1n1, MustHexDecodeString("0x002c")},
		{testJunctionV1n2, MustHexDecodeString("0x01000c010203")},
		{testJunctionV1n3, MustHexDecodeString("0x02001000000000000000")},
		{testJunctionV1n4, MustHexDecodeString("0x030300")},
		{testJunctionV1n5, MustHexDecodeString("0x0404")},
		{testJunctionV1n6, MustHexDecodeString("0x052a000000000000000000000000000000")},
		{testJunctionV1n7, MustHexDecodeString("0x06080608")},
		{testJunctionV1n8, MustHexDecodeString("0x07")},
		{testJunctionV1n9, MustHexDecodeString("0x080000")},
	})
}

func TestJunctionV1_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x002c"), testJunctionV1n1},
		{MustHexDecodeString("0x01000c010203"), testJunctionV1n2},
		{MustHexDecodeString("0x02001000000000000000"), testJunctionV1n3},
		{MustHexDecodeString("0x030300"), testJunctionV1n4},
		{MustHexDecodeString("0x0404"), testJunctionV1n5},
		{MustHexDecodeString("0x052a000000000000000000000000000000"), testJunctionV1n6},
		{MustHexDecodeString("0x06080608"), testJunctionV1n7},
		{MustHexDecodeString("0x07"), testJunctionV1n8},
		{MustHexDecodeString("0x080000"), testJunctionV1n9},
	})
}

var (
	testJunctionsV1n1 = JunctionsV1{
		IsHere: true,
	}
	testJunctionsV1n2 = JunctionsV1{
		IsX1: true,
		X1:   testJunctionV1n1,
	}
	testJunctionsV1n3 = JunctionsV1{
		IsX2: true,
		X2: [2]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
		},
	}
	testJunctionsV1n4 = JunctionsV1{
		IsX3: true,
		X3: [3]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
		},
	}
	testJunctionsV1n5 = JunctionsV1{
		IsX4: true,
		X4: [4]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
			testJunctionV1n4,
		},
	}
	testJunctionsV1n6 = JunctionsV1{
		IsX5: true,
		X5: [5]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
			testJunctionV1n4,
			testJunctionV1n5,
		},
	}
	testJunctionsV1n7 = JunctionsV1{
		IsX6: true,
		X6: [6]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
			testJunctionV1n4,
			testJunctionV1n5,
			testJunctionV1n6,
		},
	}
	testJunctionsV1n8 = JunctionsV1{
		IsX7: true,
		X7: [7]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
			testJunctionV1n4,
			testJunctionV1n5,
			testJunctionV1n6,
			testJunctionV1n7,
		},
	}
	testJunctionsV1n9 = JunctionsV1{
		IsX8: true,
		X8: [8]JunctionV1{
			testJunctionV1n1,
			testJunctionV1n2,
			testJunctionV1n3,
			testJunctionV1n4,
			testJunctionV1n5,
			testJunctionV1n6,
			testJunctionV1n7,
			testJunctionV1n8,
		},
	}

	junctionsV1FuzzOpts = CombineFuzzOpts(
		junctionV1FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(j *JunctionsV1, c fuzz.Continue) {
				switch c.Intn(9) {
				case 0:
					j.IsHere = true
				case 1:
					j.IsX1 = true

					c.Fuzz(&j.X1)
				case 2:
					j.IsX2 = true

					c.Fuzz(&j.X2)
				case 3:
					j.IsX3 = true

					c.Fuzz(&j.X3)
				case 4:
					j.IsX4 = true

					c.Fuzz(&j.X4)
				case 5:
					j.IsX5 = true

					c.Fuzz(&j.X5)
				case 6:
					j.IsX6 = true

					c.Fuzz(&j.X6)
				case 7:
					j.IsX7 = true

					c.Fuzz(&j.X7)
				case 8:
					j.IsX8 = true

					c.Fuzz(&j.X8)
				}
			}),
		},
	)
)

func TestJunctionsV1_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[JunctionsV1](t, 1000, junctionsV1FuzzOpts...)
	AssertDecodeNilData[JunctionsV1](t)
	AssertEncodeEmptyObj[JunctionsV1](t, 0)
}

func TestJunctionsV1_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testJunctionsV1n1,
			MustHexDecodeString("0x00"),
		},
		{
			testJunctionsV1n2,
			MustHexDecodeString("0x01002c"),
		},
		{
			testJunctionsV1n3,
			MustHexDecodeString("0x02002c01000c010203"),
		},
		{
			testJunctionsV1n4,
			MustHexDecodeString("0x03002c01000c01020302001000000000000000"),
		},
		{
			testJunctionsV1n5,
			MustHexDecodeString("0x04002c01000c01020302001000000000000000030300"),
		},
		{
			testJunctionsV1n6,
			MustHexDecodeString("0x05002c01000c010203020010000000000000000303000404"),
		},
		{
			testJunctionsV1n7,
			MustHexDecodeString("0x06002c01000c010203020010000000000000000303000404052a000000000000000000000000000000"),
		},
		{
			testJunctionsV1n8,
			MustHexDecodeString("0x07002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608"),
		},
		{
			testJunctionsV1n9,
			MustHexDecodeString("0x08002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"),
		},
	})
}

func TestJunctionsV1_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x00"),
			testJunctionsV1n1,
		},
		{
			MustHexDecodeString("0x01002c"),
			testJunctionsV1n2,
		},
		{
			MustHexDecodeString("0x02002c01000c010203"),
			testJunctionsV1n3,
		},
		{
			MustHexDecodeString("0x03002c01000c01020302001000000000000000"),
			testJunctionsV1n4,
		},
		{
			MustHexDecodeString("0x04002c01000c01020302001000000000000000030300"),
			testJunctionsV1n5,
		},
		{
			MustHexDecodeString("0x05002c01000c010203020010000000000000000303000404"),
			testJunctionsV1n6,
		},
		{
			MustHexDecodeString("0x06002c01000c010203020010000000000000000303000404052a000000000000000000000000000000"),
			testJunctionsV1n7,
		},
		{
			MustHexDecodeString("0x07002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608"),
			testJunctionsV1n8,
		},
		{
			MustHexDecodeString("0x08002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"),
			testJunctionsV1n9,
		},
	})
}
