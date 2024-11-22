package registry

import (
	"bytes"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// FieldDecoder is the interface implemented by all the different types that are available.
type FieldDecoder interface {
	Decode(decoder *scale.Decoder) (any, error)
}

// NoopDecoder is a FieldDecoder that does not decode anything. It comes in handy for nil tuples or variants
// with no inner types.
type NoopDecoder struct{}

func (n *NoopDecoder) Decode(_ *scale.Decoder) (any, error) {
	return nil, nil
}

// VariantDecoder holds a FieldDecoder for each variant/enum.
type VariantDecoder struct {
	FieldDecoderMap map[byte]FieldDecoder
}

type VariantDecoderResult struct {
	FieldName string
	Value     any
}

func (v *VariantDecoder) Decode(decoder *scale.Decoder) (any, error) {
	variantByte, err := decoder.ReadOneByte()

	if err != nil {
		return nil, ErrVariantByteDecoding.Wrap(err)
	}

	variantDecoder, ok := v.FieldDecoderMap[variantByte]

	if !ok {
		return nil, ErrVariantFieldDecoderNotFound.WithMsg("variant '%d'", variantByte)
	}

	if _, ok := variantDecoder.(*NoopDecoder); ok {
		return variantByte, nil
	}

	var fieldName string
	if _, ok := variantDecoder.(*CompositeDecoder); ok {
		fieldName = variantDecoder.(*CompositeDecoder).FieldName
	}

	value, err := variantDecoder.Decode(decoder)
	if err != nil {
		return nil, err
	}

	return VariantDecoderResult{fieldName, value}, nil
}

// ArrayDecoder holds information about the length of the array and the FieldDecoder used for its items.
type ArrayDecoder struct {
	Length      uint
	ItemDecoder FieldDecoder
}

func (a *ArrayDecoder) Decode(decoder *scale.Decoder) (any, error) {
	if a.ItemDecoder == nil {
		return nil, ErrArrayItemDecoderNotFound
	}

	slice := make([]any, 0, a.Length)

	for i := uint(0); i < a.Length; i++ {
		item, err := a.ItemDecoder.Decode(decoder)

		if err != nil {
			return nil, ErrArrayItemDecoding.Wrap(err)
		}

		slice = append(slice, item)
	}

	return slice, nil
}

// SliceDecoder holds a FieldDecoder for the items of a vector/slice.
type SliceDecoder struct {
	ItemDecoder FieldDecoder
}

func (s *SliceDecoder) Decode(decoder *scale.Decoder) (any, error) {
	if s.ItemDecoder == nil {
		return nil, ErrSliceItemDecoderNotFound
	}

	sliceLen, err := decoder.DecodeUintCompact()

	if err != nil {
		return nil, ErrSliceLengthDecoding.Wrap(err)
	}

	slice := make([]any, 0, sliceLen.Uint64())

	for i := uint64(0); i < sliceLen.Uint64(); i++ {
		item, err := s.ItemDecoder.Decode(decoder)

		if err != nil {
			return nil, ErrSliceItemDecoding.Wrap(err)
		}

		slice = append(slice, item)
	}

	return slice, nil
}

// CompositeDecoder holds all the information required to decoder a struct/composite.
type CompositeDecoder struct {
	FieldName string
	Fields    []*Field
}

func (e *CompositeDecoder) Decode(decoder *scale.Decoder) (any, error) {
	var decodedFields DecodedFields

	for _, field := range e.Fields {
		value, err := field.FieldDecoder.Decode(decoder)

		if err != nil {
			return nil, ErrCompositeFieldDecoding.Wrap(err)
		}

		decodedFields = append(decodedFields, &DecodedField{
			Name:        field.Name,
			Value:       value,
			LookupIndex: field.LookupIndex,
		})
	}

	return decodedFields, nil
}

// ValueDecoder decodes a primitive type.
type ValueDecoder[T any] struct{}

func (v *ValueDecoder[T]) Decode(decoder *scale.Decoder) (any, error) {
	var t T

	if err := decoder.Decode(&t); err != nil {
		return nil, ErrValueDecoding.Wrap(err)
	}

	return t, nil
}

// RecursiveDecoder is a wrapper for a FieldDecoder that is recursive.
type RecursiveDecoder struct {
	FieldDecoder FieldDecoder
}

func (r *RecursiveDecoder) Decode(decoder *scale.Decoder) (any, error) {
	if r.FieldDecoder == nil {
		return nil, ErrRecursiveFieldDecoderNotFound
	}

	return r.FieldDecoder.Decode(decoder)
}

// BitSequenceDecoder holds decoding information for a bit sequence.
type BitSequenceDecoder struct {
	FieldName string
	BitOrder  types.BitOrder
}

func (b *BitSequenceDecoder) Decode(decoder *scale.Decoder) (any, error) {
	bitVec := types.NewBitVec(b.BitOrder)

	if err := bitVec.Decode(*decoder); err != nil {
		return nil, ErrBitVecDecoding.Wrap(err)
	}

	return map[string]string{
		b.FieldName: bitVec.String(),
	}, nil
}

// TypeDecoder holds all information required to decode a particular type.
type TypeDecoder struct {
	Name   string
	Fields []*Field
}

func (t *TypeDecoder) Decode(decoder *scale.Decoder) (DecodedFields, error) {
	if t == nil {
		return nil, ErrNilTypeDecoder
	}

	var decodedFields DecodedFields

	for _, field := range t.Fields {
		decodedField, err := field.Decode(decoder)

		if err != nil {
			return nil, ErrTypeFieldDecoding.Wrap(err)
		}

		decodedFields = append(decodedFields, decodedField)
	}

	return decodedFields, nil
}

// getPrimitiveDecoder parses a primitive type definition and returns a ValueDecoder.
func getPrimitiveDecoder(primitiveTypeDef types.Si0TypeDefPrimitive) (FieldDecoder, error) {
	switch primitiveTypeDef {
	case types.IsBool:
		return &ValueDecoder[bool]{}, nil
	case types.IsChar:
		return &ValueDecoder[byte]{}, nil
	case types.IsStr:
		return &ValueDecoder[string]{}, nil
	case types.IsU8:
		return &ValueDecoder[types.U8]{}, nil
	case types.IsU16:
		return &ValueDecoder[types.U16]{}, nil
	case types.IsU32:
		return &ValueDecoder[types.U32]{}, nil
	case types.IsU64:
		return &ValueDecoder[types.U64]{}, nil
	case types.IsU128:
		return &ValueDecoder[types.U128]{}, nil
	case types.IsU256:
		return &ValueDecoder[types.U256]{}, nil
	case types.IsI8:
		return &ValueDecoder[types.I8]{}, nil
	case types.IsI16:
		return &ValueDecoder[types.I16]{}, nil
	case types.IsI32:
		return &ValueDecoder[types.I32]{}, nil
	case types.IsI64:
		return &ValueDecoder[types.I64]{}, nil
	case types.IsI128:
		return &ValueDecoder[types.I128]{}, nil
	case types.IsI256:
		return &ValueDecoder[types.I256]{}, nil
	default:
		return nil, ErrPrimitiveTypeNotSupported.WithMsg("primitive type %v", primitiveTypeDef)
	}
}

// Field represents one field of a TypeDecoder.
type Field struct {
	Name         string
	FieldDecoder FieldDecoder
	LookupIndex  int64
}

func (f *Field) Decode(decoder *scale.Decoder) (*DecodedField, error) {
	if f == nil {
		return nil, ErrNilField
	}

	if f.FieldDecoder == nil {
		return nil, ErrNilFieldDecoder
	}

	value, err := f.FieldDecoder.Decode(decoder)

	if err != nil {
		return nil, err
	}

	return &DecodedField{
		Name:        f.Name,
		Value:       value,
		LookupIndex: f.LookupIndex,
	}, nil
}

// DecodedField holds the name, value and lookup index of a field that was decoded.
type DecodedField struct {
	Name        string
	Value       any
	LookupIndex int64
}

func (d DecodedField) Encode(encoder scale.Encoder) error {
	if d.Value == nil {
		return nil
	}

	return encoder.Encode(d.Value)
}

type DecodedFields []*DecodedField

type DecodedFieldPredicateFn func(fieldIndex int, field *DecodedField) bool
type DecodedValueProcessingFn[T any] func(value any) (T, error)

// ProcessDecodedFieldValue applies the processing func to the value of the field
// that matches the provided predicate func.
func ProcessDecodedFieldValue[T any](
	decodedFields DecodedFields,
	fieldPredicateFn DecodedFieldPredicateFn,
	valueProcessingFn DecodedValueProcessingFn[T],
) (T, error) {
	var t T

	for decodedFieldIndex, decodedField := range decodedFields {
		if !fieldPredicateFn(decodedFieldIndex, decodedField) {
			continue
		}

		res, err := valueProcessingFn(decodedField.Value)

		if err != nil {
			return t, ErrDecodedFieldValueProcessingError.Wrap(err)
		}

		return res, nil
	}

	return t, ErrDecodedFieldNotFound
}

// GetDecodedFieldAsType returns the value of the field that matches the provided predicate func
// as the provided generic argument.
func GetDecodedFieldAsType[T any](
	decodedFields DecodedFields,
	fieldPredicateFn DecodedFieldPredicateFn,
) (T, error) {
	return ProcessDecodedFieldValue(
		decodedFields,
		fieldPredicateFn,
		func(value any) (T, error) {
			if res, ok := value.(T); ok {
				return res, nil
			}

			var t T

			err := fmt.Errorf("expected %T, got %T", t, value)

			return t, ErrDecodedFieldValueTypeMismatch.Wrap(err)
		},
	)
}

// GetDecodedFieldAsSliceOfType returns the value of the field that matches the provided predicate func
// as a slice of the provided generic argument.
func GetDecodedFieldAsSliceOfType[T any](
	decodedFields DecodedFields,
	fieldPredicateFn DecodedFieldPredicateFn,
) ([]T, error) {
	return ProcessDecodedFieldValue(
		decodedFields,
		fieldPredicateFn,
		func(value any) ([]T, error) {
			v, ok := value.([]any)

			if !ok {
				return nil, ErrDecodedFieldValueNotAGenericSlice
			}

			res, err := convertSliceToType[T](v)

			if err != nil {
				return nil, ErrDecodedFieldValueTypeMismatch.Wrap(err)
			}

			return res, nil
		},
	)
}

func convertSliceToType[T any](slice []any) ([]T, error) {
	res := make([]T, 0)

	for _, item := range slice {
		if v, ok := item.(T); ok {
			res = append(res, v)
			continue
		}

		var t T

		return nil, fmt.Errorf("expected %T, got %T", t, item)
	}

	return res, nil
}

// DecodedExtrinsic is the type returned when an extrinsic is decoded.
type DecodedExtrinsic struct {
	Version       byte
	DecodedFields DecodedFields
}

// IsSigned returns true if the extrinsic is signed.
func (d DecodedExtrinsic) IsSigned() bool {
	return d.Version&extrinsic.BitSigned == extrinsic.BitSigned
}

// ExtrinsicDecoder holds all the decoders for all the fields of an extrinsic.
type ExtrinsicDecoder struct {
	Fields []*Field
}

func (d *ExtrinsicDecoder) getFieldWithName(fieldName string) (*Field, error) {
	for _, field := range d.Fields {
		if field.Name == fieldName {
			return field, nil
		}
	}

	return nil, ErrExtrinsicFieldNotFound.WithMsg("expected field name '%s'", fieldName)
}

func (d *ExtrinsicDecoder) decodeField(fieldName string, decoder *scale.Decoder) (*DecodedField, error) {
	extrinsicField, err := d.getFieldWithName(fieldName)

	if err != nil {
		return nil, err
	}

	decodedField, err := extrinsicField.Decode(decoder)

	if err != nil {
		return nil, ErrExtrinsicFieldDecoding.Wrap(err).WithMsg("field name - '%s'", fieldName)
	}

	return decodedField, nil
}

func (d *ExtrinsicDecoder) DecodeHex(hexEncodedExtrinsic string) (*DecodedExtrinsic, error) {
	extrinsicBytes, err := hexutil.Decode(hexEncodedExtrinsic)

	if err != nil {
		return nil, err
	}

	decoder := scale.NewDecoder(bytes.NewReader(extrinsicBytes))

	return d.Decode(decoder)
}

// Decode is used to decode the fields of an extrinsic in the following order:
//
// 1. Address
// 2. Signature
// 3. Extra
// 4. Call
//
// NOTE - the decoding order is different from the order of the Extrinsic parameters provided in the metadata.
func (d *ExtrinsicDecoder) Decode(decoder *scale.Decoder) (*DecodedExtrinsic, error) {
	if d == nil {
		return nil, ErrNilExtrinsicDecoder
	}

	decodedExtrinsic := &DecodedExtrinsic{}

	// compact length encoding (1, 2, or 4 bytes) (may not be there for Extrinsics older than Jan 11 2019)
	_, err := decoder.DecodeUintCompact()

	if err != nil {
		return nil, ErrExtrinsicCompactLengthDecoding.Wrap(err)
	}

	if err := decoder.Decode(&decodedExtrinsic.Version); err != nil {
		return nil, ErrExtrinsicVersionDecoding.Wrap(err)
	}

	var decodedFields DecodedFields

	if decodedExtrinsic.IsSigned() {
		decodedAddress, err := d.decodeField(ExtrinsicAddressName, decoder)

		if err != nil {
			return nil, err
		}

		decodedFields = append(decodedFields, decodedAddress)

		decodedSignature, err := d.decodeField(ExtrinsicSignatureName, decoder)

		if err != nil {
			return nil, err
		}

		decodedFields = append(decodedFields, decodedSignature)

		decodedExtraField, err := d.decodeField(ExtrinsicExtraName, decoder)

		if err != nil {
			return nil, err
		}

		decodedFields = append(decodedFields, decodedExtraField)
	}

	decodedCall, err := d.decodeField(ExtrinsicCallName, decoder)

	if err != nil {
		return nil, err
	}

	decodedFields = append(decodedFields, decodedCall)

	decodedExtrinsic.DecodedFields = decodedFields

	return decodedExtrinsic, nil
}
