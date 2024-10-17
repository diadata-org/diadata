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
	"bytes"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	changesTrieSignalFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(c *ChangesTrieSignal, fc fuzz.Continue) {
			c.IsNewConfiguration = true
			fc.Fuzz(&c.AsNewConfiguration)

			return
		}),
	}
)

func TestChangesTrieSignal_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[ChangesTrieSignal](t, 100, changesTrieSignalFuzzOpts...)
	AssertDecodeNilData[ChangesTrieSignal](t)

	var c ChangesTrieSignal

	err := c.Encode(*scale.NewEncoder(bytes.NewBuffer(nil)))
	assert.NotNil(t, err)

	cc := new(ChangesTrieSignal)

	err = cc.Decode(*scale.NewDecoder(bytes.NewReader([]byte{2})))
	assert.NotNil(t, err)
}

var testDigestItem1 = DigestItem{IsOther: true, AsOther: NewBytes([]byte{0xab})}
var testDigestItem2 = DigestItem{IsChangesTrieRoot: true, AsChangesTrieRoot: NewHash([]byte{0x01, 0x02, 0x03})}

var (
	digestItemFuzzOpts = CombineFuzzOpts(
		changesTrieSignalFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(d *DigestItem, c fuzz.Continue) {
				vals := []int{0, 2, 4, 5, 6, 7}
				r := c.Intn(len(vals))

				switch vals[r] {
				case 0:
					d.IsOther = true
					c.Fuzz(&d.AsOther)
				case 2:
					d.IsChangesTrieRoot = true
					c.Fuzz(&d.AsChangesTrieRoot)
				case 4:
					d.IsConsensus = true
					c.Fuzz(&d.AsConsensus)
				case 5:
					d.IsSeal = true
					c.Fuzz(&d.AsSeal)
				case 6:
					d.IsPreRuntime = true
					c.Fuzz(&d.AsPreRuntime)
				case 7:
					d.IsChangesTrieSignal = true
					c.Fuzz(&d.AsChangesTrieSignal)
				}
			}),
		},
	)
)

func TestDigestItem_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, testDigestItem1)
	AssertRoundtrip(t, testDigestItem2)
	AssertRoundtrip(t, DigestItem{
		IsPreRuntime: true,
		AsPreRuntime: PreRuntime{},
	})
	AssertRoundtrip(t, DigestItem{
		IsConsensus: true,
		AsConsensus: Consensus{},
	})
	AssertRoundtrip(t, DigestItem{
		IsSeal: true,
		AsSeal: Seal{},
	})
	AssertRoundtrip(t, DigestItem{
		IsChangesTrieSignal: true,
		AsChangesTrieSignal: ChangesTrieSignal{IsNewConfiguration: true, AsNewConfiguration: NewBytes([]byte{1, 2, 3})},
	})

	AssertRoundTripFuzz[DigestItem](t, 100, digestItemFuzzOpts...)
	AssertDecodeNilData[DigestItem](t)
	AssertEncodeEmptyObj[DigestItem](t, 0)
}
