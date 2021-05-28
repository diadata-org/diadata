package sign


import (
"crypto/ecdsa"
"github.com/ethereum/go-ethereum/common/hexutil"
"github.com/ethereum/go-ethereum/crypto"
"os"
)

//https://etherscan.io/verifySig

type Signer struct{
	privateKey *ecdsa.PrivateKey
}

func New()(*Signer,error){
	key := os.Getenv("ORACLE_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	return &Signer{privateKey:privateKey},nil
}

func (s *Signer)Sign(data []byte) (string, error) {
	hashedData := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hashedData.Bytes(), s.privateKey)
	if err != nil {
		return "nil", err
	}
	return hexutil.Encode(signature), nil
}



