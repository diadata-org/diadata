package stackshelper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const c32table = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
const hextable = "0123456789abcdef"

func c32address(version int, hash160hex string) (string, error) {
	match, err := regexp.MatchString("^[0-9a-fA-F]{40}$", hash160hex)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("not a hash160 hex string")
	}

	c32str, err := c32checkEncode(version, hash160hex)
	if err != nil {
		return "", err
	}
	return "S" + c32str, nil
}

func c32checkEncode(version int, data string) (string, error) {
	if version < 0 || version >= 32 {
		return "", errors.New("invalid version (must be between 0 and 31)")
	}

	data = strings.ToLower(data)
	if len(data)%2 != 0 {
		data = "0" + data
	}
	versionHex := fmt.Sprintf("%02x", version)

	checksum, err := c32checksum(versionHex + data)
	if err != nil {
		return "", err
	}

	c32str, err := c32encode(data + checksum)
	if err != nil {
		return "", err
	}
	return string(c32table[version]) + c32str, nil
}

func c32checksum(dataHex string) (string, error) {
	bytes, err := hex.DecodeString(dataHex)
	if err != nil {
		return "", err
	}

	inner := sha256.Sum256(bytes)
	checksum := sha256.Sum256(inner[:])
	return hex.EncodeToString(checksum[0:4]), nil
}

func c32encode(inputHex string) (string, error) {
	if len(inputHex)%2 != 0 {
		inputHex = "0" + inputHex
	}
	inputHex = strings.ToLower(inputHex)

	inputBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		return "", err
	}

	res := make([]byte, 0)
	carry := 0

	for i := len(inputHex) - 1; i >= 0; i-- {
		if carry < 4 {
			currentCode := strings.IndexByte(hextable, inputHex[i]) >> carry
			nextCode := 0
			if i != 0 {
				nextCode = strings.IndexByte(hextable, inputHex[i-1])
			}

			nextBits := 1 + carry
			nextLowBits := nextCode % (1 << nextBits) << (5 - nextBits)
			curC32Digit := c32table[currentCode+nextLowBits]

			res = append([]byte{curC32Digit}, res...)
			carry = nextBits
		} else {
			carry = 0
		}
	}

	c32LeadingZeros := 0
	for _, d := range res {
		if d != '0' {
			break
		} else {
			c32LeadingZeros++
		}
	}
	res = res[c32LeadingZeros:]

	numLeadingZeroBytesInHex := 0
	for _, b := range inputBytes {
		if b != 0 {
			break
		}
		numLeadingZeroBytesInHex++
	}

	for i := 0; i < numLeadingZeroBytesInHex; i++ {
		res = append([]byte{c32table[0]}, res...)
	}
	return string(res), nil
}
