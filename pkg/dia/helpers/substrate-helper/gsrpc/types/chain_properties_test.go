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
)

var testChainProperties1 = ChainProperties{}

var testChainProperties2 = ChainProperties{
	IsSS58Format: true, AsSS58Format: 0x01, IsTokenDecimals: true, AsTokenDecimals: 18,
	IsTokenSymbol: true, AsTokenSymbol: "FOO",
}

func TestChainProperties_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testChainProperties1)
	AssertRoundtrip(t, testChainProperties2)
}

func TestChainProperties_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testChainProperties1, []byte{0x0, 0x0, 0x0}},
		{testChainProperties2, []byte{0x01, 0x01, 0x01, 0x12, 0x00, 0x00, 0x00, 0x01, 0x0c, 0x46, 0x4f, 0x4f}},
	})
}

func TestChainProperties_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{[]byte{0x0, 0x0, 0x0}, testChainProperties1},
		{[]byte{0x01, 0x01, 0x01, 0x12, 0x00, 0x00, 0x00, 0x01, 0x0c, 0x46, 0x4f, 0x4f}, testChainProperties2},
	})
}
