package utils

import (
	"bytes"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

type RandomnessSigner struct {
	privateKey string
	domain     TypedDataDomain
	types      Types
}

func NewRandomnessSigner(privateKey string) *RandomnessSigner {
	domain := TypedDataDomain{
		Name:              "DiaData",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(1),
		VerifyingContract: common.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
	}

	types := Types{
		"Message": []Type{
			{Name: "randomness", Type: "string"},
			{Name: "round", Type: "string"},
		},
		"EIP712Domain": []Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}
	return &RandomnessSigner{privateKey: privateKey, domain: domain, types: types}
}

func (aqs *RandomnessSigner) Sign(randomness string, round int) (string, error) {

	message := map[string]interface{}{
		"randomness": randomness,
		"round":      strconv.Itoa(round),
	}

	typedData := TypedData{
		Types:       aqs.types,
		PrimaryType: "Message",
		Domain:      aqs.domain,
		Message:     message,
	}

	typedDataHash, _, err := TypedDataAndHash(typedData)
	if err != nil {
		return "", err
	}

	pk, err := hexToPrivateKey(aqs.privateKey)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(typedDataHash, pk)
	if err != nil {
		return "", err

	}

	var buf bytes.Buffer
	buf.Write(signature)
	// buf.WriteByte(1) // recovery ID
	sigData := hexutil.Encode(buf.Bytes())
	// fmt.Println("sigdata", sigData)

	return sigData, nil
}
