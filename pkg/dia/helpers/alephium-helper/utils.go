package alephiumhelper

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcutil/base58"
)

const (
	AddressTypeP2PKH  = 0x00
	AddressTypeP2MPKH = 0x01
	AddressTypeP2SH   = 0x02
	AddressTypeP2C    = 0x03

	TotalNumberOfGroups = 4
)

func DecodeHex(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	str := string(bytes)
	return str, nil
}

func AddressFromTokenId(tokenId string) (string, error) {
	hash, err := hex.DecodeString(tokenId)
	if err != nil {
		return "", err
	}
	addressType := []byte{AddressTypeP2C}

	addressType = append(addressType, hash...)

	result := base58.Encode(addressType)

	return result, nil
}

func groupOfAddress(address string) (uint32, error) {
	decoded, err := decodeAndValidateAddress(address)
	if err != nil {
		return 0, err
	}
	addressType := decoded[0]
	addressBody := decoded[1:]

	if addressType == AddressTypeP2PKH {
		return groupOfP2pkhAddress(addressBody), nil
	} else if addressType == AddressTypeP2MPKH {
		return groupOfP2mpkhAddress(addressBody), nil
	} else if addressType == AddressTypeP2SH {
		return groupOfP2shAddress(addressBody), nil
	} else {
		id, err := groupFromAddress(address)
		return uint32(id[len(id)-1]), err
	}
}

func djb2(bytes []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(bytes); i++ {
		hash = ((hash << 5) + hash) + uint32(bytes[i]&0xff)
	}
	return hash
}

func xorByte(intValue uint32) uint8 {
	byte0 := (intValue >> 24) & 0xff
	byte1 := (intValue >> 16) & 0xff
	byte2 := (intValue >> 8) & 0xff
	byte3 := intValue & 0xff
	return uint8(byte0^byte1^byte2^byte3) & 0xff
}

func groupOfAddressBytes(bytes []byte) uint32 {
	hint := djb2(bytes) | 1
	hash := xorByte(hint)
	group := hash % TotalNumberOfGroups
	return uint32(group)
}

func groupOfP2pkhAddress(address []byte) uint32 {
	return groupOfAddressBytes(address)
}

func groupOfP2mpkhAddress(address []byte) uint32 {
	return groupOfAddressBytes(address[1:33])
}

func groupOfP2shAddress(address []byte) uint32 {
	return groupOfAddressBytes(address)
}

func groupFromAddress(address string) ([]byte, error) {
	return idFromAddress(address)
}

func idFromAddress(address string) ([]byte, error) {
	decoded := base58.Decode(address)
	if len(decoded) == 0 {
		return nil, errors.New("address string is empty")
	}
	addressType := decoded[0]
	addressBody := decoded[1:]

	if addressType == AddressTypeP2C {
		return addressBody, nil
	} else {
		return nil, errors.New("invalid contract address type")
	}
}

func decodeAndValidateAddress(address string) ([]byte, error) {
	decoded := base58.Decode(address)
	if len(decoded) == 0 {
		return nil, errors.New("decodeAndValidateAddress.Address is empty")
	}
	addressType := decoded[0]

	if addressType == AddressTypeP2MPKH && ((len(decoded)-3)%32 == 0) {
		return decoded, nil
	}
	if (addressType == AddressTypeP2PKH || addressType == AddressTypeP2SH || addressType == AddressTypeP2C) && (len(decoded) == 33) {
		return decoded, nil
	}
	return nil, errors.New("invalid address")
}
