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
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
)

var (
	testInstanceMetadata = ItemMetadata{
		Deposit:  NewU128(*big.NewInt(1234)),
		Data:     Bytes("some_data"),
		IsFrozen: true,
	}
)

func TestInstanceMetadata_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[ItemMetadata](t, 1000)
	AssertDecodeNilData[ItemMetadata](t)
	AssertEncodeEmptyObj[ItemMetadata](t, 18)
}

func TestInstanceMetadata_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testInstanceMetadata, MustHexDecodeString("0xd204000000000000000000000000000024736f6d655f6461746101")},
	})
}

func TestInstanceMetadata_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0xd204000000000000000000000000000024736f6d655f6461746101"), testInstanceMetadata},
	})
}
