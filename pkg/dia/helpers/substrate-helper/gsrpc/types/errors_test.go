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
	testDispatchError1 = DispatchError{
		IsOther: true,
	}
	testDispatchError2 = DispatchError{
		IsCannotLookup: true,
	}
	testDispatchError3 = DispatchError{
		IsBadOrigin: true,
	}
	testDispatchError4 = DispatchError{
		IsModule: true,
		ModuleError: ModuleError{
			Index: 4,
			Error: [4]U8{5, 0, 0, 0},
		},
	}
	testDispatchError5 = DispatchError{
		IsConsumerRemaining: true,
	}
	testDispatchError6 = DispatchError{
		IsNoProviders: true,
	}
	testDispatchError7 = DispatchError{
		IsTooManyConsumers: true,
	}
	testDispatchError8 = DispatchError{
		IsToken: true,
		TokenError: TokenError{
			IsUnsupported: true,
		},
	}
	testDispatchError9 = DispatchError{
		IsArithmetic: true,
		ArithmeticError: ArithmeticError{
			IsDivisionByZero: true,
		},
	}
	testDispatchError10 = DispatchError{
		IsTransactional: true,
		TransactionalError: TransactionalError{
			IsLimitReached: true,
		},
	}

	tokenErrorFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(t *TokenError, c fuzz.Continue) {
			switch c.Intn(7) {
			case 0:
				t.IsNoFunds = true
			case 1:
				t.IsWouldDie = true
			case 2:
				t.IsBelowMinimum = true
			case 3:
				t.IsCannotCreate = true
			case 4:
				t.IsUnknownAsset = true
			case 5:
				t.IsFrozen = true
			case 6:
				t.IsUnsupported = true
			}
		}),
	}

	arithmeticErrorFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(a *ArithmeticError, c fuzz.Continue) {
			switch c.Intn(3) {
			case 0:
				a.IsUnderflow = true
			case 1:
				a.IsOverflow = true
			case 2:
				a.IsDivisionByZero = true
			}
		}),
	}

	transactionalErrorFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(t *TransactionalError, c fuzz.Continue) {
			r := c.RandBool()
			t.IsLimitReached = r
			t.IsNoLayer = !r
		}),
	}

	dispatchErrorFuzzOpts = CombineFuzzOpts(
		tokenErrorFuzzOpts,
		arithmeticErrorFuzzOpts,
		transactionalErrorFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(d *DispatchError, c fuzz.Continue) {
				switch c.Intn(10) {
				case 0:
					d.IsOther = true
				case 1:
					d.IsCannotLookup = true
				case 2:
					d.IsBadOrigin = true
				case 3:
					d.IsModule = true

					c.Fuzz(&d.ModuleError)
				case 4:
					d.IsConsumerRemaining = true
				case 5:
					d.IsNoProviders = true
				case 6:
					d.IsTooManyConsumers = true
				case 7:
					d.IsToken = true

					c.Fuzz(&d.TokenError)
				case 8:
					d.IsArithmetic = true

					c.Fuzz(&d.ArithmeticError)
				case 9:
					d.IsTransactional = true

					c.Fuzz(&d.TransactionalError)
				}
			}),
		},
	)
)

func TestDispatchError_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[DispatchError](t, 1000, dispatchErrorFuzzOpts...)
	AssertDecodeNilData[DispatchError](t)
	AssertEncodeEmptyObj[DispatchError](t, 0)
}

func TestDispatchError_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testDispatchError1, MustHexDecodeString("0x00")},
		{testDispatchError2, MustHexDecodeString("0x01")},
		{testDispatchError3, MustHexDecodeString("0x02")},
		{testDispatchError4, MustHexDecodeString("0x030405000000")},
		{testDispatchError5, MustHexDecodeString("0x04")},
		{testDispatchError6, MustHexDecodeString("0x05")},
		{testDispatchError7, MustHexDecodeString("0x06")},
		{testDispatchError8, MustHexDecodeString("0x0706")},
		{testDispatchError9, MustHexDecodeString("0x0802")},
		{testDispatchError10, MustHexDecodeString("0x0900")},
	})
}

func TestDispatchError_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testDispatchError1},
		{MustHexDecodeString("0x01"), testDispatchError2},
		{MustHexDecodeString("0x02"), testDispatchError3},
		{MustHexDecodeString("0x030405000000"), testDispatchError4},
		{MustHexDecodeString("0x04"), testDispatchError5},
		{MustHexDecodeString("0x05"), testDispatchError6},
		{MustHexDecodeString("0x06"), testDispatchError7},
		{MustHexDecodeString("0x0706"), testDispatchError8},
		{MustHexDecodeString("0x0802"), testDispatchError9},
		{MustHexDecodeString("0x0900"), testDispatchError10},
	})
}
