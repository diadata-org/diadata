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
)

var (
	testCollatorID        = [32]U8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	testCollatorSignature = [64]U8{9, 8, 7, 6, 5, 4}

	testCandidateDescriptor = CandidateDescriptor{
		ParachainID:                  11,
		RelayParent:                  NewHash([]byte("hash")),
		CollatorID:                   testCollatorID,
		PersistentValidationDataHash: NewHash([]byte("hash2")),
		PoVHash:                      NewHash([]byte("hash3")),
		ErasureRoot:                  NewHash([]byte("hash4")),
		CollatorSignature:            testCollatorSignature,
		ParaHead:                     NewHash([]byte("hash5")),
		ValidationCodeHash:           NewHash([]byte("hash6")),
	}

	testCandidateReceipt = CandidateReceipt{
		Descriptor:      testCandidateDescriptor,
		CommitmentsHash: NewHash([]byte("hash7")),
	}
)

func TestCandidateReceipt_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[CandidateReceipt](t, 1000)
	AssertDecodeNilData[CandidateReceipt](t)
	AssertEncodeEmptyObj[CandidateReceipt](t, 324)
}

func TestCandidateReceipt_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testCandidateReceipt,
			MustHexDecodeString("0x0b0000006861736800000000000000000000000000000000000000000000000000000000010203040506070809000000000000000000000000000000000000000000000068617368320000000000000000000000000000000000000000000000000000006861736833000000000000000000000000000000000000000000000000000000686173683400000000000000000000000000000000000000000000000000000009080706050400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000686173683500000000000000000000000000000000000000000000000000000068617368360000000000000000000000000000000000000000000000000000006861736837000000000000000000000000000000000000000000000000000000"),
		},
	})
}

func TestCandidateReceipt_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x0b0000006861736800000000000000000000000000000000000000000000000000000000010203040506070809000000000000000000000000000000000000000000000068617368320000000000000000000000000000000000000000000000000000006861736833000000000000000000000000000000000000000000000000000000686173683400000000000000000000000000000000000000000000000000000009080706050400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000686173683500000000000000000000000000000000000000000000000000000068617368360000000000000000000000000000000000000000000000000000006861736837000000000000000000000000000000000000000000000000000000"),
			testCandidateReceipt,
		},
	})
}
