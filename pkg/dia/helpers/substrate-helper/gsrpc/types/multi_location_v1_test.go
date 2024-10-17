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
	"github.com/stretchr/testify/assert"
)

var (
	testMultiLocation = MultiLocationV1{
		Parents: 1,
		Interior: JunctionsV1{
			IsHere: true,
		},
	}

	optionMultiLocationV1FuzzOpts = CombineFuzzOpts(
		junctionsV1FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(o *OptionMultiLocationV1, c fuzz.Continue) {
				if c.RandBool() {
					*o = NewOptionMultiLocationV1Empty()
					return
				}

				var m MultiLocationV1

				c.Fuzz(&m)

				*o = NewOptionMultiLocationV1(m)
			}),
		},
	)
)

func TestOptionMultiLocation_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionMultiLocationV1](t, 1000, optionMultiLocationV1FuzzOpts...)
	AssertEncodeEmptyObj[OptionMultiLocationV1](t, 1)
}

func TestOptionMultiLocation_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewOptionMultiLocationV1(testMultiLocation), MustHexDecodeString("0x010100")},
		{NewOptionMultiLocationV1Empty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionMultiLocation_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x010100"), NewOptionMultiLocationV1(testMultiLocation)},
		{MustHexDecodeString("0x00"), NewOptionMultiLocationV1Empty()},
	})
}

func TestOptionMultiLocation_OptionMethods(t *testing.T) {
	o := NewOptionMultiLocationV1Empty()
	o.SetSome(testMultiLocationV1n1)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, MultiLocationV1{}, v)
}

var (
	testMultiLocationV1n1 = MultiLocationV1{
		Parents:  4,
		Interior: testJunctionsV1n9,
	}

	// Note that we only use the junctionsV1FuzzOpts here, since no explicit
	// opts are needed for the MultiLocationV1 as the other fields are handled by fuzz.
	multiLocationV1FuzzOpts = junctionsV1FuzzOpts
)

func TestMultiLocationV1_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiLocationV1](t, 100, multiLocationV1FuzzOpts...)
	AssertDecodeNilData[MultiLocationV1](t)
	AssertEncodeEmptyObj[MultiLocationV1](t, 1)
}

func TestMultiLocationV1_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testMultiLocationV1n1, MustHexDecodeString("0x0408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
	})
}

func TestMultiLocationV1_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testMultiLocationV1n1},
	})
}

var (
	testVersionMultiLocation1 = VersionedMultiLocation{
		IsV0:            true,
		MultiLocationV0: testMultiLocationV0n9,
	}
	testVersionMultiLocation2 = VersionedMultiLocation{
		IsV1:            true,
		MultiLocationV1: testMultiLocationV1n1,
	}

	versionedMultiLocationFuzzOpts = CombineFuzzOpts(
		multiLocationV0FuzzOpts,
		multiLocationV1FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(v *VersionedMultiLocation, c fuzz.Continue) {
				if c.RandBool() {
					v.IsV0 = true
					c.Fuzz(&v.MultiLocationV0)

					return
				}

				v.IsV1 = true
				c.Fuzz(&v.MultiLocationV1)
			}),
		},
	)
)

func TestVersionedMultiLocation_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[VersionedMultiLocation](t, 1000, versionedMultiLocationFuzzOpts...)
	AssertDecodeNilData[VersionedMultiLocation](t)
	AssertEncodeEmptyObj[VersionedMultiLocation](t, 0)
}

func TestVersionedMultiLocation_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testVersionMultiLocation1, MustHexDecodeString("0x000800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608")},
		{testVersionMultiLocation2, MustHexDecodeString("0x010408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
	})
}

func TestVersionedMultiLocation_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x000800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608"), testVersionMultiLocation1},
		{MustHexDecodeString("0x010408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testVersionMultiLocation2},
	})
}
