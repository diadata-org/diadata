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

var exampleMetadataV8 = Metadata{
	MagicNumber:  0x6174656d,
	Version:      8,
	AsMetadataV8: exampleRuntimeMetadataV8,
}

var exampleRuntimeMetadataV8 = MetadataV8{
	Modules: []ModuleMetadataV8{exampleModuleMetadataV8Empty, exampleModuleMetadataV81, exampleModuleMetadataV82},
}

var exampleModuleMetadataV8Empty = ModuleMetadataV8{
	Name:       "EmptyModule",
	HasStorage: false,
	Storage:    StorageMetadata{},
	HasCalls:   false,
	Calls:      nil,
	HasEvents:  false,
	Events:     nil,
	Constants:  nil,
	Errors:     nil,
}

var exampleModuleMetadataV81 = ModuleMetadataV8{
	Name:       "Module1",
	HasStorage: true,
	Storage:    exampleStorageMetadata,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
}

var exampleModuleMetadataV82 = ModuleMetadataV8{
	Name:       "Module2",
	HasStorage: true,
	Storage:    exampleStorageMetadata,
	HasCalls:   true,
	Calls:      []FunctionMetadataV4{exampleFunctionMetadataV4},
	HasEvents:  true,
	Events:     []EventMetadataV4{exampleEventMetadataV4},
	Constants:  []ModuleConstantMetadataV6{exampleModuleConstantMetadataV6},
	Errors:     []ErrorMetadataV8{exampleErrorMetadataV8},
}

var exampleErrorMetadataV8 = ErrorMetadataV8{
	Name:          "My Error",
	Documentation: []Text{"Error", "docs"},
}

func TestMetadataV8_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMetadataV8)
}

func TestFindEventNamesForEventIDV8(t *testing.T) {
	module, event, err := exampleMetadataV8.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV82.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestFindStorageEntryMetadataV8(t *testing.T) {
	_, err := exampleMetadataV8.FindStorageEntryMetadata("myStoragePrefix", "myStorageFunc2")
	assert.NoError(t, err)
}
