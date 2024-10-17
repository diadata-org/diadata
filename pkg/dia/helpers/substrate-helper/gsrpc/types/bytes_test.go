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
	"bytes"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestBytes_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Bytes](t, 100)
	AssertEncodeEmptyObj[Bytes](t, 1)
}

func TestBytes_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewBytes(MustHexDecodeString("0x00")), 2},
		{NewBytes(MustHexDecodeString("0xab1234")), 4},
		{NewBytes(MustHexDecodeString("0x0001")), 3},
	})
}

func TestBytes_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewBytes([]byte{0, 0, 0}), MustHexDecodeString("0x0c000000")},
		{NewBytes([]byte{171, 18, 52}), MustHexDecodeString("0x0cab1234")},
		{NewBytes([]byte{0, 1}), MustHexDecodeString("0x080001")},
		{NewBytes([]byte{18, 52, 86}), MustHexDecodeString("0x0c123456")},
	})
}

func TestBytes_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewBytes([]byte{0, 42, 254}), MustHexDecodeString(
			"0xabf7fe6eb94e0816bf2db57abb296d012f7cb9ddfe59ebf52f9c2770f49a0a46")},
		{NewBytes([]byte{0, 0}), MustHexDecodeString(
			"0xd1200120e01c48b4bbf7e1cd7ebab20087b34ea11e1e9e4ebc2f207aea77139d")},
	})
}

func TestBytes_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewBytes([]byte{0, 0, 0}), "0x0c000000"},
		{NewBytes([]byte{171, 18, 52}), "0x0cab1234"},
		{NewBytes([]byte{0, 1}), "0x080001"},
		{NewBytes([]byte{18, 52, 86}), "0x0c123456"},
	})
}

func TestBytes_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewBytes([]byte{0, 0, 0}), "[0 0 0]"},
		{NewBytes([]byte{171, 18, 52}), "[171 18 52]"},
		{NewBytes([]byte{0, 1}), "[0 1]"},
		{NewBytes([]byte{18, 52, 86}), "[18 52 86]"},
	})
}

func TestBytes_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewBytes([]byte{1, 0, 0}), NewBytes([]byte{1, 0}), false},
		{NewBytes([]byte{0, 0, 1}), NewBytes([]byte{0, 1}), false},
		{NewBytes([]byte{0, 0, 0}), NewBytes([]byte{0, 0}), false},
		{NewBytes([]byte{12, 48, 255}), NewBytes([]byte{12, 48, 255}), true},
		{NewBytes([]byte{0}), NewBytes([]byte{0}), true},
		{NewBytes([]byte{1}), NewBool(true), false},
		{NewBytes([]byte{0}), NewBool(false), false},
	})
}

func TestBytes8_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes8([8]byte{}))
	AssertRoundtrip(t, NewBytes8([8]byte{0, 1, 2, 3, 4, 5, 6, 7}))
}

func TestBytes8_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewBytes8([8]byte{}), 8},
		{NewBytes8([8]byte{7, 6, 5, 4, 3, 2, 1, 0}), 8},
	})
}

func TestBytes8_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewBytes8([8]byte{0, 0, 0}), MustHexDecodeString("0x0000000000000000")},
		{NewBytes8([8]byte{171, 18, 52}), MustHexDecodeString("0xab12340000000000")},
	})
}

func TestBytes8_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewBytes8([8]byte{0, 42, 254}), MustHexDecodeString(
			"0xb327a728de8af3187dd7eaeb74bfff9e9c37eda8c472c7459d51e1fc450faf74")},
		{NewBytes8([8]byte{0, 0}), MustHexDecodeString(
			"0x81e47a19e6b29b0a65b9591762ce5143ed30d0261e5d24a3201752506b20f15c")},
	})
}

func TestBytes8_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewBytes8([8]byte{0, 0, 0}), "0x0000000000000000"},
		{NewBytes8([8]byte{171, 18, 52}), "0xab12340000000000"},
	})
}

func TestBytes8_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewBytes8([8]byte{0, 0, 0}), "[0 0 0 0 0 0 0 0]"},
		{NewBytes8([8]byte{171, 18, 52}), "[171 18 52 0 0 0 0 0]"},
	})
}

func TestBytes8_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewBytes8([8]byte{1, 0, 0}), NewBytes8([8]byte{1, 0}), true},
		{NewBytes8([8]byte{0, 0, 1}), NewBytes8([8]byte{0, 1}), false},
		{NewBytes8([8]byte{0, 0, 0}), NewBytes8([8]byte{0, 0}), true},
		{NewBytes8([8]byte{12, 48, 255}), NewBytes8([8]byte{12, 48, 255}), true},
		{NewBytes8([8]byte{0}), NewBytes8([8]byte{0}), true},
		{NewBytes8([8]byte{1}), NewBool(true), false},
		{NewBytes8([8]byte{0}), NewBool(false), false},
	})
}

func TestBytes16_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes16([16]byte{0, 255, 0, 42}))
}

func TestBytes32_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes32([32]byte{0, 255, 0, 42}))
}

func TestBytes64_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes64([64]byte{0, 255, 0, 42}))
}

func TestBytes128_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes128([128]byte{0, 255, 0, 42}))
}

func TestBytes256_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes256([256]byte{0, 255, 0, 42}))
}

func TestBytes512_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes512([512]byte{0, 255, 0, 42}))
}

func TestBytes1024_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes1024([1024]byte{0, 255, 0, 42}))
}

func TestBytes2048_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewBytes2048([2048]byte{0, 255, 0, 42}))
}

func TestBytesBare_Decode(t *testing.T) {
	b := new(BytesBare)

	err := b.Decode(*scale.NewDecoder(bytes.NewReader(nil)))
	assert.NotNil(t, err)
}
