package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

type AssetQuotationSigner struct {
	privateKey string
	domain     TypedDataDomain
	types      Types
}

func NewAssetQuotationSigner(privateKey string) *AssetQuotationSigner {
	domain := TypedDataDomain{
		Name:              "DiaData",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(1),
		VerifyingContract: common.HexToAddress("0x0000000000000000000000000000000000000000").Hex(),
	}

	types := Types{
		"Message": []Type{
			{Name: "Symbol", Type: "string"},
			{Name: "Address", Type: "address"},
			{Name: "Blockchain", Type: "string"},
			{Name: "Price", Type: "uint256"},
			{Name: "Time", Type: "uint256"},
		},
		"EIP712Domain": []Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}
	return &AssetQuotationSigner{privateKey: privateKey, domain: domain, types: types}
}

func (aqs *AssetQuotationSigner) Sign(symbol, address, blockchain string, price float64, time time.Time) (string, error) {

	message := map[string]interface{}{
		"Symbol":     symbol,
		"Address":    common.HexToAddress(address).Hex(),
		"Blockchain": blockchain,
		"Price":      big.NewInt(int64(price * 100000000)).String(),
		"Time":       big.NewInt(time.Unix()).String(),
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

func hexToPrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	// Decode the hex string into a byte slice
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}

	// Parse the byte slice as a big integer
	keyBigInt := new(big.Int).SetBytes(keyBytes)

	// Generate a new secp256k1 elliptic curve
	curve := elliptic.P256()

	// Convert the big integer to a private key using the elliptic curve
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = curve
	privateKey.D = keyBigInt
	privateKey.PublicKey.X, privateKey.PublicKey.Y = curve.ScalarBaseMult(keyBytes)

	// Verify that the private key is valid
	if !curve.IsOnCurve(privateKey.PublicKey.X, privateKey.PublicKey.Y) {
		return nil, fmt.Errorf("invalid private key")
	}

	return privateKey, nil
}
