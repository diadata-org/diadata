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

func TestOptionHash_OptionMethods(t *testing.T) {
	o := NewOptionHashEmpty()
	o.SetSome(NewHash(hash20))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, Hash{}, v)
}

func TestOptionH160_OptionMethods(t *testing.T) {
	o := NewOptionH160Empty()
	o.SetSome(NewH160(hash20))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, H160{}, v)
}

func TestOptionH256_OptionMethods(t *testing.T) {
	o := NewOptionH256Empty()
	o.SetSome(NewH256(hash20))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, H256{}, v)
}

func TestOptionH512_OptionMethods(t *testing.T) {
	o := NewOptionH512Empty()
	o.SetSome(NewH512(hash20))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, H512{}, v)
}

func TestOptionH160_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionH160(NewH160(hash20)))
	AssertRoundtrip(t, NewOptionH160Empty())
}

func TestOptionH256_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionH256(NewH256(hash32)))
	AssertRoundtrip(t, NewOptionH256Empty())
}

func TestOptionH512_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionH512(NewH512(hash64)))
	AssertRoundtrip(t, NewOptionH512Empty())
}

func TestOptionHash_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, NewOptionHash(NewHash(hash32)))
	AssertRoundtrip(t, NewOptionHashEmpty())
}
