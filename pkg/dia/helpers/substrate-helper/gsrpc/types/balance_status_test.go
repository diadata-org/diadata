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
	balanceStatusFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(b *BalanceStatus, c fuzz.Continue) {
			*b = BalanceStatus(c.Intn(2))
		}),
	}
)

func TestBalanceStatusEncodeDecode(t *testing.T) {
	//decode error
	decoder := scale.NewDecoder(bytes.NewReader([]byte{5}))
	bs0 := BalanceStatus(0)
	err := decoder.Decode(&bs0)
	assert.Error(t, err)

	AssertRoundTripFuzz[BalanceStatus](t, 100, balanceStatusFuzzOpts...)
	AssertDecodeNilData[BalanceStatus](t)
	AssertEncodeEmptyObj[BalanceStatus](t, 1)
}
