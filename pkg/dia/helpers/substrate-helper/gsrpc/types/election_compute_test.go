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
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	electionComputeFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(e *ElectionCompute, c fuzz.Continue) {
			*e = ElectionCompute(c.Intn(3))
		}),
	}

	optionElectionComputeFuzzOpts = CombineFuzzOpts(
		electionComputeFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(o *OptionElectionCompute, c fuzz.Continue) {
				if c.RandBool() {
					*o = NewOptionElectionComputeEmpty()
					return
				}

				var ec ElectionCompute

				c.Fuzz(&ec)

				*o = NewOptionElectionCompute(ec)
			}),
		},
	)
)

func TestOptionElectionCompute_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionElectionCompute](t, 100, optionElectionComputeFuzzOpts...)
	AssertEncodeEmptyObj[ElectionCompute](t, 1)
}

func TestOptionElectionCompute_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewOptionElectionCompute(NewElectionCompute(byte(0))), MustHexDecodeString("0x0100")},
		{NewOptionElectionCompute(NewElectionCompute(byte(1))), MustHexDecodeString("0x0101")},
		{NewOptionElectionCompute(NewElectionCompute(byte(2))), MustHexDecodeString("0x0102")},
		{NewOptionBytesEmpty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionElectionCompute_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0100"), NewOptionElectionCompute(NewElectionCompute(byte(0)))},
		{MustHexDecodeString("0x0101"), NewOptionElectionCompute(NewElectionCompute(byte(1)))},
		{MustHexDecodeString("0x0102"), NewOptionElectionCompute(NewElectionCompute(byte(2)))},
		{MustHexDecodeString("0x00"), NewOptionBytesEmpty()},
	})
}

func TestOptionElectionCompute_OptionMethods(t *testing.T) {
	o := NewOptionElectionComputeEmpty()
	o.SetSome(NewElectionCompute(1))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, NewElectionCompute(0), v)
}

func TestElectionComputeEncodeDecode(t *testing.T) {
	//decode error
	decoder := scale.NewDecoder(bytes.NewReader([]byte{5}))
	ec0 := ElectionCompute(0)
	err := decoder.Decode(&ec0)
	assert.Error(t, err)

	AssertRoundTripFuzz[ElectionCompute](t, 100, electionComputeFuzzOpts...)
	AssertDecodeNilData[ElectionCompute](t)
	AssertEncodeEmptyObj[ElectionCompute](t, 1)
}
