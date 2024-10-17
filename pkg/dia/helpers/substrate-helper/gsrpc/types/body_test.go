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
	bodyID1 = BodyID{
		IsUnit: true,
	}
	bodyID2 = BodyID{
		IsNamed: true,
		Body:    []U8{4},
	}
	bodyID3 = BodyID{
		IsIndex: true,
		Index:   6,
	}
	bodyID4 = BodyID{
		IsExecutive: true,
	}
	bodyID5 = BodyID{
		IsTechnical: true,
	}
	bodyID6 = BodyID{
		IsLegislative: true,
	}
	bodyID7 = BodyID{
		IsJudicial: true,
	}

	bodyIDFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(b *BodyID, c fuzz.Continue) {
			switch c.Intn(7) {
			case 0:
				b.IsUnit = true
			case 1:
				b.IsNamed = true
				c.Fuzz(&b.Body)
			case 2:
				b.IsIndex = true
				c.Fuzz(&b.Index)
			case 3:
				b.IsExecutive = true
			case 4:
				b.IsTechnical = true
			case 5:
				b.IsLegislative = true
			case 6:
				b.IsJudicial = true
			}
		}),
	}
)

func TestBodyID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[BodyID](t, 1000, bodyIDFuzzOpts...)
	AssertDecodeNilData[BodyID](t)
	AssertEncodeEmptyObj[BodyID](t, 0)
}

func TestBodyID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{bodyID1, MustHexDecodeString("0x00")},
		{bodyID2, MustHexDecodeString("0x010404")},
		{bodyID3, MustHexDecodeString("0x0206000000")},
		{bodyID4, MustHexDecodeString("0x03")},
		{bodyID5, MustHexDecodeString("0x04")},
		{bodyID6, MustHexDecodeString("0x05")},
		{bodyID7, MustHexDecodeString("0x06")},
	})
}

func TestBodyID_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), bodyID1},
		{MustHexDecodeString("0x010404"), bodyID2},
		{MustHexDecodeString("0x0206000000"), bodyID3},
		{MustHexDecodeString("0x03"), bodyID4},
		{MustHexDecodeString("0x04"), bodyID5},
		{MustHexDecodeString("0x05"), bodyID6},
		{MustHexDecodeString("0x06"), bodyID7},
	})
}

var (
	bodyPart1 = BodyPart{
		IsVoice: true,
	}
	bodyPart2 = BodyPart{
		IsMembers:    true,
		MembersCount: 3,
	}
	bodyPart3 = BodyPart{
		IsFraction:    true,
		FractionNom:   1,
		FractionDenom: 2,
	}
	bodyPart4 = BodyPart{
		IsAtLeastProportion:    true,
		AtLeastProportionNom:   3,
		AtLeastProportionDenom: 4,
	}
	bodyPart5 = BodyPart{
		IsMoreThanProportion:    true,
		MoreThanProportionNom:   6,
		MoreThanProportionDenom: 7,
	}

	bodyPartFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(b *BodyPart, c fuzz.Continue) {
			switch c.Intn(5) {
			case 0:
				b.IsVoice = true
			case 1:
				b.IsMembers = true
				c.Fuzz(&b.MembersCount)
			case 2:
				b.IsFraction = true
				c.Fuzz(&b.FractionNom)
				c.Fuzz(&b.FractionDenom)
			case 3:
				b.IsAtLeastProportion = true
				c.Fuzz(&b.AtLeastProportionNom)
				c.Fuzz(&b.AtLeastProportionDenom)
			case 4:
				b.IsMoreThanProportion = true
				c.Fuzz(&b.MoreThanProportionNom)
				c.Fuzz(&b.MoreThanProportionDenom)
			}
		}),
	}
)

func TestBodyPart_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[BodyPart](t, 1000, bodyPartFuzzOpts...)
	AssertDecodeNilData[BodyPart](t)
	AssertEncodeEmptyObj[BodyPart](t, 0)
}

func TestBodyPart_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{bodyPart1, MustHexDecodeString("0x00")},
		{bodyPart2, MustHexDecodeString("0x0103000000")},
		{bodyPart3, MustHexDecodeString("0x020100000002000000")},
		{bodyPart4, MustHexDecodeString("0x030300000004000000")},
		{bodyPart5, MustHexDecodeString("0x040600000007000000")},
	})
}

func TestBodyPart_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), bodyPart1},
		{MustHexDecodeString("0x0103000000"), bodyPart2},
		{MustHexDecodeString("0x020100000002000000"), bodyPart3},
		{MustHexDecodeString("0x030300000004000000"), bodyPart4},
		{MustHexDecodeString("0x040600000007000000"), bodyPart5},
	})
}
