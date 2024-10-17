// Copyright 2018 Jsgenesis
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

package scale

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func hexify(bytes []byte) string {
	res := make([]string, len(bytes))
	for i, b := range bytes {
		res[i] = fmt.Sprintf("%02x", b)
	}
	return strings.Join(res, " ")
}

func encodeToBytes(t *testing.T, value interface{}) []byte {
	var buffer = bytes.Buffer{}
	err := Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	return buffer.Bytes()
}

func assertRoundtrip(t *testing.T, value interface{}) {
	var buffer = bytes.Buffer{}
	err := Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	target := reflect.New(reflect.TypeOf(value))
	err = Decoder{reader: &buffer}.Decode(target.Interface())
	assert.NoError(t, err)
	assertEqual(t, target.Elem().Interface(), value)
}

type CustomBool bool

func (c CustomBool) Encode(encoder Encoder) error {
	var encoded byte
	if c {
		encoded = 0x05
	} else {
		encoded = 0x10
	}
	err := encoder.PushByte(encoded)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomBool) Decode(decoder Decoder) error {
	b, _ := decoder.ReadOneByte()
	switch b {
	case 0x05:
		*c = true
	case 0x10:
		*c = false
	default:
		return fmt.Errorf("unknown byte prefix for encoded CustomBool: %d", b)
	}
	return nil
}

func TestTypeImplementsEncodeableDecodeableEncodedAsExpected(t *testing.T) {
	value := CustomBool(true)
	assertRoundtrip(t, value)

	var buffer = bytes.Buffer{}
	err := Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x05}, buffer.Bytes())

	var decoded CustomBool
	err = Decoder{reader: &buffer}.Decode(&decoded)
	assert.NoError(t, err)
	assert.Equal(t, CustomBool(true), decoded)
}

type CustomBytes []byte

func (c CustomBytes) Encode(encoder Encoder) error {
	for i := 0; i < len(c); i++ {
		err := encoder.PushByte(^c[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CustomBytes) Decode(decoder Decoder) error {
	for i := 0; i < len(*c); i++ {
		b, err := decoder.ReadOneByte()
		if err != nil {
			return err
		}
		(*c)[i] = ^b
	}
	return nil
}

func TestTypeImplementsEncodeableDecodeableSliceEncodedAsExpected(t *testing.T) {
	value := CustomBytes([]byte{0x01, 0x23, 0xf2})
	// assertRoundtrip(t, value)

	var buffer = bytes.Buffer{}
	err := Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0xfe, 0xdc, 0xd}, buffer.Bytes())

	decoded := make(CustomBytes, len(value))
	err = Decoder{reader: &buffer}.Decode(&decoded)
	assert.NoError(t, err)
	assert.Equal(t, value, decoded)
}

func TestSliceOfBytesEncodedAsExpected(t *testing.T) {
	value := []byte{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "28 00 01 01 02 03 05 08 0d 15 22")
}

func TestArrayOfBytesEncodedAsExpected(t *testing.T) {
	value := [10]byte{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "00 01 01 02 03 05 08 0d 15 22")
}

func TestArrayCannotBeDecodedIntoIncompatible(t *testing.T) {
	value := [3]byte{255, 254, 253}
	value2 := [5]byte{1, 2, 3, 4, 5}
	value3 := [1]byte{42}

	var buffer = bytes.Buffer{}
	err := Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	err = Decoder{reader: &buffer}.Decode(&value2)
	assert.EqualError(t, err, "expected more bytes, but could not decode any more")
	buffer.Reset()
	err = Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	err = Decoder{reader: &buffer}.Decode(&value3)
	assert.NoError(t, err)
	assert.Equal(t, [1]byte{255}, value3)
	buffer.Reset()
	err = Encoder{writer: &buffer}.Encode(value)
	assert.NoError(t, err)
	err = Decoder{reader: &buffer}.Decode(&value)
	assert.NoError(t, err)
}

func TestSliceOfInt16EncodedAsExpected(t *testing.T) {
	value := []int16{0, 1, -1, 2, -2, 3, -3}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "1c 00 00 01 00 ff ff 02 00 fe ff 03 00 fd ff")
}

func TestStructFieldByFieldEncoding(t *testing.T) {
	value := struct {
		A string
		B int16
		C bool
	}{"my", 3, true}
	assertRoundtrip(t, value)
}

// OptionInt8 is an example implementation of an "Option" type, mirroring Option<u8> in Rust version.
// Since Go does not support generics, one has to define such types manually.
// See below for ParityEncode / ParityDecode implementations.
type OptionInt8 struct {
	hasValue bool
	value    int8
}

func (o OptionInt8) Encode(encoder Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionInt8) Decode(decoder Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

func TestSliceOfOptionInt8EncodedAsExpected(t *testing.T) {
	value := []OptionInt8{{true, 1}, {true, -1}, {false, 0}}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "0c 01 01 01 ff 00")
}

func TestSliceOfOptionBoolEncodedAsExpected(t *testing.T) {
	value := []OptionBool{NewOptionBool(true), NewOptionBool(false), NewOptionBoolEmpty()}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "0c 01 02 00")
}

func TestSliceOfStringEncodedAsExpected(t *testing.T) {
	value := []string{
		"Hamlet",
		"Война и мир",
		"三国演义",
		"أَلْف لَيْلَة وَلَيْلَة‎"}
	assertRoundtrip(t, value)
	assertEqual(t, hexify(encodeToBytes(t, value)), "10 18 48 61 6d 6c 65 74 50 d0 92 d0 be d0 b9 d0 bd d0 b0 20 d0 "+
		"b8 20 d0 bc d0 b8 d1 80 30 e4 b8 89 e5 9b bd e6 bc 94 e4 b9 89 bc d8 a3 d9 8e d9 84 d9 92 "+
		"d9 81 20 d9 84 d9 8e d9 8a d9 92 d9 84 d9 8e d8 a9 20 d9 88 d9 8e d9 84 d9 8e d9 8a d9 92 "+
		"d9 84 d9 8e d8 a9 e2 80 8e")
}

func TestCompactIntegersEncodedAsExpected(t *testing.T) {
	tests := map[uint64]string{
		0:              "00",
		63:             "fc",
		64:             "01 01",
		16383:          "fd ff",
		16384:          "02 00 01 00",
		1073741823:     "fe ff ff ff",
		1073741824:     "03 00 00 00 40",
		1<<32 - 1:      "03 ff ff ff ff",
		1 << 32:        "07 00 00 00 00 01",
		1 << 40:        "0b 00 00 00 00 00 01",
		1 << 48:        "0f 00 00 00 00 00 00 01",
		1<<56 - 1:      "0f ff ff ff ff ff ff ff",
		1 << 56:        "13 00 00 00 00 00 00 00 01",
		math.MaxUint64: "13 ff ff ff ff ff ff ff ff"}
	for value, expectedHex := range tests {
		var buffer = bytes.Buffer{}
		valueBig := new(big.Int).SetUint64(value)
		err := Encoder{writer: &buffer}.EncodeUintCompact(*valueBig)
		assert.NoError(t, err)
		assertEqual(t, hexify(buffer.Bytes()), expectedHex)
		decoded, _ := Decoder{reader: &buffer}.DecodeUintCompact()
		assertEqual(t, decoded, big.NewInt(0).SetUint64(value))
	}
}
