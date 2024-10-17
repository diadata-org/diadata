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
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var examplePhaseApp = Phase{
	IsApplyExtrinsic: true,
	AsApplyExtrinsic: 42,
}

var examplePhaseFin = Phase{
	IsFinalization: true,
}

var exampleEventApp = EventSystemExtrinsicSuccess{
	Phase:        examplePhaseApp,
	DispatchInfo: DispatchInfo{Weight: testWeight, Class: DispatchClass{IsNormal: true}, PaysFee: Pays{IsYes: true}},
	Topics:       []Hash{{1, 2}},
}

var exampleEventFin = EventSystemExtrinsicSuccess{
	Phase:        examplePhaseFin,
	DispatchInfo: DispatchInfo{Weight: testWeight, Class: DispatchClass{IsNormal: true}, PaysFee: Pays{IsYes: true}},
	Topics:       []Hash{{1, 2}},
}

var exampleEventAppEnc = []byte{0x0, 0x2a, 0x0, 0x0, 0x0, 0x2c, 0xe9, 0x9, 0x0, 0x0, 0x4, 0x1, 0x2, 0x0, 0x0, 0x0, 0x0,
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0} //nolint:lll

var exampleEventFinEnc = []byte{0x1, 0x2c, 0xe9, 0x9, 0x0, 0x0, 0x4, 0x1, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0} //nolint:lll

var (
	eventSystemExtrinsicSuccessFuzzOpts = CombineFuzzOpts(
		phaseFuzzOpts,
		dispatchInfoFuzzOpts,
	)
)

func TestEventSystemExtrinsicSuccess_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[EventSystemExtrinsicSuccess](t, 100, eventSystemExtrinsicSuccessFuzzOpts...)
	AssertDecodeNilData[EventSystemExtrinsicSuccess](t)
	AssertEncodeEmptyObj[EventSystemExtrinsicSuccess](t, 3)
}

func TestEventSystemExtrinsicSuccess_Encode(t *testing.T) {
	encoded, err := Encode(exampleEventFin)
	assert.NoError(t, err)
	assert.Equal(t, exampleEventFinEnc, encoded)

	encoded, err = Encode(exampleEventApp)
	assert.NoError(t, err)
	assert.Equal(t, exampleEventAppEnc, encoded)
}

func TestEventSystemExtrinsicSuccess_Decode(t *testing.T) {
	decoded := EventSystemExtrinsicSuccess{}
	err := Decode(exampleEventFinEnc, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, exampleEventFin, decoded)

	decoded = EventSystemExtrinsicSuccess{}
	err = Decode(exampleEventAppEnc, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, exampleEventApp, decoded)
}

func TestDispatchError(t *testing.T) {
	AssertRoundTripFuzz[DispatchError](t, 1000, dispatchErrorFuzzOpts...)
	AssertDecodeNilData[DispatchError](t)
	AssertEncodeEmptyObj[DispatchError](t, 0)
}

var (
	phaseFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(p *Phase, c fuzz.Continue) {
			switch c.Intn(3) {
			case 0:
				p.IsApplyExtrinsic = true
				c.Fuzz(&p.AsApplyExtrinsic)
			case 1:
				p.IsFinalization = true
			case 2:
				p.IsInitialization = true
			}
		}),
	}
)

func TestPhase(t *testing.T) {
	AssertRoundTripFuzz[Phase](t, 100, phaseFuzzOpts...)
	AssertDecodeNilData[Phase](t)
	AssertEncodeEmptyObj[Phase](t, 0)
}
