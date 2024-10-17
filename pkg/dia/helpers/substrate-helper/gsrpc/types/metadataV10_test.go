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

var exampleMetadataV10 = Metadata{
	MagicNumber:   0x6174656d,
	Version:       10,
	AsMetadataV10: exampleRuntimeMetadataV10,
}

var exampleRuntimeMetadataV10 = MetadataV10{
	Modules: []ModuleMetadataV10{exampleModuleMetadataV10Empty, exampleModuleMetadataV101, exampleModuleMetadataV102},
}

var exampleModuleMetadataV10Empty = ModuleMetadataV10{
	Name:       "EmptyModule",
	HasStorage: false,
	Storage:    StorageMetadataV10{},
	HasCalls:   false,
	Calls:      nil,
	HasEvents:  false,
	Events:     nil,
	Constants:  nil,
	Errors:     nil,
}

var exampleModuleMetadataV101 = ModuleMetadataV10{
	Name:       "Module1",
	HasStorage: true,
	Storage:    exampleStorageMetadataV10,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
}

var exampleModuleMetadataV102 = ModuleMetadataV10{
	Name:       "Module2",
	HasStorage: true,
	Storage:    exampleStorageMetadataV10,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
}

var exampleStorageMetadataV10 = StorageMetadataV10{
	Prefix: "myStoragePrefix",
	Items: []StorageFunctionMetadataV10{exampleStorageFunctionMetadataV10Type, exampleStorageFunctionMetadataV10Map,
		exampleStorageFunctionMetadataV10DoubleMap},
}

var exampleStorageFunctionMetadataV10Type = StorageFunctionMetadataV10{
	Name:          "myStorageFunc",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV10{IsType: true, AsType: "U8"},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV10Map = StorageFunctionMetadataV10{
	Name:          "myStorageFunc2",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV10{IsMap: true, AsMap: exampleMapTypeV10},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV10DoubleMap = StorageFunctionMetadataV10{
	Name:          "myStorageFunc3",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV10{IsDoubleMap: true, AsDoubleMap: exampleDoubleMapTypeV10},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleMapTypeV10 = MapTypeV10{
	Hasher: StorageHasherV10{IsBlake2_256: true},
	Key:    "my key",
	Value:  "and my value",
	Linked: false,
}

var exampleDoubleMapTypeV10 = DoubleMapTypeV10{
	Hasher:     StorageHasherV10{IsBlake2_256: true},
	Key1:       "myKey",
	Key2:       "otherKey",
	Value:      "and a value",
	Key2Hasher: StorageHasherV10{IsTwox256: true},
}

func TestMetadataV10_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMetadataV10)
}

func TestFindEventNamesForEventIDV10(t *testing.T) {
	module, event, err := exampleMetadataV10.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV102.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestFindStorageEntryMetadataV10(t *testing.T) {
	_, err := exampleMetadataV10.FindStorageEntryMetadata("myStoragePrefix", "myStorageFunc2")
	assert.NoError(t, err)
}
