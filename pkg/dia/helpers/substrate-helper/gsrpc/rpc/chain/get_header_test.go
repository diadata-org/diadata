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

package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain_GetHeaderLatest(t *testing.T) {
	header, err := testChain.GetHeaderLatest()
	assert.NoError(t, err)
	assert.NotEmpty(t, header.Number)
}

func TestChain_GetHeader(t *testing.T) {
	res, err := testChain.GetFinalizedHead()
	assert.NoError(t, err)

	header, err := testChain.GetHeader(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}
