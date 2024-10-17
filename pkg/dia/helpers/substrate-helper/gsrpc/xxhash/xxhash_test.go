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

package xxhash_test

import (
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/xxhash"
	"github.com/stretchr/testify/assert"
)

func Test64(t *testing.T) {
	h := New64([]byte("abc"))
	assert.Equal(t, []byte{0x99, 0x09, 0x77, 0xad, 0xf5, 0x2c, 0xbc, 0x44}, h.Sum(nil))
}

func Test64Concat(t *testing.T) {
	h := New64Concat([]byte("abc"))
	assert.Equal(t, []byte{0x99, 0x09, 0x77, 0xad, 0xf5, 0x2c, 0xbc, 0x44, 0x61, 0x62, 0x63}, h.Sum(nil))
}

func Test128(t *testing.T) {
	h := New128([]byte("abc"))
	assert.Equal(t, []byte{
		0x99, 0x09, 0x77, 0xad, 0xf5, 0x2c, 0xbc, 0x44, 0x08, 0x89, 0x32, 0x99, 0x81, 0xca, 0xa9, 0xbe,
	}, h.Sum(nil))
}

func Test256(t *testing.T) {
	h := New256([]byte("abc"))
	assert.Equal(t, []byte{
		0x99, 0x09, 0x77, 0xad, 0xf5, 0x2c, 0xbc, 0x44, 0x08, 0x89, 0x32, 0x99, 0x81, 0xca, 0xa9, 0xbe,
		0xf7, 0xda, 0x57, 0x70, 0xb2, 0xb8, 0xa0, 0x53, 0x03, 0xb7, 0x5d, 0x95, 0x36, 0x0d, 0xd6, 0x2b,
	}, h.Sum(nil))
}

func Test64OtherMethods(t *testing.T) {
	h := New64(nil)

	assert.Equal(t, 8, h.Size())
	assert.Equal(t, []byte{0x99, 0xe9, 0xd8, 0x51, 0x37, 0xdb, 0x46, 0xef}, h.Sum(nil))

	n, err := h.Write([]byte("ab"))
	assert.NoError(t, err)
	assert.Equal(t, 2, n)

	assert.Equal(t, []byte{0x61, 0x4a, 0xd0, 0x92, 0xca, 0x8, 0xf7, 0x65}, h.Sum(nil))

	n, err = h.Write([]byte("c"))
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	assert.Equal(t, 8, h.Size())
	assert.Equal(t, []byte{0x99, 0x09, 0x77, 0xad, 0xf5, 0x2c, 0xbc, 0x44}, h.Sum(nil))
	assert.Equal(t, 64, h.BlockSize())

	h.Reset()
	assert.Equal(t, 64, h.BlockSize())
	assert.Equal(t, 8, h.Size())
	assert.Equal(t, []byte{0x99, 0xe9, 0xd8, 0x51, 0x37, 0xdb, 0x46, 0xef}, h.Sum(nil))
}
