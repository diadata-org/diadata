package stackshelper

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
	"sort"
)

const (
	clarityIntByteSize       = 16
	clarityPrincipalByteSize = 20
)

const (
	uintCV            = 0x01
	boolTrueCV        = 0x03
	boolFalseCV       = 0x04
	principalStandard = 0x05
	principalContract = 0x06
	responseOkCV      = 0x07
	responseErrCV     = 0x08
	optionNoneCV      = 0x09
	optionSomeCV      = 0x0a
	tupleCV           = 0x0c
	stringASCII       = 0x0d
	stringUTF8        = 0x0e
)

type CVTuple map[string][]byte

// SerializeCVUint converts a `big.Int` instance into Clarity value
// binary representation.
func SerializeCVUint(value *big.Int) []byte {
	result := make([]byte, clarityIntByteSize+1)
	result[0] = uintCV
	value.FillBytes(result[1:])
	return result
}

// DeserializeCVUint converts a clarity 128-bit uint value into a `big.Int`.
func DeserializeCVUint(src []byte) (*big.Int, error) {
	if src[0] != uintCV {
		err := errors.New("value is not a CV uint")
		return nil, err
	}

	value := new(big.Int)
	value.SetBytes(src[1:])
	return value, nil
}

func DeserializeCVPrincipal(src []byte) (string, error) {
	switch src[0] {
	case principalStandard:
		version, hash160 := deserializeAddress(src[1:])
		return c32address(version, hash160)
	case principalContract:
		version, hash160 := deserializeAddress(src[1:])
		cAddress, err := c32address(version, hash160)
		if err != nil {
			return "", err
		}
		contractName, _ := deserializeLPString(src[clarityPrincipalByteSize+2:])
		return cAddress + "." + contractName, nil
	default:
		return "", errors.New("value is not a CV principal")
	}
}

func DeserializeCVResponse(src []byte) ([]byte, bool) {
	switch src[0] {
	case responseOkCV:
		return src[1:], true
	case responseErrCV:
		return src[1:], false
	default:
		return nil, false
	}
}

func DeserializeCVString(src []byte) (string, error) {
	if src[0] != stringASCII && src[0] != stringUTF8 {
		err := errors.New("value is not an ASCII/UTF8 encoded string")
		return "", err
	}

	size := readClarityTypeSize(src[1:])
	start := 5
	end := start + int(size)
	return string(src[start:end]), nil
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
// IMPORTANT: this function supports a limited amount of Clarity types at the
// moment, therefore it should NOT be used as a complete solution to
// deserialize any arbitrary Clarity tuple.
func DeserializeCVTuple(src []byte) (CVTuple, error) {
	if src[0] != tupleCV {
		err := errors.New("value is not a CV tuple")
		return nil, err
	}

	length := readClarityTypeSize(src[1:])
	result := make(CVTuple, length)
	offset := 5

	for i := 0; i < int(length); i++ {
		key, keySize := deserializeLPString(src[offset:])
		offset += keySize + 1

		valueSize := 1

		switch src[offset] {
		case uintCV:
			valueSize += clarityIntByteSize
		case principalStandard:
			valueSize += clarityPrincipalByteSize + 1
		case principalContract:
			principalSize := clarityPrincipalByteSize + 1
			valueSize += principalSize + int(src[offset+principalSize+1]) + 1
		case tupleCV:
			tuple, err := DeserializeCVTuple(src[offset:])
			if err != nil {
				return nil, err
			}
			valueSize += 4

			for k, v := range tuple {
				entrySize := len(serializeLPString(k)) + len(v)
				valueSize += entrySize
			}
		case stringASCII, stringUTF8:
			size := readClarityTypeSize(src[offset+1:])
			valueSize += 4 + int(size)
		default:
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

func deserializeAddress(src []byte) (int, string) {
	version := int(src[0])
	hash160 := hex.EncodeToString(src[1 : clarityPrincipalByteSize+1])
	return version, hash160
}

func readClarityTypeSize(src []byte) uint32 {
	return binary.BigEndian.Uint32(src[0:4])
}
