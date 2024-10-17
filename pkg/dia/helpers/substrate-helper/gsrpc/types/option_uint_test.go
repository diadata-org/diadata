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
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	optionU128FuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *OptionU128, c fuzz.Continue) {
			if c.RandBool() {
				*o = NewOptionU128Empty()
				return
			}

			var u U128

			c.Fuzz(&u)

			*o = NewOptionU128(u)
		}),
	}
)

func TestOptionU8_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionU128](t, 100, optionU128FuzzOpts...)
	AssertEncodeEmptyObj[OptionU128](t, 1)
}

func TestOptionU8_OptionMethods(t *testing.T) {
	o := NewOptionU8Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, U8(0), v)
}

func TestOptionU16_OptionMethods(t *testing.T) {
	o := NewOptionU16Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, U16(0), v)
}

func TestOptionU32_OptionMethods(t *testing.T) {
	o := NewOptionU32Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, U32(0), v)
}

func TestOptionU64_OptionMethods(t *testing.T) {
	o := NewOptionU64Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, U64(0), v)
}

func TestOptionU128_OptionMethods(t *testing.T) {
	o := NewOptionU128Empty()
	o.SetSome(NewU128(*big.NewInt(11)))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, NewU128(*big.NewInt(0)), v)
}

func TestOptionU16_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionU16(NewU16(14)))
	AssertRoundtrip(t, NewOptionU16(NewU16(0)))
	AssertRoundtrip(t, NewOptionU16Empty())
}

func TestOptionU32_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionU32(NewU32(21)))
	AssertRoundtrip(t, NewOptionU32(NewU32(0)))
	AssertRoundtrip(t, NewOptionU32Empty())
}

func TestOptionU64_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionU64(NewU64(28)))
	AssertRoundtrip(t, NewOptionU64(NewU64(0)))
	AssertRoundtrip(t, NewOptionU64Empty())
}

func TestOptionU128_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionU128(NewU128(*big.NewInt(123))))
	AssertRoundtrip(t, NewOptionU128(NewU128(*big.NewInt(0))))
	AssertRoundtrip(t, NewOptionU128Empty())
}
