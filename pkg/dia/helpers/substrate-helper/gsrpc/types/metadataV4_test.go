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

var exampleMetadataV4 = Metadata{
	MagicNumber:  0x6174656d,
	Version:      4,
	AsMetadataV4: exampleRuntimeMetadataV4,
}

var exampleRuntimeMetadataV4 = MetadataV4{
	Modules: []ModuleMetadataV4{exampleModuleMetadataV4Empty, exampleModuleMetadataV41, exampleModuleMetadataV42},
}

var exampleCallIndex = CallIndex{
	SectionIndex: 123,
	MethodIndex:  234,
}

var exampleModuleMetadataV4Empty = ModuleMetadataV4{
	Name:       "EmptyModule",
	Prefix:     "EmptyModule",
	HasStorage: false,
	Storage:    nil,
	HasCalls:   false,
	Calls:      nil,
	HasEvents:  false,
	Events:     nil,
}

var exampleModuleMetadataV41 = ModuleMetadataV4{
	Name:       "Module1",
	Prefix:     "Module1",
	HasStorage: true,
	Storage:    []StorageFunctionMetadataV4{exampleStorageFunctionMetadataV4DoubleMap},
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
}

var exampleModuleMetadataV42 = ModuleMetadataV4{
	Name:       "Module2",
	Prefix:     "Module2",
	HasStorage: true,
	Storage:    []StorageFunctionMetadataV4{exampleStorageFunctionMetadataV4DoubleMap},
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
}

var exampleStorageFunctionMetadataV4Type = StorageFunctionMetadataV4{
	Name:          "myStorageFunc",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV4{IsType: true, AsType: "U8"},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV4Map = StorageFunctionMetadataV4{
	Name:          "myStorageFunc2",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV4{IsMap: true, AsMap: exampleMapTypeV4},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleStorageFunctionMetadataV4DoubleMap = StorageFunctionMetadataV4{
	Name:          "myStorageFunc3",
	Modifier:      StorageFunctionModifierV0{IsOptional: true},
	Type:          StorageFunctionTypeV4{IsDoubleMap: true, AsDoubleMap: exampleDoubleMapTypeV4},
	Fallback:      []byte{23, 14},
	Documentation: []Text{"My", "storage func", "doc"},
}

var exampleFunctionMetadataV4 = FunctionMetadataV4{
	Name:          "my function",
	Args:          []FunctionArgumentMetadata{exampleFunctionArgumentMetadata},
	Documentation: []Text{"My", "doc"},
}

var exampleEventMetadataV4 = EventMetadataV4{
	Name:          "myEvent",
	Args:          []Type{"arg1", "arg2"},
	Documentation: []Text{"My", "doc"},
}

var exampleMapTypeV4 = MapTypeV4{
	Hasher: StorageHasher{IsBlake2_256: true},
	Key:    "my key",
	Value:  "and my value",
	Linked: false,
}

var exampleDoubleMapTypeV4 = DoubleMapTypeV4{
	Hasher:     StorageHasher{IsBlake2_256: true},
	Key1:       "myKey",
	Key2:       "otherKey",
	Value:      "and a value",
	Key2Hasher: "and a hasher",
}

var exampleFunctionArgumentMetadata = FunctionArgumentMetadata{Name: "myFunctionName", Type: "myType"}

func TestMetadataV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMetadataV4)
}

func TestCallIndex_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleCallIndex)
}

func TestModuleMetadataV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleModuleMetadataV42)
}

func TestStorageFunctionMetadataV4Type_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleStorageFunctionMetadataV4Type)
}

func TestStorageFunctionMetadataV4Map_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleStorageFunctionMetadataV4Map)
}

func TestStorageFunctionMetadataV4DoubleMap_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleStorageFunctionMetadataV4DoubleMap)
}

func TestFunctionMetadataV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleFunctionMetadataV4)
}

func TestEventMetadataV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleEventMetadataV4)
}

func TestMapTypeV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMapTypeV4)
}

func TestDoubleMapTypeV4_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleDoubleMapTypeV4)
}

func TestFunctionArgumentMetadata_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleFunctionArgumentMetadata)
}

func TestFindEventNamesForEventIDV4(t *testing.T) {
	module, event, err := exampleMetadataV4.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV42.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}
