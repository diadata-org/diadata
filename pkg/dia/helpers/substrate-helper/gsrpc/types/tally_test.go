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
	testTally = Tally{
		Votes: NewU128(*big.NewInt(123)),
		Total: NewU128(*big.NewInt(456)),
	}
)

func TestTally_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Tally](t, 100)
	AssertDecodeNilData[Tally](t)
	AssertEncodeEmptyObj[Tally](t, 32)
}

func TestTally_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testTally, MustHexDecodeString("0x7b000000000000000000000000000000c8010000000000000000000000000000")},
	})
}

func TestTally_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x7b000000000000000000000000000000c8010000000000000000000000000000"), testTally},
	})
}
