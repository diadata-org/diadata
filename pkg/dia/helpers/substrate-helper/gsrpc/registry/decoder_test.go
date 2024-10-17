package registry

import (
	"bytes"
	"errors"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/test"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic/extensions"
	"github.com/stretchr/testify/assert"
)

func Test_getPrimitiveType_UnsupportedTypeError(t *testing.T) {
	primitiveTypeDef := types.Si0TypeDefPrimitive(32)

	res, err := getPrimitiveDecoder(primitiveTypeDef)
	assert.ErrorIs(t, err, ErrPrimitiveTypeNotSupported)
	assert.Nil(t, res)
}

func Test_TypeDecoder(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	typeDecoder := &TypeDecoder{
		Name: "test_decoder_1",
		Fields: []*Field{
			{
				Name:         "field_1",
				FieldDecoder: &ValueDecoder[types.U8]{},
				LookupIndex:  0,
			},
			{
				Name:         "field_2",
				FieldDecoder: &ValueDecoder[types.U16]{},
				LookupIndex:  1,
			},
			{
				Name:         "field_3",
				FieldDecoder: &ValueDecoder[types.U32]{},
				LookupIndex:  2,
			},
		},
	}

	encodedTestData, err := encodeTestData(testData)
	assert.NoError(t, err)

	decoder := scale.NewDecoder(bytes.NewReader(encodedTestData))

	res, err := typeDecoder.Decode(decoder)
	assert.NoError(t, err)
	assert.Len(t, res, len(testData))
	assert.Equal(t, typeDecoder.Fields[0].Name, res[0].Name)
	assert.Equal(t, typeDecoder.Fields[0].LookupIndex, res[0].LookupIndex)
	assert.Equal(t, testData[0], res[0].Value)
}

func Test_TypeDecoder_FieldDecodingError(t *testing.T) {
	testData := []any{
		types.U32(3),
	}

	typeDecoder := &TypeDecoder{
		Name: "test_decoder_1",
		Fields: []*Field{
			{
				Name:         "field_1",
				FieldDecoder: &ValueDecoder[[32]types.U8]{},
				LookupIndex:  0,
			},
		},
	}

	encodedTestData, err := encodeTestData(testData)
	assert.NoError(t, err)

	decoder := scale.NewDecoder(bytes.NewReader(encodedTestData))

	res, err := typeDecoder.Decode(decoder)
	assert.ErrorIs(t, err, ErrTypeFieldDecoding)
	assert.Nil(t, res)
}

func Test_ProcessDecodedFieldValue(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	// Field index match
	res, err := ProcessDecodedFieldValue[types.U32](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
		func(value any) (types.U32, error) {
			res, ok := value.(types.U32)
			assert.True(t, ok)

			return res, nil
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, testData[2], res)

	// Field name match
	res, err = ProcessDecodedFieldValue[types.U32](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return field.Name == "decoded_field_2"
		},
		func(value any) (types.U32, error) {
			res, ok := value.(types.U32)
			assert.True(t, ok)

			return res, nil
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, testData[2], res)

	// Field lookup index match
	res, err = ProcessDecodedFieldValue[types.U32](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return field.LookupIndex == 2
		},
		func(value any) (types.U32, error) {
			res, ok := value.(types.U32)
			assert.True(t, ok)

			return res, nil
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, testData[2], res)
}

func Test_ProcessDecodedFieldValue_FieldNotFoundError(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := ProcessDecodedFieldValue[types.U32](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return false
		},
		func(value any) (types.U32, error) {
			res, ok := value.(types.U32)
			assert.True(t, ok)

			return res, nil
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldNotFound)
	assert.Equal(t, types.U32(0), res)
}

func Test_ProcessDecodedFieldValue_FieldValueProcessingError(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := ProcessDecodedFieldValue[types.U32](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
		func(value any) (types.U32, error) {
			return 0, errors.New("error")
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldValueProcessingError)
	assert.Equal(t, types.U32(0), res)
}

func Test_GetDecodedFieldAsType(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsType[types.U8](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == 0
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, testData[0], res)
}

func Test_GetDecodedFieldAsType_FieldNotFound(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsType[types.U8](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldNotFound)
	assert.Equal(t, types.U8(0), res)
}

func Test_GetDecodedFieldAsType_ValueTypeMismatch(t *testing.T) {
	testData := []any{
		types.U8(1),
		types.U16(2),
		types.U32(3),
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsType[types.U8](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldValueTypeMismatch)
	assert.Equal(t, types.U8(0), res)
}

func Test_GetDecodedFieldAsSliceOfType(t *testing.T) {
	testData := []any{
		types.U8(1),
		[]any{
			types.U16(0),
			types.U16(1),
			types.U16(2),
		},
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsSliceOfType[types.U16](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, []types.U16{0, 1, 2}, res)
}

func Test_GetDecodedFieldAsSliceOfType_DecodedFieldNotFound(t *testing.T) {
	testData := []any{
		types.U8(1),
		[]any{
			types.U16(0),
			types.U16(1),
			types.U16(2),
		},
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsSliceOfType[types.U16](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldNotFound)
	assert.Nil(t, res)
}

func Test_GetDecodedFieldAsSliceOfType_NotAGenericSlice(t *testing.T) {
	testData := []any{
		types.U8(1),
		// Slices in decoded fields are expected to be []any
		[]types.U16{
			0,
			1,
			2,
		},
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsSliceOfType[types.U16](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldValueNotAGenericSlice)
	assert.Nil(t, res)
}

func Test_GetDecodedFieldAsSliceOfType_SliceItemTypeMismatch(t *testing.T) {
	testData := []any{
		types.U8(1),
		[]any{
			types.U16(0),
			types.U16(1),
			types.U16(2),
		},
	}

	decodedFields := testDataToDecodedFields(testData)

	res, err := GetDecodedFieldAsSliceOfType[types.U8](
		decodedFields,
		func(fieldIndex int, field *DecodedField) bool {
			return fieldIndex == len(testData)-1
		},
	)

	assert.ErrorIs(t, err, ErrDecodedFieldValueTypeMismatch)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_DecodeHex(t *testing.T) {
	// NOTE - The following test relies on the following data from the Centrifuge development chain:
	//
	// Metadata from block 517393
	// Centrifuge development block 480514 - 0xacd5bd66bb3144f559be2bed09d0a1ae36e650b31f9eb8f5833eeb2c545d07cd
	// System remark extrinsic - 0xb10184008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a480118346322ed93ad7d2583ab3e4b71acd66cc1fce77cb225624c8eb00977681468aec33b933606ed8c2eaa75b84278c42415d491f89c5e79db6910986c1b95f486e401e0000000000431

	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	res, err := extrinsicDecoder.DecodeHex("0xb10184008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a480118346322ed93ad7d2583ab3e4b71acd66cc1fce77cb225624c8eb00977681468aec33b933606ed8c2eaa75b84278c42415d491f89c5e79db6910986c1b95f486e401e0000000000431")
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func Test_ExtrinsicDecoder_NilDecoder(t *testing.T) {
	var extDec *ExtrinsicDecoder

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(nil))

	res, err := extDec.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrNilExtrinsicDecoder)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_CompactLengthDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(nil))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicCompactLengthDecoding)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_VersionDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	encodedUcompact, err := codec.Encode(types.NewUCompactFromUInt(2))
	assert.NoError(t, err)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(encodedUcompact))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicVersionDecoding)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_AddressDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	encodedUcompact, err := codec.Encode(types.NewUCompactFromUInt(2))
	assert.NoError(t, err)

	extrinsicVersion := byte(extrinsic.Version4 | extrinsic.BitSigned)

	encodedVersion, err := codec.Encode(extrinsicVersion)
	assert.NoError(t, err)

	decoderBytes := append(encodedUcompact, encodedVersion...)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(decoderBytes))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicFieldDecoding)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_SignatureDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	encodedUcompact, err := codec.Encode(types.NewUCompactFromUInt(2))
	assert.NoError(t, err)

	extrinsicVersion := byte(extrinsic.Version4 | extrinsic.BitSigned)

	encodedVersion, err := codec.Encode(extrinsicVersion)
	assert.NoError(t, err)

	decoderBytes := append(encodedUcompact, encodedVersion...)

	testMultiAdress, err := types.NewMultiAddressFromHexAccountID("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	assert.NoError(t, err)

	encodedMultiAddress, err := codec.Encode(testMultiAdress)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedMultiAddress...)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(decoderBytes))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicFieldDecoding)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_ExtraDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	encodedUcompact, err := codec.Encode(types.NewUCompactFromUInt(2))
	assert.NoError(t, err)

	extrinsicVersion := byte(extrinsic.Version4 | extrinsic.BitSigned)

	encodedVersion, err := codec.Encode(extrinsicVersion)
	assert.NoError(t, err)

	decoderBytes := append(encodedUcompact, encodedVersion...)

	testMultiAdress, err := types.NewMultiAddressFromHexAccountID("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	assert.NoError(t, err)

	encodedMultiAddress, err := codec.Encode(testMultiAdress)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedMultiAddress...)

	testSignature := types.MultiSignature{
		IsSr25519: true,
		AsSr25519: types.SignatureHash{},
	}

	encodedSignature, err := codec.Encode(testSignature)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedSignature...)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(decoderBytes))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicFieldDecoding)
	assert.Nil(t, res)
}

func Test_ExtrinsicDecoder_CallDecodingError(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	extrinsicDecoder, err := NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	encodedUcompact, err := codec.Encode(types.NewUCompactFromUInt(2))
	assert.NoError(t, err)

	extrinsicVersion := byte(extrinsic.Version4 | extrinsic.BitSigned)

	encodedVersion, err := codec.Encode(extrinsicVersion)
	assert.NoError(t, err)

	decoderBytes := append(encodedUcompact, encodedVersion...)

	testMultiAdress, err := types.NewMultiAddressFromHexAccountID("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	assert.NoError(t, err)

	encodedMultiAddress, err := codec.Encode(testMultiAdress)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedMultiAddress...)

	testSignature := types.MultiSignature{
		IsSr25519: true,
		AsSr25519: types.SignatureHash{},
	}

	encodedSignature, err := codec.Encode(testSignature)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedSignature...)

	encodedEra, err := codec.Encode(types.ExtrinsicEra{
		IsImmortalEra: true,
	})
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedEra...)

	encodedNonce, err := codec.Encode(types.NewUCompactFromUInt(0))
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedNonce...)

	encodedTip, err := codec.Encode(types.NewUCompactFromUInt(0))
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedTip...)

	encodedMetadataMode, err := codec.Encode(extensions.CheckMetadataModeDisabled)
	assert.NoError(t, err)

	decoderBytes = append(decoderBytes, encodedMetadataMode...)

	scaleDecoder := scale.NewDecoder(bytes.NewBuffer(decoderBytes))

	res, err := extrinsicDecoder.Decode(scaleDecoder)
	assert.ErrorIs(t, err, ErrExtrinsicFieldDecoding)
	assert.Nil(t, res)
}
