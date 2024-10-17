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

var exampleMetadataV9 = Metadata{
	MagicNumber:  0x6174656d,
	Version:      9,
	AsMetadataV9: exampleRuntimeMetadataV9,
}

var exampleRuntimeMetadataV9 = MetadataV9{
	Modules: []ModuleMetadataV8{exampleModuleMetadataV8Empty, exampleModuleMetadataV81, exampleModuleMetadataV82},
}

func TestMetadataV9_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, exampleMetadataV9)
}

func TestFindEventNamesForEventIDV9(t *testing.T) {
	module, event, err := exampleMetadataV9.FindEventNamesForEventID(EventID([2]byte{1, 0}))

	assert.NoError(t, err)
	assert.Equal(t, exampleModuleMetadataV82.Name, module)
	assert.Equal(t, exampleEventMetadataV4.Name, event)
}

func TestFindStorageEntryMetadataV9(t *testing.T) {
	_, err := exampleMetadataV9.FindStorageEntryMetadata("myStoragePrefix", "myStorageFunc2")
	assert.NoError(t, err)
}

func TestMetadataV9_ExistsModuleMetadata(t *testing.T) {
	assert.True(t, exampleMetadataV9.ExistsModuleMetadata("EmptyModule"))
	assert.False(t, exampleMetadataV9.ExistsModuleMetadata("NotExistModule"))
}
