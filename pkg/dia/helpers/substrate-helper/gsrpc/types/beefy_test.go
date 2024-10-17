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
	"fmt"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var sig1 = [65]byte{85, 132, 85, 173, 129, 39, 157, 240, 121, 92, 201, 133, 88, 14, 79, 183, 93, 114, 217, 72, 209, 16, 123, 42, 200, 10, 9, 171, 237, 77, 168, 72, 12, 116, 108, 195, 33, 242, 49, 154, 94, 153, 168, 48, 227, 20, 209, 13, 211, 205, 104, 206, 61, 192, 195, 60, 134, 233, 155, 203, 120, 22, 249, 186, 1}
var sig2 = [65]byte{45, 110, 31, 129, 5, 195, 55, 168, 108, 221, 154, 170, 205, 196, 150, 87, 127, 61, 184, 197, 94, 249, 230, 253, 72, 242, 197, 192, 90, 34, 116, 112, 116, 145, 99, 93, 139, 163, 223, 100, 243, 36, 87, 91, 123, 42, 52, 72, 123, 202, 35, 36, 182, 160, 4, 99, 149, 167, 22, 129, 190, 61, 12, 42, 0}

func TestBeefySignature(t *testing.T) {
	empty := NewOptionBeefySignatureEmpty()
	assert.True(t, empty.IsNone())
	assert.False(t, empty.IsSome())

	sig := NewOptionBeefySignature(BeefySignature{})
	sig.SetNone()
	assert.True(t, sig.IsNone())
	sig.SetSome(BeefySignature{})
	assert.True(t, sig.IsSome())
	ok, _ := sig.Unwrap()
	assert.True(t, ok)
	AssertRoundtrip(t, sig)
}

func makeCommitment() (*Commitment, error) {
	data, err := Encode([]byte("Hello World!"))
	if err != nil {
		return nil, err
	}

	payloadItem := PayloadItem{
		ID:   [2]byte{'m', 'h'},
		Data: data,
	}

	commitment := Commitment{
		Payload:        []PayloadItem{payloadItem},
		BlockNumber:    5,
		ValidatorSetID: 0,
	}

	return &commitment, nil
}

func makeLargeCommitment() (*Commitment, error) {
	data := MustHexDecodeString("0xb5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c")

	payloadItem := PayloadItem{
		ID:   [2]byte{'m', 'h'},
		Data: data,
	}

	commitment := Commitment{
		Payload:        []PayloadItem{payloadItem},
		BlockNumber:    5,
		ValidatorSetID: 3,
	}

	return &commitment, nil
}

func TestCommitment_Encode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)
	AssertEncode(t, []EncodingAssert{
		{c, MustHexDecodeString("0x046d68343048656c6c6f20576f726c6421050000000000000000000000")},
	})
}

func TestLargeCommitment_Encode(t *testing.T) {
	c, err := makeLargeCommitment()
	assert.NoError(t, err)
	fmt.Println(len(c.Payload[0].Data))
	fmt.Println(EncodeToHex(c))
}

func TestCommitment_Decode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	AssertDecode(t, []DecodingAssert{
		{
			Input:    MustHexDecodeString("0x046d68343048656c6c6f20576f726c6421050000000000000000000000"),
			Expected: *c,
		},
	})
}

func TestCommitment_EncodeDecode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	AssertRoundtrip(t, *c)
}

func TestSignedCommitment_Decode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	s := SignedCommitment{
		Commitment: *c,
		Signatures: []OptionBeefySignature{
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignature(sig2),
		},
	}

	AssertDecode(t, []DecodingAssert{
		{
			Input:    MustHexDecodeString("0x046d68343048656c6c6f20576f726c642105000000000000000000000004300400000008558455ad81279df0795cc985580e4fb75d72d948d1107b2ac80a09abed4da8480c746cc321f2319a5e99a830e314d10dd3cd68ce3dc0c33c86e99bcb7816f9ba012d6e1f8105c337a86cdd9aaacdc496577f3db8c55ef9e6fd48f2c5c05a2274707491635d8ba3df64f324575b7b2a34487bca2324b6a0046395a71681be3d0c2a00"),
			Expected: s,
		},
	})
}

func TestSignedCommitment_EncodeDecode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	s := SignedCommitment{
		Commitment: *c,
		Signatures: []OptionBeefySignature{
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
		},
	}

	AssertRoundtrip(t, s)
}

func TestBeefySignature_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[BeefySignature](t, 100)
	AssertDecodeNilData[BeefySignature](t)
	AssertEncodeEmptyObj[BeefySignature](t, 65)
}

var (
	optionBeefySignatureFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *OptionBeefySignature, c fuzz.Continue) {
			if c.RandBool() {
				*o = NewOptionBeefySignatureEmpty()
				return
			}

			var b BeefySignature
			c.Fuzz(&b)

			*o = NewOptionBeefySignature(b)
		}),
	}
)

func TestOptionBeefySignature_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionBeefySignature](t, 100, optionBeefySignatureFuzzOpts...)
	AssertEncodeEmptyObj[OptionBeefySignature](t, 1)
}
