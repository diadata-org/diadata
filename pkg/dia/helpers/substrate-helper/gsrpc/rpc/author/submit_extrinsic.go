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

package author

import (
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic"
)

// SubmitExtrinsic will submit a fully formatted extrinsic for block inclusion
func (a *author) SubmitExtrinsic(xt extrinsic.Extrinsic) (types.Hash, error) {
	enc, err := codec.EncodeToHex(xt)
	if err != nil {
		return types.Hash{}, err
	}

	var res string
	err = a.client.Call(&res, "author_submitExtrinsic", enc)
	if err != nil {
		return types.Hash{}, err
	}

	return types.NewHashFromHexString(res)
}
