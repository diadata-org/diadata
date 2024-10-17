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
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var testExtrinsicStatus0 = ExtrinsicStatus{IsFuture: true}
var testExtrinsicStatus1 = ExtrinsicStatus{IsReady: true}
var testExtrinsicStatus2 = ExtrinsicStatus{IsBroadcast: true, AsBroadcast: []Text{"This", "is", "broadcast"}}
var testExtrinsicStatus3 = ExtrinsicStatus{IsInBlock: true, AsInBlock: NewHash([]byte{0xaa})}
var testExtrinsicStatus4 = ExtrinsicStatus{IsRetracted: true, AsRetracted: NewHash([]byte{0xbb})}
var testExtrinsicStatus5 = ExtrinsicStatus{IsFinalityTimeout: true, AsFinalityTimeout: NewHash([]byte{0xcc})}
var testExtrinsicStatus6 = ExtrinsicStatus{IsFinalized: true, AsFinalized: NewHash([]byte{0xdd})}
var testExtrinsicStatus7 = ExtrinsicStatus{IsUsurped: true, AsUsurped: NewHash([]byte{0xee})}
var testExtrinsicStatus8 = ExtrinsicStatus{IsDropped: true}
var testExtrinsicStatus9 = ExtrinsicStatus{IsInvalid: true}

var (
	extrinsicStatusFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(e *ExtrinsicStatus, c fuzz.Continue) {
			switch c.Intn(10) {
			case 0:
				e.IsFuture = true
			case 1:
				e.IsReady = true
			case 2:
				e.IsBroadcast = true
				c.Fuzz(&e.AsBroadcast)
			case 3:
				e.IsInBlock = true
				c.Fuzz(&e.AsInBlock)
			case 4:
				e.IsRetracted = true
				c.Fuzz(&e.AsRetracted)
			case 5:
				e.IsFinalityTimeout = true
				c.Fuzz(&e.AsFinalityTimeout)
			case 6:
				e.IsFinalized = true
				c.Fuzz(&e.AsFinalized)
			case 7:
				e.IsUsurped = true
				c.Fuzz(&e.AsUsurped)
			case 8:
				e.IsDropped = true
			case 9:
				e.IsInvalid = true
			}
		}),
	}
)

func TestExtrinsicStatus_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testExtrinsicStatus0)
	AssertRoundtrip(t, testExtrinsicStatus1)
	AssertRoundtrip(t, testExtrinsicStatus2)
	AssertRoundtrip(t, testExtrinsicStatus3)
	AssertRoundtrip(t, testExtrinsicStatus4)
	AssertRoundtrip(t, testExtrinsicStatus5)
	AssertRoundtrip(t, testExtrinsicStatus6)
	AssertRoundtrip(t, testExtrinsicStatus7)
	AssertRoundtrip(t, testExtrinsicStatus8)
	AssertRoundtrip(t, testExtrinsicStatus9)
	AssertRoundTripFuzz[ExtrinsicStatus](t, 1000, extrinsicStatusFuzzOpts...)
	AssertDecodeNilData[ExtrinsicStatus](t)
	AssertEncodeEmptyObj[ExtrinsicStatus](t, 0)
}

func TestExtrinsicStatus_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testExtrinsicStatus0, []byte{0x00}},
		{testExtrinsicStatus1, []byte{0x01}},
		{testExtrinsicStatus2, MustHexDecodeString("0x020c10546869730869732462726f616463617374")},
		{testExtrinsicStatus3, MustHexDecodeString("0x03aa00000000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
		{testExtrinsicStatus4, MustHexDecodeString("0x04bb00000000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
		{testExtrinsicStatus5, MustHexDecodeString("0x05cc00000000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
		{testExtrinsicStatus6, MustHexDecodeString("0x06dd00000000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
		{testExtrinsicStatus7, MustHexDecodeString("0x07ee00000000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
		{testExtrinsicStatus8, []byte{0x08}},
		{testExtrinsicStatus9, []byte{0x09}},
	})
}

func TestExtrinsicStatus_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{[]byte{0x00}, testExtrinsicStatus0},
		{[]byte{0x01}, testExtrinsicStatus1},
		{MustHexDecodeString("0x020c10546869730869732462726f616463617374"), testExtrinsicStatus2},
		{MustHexDecodeString("0x03aa00000000000000000000000000000000000000000000000000000000000000"), testExtrinsicStatus3}, //nolint:lll
		{MustHexDecodeString("0x04bb00000000000000000000000000000000000000000000000000000000000000"), testExtrinsicStatus4}, //nolint:lll
		{MustHexDecodeString("0x05cc00000000000000000000000000000000000000000000000000000000000000"), testExtrinsicStatus5}, //nolint:lll
		{MustHexDecodeString("0x06dd00000000000000000000000000000000000000000000000000000000000000"), testExtrinsicStatus6}, //nolint:lll
		{MustHexDecodeString("0x07ee00000000000000000000000000000000000000000000000000000000000000"), testExtrinsicStatus7}, //nolint:lll
		{[]byte{0x08}, testExtrinsicStatus8},
		{[]byte{0x09}, testExtrinsicStatus9},
	})
}

var testExtrinsicStatusTestCases = []struct {
	encoded []byte
	decoded ExtrinsicStatus
}{
	{
		[]byte("\"future\""),
		ExtrinsicStatus{IsFuture: true},
	}, {
		[]byte("\"ready\""),
		ExtrinsicStatus{IsReady: true},
	}, {
		[]byte("{\"broadcast\":[\"hello\",\"world\"]}"),
		ExtrinsicStatus{IsBroadcast: true, AsBroadcast: []Text{"hello", "world"}},
	}, {
		[]byte("{\"inBlock\":\"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8\"}"),
		ExtrinsicStatus{IsInBlock: true, AsInBlock: NewHash(MustHexDecodeString(
			"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8"))},
	}, {
		[]byte("{\"retracted\":\"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8\"}"),
		ExtrinsicStatus{IsRetracted: true, AsRetracted: NewHash(MustHexDecodeString(
			"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8"))},
	}, {
		[]byte("{\"finalityTimeout\":\"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8\"}"),
		ExtrinsicStatus{IsFinalityTimeout: true, AsFinalityTimeout: NewHash(MustHexDecodeString(
			"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8"))},
	}, {
		[]byte("{\"finalized\":\"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8\"}"),
		ExtrinsicStatus{IsFinalized: true, AsFinalized: NewHash(MustHexDecodeString(
			"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8f8"))},
	}, {
		[]byte("{\"usurped\":\"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8ab\"}"),
		ExtrinsicStatus{IsUsurped: true, AsUsurped: NewHash(MustHexDecodeString(
			"0x95e3b7f86541d06306691a2fe8cbd935d0bdd28ea14fe515e2db0fa87847f8ab"))},
	}, {
		[]byte("\"dropped\""),
		ExtrinsicStatus{IsDropped: true},
	}, {
		[]byte("\"invalid\""),
		ExtrinsicStatus{IsInvalid: true},
	},
}

func TestExtrinsicStatus_UnmarshalJSON(t *testing.T) {
	for _, test := range testExtrinsicStatusTestCases {
		var actual ExtrinsicStatus
		err := json.Unmarshal(test.encoded, &actual)
		assert.NoError(t, err)
		assert.Equal(t, test.decoded, actual)
	}
}

func TestExtrinsicStatus_MarshalJSON(t *testing.T) {
	for _, test := range testExtrinsicStatusTestCases {
		actual, err := json.Marshal(test.decoded)
		assert.NoError(t, err)
		assert.Equal(t, test.encoded, actual)
	}
}
