// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based system RPC calls
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

package system

import (
	"os"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/client"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpcmocksrv"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

var testSystem System

func TestMain(m *testing.M) {
	s := rpcmocksrv.New()
	err := s.RegisterName("system", &mockSrv)
	if err != nil {
		panic(err)
	}

	cl, err := client.Connect(s.URL)
	// cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testSystem = NewSystem(cl)

	os.Exit(m.Run())
}

// MockSrv holds data and methods exposed by the RPC Mock Server used in integration tests
type MockSrv struct {
	chain        types.Text
	health       types.Health
	name         types.Text
	networkState types.NetworkState
	peers        []types.PeerInfo
	properties   types.ChainProperties
	version      types.Text
}

func (s *MockSrv) Chain() types.Text {
	return mockSrv.chain
}

func (s *MockSrv) Health() types.Health {
	return mockSrv.health
}

func (s *MockSrv) Name() types.Text {
	return mockSrv.name
}

func (s *MockSrv) NetworkState() types.NetworkState {
	return mockSrv.networkState
}

func (s *MockSrv) Peers() []types.PeerInfo {
	return mockSrv.peers
}

func (s *MockSrv) Properties() types.ChainProperties {
	return mockSrv.properties
}

func (s *MockSrv) Version() types.Text {
	return mockSrv.version
}

// mockSrv sets default data used in tests. This data might become stale when substrate is updated – just run the tests
// against real servers and update the values stored here. To do that, replace s.URL with
// config.Default().RPCURL
var mockSrv = MockSrv{
	chain:        "test-chain",
	health:       types.Health{Peers: 2, IsSyncing: false, ShouldHavePeers: true},
	name:         "test-node",
	networkState: types.NetworkState{PeerID: "my-peer-id"},
	peers: []types.PeerInfo{{PeerID: "another-peer-id", Roles: "Role", ProtocolVersion: 42,
		BestHash: types.NewHash(codec.MustHexDecodeString("0xabcd")), BestNumber: 420}},
	properties: types.ChainProperties{IsTokenDecimals: true, AsTokenDecimals: 18,
		IsTokenSymbol: true, AsTokenSymbol: "GSRPCCOIN"},
	version: "My version",
}
