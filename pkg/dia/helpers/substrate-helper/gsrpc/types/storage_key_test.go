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
	"fmt"
	"strings"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/hash"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/xxhash"
	"github.com/stretchr/testify/assert"
)

const (
	AlicePubKey = "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d"
)

type KeyID struct {
	Hash       Hash
	KeyPurpose uint
}

func (k KeyID) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(k.Hash); err != nil {
		return err
	}

	return encoder.PushByte(byte(k.KeyPurpose))
}

func TestCreateStorageKeyArgValidationForDoubleMapKey(t *testing.T) {
	for _, m := range []*Metadata{DecodedMetadataV14Example()} {
		fmt.Println("Testing against metadata v", m.Version)

		_, err := CreateStorageKey(m, "Keystore", "Keys")
		assert.EqualError(t, err, "Keystore:Keys is a map, therefore requires that number of arguments should exactly match number of hashers in metadata. Expected: 2, received: 0")

		_, err = CreateStorageKey(m, "Keystore", "Keys", nil)
		assert.EqualError(t, err, "Keystore:Keys is a map, therefore requires that number of arguments should exactly match number of hashers in metadata. Expected: 2, received: 0")

		_, err = CreateStorageKey(m, "Keystore", "Keys", nil, []byte{})
		assert.EqualError(t, err, "Keystore:Keys is a map, therefore requires that number of arguments should exactly match number of hashers in metadata. Expected: 2, received: 0")

		_, err = CreateStorageKey(m, "Keystore", "Keys", nil, []byte{0x01})
		assert.EqualError(t, err, "non-nil arguments cannot be preceded by nil arguments")

		_, err = CreateStorageKey(m, "Keystore", "Keys", []byte{0x01})
		assert.EqualError(t, err, "Keystore:Keys is a map, therefore requires that number of arguments should exactly match number of hashers in metadata. Expected: 2, received: 1")

		// Serialize EraIndex and AccountId
		accountID, err := NewAccountIDFromHexString(AlicePubKey)
		assert.NoError(t, err)
		encodedAccountID, err := Encode(accountID)
		assert.NoError(t, err)

		keyID := KeyID{
			Hash:       Hash{1, 2, 3, 4},
			KeyPurpose: 5,
		}
		encodedKeyID, err := Encode(keyID)
		assert.NoError(t, err)

		// Build expected answer
		expectedKeyBuilder := strings.Builder{}

		hexStr, err := Hex(xxhash.New128([]byte("Keystore")).Sum(nil))
		assert.NoError(t, err)
		expectedKeyBuilder.WriteString(hexStr)

		hexStr, err = Hex(xxhash.New128([]byte("Keys")).Sum(nil))
		assert.NoError(t, err)
		expectedKeyBuilder.WriteString(strings.TrimPrefix(hexStr, "0x"))

		accountIDHash, err := hash.NewBlake2b128Concat(nil)
		assert.NoError(t, err)

		_, err = accountIDHash.Write(encodedAccountID)
		assert.NoError(t, err)

		hexStr, err = Hex(accountIDHash.Sum(nil))
		assert.NoError(t, err)
		expectedKeyBuilder.WriteString(strings.TrimPrefix(hexStr, "0x"))

		keyIDHash, err := hash.NewBlake2b128Concat(nil)
		assert.NoError(t, err)

		_, err = keyIDHash.Write(encodedKeyID)
		assert.NoError(t, err)

		hexStr, err = Hex(keyIDHash.Sum(nil))
		assert.NoError(t, err)
		expectedKeyBuilder.WriteString(strings.TrimPrefix(hexStr, "0x"))

		key, err := CreateStorageKey(m, "Keystore", "Keys", encodedAccountID, encodedKeyID)
		assert.NoError(t, err)
		hex, err := Hex(key)
		assert.NoError(t, err)

		assert.Equal(t, expectedKeyBuilder.String(), hex)
	}
}

func TestCreateStorageKeyPlainV14(t *testing.T) {
	m := DecodedMetadataV14Example()

	key, err := CreateStorageKey(m, "Timestamp", "Now")
	assert.NoError(t, err)
	hex, err := Hex(key)
	assert.NoError(t, err)
	assert.Equal(t, "0xf0c365c3cf59d671eb72da0e7a4113c49f1f0515f462cdcf84e0f1d6045dfcbb", hex)
}

func TestCreateStorageKeyMapV14(t *testing.T) {
	m := DecodedMetadataV14Example()

	b := MustHexDecodeString(AlicePubKey)
	key, err := CreateStorageKey(m, "System", "Account", b)
	assert.NoError(t, err)
	hex, err := Hex(key)
	assert.NoError(t, err)
	assert.Equal(t, "0x26aa394eea5630e07c48ae0c9558cef7b99d880ec681799c0cf30e8886371da9de1e86a9a8c739864cf3cc5ec2bea59fd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d", hex)
}

func TestStorageKey_EncodedLength(t *testing.T) {
	AssertEncodedLength(t, []EncodedLengthAssert{
		{NewStorageKey(MustHexDecodeString("0x00")), 1},
		{NewStorageKey(MustHexDecodeString("0xab1234")), 3},
		{NewStorageKey(MustHexDecodeString("0x0001")), 2},
	})
}

func TestStorageKey_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{NewStorageKey([]byte{171, 18, 52}), MustHexDecodeString("0xab1234")},
		{NewStorageKey([]byte{}), MustHexDecodeString("0x")},
	})
}

func TestStorageKey_Decode(t *testing.T) {
	bz := []byte{12, 251, 42}
	decoded := make(StorageKey, len(bz))
	err := Decode(bz, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, StorageKey(bz), decoded)
}

func TestStorageKey_Hash(t *testing.T) {
	AssertHash(t, []HashAssert{
		{NewStorageKey([]byte{0, 42, 254}), MustHexDecodeString(
			"0x537db36f5b5970b679a28a3df8d219317d658014fb9c3d409c0c799d8ecf149d")},
		{NewStorageKey([]byte{0, 0}), MustHexDecodeString(
			"0x9ee6dfb61a2fb903df487c401663825643bb825d41695e63df8af6162ab145a6")},
	})
}

func TestStorageKey_Hex(t *testing.T) {
	AssertEncodeToHex(t, []EncodeToHexAssert{
		{NewStorageKey([]byte{0, 0, 0}), "0x000000"},
		{NewStorageKey([]byte{171, 18, 52}), "0xab1234"},
		{NewStorageKey([]byte{0, 1}), "0x0001"},
		{NewStorageKey([]byte{18, 52, 86}), "0x123456"},
	})
}

func TestStorageKey_String(t *testing.T) {
	AssertString(t, []StringAssert{
		{NewStorageKey([]byte{0, 0, 0}), "[0 0 0]"},
		{NewStorageKey([]byte{171, 18, 52}), "[171 18 52]"},
		{NewStorageKey([]byte{0, 1}), "[0 1]"},
		{NewStorageKey([]byte{18, 52, 86}), "[18 52 86]"},
	})
}

func TestStorageKey_Eq(t *testing.T) {
	AssertEq(t, []EqAssert{
		{NewStorageKey([]byte{1, 0, 0}), NewStorageKey([]byte{1, 0}), false},
		{NewStorageKey([]byte{0, 0, 1}), NewStorageKey([]byte{0, 1}), false},
		{NewStorageKey([]byte{0, 0, 0}), NewStorageKey([]byte{0, 0}), false},
		{NewStorageKey([]byte{12, 48, 255}), NewStorageKey([]byte{12, 48, 255}), true},
		{NewStorageKey([]byte{0}), NewStorageKey([]byte{0}), true},
		{NewStorageKey([]byte{1}), NewBool(true), false},
		{NewStorageKey([]byte{0}), NewBool(false), false},
	})
}

func DecodedMetadataV14Example() *Metadata {
	var metadata Metadata
	err := DecodeFromHex(MetadataV14Data, &metadata)
	if err != nil {
		panic("failed to decode the example metadata v14")
	}

	return &metadata
}
