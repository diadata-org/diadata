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
	testDisputeLocation1 = DisputeLocation{
		IsLocal: true,
	}
	testDisputeLocation2 = DisputeLocation{
		IsRemote: true,
	}

	testDisputeResult1 = DisputeResult{
		IsValid: true,
	}

	testDisputeResult2 = DisputeResult{
		IsInvalid: true,
	}
)

var (
	disputeLocationFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(d *DisputeLocation, c fuzz.Continue) {
			if c.RandBool() {
				d.IsLocal = true
				return
			}

			d.IsRemote = true
		}),
	}
)

func TestDisputeLocation_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testDisputeLocation1)
	AssertRoundtrip(t, testDisputeLocation2)
	AssertRoundTripFuzz[DisputeLocation](t, 100, disputeLocationFuzzOpts...)
	AssertDecodeNilData[DisputeLocation](t)
	AssertEncodeEmptyObj[DisputeLocation](t, 0)
}

func TestDisputeLocation_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testDisputeLocation1, MustHexDecodeString("0x00")},
		{testDisputeLocation2, MustHexDecodeString("0x01")},
	})
}

func TestDisputeLocation_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testDisputeLocation1},
		{MustHexDecodeString("0x01"), testDisputeLocation2},
	})
}

var (
	disputeResultFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(d *DisputeResult, c fuzz.Continue) {
			if c.RandBool() {
				d.IsValid = true
				return
			}

			d.IsInvalid = true
			return
		}),
	}
)

func TestDisputeResult_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testDisputeResult1)
	AssertRoundtrip(t, testDisputeResult2)
	AssertRoundTripFuzz[DisputeResult](t, 100, disputeResultFuzzOpts...)
	AssertDecodeNilData[DisputeResult](t)
	AssertEncodeEmptyObj[DisputeResult](t, 0)
}

func TestDisputeResult_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testDisputeResult1, MustHexDecodeString("0x00")},
		{testDisputeResult2, MustHexDecodeString("0x01")},
	})
}

func TestDisputeResult_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testDisputeResult1},
		{MustHexDecodeString("0x01"), testDisputeResult2},
	})
}
