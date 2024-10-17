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
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/stretchr/testify/assert"
)

var exampleMetadataV7 = Metadata{
	MagicNumber:  0x6174656d,
	Version:      7,
	AsMetadataV7: exampleRuntimeMetadataV7,
}

var exampleRuntimeMetadataV7 = MetadataV7{
	Modules: []ModuleMetadataV7{exampleModuleMetadataV7Empty, exampleModuleMetadataV71, exampleModuleMetadataV72},
}

var exampleModuleMetadataV7Empty = ModuleMetadataV7{
	Name:       "EmptyModule",
	HasStorage: false,
	Storage:    StorageMetadata{},
	HasCalls:   false,
	Calls:      nil,
	HasEvents:  false,
	Events:     nil,
	Constants:  nil,
}

var exampleModuleMetadataV71 = ModuleMetadataV7{
	Name:       "Module1",
	HasStorage: true,
	Storage:    exampleStorageMetadata,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
}

var exampleModuleMetadataV72 = ModuleMetadataV7{
	Name:       "Module2",
	HasStorage: true,
	Storage:    exampleStorageMetadata,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
}

var exampleStorageMetadata = StorageMetadata{
	Prefix: "myStoragePrefix",
	Items: []StorageFunctionMetadataV5{exampleStorageFunctionMetadataV5Type, exampleStorageFunctionMetadataV5Map,
		exampleStorageFunctionMetadataV5DoubleMap},
}

var exampleStorageFunctionMetadataV5Type = StorageFunctionMetadataV5{
	Name:          "myStorageFunc",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV5{IsType: true, AsType: "U8"},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV5Map = StorageFunctionMetadataV5{
	Name:          "myStorageFunc2",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV5{IsMap: true, AsMap: exampleMapTypeV4},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV5DoubleMap = StorageFunctionMetadataV5{
	Name:          "myStorageFunc3",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV5{IsDoubleMap: true, AsDoubleMap: exampleDoubleMapTypeV5},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleDoubleMapTypeV5 = DoubleMapTypeV5{
	Hasher:     StorageHasher{IsBlake2_256: true},
	Key1:       "myKey",
	Key2:       "otherKey",
	Value:      "and a value",
	Key2Hasher: StorageHasher{IsTwox256: true},
}

var exampleModuleConstantMetadataV6 = ModuleConstantMetadataV6{
	Name:          Text("My name"),
	Type:          Type("My Type"),
	Value:         Bytes{123, 23},
	Documentation: []Text{"Doca", "Docb"},
}

func TestMetadataV7_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMetadataV7)
}

func TestFindEventNamesForEventIDV7(t *testing.T) {
	module, event, err := exampleMetadataV7.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV72.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestFindStorageEntryMetadataV7(t *testing.T) {
	_, err := exampleMetadataV7.FindStorageEntryMetadata("myStoragePrefix", "myStorageFunc2")
	assert.NoError(t, err)
}
