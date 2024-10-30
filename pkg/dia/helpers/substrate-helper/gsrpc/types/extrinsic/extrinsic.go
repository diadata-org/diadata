package extrinsic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	libErr "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/error"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

const (
	ErrEncodeToHex          = libErr.Error("encode to hex")
	ErrScaleEncode          = libErr.Error("scale encode")
	ErrInvalidVersion       = libErr.Error("invalid version")
	ErrPayloadCreation      = libErr.Error("payload creation")
	ErrPayloadMutation      = libErr.Error("payload mutation")
	ErrMultiAddressCreation = libErr.Error("multi address creation")
	ErrPayloadSigning       = libErr.Error("payload signing")
)

const (
	BitSigned   = 0x80
	BitUnsigned = 0

	UnmaskVersion = 0x7f

	DefaultVersion = 1
	VersionUnknown = 0 // v0 is unknown
	Version1       = 1
	Version2       = 2
	Version3       = 3
	Version4       = 4
)

// Extrinsic is an extrinsic type that can be used on chains that
// have a custom signed extension logic.
type Extrinsic struct {
	// Version is the encoded version flag (which encodes the raw transaction version
	// and signing information in one byte).
	Version byte
	// Signature is the extrinsic signature.
	Signature *Signature
	// Method is the call this extrinsic wraps
	Method types.Call
}

// NewExtrinsic creates a new Extrinsic from the provided Call.
func NewExtrinsic(c types.Call) Extrinsic {
	return Extrinsic{
		Version: Version4,
		Method:  c,
	}
}

// MarshalJSON returns a JSON encoded byte array of Extrinsic.
func (e Extrinsic) MarshalJSON() ([]byte, error) {
	s, err := codec.EncodeToHex(e)
	if err != nil {
		return nil, ErrEncodeToHex.Wrap(err)
	}
	return json.Marshal(s)
}

// IsSigned returns true if the extrinsic is signed
func (e Extrinsic) IsSigned() bool {
	return e.Version&BitSigned == BitSigned
}

// Type returns the raw transaction version (not flagged with signing information)
func (e Extrinsic) Type() uint8 {
	return e.Version & UnmaskVersion
}

func (e Extrinsic) Encode(encoder scale.Encoder) error {
	if e.Type() != Version4 {
		return fmt.Errorf("unsupported extrinsic version: %v (isSigned: %v, type: %v)", e.Version, e.IsSigned(),
			e.Type())
	}

	var bb = bytes.Buffer{}
	tempEnc := scale.NewEncoder(&bb)

	err := tempEnc.Encode(e.Version)
	if err != nil {
		return err
	}

	if e.IsSigned() {
		err = tempEnc.Encode(e.Signature)
		if err != nil {
			return err
		}
	}

	// encode the method
	err = tempEnc.Encode(e.Method)
	if err != nil {
		return err
	}

	// take the temporary buffer to determine length, write that as prefix
	eb := bb.Bytes()
	err = encoder.EncodeUintCompact(*big.NewInt(0).SetUint64(uint64(len(eb))))
	if err != nil {
		return err
	}

	// write the actual encoded transaction
	err = encoder.Write(eb)
	if err != nil {
		return err
	}

	return nil
}
