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
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestOptionI8_OptionMethods(t *testing.T) {
	o := NewOptionI8Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, I8(0), v)
}

func TestOptionI16_OptionMethods(t *testing.T) {
	o := NewOptionI16Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, I16(0), v)
}

func TestOptionI32_OptionMethods(t *testing.T) {
	o := NewOptionI32Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, I32(0), v)
}

func TestOptionI64_OptionMethods(t *testing.T) {
	o := NewOptionI64Empty()
	o.SetSome(11)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, I64(0), v)
}

func TestOptionI8_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionI8(NewI8(7)))
	AssertRoundtrip(t, NewOptionI8(NewI8(0)))
	AssertRoundtrip(t, NewOptionI8Empty())
}

func TestOptionI16_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionI16(NewI16(14)))
	AssertRoundtrip(t, NewOptionI16(NewI16(0)))
	AssertRoundtrip(t, NewOptionI16Empty())
}

func TestOptionI32_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionI32(NewI32(21)))
	AssertRoundtrip(t, NewOptionI32(NewI32(0)))
	AssertRoundtrip(t, NewOptionI32Empty())
}

func TestOptionI64_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionI64(NewI64(28)))
	AssertRoundtrip(t, NewOptionI64(NewI64(0)))
	AssertRoundtrip(t, NewOptionI64Empty())
}
