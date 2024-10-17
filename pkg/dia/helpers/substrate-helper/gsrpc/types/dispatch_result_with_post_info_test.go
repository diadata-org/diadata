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
	testDispatchResultWithPostInfo1 = DispatchResultWithPostInfo{
		IsOk: true,
		Ok: PostDispatchInfo{
			ActualWeight: NewOption(testWeight),
			PaysFee: Pays{
				IsYes: true,
			},
		},
	}
	testDispatchResultWithPostInfo2 = DispatchResultWithPostInfo{
		IsError: true,
		Error: DispatchErrorWithPostInfo{
			PostInfo: PostDispatchInfo{
				ActualWeight: NewOption(testWeight),
				PaysFee: Pays{
					IsNo: true,
				},
			},
			Error: DispatchError{
				IsOther: true,
			},
		},
	}

	dispatchResultWithPostInfoFuzzOpts = CombineFuzzOpts(
		optionFuzzOpts,
		paysFuzzOpts,
		dispatchErrorFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(d *DispatchResultWithPostInfo, c fuzz.Continue) {
				if c.RandBool() {
					d.IsOk = true
					c.Fuzz(&d.Ok)
					return
				}

				d.IsError = true
				c.Fuzz(&d.Error)
			}),
		},
	)
)

func TestDispatchResultWithPostInfo_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[DispatchResultWithPostInfo](t, 1000, dispatchResultWithPostInfoFuzzOpts...)
	AssertDecodeNilData[DispatchResultWithPostInfo](t)
	AssertEncodeEmptyObj[DispatchResultWithPostInfo](t, 0)
}

func TestDispatchResultWithPostInfo_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testDispatchResultWithPostInfo1, MustHexDecodeString("0x00012ce90900")},
		{testDispatchResultWithPostInfo2, MustHexDecodeString("0x01012ce9090100")},
	})
}

func TestDispatchResultWithPostInfo_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00012ce90900"), testDispatchResultWithPostInfo1},
		{MustHexDecodeString("0x01012ce9090100"), testDispatchResultWithPostInfo2},
	})
}
