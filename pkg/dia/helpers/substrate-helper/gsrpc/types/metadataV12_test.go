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

var exampleMetadataV12 = Metadata{
	MagicNumber:   0x6174656d,
	Version:       12,
	AsMetadataV12: exampleRuntimeMetadataV12,
}

var exampleRuntimeMetadataV12 = MetadataV12{
	Modules: []ModuleMetadataV12{exampleModuleMetadataV12Empty, exampleModuleMetadataV121, exampleModuleMetadataV122},
}

var exampleModuleMetadataV12Empty = ModuleMetadataV12{
	ModuleMetadataV10: ModuleMetadataV10{
		Name:       "EmptyModule",
		HasStorage: false,
		Storage:    StorageMetadataV10{},
		HasCalls:   false,
		Calls:      nil,
		HasEvents:  false,
		Events:     nil,
		Constants:  nil,
		Errors:     nil,
	},
	Index: 0,
}

var exampleModuleMetadataV121 = ModuleMetadataV12{
	ModuleMetadataV10: ModuleMetadataV10{
		Name:       "Module1",
		HasStorage: true,
		Storage:    exampleStorageMetadataV10,
		HasCalls:   true,
		Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
		HasEvents:  true,
		Events:     []EventMetadataV4{exampleEventMetadataV4},
		Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
		Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
	},
	Index: 1,
}

var exampleModuleMetadataV122 = ModuleMetadataV12{
	ModuleMetadataV10: ModuleMetadataV10{
		Name:       "Module2",
		HasStorage: true,
		Storage:    exampleStorageMetadataV10,
		HasCalls:   true,
		Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
		HasEvents:  true,
		Events:     []EventMetadataV4{exampleEventMetadataV4},
		Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
		Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
	},
	Index: 2,
}

func TestMetadataV12_ExistsModuleMetadata(t *testing.T) {
	assert.True(t, exampleMetadataV12.ExistsModuleMetadata("EmptyModule"))
	assert.False(t, exampleMetadataV12.ExistsModuleMetadata("NotExistModule"))
}

func TestMetadataV12_FindEventNamesForEventID(t *testing.T) {
	module, event, err := exampleMetadataV12.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV121.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestMetadataV12_FindEventNamesForUnknownModule(t *testing.T) {
	_, _, err := exampleMetadataV12.FindEventNamesForEventID(EventID([2]byte{1, 18}))

	assert.Error(t, err)
}

func TestMetadataV12_TestFindStorageEntryMetadata(t *testing.T) {
	_, err := exampleMetadataV12.FindStorageEntryMetadata("myStoragePrefix", "myStorageFunc2")
	assert.NoError(t, err)
}

func TestMetadataV12_TestFindCallIndex(t *testing.T) {
	callIndex, err := exampleMetadataV12.FindCallIndex("Module2.my function")
	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV122.Index, callIndex.SectionIndex)
	assert.Equal(t, uint8(0), callIndex.MethodIndex)
}

func TestMetadataV12_TestFindCallIndexWithUnknownModule(t *testing.T) {
	_, err := exampleMetadataV12.FindCallIndex("UnknownModule.my function")
	assert.Error(t, err)
}

func TestMetadataV12_TestFindCallIndexWithUnknownFunction(t *testing.T) {
	_, err := exampleMetadataV12.FindCallIndex("Module2.unknownFunction")
	assert.Error(t, err)
}
