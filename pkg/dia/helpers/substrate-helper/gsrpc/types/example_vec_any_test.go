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
	"fmt"
	"reflect"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

// MyVal is a custom type that is used to hold arbitrarily encoded data. In this example, we encode uint8s with a 0x00
// and strings with 0x01 as the first byte.
type MyVal struct {
	Value interface{}
}

func (a *MyVal) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		var u uint8
		err = decoder.Decode(&u)
		a.Value = u
	} else if b == 1 {
		var s string
		err = decoder.Decode(&s)
		a.Value = s
	}

	if err != nil {
		return err
	}

	return nil
}

func (a MyVal) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch v := a.Value.(type) {
	case uint8:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(v)
	case string:
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(v)
	default:
		return fmt.Errorf("unknown type %T", v)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

func ExampleExampleVecAny() {
	myValSlice := []MyVal{{uint8(12)}, {"Abc"}}

	encoded, err := Encode(myValSlice)
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)

	var decoded []MyVal
	err = Decode(encoded, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.DeepEqual(myValSlice, decoded))
	// Output: [8 0 12 1 12 65 98 99]
	// true
}
