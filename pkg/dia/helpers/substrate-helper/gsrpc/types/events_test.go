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
	dispatchInfoFuzzOpts = CombineFuzzOpts(
		dispatchClassFuzzOpts,
		paysFuzzOpts,
	)
)

func TestDispatchInfo_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[DispatchInfo](t, 100, dispatchInfoFuzzOpts...)
	AssertDecodeNilData[DispatchInfo](t)
	AssertEncodeEmptyObj[DispatchInfo](t, 2)
}

func TestVoteThreshold_Decoder(t *testing.T) {
	// SuperMajorityAgainst
	decoder := scale.NewDecoder(bytes.NewReader([]byte{1}))
	vt := VoteThreshold(0)
	err := decoder.Decode(&vt)
	assert.NoError(t, err)
	assert.Equal(t, vt, SuperMajorityAgainst)

	// Error
	decoder = scale.NewDecoder(bytes.NewReader([]byte{3}))
	err = decoder.Decode(&vt)
	assert.Error(t, err)
}

func TestVoteThreshold_Encode(t *testing.T) {
	vt := SuperMajorityAgainst
	var buf bytes.Buffer
	encoder := scale.NewEncoder(&buf)
	assert.NoError(t, encoder.Encode(vt))
	assert.Equal(t, buf.Len(), 1)
	assert.Equal(t, buf.Bytes(), []byte{1})
}

var (
	dispatchResultFuzzOpts = CombineFuzzOpts(
		dispatchErrorFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(d *DispatchResult, c fuzz.Continue) {
				if c.RandBool() {
					d.Ok = true
					return
				}

				c.Fuzz(&d.Error)
			}),
		},
	)
)

func TestDispatchResult_Decode(t *testing.T) {
	AssertRoundTripFuzz[DispatchResult](t, 100, dispatchResultFuzzOpts...)
	AssertDecodeNilData[DispatchResult](t)
	AssertEncodeEmptyObj[DispatchResult](t, 1)
}

var (
	paysFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(p *Pays, c fuzz.Continue) {
			r := c.RandBool()
			p.IsYes = r
			p.IsNo = !r
		}),
	}
)

func TestPaysEncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Pays](t, 1000, paysFuzzOpts...)
	AssertDecodeNilData[Pays](t)
	AssertEncodeEmptyObj[Pays](t, 0)
}

var (
	dispatchClassFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(d *DispatchClass, c fuzz.Continue) {
			switch c.Intn(3) {
			case 0:
				d.IsNormal = true
			case 1:
				d.IsOperational = true
			case 2:
				d.IsMandatory = true
			}
		}),
	}
)

func TestDispatchClassEncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[DispatchClass](t, 100, dispatchClassFuzzOpts...)
	AssertDecodeNilData[DispatchClass](t)
	AssertEncodeEmptyObj[DispatchClass](t, 0)
}

var (
	democracyConvictionFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(d *DemocracyConviction, c fuzz.Continue) {
			*d = DemocracyConviction(c.Intn(7))
		}),
	}
)

func TestDemocracyConviction_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[DemocracyConviction](t, 100, democracyConvictionFuzzOpts...)
	AssertDecodeNilData[DemocracyConviction](t)
	AssertEncodeEmptyObj[DemocracyConviction](t, 1)
}

var (
	voteAccountVoteFuzzOpts = CombineFuzzOpts(
		democracyConvictionFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(v *VoteAccountVote, c fuzz.Continue) {
				if c.RandBool() {
					v.IsStandard = true
					c.Fuzz(&v.AsStandard)
					return
				}

				v.IsSplit = true
				c.Fuzz(&v.AsSplit)
			}),
		},
	)
)

func TestVoteAccountVote_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[VoteAccountVote](t, 100, voteAccountVoteFuzzOpts...)
	AssertDecodeNilData[VoteAccountVote](t)
	AssertEncodeEmptyObj[VoteAccountVote](t, 0)
}

var (
	schedulerLookupErrorFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(e *SchedulerLookupError, c fuzz.Continue) {
			*e = SchedulerLookupError(c.Intn(2))
		}),
	}
)

func TestSchedulerLookupError_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[SchedulerLookupError](t, 100, schedulerLookupErrorFuzzOpts...)
	AssertDecodeNilData[SchedulerLookupError](t)
	AssertEncodeEmptyObj[SchedulerLookupError](t, 1)
}
