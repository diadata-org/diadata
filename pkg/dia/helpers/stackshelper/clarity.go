package stackshelper

import (
	"encoding/binary"
	"errors"
	"math/big"
	"sort"
)

const (
	uintCV       = 0x01
	boolTrueCV   = 0x03
	boolFalseCV  = 0x04
	optionNoneCV = 0x09
	optionSomeCV = 0x0a
	tupleCV      = 0x0c
)

type CVTuple map[string][]byte

// DeserializeCVUint converts a clarity 128-bit uint value into a `big.Int`.
func DeserializeCVUint(src []byte) (*big.Int, error) {
	if src[0] != uintCV {
		err := errors.New("value is not a CV tuple")
		return nil, err
	}

	value := new(big.Int)
	value.SetBytes(src[1:])
	return value, nil
}

// SerializeCVTuple converts a clarity value tuple into its binary representation
// that can be used to call stacks smart contract functions.
func SerializeCVTuple(tuple CVTuple) []byte {
	result := make([]byte, 5)
	result[0] = tupleCV

	length := len(tuple)
	binary.BigEndian.PutUint32(result[1:], uint32(length))

	i := 0
	keys := make([]string, length)
	for k := range tuple {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		key := serializeLPString(k)
		entry := append(key, tuple[k]...)
		result = append(result, entry...)
	}
	return result
}

// DeserializeCVTuple converts binary representation of a clarity value tuple
// into a `CVTuple` map.
// IMPORTANT: this function only supports uint and bool values at the moment,
// therefore it should NOT be used as a complete solution to deserialize any
// arbitrary clarity tuple.
func DeserializeCVTuple(src []byte) (CVTuple, error) {
	if src[0] != tupleCV {
		err := errors.New("value is not a CV tuple")
		return nil, err
	}

	length := binary.BigEndian.Uint32(src[1:5])
	result := make(CVTuple, length)

	offset := 5

	for i := 0; i < int(length); i++ {
		key, keySize := deserializeLPString(src[offset:])
		offset += keySize + 1

		valueSize := 1
		if src[offset] == uintCV {
			valueSize += 16
		}

		result[key] = src[offset : offset+valueSize]
		offset += valueSize
	}

	return result, nil
}

func deserializeCVOption(src []byte) ([]byte, bool) {
	switch src[0] {
	case optionNoneCV:
		return nil, false
	case optionSomeCV:
		return src[1:], true
	default:
		return nil, false
	}
}

func serializeLPString(val string) []byte {
	content := []byte(val)
	size := byte(len(content))
	return append([]byte{size}, content...)
}

func deserializeLPString(val []byte) (string, int) {
	size := int(val[0])
	return string(val[1 : size+1]), size
}
