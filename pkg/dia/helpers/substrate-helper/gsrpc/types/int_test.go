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
	"github.com/stretchr/testify/assert"
)

func TestI8_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI8(0))
	AssertRoundtrip(t, NewI8(12))
	AssertRoundtrip(t, NewI8(-12))
}

func TestI8_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI8(-13), 1}})
}

func TestI8_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewI8(-29), MustHexDecodeString("0xe3")},
	})
}

func TestI8_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI8(-29), MustHexDecodeString("0xb683f1b6c99388ff3443b35a0051eeaafdc5e364e771bdfc72c7fd5d2be800bc")},
	})
}

func TestI8_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI8(-29), "0xe3"},
	})
}

func TestI8_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI8(-29), "-29"},
	})
}

func TestI8_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI8(23), NewI8(23), true},
		{NewI8(-23), NewI8(23), false},
		{NewI8(23), NewU8(23), false},
		{NewI8(23), NewBool(false), false},
	})
}

func TestI8_JSONMarshalUnmarshal(t *testing.T) {
	i := NewI8(1)
	AssertJSONRoundTrip(t, &i)
}

func TestI16_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI16(0))
	AssertRoundtrip(t, NewI16(12))
	AssertRoundtrip(t, NewI16(-12))
}

func TestI16_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI16(-13), 2}})
}

func TestI16_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewI16(-29), MustHexDecodeString("0xe3ff")},
	})
}

func TestI16_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI16(-29), MustHexDecodeString("0x39fbf34f574b72d1815c602a2fe95b7af4b5dfd7bc92a2fc0824aa55f8b9d7b2")},
	})
}

func TestI16_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI16(-29), "0xe3ff"},
	})
}

func TestI16_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI16(-29), "-29"},
	})
}

func TestI16_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI16(23), NewI16(23), true},
		{NewI16(-23), NewI16(23), false},
		{NewI16(23), NewU16(23), false},
		{NewI16(23), NewBool(false), false},
	})
}

func TestI16_JSONMarshalUnmarshal(t *testing.T) {
	i := NewI16(1)
	AssertJSONRoundTrip(t, &i)
}

func TestI32_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI32(0))
	AssertRoundtrip(t, NewI32(12))
	AssertRoundtrip(t, NewI32(-12))
}

func TestI32_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI32(-13), 4}})
}

func TestI32_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewI32(-29), MustHexDecodeString("0xe3ffffff")},
	})
}

func TestI32_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI32(-29), MustHexDecodeString("0x6ef9d4772b9d657bfa727862d9690d5bf8b9045943279e95d3ae0743684f1b95")},
	})
}

func TestI32_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI32(-29), "0xe3ffffff"},
	})
}

func TestI32_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI32(-29), "-29"},
	})
}

func TestI32_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI32(23), NewI32(23), true},
		{NewI32(-23), NewI32(23), false},
		{NewI32(23), NewU32(23), false},
		{NewI32(23), NewBool(false), false},
	})
}

func TestI32_JSONMarshalUnmarshal(t *testing.T) {
	i := NewI32(1)
	AssertJSONRoundTrip(t, &i)
}

func TestI64_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI64(0))
	AssertRoundtrip(t, NewI64(12))
	AssertRoundtrip(t, NewI64(-12))
}

func TestI64_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI64(-13), 8}})
}

func TestI64_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewI64(-29), MustHexDecodeString("0xe3ffffffffffffff")},
	})
}

func TestI64_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI64(-29), MustHexDecodeString("0x4d42db2aa4a23bde81a3ad3705220affaa457c56a0135080c71db7783fec8f44")},
	})
}

func TestI64_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI64(-29), "0xe3ffffffffffffff"},
	})
}

func TestI64_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI64(-29), "-29"},
	})
}

func TestI64_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI64(23), NewI64(23), true},
		{NewI64(-23), NewI64(23), false},
		{NewI64(23), NewU64(23), false},
		{NewI64(23), NewBool(false), false},
	})
}

func TestI64_JSONMarshalUnmarshal(t *testing.T) {
	i := NewI64(1)
	AssertJSONRoundTrip(t, &i)
}

func TestI128_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI128(*big.NewInt(0)))
	AssertRoundtrip(t, NewI128(*big.NewInt(12)))
	AssertRoundtrip(t, NewI128(*big.NewInt(-12)))

	bigPos := big.NewInt(0)
	bigPos.SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	AssertRoundtrip(t, NewI128(*bigPos))

	bigNeg := big.NewInt(0)
	bigNeg.SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	bigNeg.Neg(bigNeg)
	AssertRoundtrip(t, NewI128(*bigNeg))
}

func TestI128_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI128(*big.NewInt(-13)), 16}})
}

func TestI128_Encode(t *testing.T) {
	a := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})

	b := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})
	b.Neg(b)

	c := big.NewInt(0).SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	d := big.NewInt(0).SetBytes([]byte{127, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15})
	d.Neg(d)

	AssertEncode(t, []EncodingAssert{
		{NewI128(*big.NewInt(0)), MustHexDecodeString("0x00000000000000000000000000000000")},
		{NewI128(*big.NewInt(29)), MustHexDecodeString("0x1d000000000000000000000000000000")},
		{NewI128(*big.NewInt(-29)), MustHexDecodeString("0xe3ffffffffffffffffffffffffffffff")},
		{NewI128(*a), MustHexDecodeString("0x34120000000000000000000000000000")},
		{NewI128(*b), MustHexDecodeString("0xccedffffffffffffffffffffffffffff")},
		{NewI128(*c), MustHexDecodeString("0x100f0e0d0c0b0a090807060504030201")},
		{NewI128(*d), MustHexDecodeString("0xf1f0f1f2f3f4f5f6f7f8f9fafbfcfd80")},
	})
}

func TestI128_Decode(t *testing.T) {
	a := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})

	b := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})
	b.Neg(b)

	c := big.NewInt(0).SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	d := big.NewInt(0).SetBytes([]byte{127, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15})
	d.Neg(d)

	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00000000000000000000000000000000"), NewI128(*big.NewInt(0))},
		{MustHexDecodeString("0x1d000000000000000000000000000000"), NewI128(*big.NewInt(29))},
		{MustHexDecodeString("0xe3ffffffffffffffffffffffffffffff"), NewI128(*big.NewInt(-29))},
		{MustHexDecodeString("0x34120000000000000000000000000000"), NewI128(*a)},
		{MustHexDecodeString("0xccedffffffffffffffffffffffffffff"), NewI128(*b)},
		{MustHexDecodeString("0x100f0e0d0c0b0a090807060504030201"), NewI128(*c)},
		{MustHexDecodeString("0xf1f0f1f2f3f4f5f6f7f8f9fafbfcfd80"), NewI128(*d)},
	})
}

func TestI128_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI128(*big.NewInt(-29)), MustHexDecodeString(
			"0x7f8f93dd36321a50796a2e88df3bc7238abad58361c2051009bc457a000c4de9")},
	})
}

func TestI128_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI128(*big.NewInt(-29)), "0xe3ffffffffffffffffffffffffffffff"},
	})
}

func TestI128_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI128(*big.NewInt(-29)), "-29"},
	})
}

func TestI128_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI128(*big.NewInt(23)), NewI128(*big.NewInt(23)), true},
		{NewI128(*big.NewInt(-23)), NewI128(*big.NewInt(23)), false},
		{NewI128(*big.NewInt(23)), NewU64(23), false},
		{NewI128(*big.NewInt(23)), NewBool(false), false},
	})
}

func TestI256_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewI256(*big.NewInt(0)))
	AssertRoundtrip(t, NewI256(*big.NewInt(12)))
	AssertRoundtrip(t, NewI256(*big.NewInt(-12)))

	bigPos := big.NewInt(0)
	bigPos.SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})
	AssertRoundtrip(t, NewI256(*bigPos))

	bigNeg := big.NewInt(0)
	bigNeg.SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})
	bigNeg.Neg(bigNeg)
	AssertRoundtrip(t, NewI256(*bigNeg))
}

func TestI256_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{{NewI256(*big.NewInt(-13)), 32}})
}

func TestI256_Encode(t *testing.T) {
	a := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})

	b := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})
	b.Neg(b)

	c := big.NewInt(0).SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})

	d := big.NewInt(0).SetBytes([]byte{127, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})
	d.Neg(d)

	AssertEncode(t, []EncodingAssert{
		{NewI256(*big.NewInt(0)), MustHexDecodeString(
			"0x0000000000000000000000000000000000000000000000000000000000000000")},
		{NewI256(*big.NewInt(29)), MustHexDecodeString(
			"0x1d00000000000000000000000000000000000000000000000000000000000000")},
		{NewI256(*big.NewInt(-29)), MustHexDecodeString(
			"0xe3ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")},
		{NewI256(*a), MustHexDecodeString("0x3412000000000000000000000000000000000000000000000000000000000000")},
		{NewI256(*b), MustHexDecodeString("0xccedffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")},
		{NewI256(*c), MustHexDecodeString("0x201f1e1d1c1b1a1817161514131211100f0e0d0c0b0a09080706050403020100")},
		{NewI256(*d), MustHexDecodeString("0xe0e0e1e2e3e4e5e7e8e9eaebecedeef0f0f1f2f3f4f5f6f7f8f9fafbfcfd80ff")},
	})
}

func TestI256_Decode(t *testing.T) {
	a := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})

	b := big.NewInt(0).SetBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 52})
	b.Neg(b)

	c := big.NewInt(0).SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})

	d := big.NewInt(0).SetBytes([]byte{127, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15,
		17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 32})
	d.Neg(d)

	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0000000000000000000000000000000000000000000000000000000000000000"),
			NewI256(*big.NewInt(0))},
		{MustHexDecodeString("0x1d00000000000000000000000000000000000000000000000000000000000000"),
			NewI256(*big.NewInt(29))},
		{MustHexDecodeString("0xe3ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
			NewI256(*big.NewInt(-29))},
		{MustHexDecodeString("0x3412000000000000000000000000000000000000000000000000000000000000"), NewI256(*a)},
		{MustHexDecodeString("0xccedffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), NewI256(*b)},
		{MustHexDecodeString("0x201f1e1d1c1b1a1817161514131211100f0e0d0c0b0a09080706050403020100"), NewI256(*c)},
		{MustHexDecodeString("0xe0e0e1e2e3e4e5e7e8e9eaebecedeef0f0f1f2f3f4f5f6f7f8f9fafbfcfd80ff"), NewI256(*d)},
	})
}

func TestI256_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewI256(*big.NewInt(-29)), MustHexDecodeString(
			"0xca6ae1636199279abc3e15e366ea463cf06829e7816e1ad08c0c15c158dfeba6")},
	})
}

func TestI256_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewI256(*big.NewInt(-29)), "0xe3ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
	})
}

func TestI256_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewI256(*big.NewInt(-29)), "-29"},
	})
}

func TestI256_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewI256(*big.NewInt(23)), NewI256(*big.NewInt(23)), true},
		{NewI256(*big.NewInt(-23)), NewI256(*big.NewInt(23)), false},
		{NewI256(*big.NewInt(23)), NewI128(*big.NewInt(23)), false},
		{NewI256(*big.NewInt(23)), NewU64(23), false},
		{NewI256(*big.NewInt(23)), NewBool(false), false},
	})
}

func TestBigIntToIntBytes(t *testing.T) {
	res, err := BigIntToIntBytes(big.NewInt(4), 2)
	assert.NoError(t, err)
	assert.Equal(t, MustHexDecodeString("0x0004"), res)

	res, err = BigIntToIntBytes(big.NewInt(0).Neg(big.NewInt(4)), 2)
	assert.NoError(t, err)
	assert.Equal(t, MustHexDecodeString("0xfffc"), res)

	_, err = BigIntToIntBytes(big.NewInt(128), 1)
	assert.EqualError(t, err, "cannot encode big.Int to []byte: given big.Int exceeds highest positive number "+
		"127 for an int with 8 bits")

	c := big.NewInt(129)
	_, err = BigIntToIntBytes(c.Neg(c), 1)
	assert.EqualError(t, err, "cannot encode big.Int to []byte: given big.Int exceeds highest negative number "+
		"-128 for an int with 8 bits")
}

func TestIntBytesToBigInt(t *testing.T) {
	res, err := IntBytesToBigInt(MustHexDecodeString("0x0004"))
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(4), res)

	res, err = IntBytesToBigInt(MustHexDecodeString("0xfffc"))
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(0).Neg(big.NewInt(4)), res)

	_, err = IntBytesToBigInt([]byte{})
	assert.EqualError(t, err, "cannot decode an empty byte slice")
}

func TestBigIntToIntBytes_128(t *testing.T) {
	b := big.NewInt(0)
	b.SetBytes([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x12, 0x34})

	res, err := BigIntToIntBytes(b, 16)
	assert.NoError(t, err)
	assert.Equal(t, MustHexDecodeString("0x00000000000000000000000000001234"), res)

	b = b.Neg(b)
	res, err = BigIntToIntBytes(b, 16)
	assert.NoError(t, err)
	assert.Equal(t, MustHexDecodeString("0xffffffffffffffffffffffffffffedcc"), res)
}
