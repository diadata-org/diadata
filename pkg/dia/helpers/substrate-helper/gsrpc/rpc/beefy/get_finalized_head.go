// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2021 Snowfork
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

package beefy

import (
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

// GetFinalizedHead returns the hash of the latest BEEFY block
func (b *beefy) GetFinalizedHead() (types.Hash, error) {
	var res string

	err := b.client.Call(&res, "beefy_getFinalizedHead")
	if err != nil {
		return types.Hash{}, err
	}

	return types.NewHashFromHexString(res)
}
