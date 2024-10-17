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

package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlake2_128Concat(t *testing.T) {
	h, err := NewBlake2b128Concat(nil)
	assert.NoError(t, err)
	n, err := h.Write([]byte("abc"))
	assert.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, []byte{
		0xcf, 0x4a, 0xb7, 0x91, 0xc6, 0x2b, 0x8d, 0x2b, 0x21, 0x9, 0xc9, 0x2, 0x75, 0x28, 0x78, 0x16, 0x61, 0x62, 0x63,
	}, h.Sum(nil))
	assert.Equal(t, 128, h.BlockSize())
	assert.Equal(t, 19, h.Size())
	h.Reset()
	assert.Equal(t, 16, h.Size())
}

func TestBlake2b_512(t *testing.T) {
	h, err := NewBlake2b512(nil)
	assert.NoError(t, err)
	n, err := h.Write([]byte("abc"))
	assert.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, []byte{
		0xba, 0x80, 0xa5, 0x3f, 0x98, 0x1c, 0x4d, 0xd, 0x6a, 0x27, 0x97, 0xb6, 0x9f, 0x12, 0xf6, 0xe9, 0x4c, 0x21, 0x2f,
		0x14, 0x68, 0x5a, 0xc4, 0xb7, 0x4b, 0x12, 0xbb, 0x6f, 0xdb, 0xff, 0xa2, 0xd1, 0x7d, 0x87, 0xc5, 0x39, 0x2a, 0xab,
		0x79, 0x2d, 0xc2, 0x52, 0xd5, 0xde, 0x45, 0x33, 0xcc, 0x95, 0x18, 0xd3, 0x8a, 0xa8, 0xdb, 0xf1, 0x92, 0x5a, 0xb9,
		0x23, 0x86, 0xed, 0xd4, 0x0, 0x99, 0x23,
	}, h.Sum(nil))
	assert.Equal(t, 128, h.BlockSize())
	assert.Equal(t, 64, h.Size())
}
