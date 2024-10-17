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
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/stretchr/testify/assert"
)

var exampleRuntimeVersion = RuntimeVersion{
	APIs:             []RuntimeVersionAPI{exampleRuntimeVersionAPI},
	AuthoringVersion: 13,
	ImplName:         "My impl",
	ImplVersion:      21,
	SpecName:         "My spec",
	SpecVersion:      39,
}

var exampleRuntimeVersionAPI = RuntimeVersionAPI{
	APIID:   "0x37e397fc7c91f5e4",
	Version: 23,
}

func TestRuntimeVersion_Encode_Decode(t *testing.T) {
	enc, err := Encode(exampleRuntimeVersion)
	assert.NoError(t, err)

	var output RuntimeVersion
	err = Decode(enc, &output)
	assert.NoError(t, err)

	assert.Equal(t, exampleRuntimeVersion, output)
}

func TestRuntimeVersionAPI_Encode_Decode(t *testing.T) {
	enc, err := Encode(exampleRuntimeVersionAPI)
	assert.NoError(t, err)

	var output RuntimeVersionAPI
	err = Decode(enc, &output)
	assert.NoError(t, err)

	assert.Equal(t, exampleRuntimeVersionAPI, output)
}

func TestRuntimeVersionAPI_JSONMarshalUnmarshal(t *testing.T) {
	r := exampleRuntimeVersionAPI
	AssertJSONRoundTrip(t, &r)
}
