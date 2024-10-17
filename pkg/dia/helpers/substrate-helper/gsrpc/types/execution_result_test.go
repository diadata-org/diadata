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
	testExecutionResult = ExecutionResult{
		Outcome: 1,
		Error: XCMError{
			IsOverflow: true,
		},
	}

	executionResultFuzzOpts = xcmErrorFuzzOpts

	optionExecutionResultFuzzOpts = CombineFuzzOpts(
		executionResultFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(o *OptionExecutionResult, c fuzz.Continue) {
				if c.RandBool() {
					*o = NewOptionExecutionResultEmpty()
					return
				}

				var r ExecutionResult

				c.Fuzz(&r)

				*o = NewOptionExecutionResult(r)
			}),
		},
	)
)

func TestOptionExecutionResult_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionExecutionResult](t, 1000, optionExecutionResultFuzzOpts...)
}

func TestOptionExecutionResult_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewOptionExecutionResult(testExecutionResult), MustHexDecodeString("0x010100000000")},
		{NewOptionExecutionResultEmpty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionExecutionResult_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x010100000000"), NewOptionExecutionResult(testExecutionResult)},
		{MustHexDecodeString("0x00"), NewOptionBytesEmpty()},
	})
}

func TestOptionExecutionResult_OptionMethods(t *testing.T) {
	o := NewOptionExecutionResultEmpty()
	o.SetSome(testExecutionResult)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, ExecutionResult{}, v)
}

func TestExecutionResult_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[ExecutionResult](t, 1000, executionResultFuzzOpts...)
	AssertDecodeNilData[ExecutionResult](t)
	AssertEncodeEmptyObj[ExecutionResult](t, 4)
}

func TestExecutionResult_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testExecutionResult, MustHexDecodeString("0x0100000000")},
	})
}

func TestExecutionResult_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0100000000"), testExecutionResult},
	})
}
