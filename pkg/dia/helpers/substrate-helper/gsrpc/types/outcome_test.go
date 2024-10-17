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
	testOutcome1 = Outcome{
		IsComplete:     true,
		CompleteWeight: testWeight,
	}
	testOutcome2 = Outcome{
		IsIncomplete:     true,
		IncompleteWeight: testWeight,
		IncompleteError: XCMError{
			IsOverflow: true,
		},
	}
	testOutcome3 = Outcome{
		IsError: true,
		Error: XCMError{
			IsUnimplemented: true,
		},
	}

	outcomeFuzzOpts = CombineFuzzOpts(
		xcmErrorFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(o *Outcome, c fuzz.Continue) {
				switch c.Intn(3) {
				case 0:
					o.IsComplete = true

					c.Fuzz(&o.CompleteWeight)
				case 1:
					o.IsIncomplete = true

					c.Fuzz(&o.IncompleteWeight)

					c.Fuzz(&o.IncompleteError)
				case 2:
					o.IsError = true

					c.Fuzz(&o.Error)
				}
			}),
		},
	)
)

func TestOutcome_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Outcome](t, 100, outcomeFuzzOpts...)
	AssertDecodeNilData[Outcome](t)
	AssertEncodeEmptyObj[Outcome](t, 0)
}

func TestOutcome_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testOutcome1, MustHexDecodeString("0x002ce909")},
		{testOutcome2, MustHexDecodeString("0x012ce90900")},
		{testOutcome3, MustHexDecodeString("0x0201")},
	})
}

func TestOutcome_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x002ce909"), testOutcome1},
		{MustHexDecodeString("0x012ce90900"), testOutcome2},
		{MustHexDecodeString("0x0201"), testOutcome3},
	})
}
