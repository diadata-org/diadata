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
	"github.com/stretchr/testify/assert"
)

var exampleMetadataV13 = Metadata{
	MagicNumber:   0x6174656d,
	Version:       13,
	AsMetadataV13: exampleRuntimeMetadataV13,
}

var exampleRuntimeMetadataV13 = MetadataV13{
	Modules: []ModuleMetadataV13{exampleModuleMetadataV13Empty, exampleModuleMetadataV131, exampleModuleMetadataV132},
}

var exampleModuleMetadataV13Empty = ModuleMetadataV13{
	Name:       "EmptyModule_13",
	HasStorage: false,
	Storage:    StorageMetadataV13{},
	HasCalls:   false,
	Calls:      nil,
	HasEvents:  false,
	Events:     nil,
	Constants:  nil,
	Errors:     nil,
	Index:      0,
}

var exampleModuleMetadataV131 = ModuleMetadataV13{
	Name:       "Module1_13",
	HasStorage: true,
	Storage:    exampleStorageMetadataV13,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
	Index:      1,
}

var exampleModuleMetadataV132 = ModuleMetadataV13{
	Name:       "Module2_13",
	HasStorage: true,
	Storage:    exampleStorageMetadataV13,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
	Index:      2,
}

var exampleStorageMetadataV13 = StorageMetadataV13{
	Prefix: "myStoragePrefix_13",
	Items: []StorageFunctionMetadataV13{exampleStorageFunctionMetadataV13Type, exampleStorageFunctionMetadataV13Map,
		exampleStorageFunctionMetadataV13DoubleMap, exampleStorageFunctionMetadataV13NMap},
}

var exampleStorageFunctionMetadataV13Type = StorageFunctionMetadataV13{
	Name:          "myStorageFunc_13",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV13{IsType: true, AsType: "U8"},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV13Map = StorageFunctionMetadataV13{
	Name:          "myStorageFunc2_13",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV13{IsMap: true, AsMap: exampleMapTypeV10},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV13DoubleMap = StorageFunctionMetadataV13{
	Name:          "myStorageFunc3_13",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV13{IsDoubleMap: true, AsDoubleMap: exampleDoubleMapTypeV10},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV13NMap = StorageFunctionMetadataV13{
	Name:          "myStorageFunc4_13",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV13{IsNMap: true, AsNMap: exampleNMapTypeV13},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleNMapTypeV13 = NMapTypeV13{
	Hashers: []StorageHasherV10{{IsBlake2_256: true}, {IsBlake2_128Concat: true}, {IsIdentity: true}},
	Keys:    []Type{"myKey1", "myKey2", "myKey3"},
	Value:   "and a value",
}

func TestMetadataV13_ExistsModuleMetadata(t *testing.T) {
	assert.True(t, exampleMetadataV13.ExistsModuleMetadata("EmptyModule_13"))
	assert.False(t, exampleMetadataV13.ExistsModuleMetadata("NotExistModule"))
}

func TestMetadataV13_FindEventNamesForEventID(t *testing.T) {
	module, event, err := exampleMetadataV13.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV131.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestMetadataV13_FindEventNamesForUnknownModule(t *testing.T) {
	_, _, err := exampleMetadataV13.FindEventNamesForEventID(EventID([2]byte{1, 18}))

	assert.Error(t, err)
}

func TestMetadataV13_TestFindStorageEntryMetadata(t *testing.T) {
	_, err := exampleMetadataV13.FindStorageEntryMetadata("myStoragePrefix_13", "myStorageFunc2_13")
	assert.NoError(t, err)
}

func TestMetadataV13_TestFindCallIndex(t *testing.T) {
	callIndex, err := exampleMetadataV13.FindCallIndex("Module2_13.my function")
	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV132.Index, callIndex.SectionIndex)
	assert.Equal(t, uint8(0), callIndex.MethodIndex)
}

func TestMetadataV13_TestFindCallIndexWithUnknownModule(t *testing.T) {
	_, err := exampleMetadataV13.FindCallIndex("UnknownModule.my function")
	assert.Error(t, err)
}

func TestMetadataV13_TestFindCallIndexWithUnknownFunction(t *testing.T) {
	_, err := exampleMetadataV13.FindCallIndex("Module2_13.unknownFunction")
	assert.Error(t, err)
}
