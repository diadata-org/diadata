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
	testClassDetails = CollectionDetails{
		Owner:             newTestAccountID(),
		Issuer:            newTestAccountID(),
		Admin:             newTestAccountID(),
		Freezer:           newTestAccountID(),
		TotalDeposit:      NewU128(*big.NewInt(123)),
		FreeHolding:       true,
		Instances:         4,
		InstanceMetadatas: 5,
		Attributes:        6,
		IsFrozen:          true,
	}
)

func TestClassDetails_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[CollectionDetails](t, 1000)
	AssertDecodeNilData[CollectionDetails](t)
	AssertEncodeEmptyObj[CollectionDetails](t, 158)
}

func TestClassDetails_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testClassDetails,
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f207b0000000000000000000000000000000104000000050000000600000001"),
		},
	})
}

func TestClassDetails_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f207b0000000000000000000000000000000104000000050000000600000001"),
			testClassDetails,
		},
	})
}
