package utils

import (
	"log"
	"os"

	zkscrypto "github.com/zksync-sdk/zksync-sdk-go"
)

type SignatureUtil struct {
	privatekey *zkscrypto.PrivateKey
}

func NewSignatureUtil() (*SignatureUtil, error) {
	pk := os.Getenv("PRIVATE_KEY")

	privateKey, err := zkscrypto.NewPrivateKey([]byte(pk)[0:32])
	if err != nil {
		return nil, err
	}

	return &SignatureUtil{privatekey: privateKey}, nil

}

func (su *SignatureUtil) Sign(message []byte) string {

	signature, err := su.privatekey.Sign(message)
	if err != nil {
		log.Fatalf("error signing message: %s", err.Error())
	}
	return signature.HexString()
}
