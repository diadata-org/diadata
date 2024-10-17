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
	"golang.org/x/crypto/blake2b"
)

func ExampleExampleTuple() {
	// This represents a document tuple of types [uint64, hash]
	type Doc struct {
		ID   U64
		Hash Hash
	}

	doc := Doc{12, blake2b.Sum256([]byte("My document"))}

	encoded, err := EncodeToHex(doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded)

	var decoded Doc
	err = DecodeFromHex(encoded, &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
	// Output: 0x0c000000000000009199a254aedc9d92a3157cd27bd21ceccc1e2ecee5760788663a3e523bc1a759
	// {12 [145 153 162 84 174 220 157 146 163 21 124 210 123 210 28 236 204 30 46 206 229 118 7 136 102 58 62 82 59 193 167 89]}
}
