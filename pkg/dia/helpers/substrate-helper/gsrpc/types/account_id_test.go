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
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	optionAccountIDFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *OptionAccountID, c fuzz.Continue) {
			if c.RandBool() {
				*o = NewOptionAccountIDEmpty()
				return
			}
			var accId AccountID
			c.Fuzz(&accId)

			*o = NewOptionAccountID(accId)
		}),
	}
)

func newTestAccountID() AccountID {
	accID, err := NewAccountID(testAccountIDBytes)

	if err != nil {
		panic(fmt.Errorf("couldn't create test account ID: %w", err))
	}

	return *accID
}

func TestOptionAccountID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionAccountID](t, 100, optionAccountIDFuzzOpts...)
	AssertEncodeEmptyObj[OptionAccountID](t, 1)
}

var testAccountIDBytes = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

func TestOptionAccountID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewOptionAccountID(newTestAccountID()), MustHexDecodeString("0x010102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")},
		{NewOptionAccountIDEmpty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionAccountID_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x010102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"), NewOptionAccountID(newTestAccountID())},
		{MustHexDecodeString("0x00"), NewOptionAccountIDEmpty()},
	})
}

func TestOptionAccountID_OptionMethods(t *testing.T) {
	o := NewOptionAccountIDEmpty()

	o.SetSome(newTestAccountID())

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	var a AccountID

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, a, v)
}

func TestAccountID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AccountID](t, 100, WithNilChance(0.01))
	AssertDecodeNilData[AccountID](t)
	AssertEncodeEmptyObj[AccountID](t, 32)
}

func TestAccountID_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{newTestAccountID(), 32},
	})
}

func TestAccountID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{newTestAccountID(), MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")}, //nolint:lll
	})
}

func TestAccountID_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{
			newTestAccountID(),
			MustHexDecodeString("0x441edc56cebc8e285d02267aa650819f15add7b06ef9b41b2690128dce655924"),
		},
	})
}

func TestAccountID_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{newTestAccountID(), "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"},
	})
}

func TestAccountID_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{newTestAccountID(), "[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32]"},
	})
}

func TestAccountID_Eq(t *testing.T) {
	b := testAccountIDBytes[:31]
	b = append(b, 11)

	accID, err := NewAccountID(b)
	assert.NoError(t, err)

	AssertEq(t, []EqAssert{
		{newTestAccountID(), newTestAccountID(), true},
		{newTestAccountID(), accID, false},
	})
}

func TestAccountID_ToBytes(t *testing.T) {
	accID, err := NewAccountID(testAccountIDBytes)
	assert.NoError(t, err)

	assert.Equal(t, accID.ToBytes(), testAccountIDBytes)
}

func TestAccountID_ToHexString(t *testing.T) {
	accID, err := NewAccountID(testAccountIDBytes)
	assert.NoError(t, err)

	hb, err := Hex(testAccountIDBytes)
	assert.NoError(t, err)

	assert.Equal(t, hb, accID.ToHexString())
}

func TestAccountID_JSONMarshalUnmarshal(t *testing.T) {
	accID, err := NewAccountID(testAccountIDBytes)
	assert.NoError(t, err)

	b, err := json.Marshal(accID)
	assert.NoError(t, err)

	var res AccountID

	err = json.Unmarshal(b, &res)
	assert.NoError(t, err)

	assert.True(t, accID.Equal(&res))
}
