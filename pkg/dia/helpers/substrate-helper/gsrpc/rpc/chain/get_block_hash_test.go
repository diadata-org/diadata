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

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/stretchr/testify/assert"
)

func TestChain_GetBlockHash(t *testing.T) {
	res, err := testChain.GetBlockHash(1)
	assert.NoError(t, err)

	blk, err := testChain.GetBlock(res)
	assert.NoError(t, err)
	assert.Equal(t, types.BlockNumber(1), blk.Block.Header.Number)
}

func TestChain_GetBlockHashLatest(t *testing.T) {
	res, err := testChain.GetBlockHashLatest()
	assert.NoError(t, err)

	blk, err := testChain.GetBlock(res)
	assert.NoError(t, err)
	assert.True(t, blk.Block.Header.Number > 0)
}
