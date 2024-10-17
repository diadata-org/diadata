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

package rpc

import (
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/client"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/author"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/beefy"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/chain"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/mmr"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/offchain"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/system"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

type RPC struct {
	Author   author.Author
	Beefy    beefy.Beefy
	Chain    chain.Chain
	MMR      mmr.MMR
	Offchain offchain.Offchain
	State    state.State
	System   system.System
	client   client.Client
}

func NewRPC(cl client.Client) (*RPC, error) {
	st := state.NewState(cl)
	meta, err := st.GetMetadataLatest()
	if err != nil {
		return nil, err
	}

	opts := types.SerDeOptionsFromMetadata(meta)
	types.SetSerDeOptions(opts)

	return &RPC{
		Author:   author.NewAuthor(cl),
		Beefy:    beefy.NewBeefy(cl),
		Chain:    chain.NewChain(cl),
		MMR:      mmr.NewMMR(cl),
		Offchain: offchain.NewOffchain(cl),
		State:    st,
		System:   system.NewSystem(cl),
		client:   cl,
	}, nil
}
