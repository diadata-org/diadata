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

package testutils

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/blake2b"
)

type FuzzOpt func(f *fuzz.Fuzzer)

func WithNilChance(p float64) FuzzOpt {
	return func(f *fuzz.Fuzzer) {
		f.NilChance(p)
	}
}

func WithFuzzFuncs(fuzzFns ...any) FuzzOpt {
	return func(f *fuzz.Fuzzer) {
		f.Funcs(fuzzFns...)
	}
}

func WithNumElements(atLeast, atMost int) FuzzOpt {
	return func(f *fuzz.Fuzzer) {
		f.NumElements(atLeast, atMost)
	}
}

func WithMaxDepth(depth int) FuzzOpt {
	return func(f *fuzz.Fuzzer) {
		f.MaxDepth(depth)
	}
}

func CombineFuzzOpts(optsSlice ...[]FuzzOpt) []FuzzOpt {
	var o []FuzzOpt

	for _, opts := range optsSlice {
		o = append(o, opts...)
	}

	return o
}

func AssertRoundTripFuzz[T any](t *testing.T, fuzzCount int, fuzzOpts ...FuzzOpt) {
	f := fuzz.New().NilChance(0)

	for _, opt := range fuzzOpts {
		opt(f)
	}

	for i := 0; i < fuzzCount; i++ {
		var fuzzData T

		f.Fuzz(&fuzzData)

		AssertRoundtrip(t, fuzzData)
	}
}

func AssertDecodeNilData[T any](t *testing.T) {
	var obj T

	err := scale.NewDecoder(bytes.NewReader(nil)).Decode(&obj)
	assert.NotNil(t, err)
}

func AssertEncodeEmptyObj[T any](t *testing.T, expectedByteLen int) {
	var obj T

	var buffer = bytes.Buffer{}

	err := scale.NewEncoder(&buffer).Encode(obj)
	assert.Nil(t, err)
	assert.Len(t, buffer.Bytes(), expectedByteLen)
}

type EncodedLengthAssert struct {
	Input    interface{}
	Expected int
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}
	t.Errorf("Received %#v (type %v), expected %#v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func AssertRoundtrip(t *testing.T, value interface{}) {
	var buffer = bytes.Buffer{}
	err := scale.NewEncoder(&buffer).Encode(value)
	assert.NoError(t, err)
	target := reflect.New(reflect.TypeOf(value))
	err = scale.NewDecoder(&buffer).Decode(target.Interface())
	assert.NoError(t, err)
	AssertEqual(t, target.Elem().Interface(), value)
}

func AssertEncodedLength(t *testing.T, encodedLengthAsserts []EncodedLengthAssert) {
	for _, test := range encodedLengthAsserts {
		result, err := codec.EncodedLength(test.Input)
		if err != nil {
			t.Errorf("Encoded length error for input %v: %v\n", test.Input, err)
		}
		if result != test.Expected {
			t.Errorf("Fail, input %v, expected %v, result %v\n", test.Input, test.Expected, result)
		}
	}
}

type EncodingAssert struct {
	Input    interface{}
	Expected []byte
}

func AssertEncode(t *testing.T, encodingAsserts []EncodingAssert) {
	for _, test := range encodingAsserts {
		result, err := codec.Encode(test.Input)
		if err != nil {
			t.Errorf("Encoding error for input %v: %v\n", test.Input, err)
		}

		if !bytes.Equal(result, test.Expected) {
			t.Errorf("Fail, input %v, expected %#x, result %#x\n", test.Input, test.Expected, result)
		}
	}
}

type DecodingAssert struct {
	Input    []byte
	Expected interface{}
}

func AssertDecode(t *testing.T, decodingAsserts []DecodingAssert) {
	for _, test := range decodingAsserts {
		target := reflect.New(reflect.TypeOf(test.Expected))
		err := codec.Decode(test.Input, target.Interface())
		if err != nil {
			t.Errorf("Encoding error for input %v: %v\n", test.Input, err)
		}
		AssertEqual(t, target.Elem().Interface(), test.Expected)
	}
}

type HashAssert struct {
	Input    interface{}
	Expected []byte
}

func AssertHash(t *testing.T, hashAsserts []HashAssert) {
	for _, test := range hashAsserts {
		enc, err := codec.Encode(test.Input)

		if err != nil {
			t.Errorf("Couldn't encode input %v: %s", test.Input, err)
		}

		result := blake2b.Sum256(enc)

		if !bytes.Equal(result[:], test.Expected) {
			t.Errorf("Fail, input %v, expected %#x, result %#x\n", test.Input, test.Expected, result)
		}
	}
}

type EncodeToHexAssert struct {
	Input    interface{}
	Expected string
}

func AssertEncodeToHex(t *testing.T, encodeToHexAsserts []EncodeToHexAssert) {
	for _, test := range encodeToHexAsserts {
		result, err := codec.EncodeToHex(test.Input)
		if err != nil {
			t.Errorf("Hex error for input %v: %v\n", test.Input, err)
		}
		if result != test.Expected {
			t.Errorf("Fail, input %v, expected %v, result %v\n", test.Input, test.Expected, result)
		}
	}
}

type StringAssert struct {
	Input    interface{}
	Expected string
}

func AssertString(t *testing.T, stringAsserts []StringAssert) {
	for _, test := range stringAsserts {
		result := fmt.Sprintf("%v", test.Input)
		if result != test.Expected {
			t.Errorf("Fail, input %v, expected %v, result %v\n", test.Input, test.Expected, result)
		}
	}
}

type EqAssert struct {
	Input    interface{}
	Other    interface{}
	Expected bool
}

func AssertEq(t *testing.T, eqAsserts []EqAssert) {
	for _, test := range eqAsserts {
		result := codec.Eq(test.Input, test.Other)
		if result != test.Expected {
			t.Errorf("Fail, input %v, other %v, expected %v, result %v\n", test.Input, test.Other, test.Expected, result)
		}
	}
}

type jsonMarshalerUnmarshaler[T any] interface {
	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)

	*T // helper type that allows us to instantiate a non-nil T
}

func AssertJSONRoundTrip[T any, U jsonMarshalerUnmarshaler[T]](t *testing.T, value U) {
	b, err := value.MarshalJSON()
	assert.NoError(t, err)

	tt := U(new(T))
	err = tt.UnmarshalJSON(b)
	assert.NoError(t, err)

	AssertEqual(t, value, tt)
}
