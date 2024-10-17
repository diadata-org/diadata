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

package state

import (
	"os"
	"strings"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/client"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpcmocksrv"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

var testState State

func TestMain(m *testing.M) {
	s := rpcmocksrv.New()
	err := s.RegisterName("state", &mockSrv)
	if err != nil {
		panic(err)
	}

	cl, err := client.Connect(s.URL)
	// cl, err := client.Connect(config.Default().RPCURL)
	// cl, err := client.Connect("ws://127.0.0.1:9944")
	if err != nil {
		panic(err)
	}
	testState = NewState(cl)

	os.Exit(m.Run())
}

// MockSrv holds data and methods exposed by the RPC Mock Server used in integration tests
type MockSrv struct {
	blockHashLatest          types.Hash
	metadataString           string
	runtimeVersion           types.RuntimeVersion
	storageKeyHex            string
	storageKeyHexEmpty       string
	storageChangeSets        []types.StorageChangeSet
	storageDataHex           string
	storageSize              types.U64
	storageHashHex           string
	childStorageKeyHex       string // the key pointing to the child storage trie
	childStorageTrieKeyHex   string // a key within the child storage trie
	childStorageTrieValueHex string // a value stored int he child storage trie
	childStorageTrieValue    ChildStorageTrieTestVal
	childStorageTrieSize     types.U64
	childStorageTrieHashHex  string
}

func (s *MockSrv) GetMetadata(hash *string) string {
	return mockSrv.metadataString
}

func (s *MockSrv) GetRuntimeVersion(hash *string) types.RuntimeVersion {
	return mockSrv.runtimeVersion
}

func (s *MockSrv) GetKeys(key string, hash *string) []string {
	if !strings.HasPrefix(mockSrv.storageKeyHex, key) {
		panic("key not found")
	}
	return []string{mockSrv.storageKeyHex}
}

func (s *MockSrv) GetStorage(key string, hash *string) string {
	if key != s.storageKeyHex {
		return ""
	}
	return mockSrv.storageDataHex
}

func (s *MockSrv) GetStorageSize(key string, hash *string) types.U64 {
	return mockSrv.storageSize
}

func (s *MockSrv) GetStorageHash(key string, hash *string) string {
	return mockSrv.storageHashHex
}

func (s *MockSrv) GetChildKeys(childStorageKey, prefix string, hash *string) []string {
	if childStorageKey != mockSrv.childStorageKeyHex {
		panic("childStorageKey not found")
	}
	if !strings.HasPrefix(mockSrv.childStorageTrieKeyHex, prefix) {
		panic("no keys for prefix found")
	}
	return []string{mockSrv.childStorageTrieKeyHex}
}

func (s *MockSrv) GetChildStorage(childStorageKey, key string, hash *string) string {
	if childStorageKey != mockSrv.childStorageKeyHex {
		panic("childStorageKey not found")
	}
	if key != mockSrv.childStorageTrieKeyHex {
		panic("key not found")
	}
	return mockSrv.childStorageTrieValueHex
}

func (s *MockSrv) GetChildStorageSize(childStorageKey, key string, hash *string) types.U64 {
	if childStorageKey != mockSrv.childStorageKeyHex {
		panic("childStorageKey not found")
	}
	if key != mockSrv.childStorageTrieKeyHex {
		panic("key not found")
	}
	return mockSrv.childStorageTrieSize
}

func (s *MockSrv) GetChildStorageHash(childStorageKey, key string, hash *string) string {
	if childStorageKey != mockSrv.childStorageKeyHex {
		panic("childStorageKey not found")
	}
	if key != mockSrv.childStorageTrieKeyHex {
		panic("key not found")
	}
	return mockSrv.childStorageTrieHashHex
}

func (s *MockSrv) QueryStorage(keys []string, startBlock string, block *string) []types.StorageChangeSet {
	if len(keys) != 1 {
		panic("keys need to have len of 1 in tests")
	}
	if keys[0] != mockSrv.storageKeyHex {
		panic("key not found")
	}
	if startBlock != "0xdd1816b6f6889f46e23b0d6750bc441af9dad0fda8bae90677c1708d01035fbe" {
		panic("startBlock must be 0xdd1816b6f6889f46e23b0d6750bc441af9dad0fda8bae90677c1708d01035fbe in tests")
	}

	return mockSrv.storageChangeSets
}

func (s *MockSrv) QueryStorageAt(keys []string, hash *string) []types.StorageChangeSet {
	if len(keys) != 1 {
		panic("keys need to have len of 1 in tests")
	}
	if keys[0] != mockSrv.storageKeyHex {
		panic("key not found")
	}

	return mockSrv.storageChangeSets
}

// func (s *MockSrv) SubscribeStorage(args []string) {
// 	fmt.Println("Hit")
// }

type ChildStorageTrieTestVal struct {
	Key     types.Hash
	OtherID types.Hash
	Value   types.U32
}

// mockSrv sets default data used in tests. This data might become stale when substrate is updated – just run the tests
// against real servers and update the values stored here. To do that, replace s.URL with
// config.Default().RPCURL
var mockSrv = MockSrv{
	blockHashLatest:          types.Hash{1, 2, 3},
	metadataString:           types.MetadataV14Data,
	runtimeVersion:           types.RuntimeVersion{APIs: []types.RuntimeVersionAPI{{APIID: "0xdf6acb689907609b", Version: 0x2}, {APIID: "0x37e397fc7c91f5e4", Version: 0x1}, {APIID: "0x40fe3ad401f8959a", Version: 0x3}, {APIID: "0xd2bc9897eed08f15", Version: 0x1}, {APIID: "0xf78b278be53f454c", Version: 0x1}, {APIID: "0xed99c5acb25eedf5", Version: 0x2}, {APIID: "0xdd718d5cc53262d4", Version: 0x1}, {APIID: "0x7801759919ee83e5", Version: 0x1}}, AuthoringVersion: 0xa, ImplName: "substrate-node", ImplVersion: 0x3e, SpecName: "node", SpecVersion: 0x3c}, //nolint:lll
	storageKeyHex:            "0x0e4944cfd98d6f4cc374d16f5a4e3f9c",
	storageKeyHexEmpty:       "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
	storageChangeSets:        []types.StorageChangeSet{{Block: types.Hash{0xdd, 0x18, 0x16, 0xb6, 0xf6, 0x88, 0x9f, 0x46, 0xe2, 0x3b, 0xd, 0x67, 0x50, 0xbc, 0x44, 0x1a, 0xf9, 0xda, 0xd0, 0xfd, 0xa8, 0xba, 0xe9, 0x6, 0x77, 0xc1, 0x70, 0x8d, 0x1, 0x3, 0x5f, 0xbe}, Changes: []types.KeyValueOption{{StorageKey: types.StorageKey{0xe, 0x49, 0x44, 0xcf, 0xd9, 0x8d, 0x6f, 0x4c, 0xc3, 0x74, 0xd1, 0x6f, 0x5a, 0x4e, 0x3f, 0x9c}, HasStorageData: true, StorageData: types.StorageDataRaw{0x88, 0x2, 0x66, 0x9f, 0x6e, 0x1, 0x0, 0x0}}}}, {Block: types.Hash{0x82, 0x14, 0xa1, 0x80, 0x8b, 0xd6, 0xb0, 0x46, 0xc8, 0x77, 0xa6, 0x4f, 0xce, 0xad, 0xb4, 0xa2, 0xa7, 0x3a, 0x65, 0x76, 0x9f, 0x61, 0x4, 0xc0, 0x20, 0xd7, 0x59, 0xad, 0x8f, 0x61, 0xc0, 0xd8}, Changes: []types.KeyValueOption{{StorageKey: types.StorageKey{0xe, 0x49, 0x44, 0xcf, 0xd9, 0x8d, 0x6f, 0x4c, 0xc3, 0x74, 0xd1, 0x6f, 0x5a, 0x4e, 0x3f, 0x9c}, HasStorageData: true, StorageData: types.StorageDataRaw{0x40, 0xe, 0x66, 0x9f, 0x6e, 0x1, 0x0, 0x0}}}}}, //nolint:lll
	storageDataHex:           "0xb82d895d00000000",
	storageSize:              926778,
	storageHashHex:           "0xdf0e877ee1cb973b9a566f53707d365b269d7131b55e65b9790994e4e63b95e1",
	childStorageKeyHex:       "0x3a6368696c645f73746f726167653a64656661756c743a05470000", //nolint:lll // beginning with hex encoded `:child_storage:` as per https://github.com/paritytech/substrate/blob/master/core/primitives/storage/src/lib.rs#L71
	childStorageTrieKeyHex:   "0x81914b11321c39f8728981888024196b616142cc0369234775b20b539aaf29d0",
	childStorageTrieValueHex: "0x81914b11321c39f8728981888024196b616142cc0369234775b20b539aaf29d09c1705d98d059a2d7f5faa89277ee5d0a38cc455f8b5fdf38fda471e988cb8a921000000", //nolint:lll
	childStorageTrieValue: ChildStorageTrieTestVal{
		Key:     types.NewHash(codec.MustHexDecodeString("0x81914b11321c39f8728981888024196b616142cc0369234775b20b539aaf29d0")), //nolint:lll
		OtherID: types.NewHash(codec.MustHexDecodeString("0x9c1705d98d059a2d7f5faa89277ee5d0a38cc455f8b5fdf38fda471e988cb8a9")), //nolint:lll
		Value:   types.NewU32(0x21),
	},
	childStorageTrieSize:    68,
	childStorageTrieHashHex: "0x20e3fc48a91087d091c17de08a5c470de53ccdaebd361025b0e5b7c65b9a0d30", //nolint:lll
}
