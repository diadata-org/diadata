package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	// "github.com/storyicon/sigverify"
)

func NewKeyPair() (publickey, privatekey string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")

	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, hexutil.Encode(privateKeyBytes)[:]

}

func GetSigner(chainID, creator, oracleaddress, signed string) (common.Address, error) {
	typedatastring := `{
		"types": {
			"EIP712Domain": [{
				"name": "name",
				"type": "string"
			}, {
				"name": "version",
				"type": "string"
			}, {
				"name": "chainId",
				"type": "uint256"
			}, {
				"name": "verifyingContract",
				"type": "address"
			}],
			"Oracle": [{
				"name": "contents",
				"type": "string"
			}, {
				"name": "creator",
				"type": "address"
			}, {
				"name": "oracleaddress",
				"type": "address"
			}]
		},
		"primaryType": "Oracle",
		"domain": {
			"name": "Oracle Builder",
			"version": "1",
			"chainId": "` + chainID + `",
			"verifyingContract": "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC",
			"salt": ""
		},
		"message": {
			"contents": "Verify its your address to call oracle builder",
			"creator": "` + creator + `",
			"oracleaddress": "` + oracleaddress + `"
		}
	}`
	var typedData TypedData

	if err := json.Unmarshal([]byte(typedatastring), &typedData); err != nil {
	}

	signature, _ := HexDecode(signed)

	return VerifyTypedData("Oracle", typedData.Domain, typedData.Types, typedData.Message, signature)

}
func Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func HexDecode(s string) ([]byte, error) {
	if Has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}

func CopyBytes(data []byte) []byte {
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret
}

func VerifyTypedData(primaryType string, domain TypedDataDomain, types Types, message TypedDataMessage, signature []byte) (common.Address, error) {
	signature = CopyBytes(signature)
	if len(signature) != crypto.SignatureLength {
		return common.Address{}, errors.Errorf("invalid signature length: %d", crypto.SignatureLength)
	}

	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27 // Transform V from 27/28 to 0/1
	}

	// if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
	// 	signature[crypto.RecoveryIDOffset] += 27
	// }
	challengeHash, err := hashTypedData(primaryType, domain, types, message)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "hash typed data")
	}

	pubKeyRaw, err := crypto.Ecrecover(challengeHash, signature)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "invalid signature")
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "unmarshal pubkey")
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr, nil
}

func hashTypedData(primaryType string, domain TypedDataDomain, types Types, message TypedDataMessage) ([]byte, error) {
	types["EIP712Domain"] = []Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	}

	signerData := TypedData{
		Types:       types,
		PrimaryType: primaryType,
		Domain:      domain,
		Message:     message,
	}

	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return nil, errors.Wrap(err, "hash typed data")
	}

	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return nil, errors.Wrap(err, "hash domain separator")
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	challengeHash := crypto.Keccak256Hash(rawData)

	return challengeHash.Bytes(), nil
}
