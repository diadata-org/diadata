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
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/hash"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	//verify checksum from ss58
	contextPrefix := []byte("SS58PRE")
	ss58d := base58.Decode("4ecQzsMCwbJXu6Cad597T7gZx1MTZWQi8jZZC2DmsQq72knj")
	assert.Equal(t, uint8(36), ss58d[0]) // Centrifuge network version check
	noSum := ss58d[:len(ss58d)-2]
	all := append(contextPrefix, noSum...)
	checksum, err := hash.NewBlake2b512(nil)
	assert.NoError(t, err)
	checksum.Write(all)
	res := checksum.Sum(nil)
	assert.Equal(t, ss58d[len(ss58d)-2:], res[:2]) // Verified checksum
}

var (
	addressFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(a *Address, c fuzz.Continue) {
			if c.RandBool() {
				a.IsAccountID = true
				c.Fuzz(&a.AsAccountID)

				return
			}

			a.IsAccountIndex = true
			c.Fuzz(&a.AsAccountIndex)
		}),
	}
)

func newTestAddress() Address {
	addr, err := NewAddressFromAccountID(testAccountIDBytes)

	if err != nil {
		panic(fmt.Errorf("couldn't create test address: %w", err))
	}

	return addr
}

func TestAddress_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Address](t, 100, addressFuzzOpts...)
	AssertDecodeNilData[Address](t)
	AssertEncodeEmptyObj[Address](t, 1)
}

func TestAddress_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{newTestAddress(),
			MustHexDecodeString("0xff0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f0b"),
		},
		{NewAddressFromAccountIndex(binary.BigEndian.Uint32([]byte{
			17, 18, 19, 20,
		})), []byte{
			253, 20, 19, 18, 17, // order is reversed because scale uses little endian
		}},
		{NewAddressFromAccountIndex(uint32(binary.BigEndian.Uint16([]byte{
			21, 22,
		}))), []byte{
			252, 22, 21, // order is reversed because scale uses little endian
		}},
		{NewAddressFromAccountIndex(uint32(23)), []byte{23}},
	})
}

func TestAddress_EncodeWithOptions(t *testing.T) {
	SetSerDeOptions(SerDeOptions{NoPalletIndices: true})
	defer SetSerDeOptions(SerDeOptions{NoPalletIndices: false})
	AssertEncode(t, []EncodingAssert{
		{newTestAddress(),
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f0b"),
		},
		{NewAddressFromAccountIndex(binary.BigEndian.Uint32([]byte{
			17, 18, 19, 20,
		})), []byte{
			253, 20, 19, 18, 17, // order is reversed because scale uses little endian
		}},
	})
}

func TestAddress_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0xff0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f0b"),
			newTestAddress(),
		},
		{[]byte{
			253, 20, 19, 18, 17, // order is reversed because scale uses little endian
		}, NewAddressFromAccountIndex(binary.BigEndian.Uint32([]byte{
			17, 18, 19, 20,
		}))},
		{[]byte{
			252, 22, 21, // order is reversed because scale uses little endian
		}, NewAddressFromAccountIndex(uint32(binary.BigEndian.Uint16([]byte{
			21, 22,
		})))},
		{[]byte{23}, NewAddressFromAccountIndex(uint32(23))},
	})

	_, err := NewAddressFromHexAccountID("zzz")
	assert.NotNil(t, err)

	a := new(Address)
	err = a.Decode(*scale.NewDecoder(bytes.NewReader([]byte{0xfe})))
	assert.NotNil(t, err)
}

func TestAddress_DecodeWithOptions(t *testing.T) {
	SetSerDeOptions(SerDeOptions{NoPalletIndices: true})
	defer SetSerDeOptions(SerDeOptions{NoPalletIndices: false})
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f0b"),
			newTestAddress(),
		},
	})
}
