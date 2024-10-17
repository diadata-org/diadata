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
	testJunction1 = JunctionV0{
		IsParent: true,
	}
	testJunction2 = JunctionV0{
		IsParachain: true,
		ParachainID: 11,
	}
	testJunction3 = JunctionV0{
		IsAccountID32: true,
		AccountID32NetworkID: NetworkID{
			IsAny: true,
		},
		AccountID: []U8{1, 2, 3},
	}
	testJunction4 = JunctionV0{
		IsAccountIndex64: true,
		AccountIndex64NetworkID: NetworkID{
			IsAny: true,
		},
		AccountIndex: 16,
	}
	testJunction5 = JunctionV0{
		IsAccountKey20: true,
		AccountKey20NetworkID: NetworkID{
			IsKusama: true,
		},
	}
	testJunction6 = JunctionV0{
		IsPalletInstance: true,
		PalletIndex:      4,
	}
	testJunction7 = JunctionV0{
		IsGeneralIndex: true,
		GeneralIndex:   NewU128(*big.NewInt(42)),
	}
	testJunction8 = JunctionV0{
		IsGeneralKey: true,
		GeneralKey:   []U8{6, 8},
	}
	testJunction9 = JunctionV0{
		IsOnlyChild: true,
	}
	testJunction10 = JunctionV0{
		IsPlurality: true,
		PluralityID: BodyID{
			IsUnit: true,
		},
		PluralityPart: BodyPart{
			IsVoice: true,
		},
	}

	junctionV0FuzzOpts = CombineFuzzOpts(
		networkIDFuzzOpts,
		bodyIDFuzzOpts,
		bodyPartFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(j *JunctionV0, c fuzz.Continue) {
				switch c.Intn(10) {
				case 0:
					j.IsParent = true
				case 1:
					j.IsParachain = true

					c.Fuzz(&j.ParachainID)
				case 2:
					j.IsAccountID32 = true

					c.Fuzz(&j.AccountID32NetworkID)

					c.Fuzz(&j.AccountID)
				case 3:
					j.IsAccountIndex64 = true

					c.Fuzz(&j.AccountIndex64NetworkID)

					c.Fuzz(&j.AccountIndex)
				case 4:
					j.IsAccountKey20 = true

					c.Fuzz(&j.AccountKey20NetworkID)

					c.Fuzz(&j.AccountKey)
				case 5:
					j.IsPalletInstance = true

					c.Fuzz(&j.PalletIndex)
				case 6:
					j.IsGeneralIndex = true

					c.Fuzz(&j.GeneralIndex)
				case 7:
					j.IsGeneralKey = true

					c.Fuzz(&j.GeneralKey)
				case 8:
					j.IsOnlyChild = true
				case 9:
					j.IsPlurality = true

					c.Fuzz(&j.PluralityID)

					c.Fuzz(&j.PluralityPart)
				}
			}),
		},
	)
)

func TestJunctionV0_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[JunctionV0](t, 1000, junctionV0FuzzOpts...)
	AssertDecodeNilData[JunctionV0](t)
	AssertEncodeEmptyObj[JunctionV0](t, 0)
}

func TestJunctionV0_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testJunction1, MustHexDecodeString("0x00")},
		{testJunction2, MustHexDecodeString("0x010b000000")},
		{testJunction3, MustHexDecodeString("0x02000c010203")},
		{testJunction4, MustHexDecodeString("0x03001000000000000000")},
		{testJunction5, MustHexDecodeString("0x040300")},
		{testJunction6, MustHexDecodeString("0x0504")},
		{testJunction7, MustHexDecodeString("0x062a000000000000000000000000000000")},
		{testJunction8, MustHexDecodeString("0x07080608")},
		{testJunction9, MustHexDecodeString("0x08")},
		{testJunction10, MustHexDecodeString("0x090000")},
	})
}

func TestJunctionV0_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testJunction1},
		{MustHexDecodeString("0x010b000000"), testJunction2},
		{MustHexDecodeString("0x02000c010203"), testJunction3},
		{MustHexDecodeString("0x03001000000000000000"), testJunction4},
		{MustHexDecodeString("0x040300"), testJunction5},
		{MustHexDecodeString("0x0504"), testJunction6},
		{MustHexDecodeString("0x062a000000000000000000000000000000"), testJunction7},
		{MustHexDecodeString("0x07080608"), testJunction8},
		{MustHexDecodeString("0x08"), testJunction9},
		{MustHexDecodeString("0x090000"), testJunction10},
	})
}
