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

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

func ExampleExampleStruct() {
	type Animal struct {
		Name     string
		Legs     U8
		Children []string
	}

	dog := Animal{Name: "Bello", Legs: 2, Children: []string{"Sam"}}

	encoded, err := EncodeToHex(dog)
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)

	var decoded Animal
	err = DecodeFromHex(encoded, &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
	// Output: 0x1442656c6c6f02040c53616d
	// {Bello 2 [Sam]}
}
