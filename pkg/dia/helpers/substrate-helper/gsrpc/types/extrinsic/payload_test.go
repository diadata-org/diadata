package extrinsic

import (
	"bytes"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestPayload_Encode(t *testing.T) {
	b := bytes.NewBuffer(nil)

	encoder := scale.NewEncoder(b)

	call := types.BytesBare([]byte{1, 2, 3})
	signedFieldValue := uint64(1)
	signedExtraFieldValue := "test"

	encodedCall, err := codec.Encode(call)
	assert.NoError(t, err)
	encodedSignedFieldValue, err := codec.Encode(signedFieldValue)
	assert.NoError(t, err)
	encodedSignedExtraFieldValue, err := codec.Encode(signedExtraFieldValue)
	assert.NoError(t, err)

	expectedResult := append(encodedCall, append(encodedSignedFieldValue, encodedSignedExtraFieldValue...)...)

	payload := Payload{
		EncodedCall: call,
		SignedFields: []*SignedField{
			{
				Name:    "signed field 1",
				Value:   signedFieldValue,
				Mutated: true,
			},
		},
		SignedExtraFields: []*SignedField{
			{
				Name:    "signed extra field 1",
				Value:   signedExtraFieldValue,
				Mutated: true,
			},
		},
	}

	err = payload.Encode(*encoder)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, b.Bytes())
}

func TestPayload_Encode_SignedFieldNotMutatedError(t *testing.T) {
	b := bytes.NewBuffer(nil)

	encoder := scale.NewEncoder(b)

	call := types.BytesBare([]byte{1, 2, 3})
	signedFieldValue := uint64(1)
	signedExtraFieldValue := "test"

	payload := Payload{
		EncodedCall: call,
		SignedFields: []*SignedField{
			{
				Name:    "signed field 1",
				Value:   signedFieldValue,
				Mutated: false,
			},
		},
		SignedExtraFields: []*SignedField{
			{
				Name:    "signed extra field 1",
				Value:   signedExtraFieldValue,
				Mutated: true,
			},
		},
	}

	err := payload.Encode(*encoder)
	assert.ErrorIs(t, err, ErrSignedFieldNotMutated)
}

func TestPayload_Encode_SignedExtraFieldNotMutatedError(t *testing.T) {
	b := bytes.NewBuffer(nil)

	encoder := scale.NewEncoder(b)

	call := types.BytesBare([]byte{1, 2, 3})
	signedFieldValue := uint64(1)
	signedExtraFieldValue := "test"

	payload := Payload{
		EncodedCall: call,
		SignedFields: []*SignedField{
			{
				Name:    "signed field 1",
				Value:   signedFieldValue,
				Mutated: true,
			},
		},
		SignedExtraFields: []*SignedField{
			{
				Name:    "signed extra field 1",
				Value:   signedExtraFieldValue,
				Mutated: false,
			},
		},
	}

	err := payload.Encode(*encoder)
	assert.ErrorIs(t, err, ErrSignedExtraFieldNotMutated)
}

func TestPayload_MutateSignedFields(t *testing.T) {
	var payload *Payload

	err := payload.MutateSignedFields(nil)
	assert.ErrorIs(t, err, ErrPayloadNil)

	signedFieldName := SignedFieldName("signed field 1")
	signedFieldValue := uint64(1)

	signedExtraFieldName := SignedFieldName("signed extra field 1")
	signedExtraFieldValue := "test"

	payload = &Payload{
		SignedFields: []*SignedField{
			{
				Name:    signedFieldName,
				Value:   signedFieldValue,
				Mutated: true,
			},
		},
		SignedExtraFields: []*SignedField{
			{
				Name:    signedExtraFieldName,
				Value:   signedExtraFieldValue,
				Mutated: true,
			},
		},
	}
	assert.Equal(t, signedFieldValue, payload.SignedFields[0].Value)
	assert.Equal(t, signedExtraFieldValue, payload.SignedExtraFields[0].Value)

	newSignedFieldValue := uint64(2)
	newSignedExtraFieldValue := "testtest"

	vals := map[SignedFieldName]any{
		signedFieldName:      newSignedFieldValue,
		signedExtraFieldName: newSignedExtraFieldValue,
	}

	err = payload.MutateSignedFields(vals)
	assert.NoError(t, err)

	assert.Equal(t, newSignedFieldValue, payload.SignedFields[0].Value)
	assert.Equal(t, newSignedExtraFieldValue, payload.SignedExtraFields[0].Value)
}

func TestPayload_createPayload(t *testing.T) {
	call := types.BytesBare([]byte{1, 2, 3})

	var meta types.Metadata

	err := codec.DecodeFromHex(types.MetadataV14Data, &meta)
	assert.NoError(t, err)

	payload, err := createPayload(&meta, call)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
}

func TestPayload_createPayload_SignedExtensionNotDefinedError(t *testing.T) {
	call := types.BytesBare([]byte{1, 2, 3})

	var meta types.Metadata

	err := codec.DecodeFromHex(types.MetadataV14Data, &meta)
	assert.NoError(t, err)

	delete(meta.AsMetadataV14.EfficientLookup, meta.AsMetadataV14.Extrinsic.SignedExtensions[0].Type.Int64())

	payload, err := createPayload(&meta, call)
	assert.ErrorIs(t, err, ErrSignedExtensionTypeNotDefined)
	assert.Nil(t, payload)
}

func TestPayload_createPayload_SignedExtensionNotSupportedError(t *testing.T) {
	call := types.BytesBare([]byte{1, 2, 3})

	var meta types.Metadata

	err := codec.DecodeFromHex(types.MetadataV14Data, &meta)
	assert.NoError(t, err)

	meta.AsMetadataV14.Extrinsic.SignedExtensions = append(meta.AsMetadataV14.Extrinsic.SignedExtensions, types.SignedExtensionMetadataV14{
		Identifier:       "unsupported_extension",
		Type:             types.Si1LookupTypeID{},
		AdditionalSigned: types.Si1LookupTypeID{},
	})

	payload, err := createPayload(&meta, call)
	assert.ErrorIs(t, err, ErrSignedExtensionTypeNotSupported)
	assert.Nil(t, payload)
}
