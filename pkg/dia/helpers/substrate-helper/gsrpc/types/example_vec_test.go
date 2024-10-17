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

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

func ExampleExampleVec_simple() {
	ingredients := []string{"salt", "sugar"}

	encoded, err := EncodeToHex(ingredients)
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)

	var decoded []string
	err = DecodeFromHex(encoded, &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
	// Output: 0x081073616c74147375676172
	// [salt sugar]
}

func ExampleExampleVec_struct() {
	type Votes struct {
		Options     [2]string
		Yay         []string
		Nay         []string
		Outstanding []string
	}

	votes := Votes{
		Options:     [2]string{"no deal", "muddle through"},
		Yay:         []string{"Alice"},
		Nay:         nil,
		Outstanding: []string{"Bob", "Carol"},
	}

	encoded, err := Encode(votes)
	if err != nil {
		panic(err)
	}
	var decoded Votes
	err = Decode(encoded, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.DeepEqual(votes, decoded))
	// Output: true
}

// type MyOption struct {
// 	Woohoo *bool
// }

// type MyOptionNoPoointer struct {
// 	Woohoo bool
// }

// func NewMyOption(b bool) MyOption {
// 	return MyOption{&b}
// }

// func Example2() {
// 	myopt := NewMyOption(true)
// 	// myopt := NewBool(true)

// 	encoded, err := EncodeToHex(myopt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(encoded)

// 	var decoded MyOption
// 	err = DecodeFromHex(encoded, &decoded)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(decoded)
// 	// Output: 0x081073616c74147375676172
// 	// [salt sugar]
// }
