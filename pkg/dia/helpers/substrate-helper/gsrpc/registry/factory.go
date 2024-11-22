package registry

import (
	"errors"
	"fmt"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

//go:generate mockery --name Factory --structname FactoryMock --filename factory_mock.go --inpackage

// Factory is the interface responsible for generating the according registries from the metadata.
type Factory interface {
	CreateCallRegistry(meta *types.Metadata) (CallRegistry, error)
	CreateErrorRegistry(meta *types.Metadata) (ErrorRegistry, error)
	CreateEventRegistry(meta *types.Metadata) (EventRegistry, error)
	CreateExtrinsicDecoder(meta *types.Metadata) (*ExtrinsicDecoder, error)
}

// CallRegistry maps a call name to its TypeDecoder.
type CallRegistry map[types.CallIndex]*TypeDecoder

// ErrorID is the type using for identifying an error in the metadata.
type ErrorID struct {
	ModuleIndex types.U8
	ErrorIndex  [4]types.U8
}

// ErrorRegistry maps an error name to its TypeDecoder.
type ErrorRegistry map[ErrorID]*TypeDecoder

// EventRegistry maps an event ID to its TypeDecoder.
type EventRegistry map[types.EventID]*TypeDecoder

// FieldOverride is used to override the default FieldDecoder for a particular type.
type FieldOverride struct {
	FieldLookupIndex int64
	FieldDecoder     FieldDecoder
}

type factory struct {
	fieldStorage          map[int64]FieldDecoder
	recursiveFieldStorage map[int64]*RecursiveDecoder
	fieldOverrides        []FieldOverride
}

// NewFactory creates a new Factory using the provided overrides, if any.
func NewFactory(fieldOverrides ...FieldOverride) Factory {
	f := &factory{}
	f.fieldOverrides = fieldOverrides

	return f
}

func (f *factory) resetStorages() {
	f.fieldStorage = make(map[int64]FieldDecoder)
	f.recursiveFieldStorage = make(map[int64]*RecursiveDecoder)

	for _, fieldOverride := range f.fieldOverrides {
		f.fieldStorage[fieldOverride.FieldLookupIndex] = fieldOverride.FieldDecoder
	}
}

// CreateErrorRegistry creates the registry that contains the types for errors.
// nolint:dupl
func (f *factory) CreateErrorRegistry(meta *types.Metadata) (ErrorRegistry, error) {
	f.resetStorages()

	errorRegistry := make(map[ErrorID]*TypeDecoder)

	for _, mod := range meta.AsMetadataV14.Pallets {
		if !mod.HasErrors {
			continue
		}

		errorsType, ok := meta.AsMetadataV14.EfficientLookup[mod.Errors.Type.Int64()]

		if !ok {
			return nil, ErrErrorsTypeNotFound.WithMsg("errors type '%d', module '%s'", mod.Errors.Type.Int64(), mod.Name)
		}

		if !errorsType.Def.IsVariant {
			return nil, ErrErrorsTypeNotVariant.WithMsg("errors type '%d', module '%s'", mod.Errors.Type.Int64(), mod.Name)
		}

		for _, errorVariant := range errorsType.Def.Variant.Variants {
			errorName := fmt.Sprintf("%s.%s", mod.Name, errorVariant.Name)

			errorFields, err := f.getTypeFields(meta, errorVariant.Fields)

			if err != nil {
				return nil, ErrErrorFieldsRetrieval.WithMsg(errorName).Wrap(err)
			}

			errorID := ErrorID{
				ModuleIndex: mod.Index,
				ErrorIndex:  [4]types.U8{errorVariant.Index},
			}

			errorRegistry[errorID] = &TypeDecoder{
				Name:   errorName,
				Fields: errorFields,
			}
		}
	}

	if err := f.resolveRecursiveDecoders(); err != nil {
		return nil, ErrRecursiveDecodersResolving.Wrap(err)
	}

	return errorRegistry, nil
}

// CreateCallRegistry creates the registry that contains the types for calls.
// nolint:dupl
func (f *factory) CreateCallRegistry(meta *types.Metadata) (CallRegistry, error) {
	f.resetStorages()

	callRegistry := make(map[types.CallIndex]*TypeDecoder)

	for _, mod := range meta.AsMetadataV14.Pallets {
		if !mod.HasCalls {
			continue
		}

		callsType, ok := meta.AsMetadataV14.EfficientLookup[mod.Calls.Type.Int64()]

		if !ok {
			return nil, ErrCallsTypeNotFound.WithMsg("calls type '%d', module '%s'", mod.Calls.Type.Int64(), mod.Name)
		}

		if !callsType.Def.IsVariant {
			return nil, ErrCallsTypeNotVariant.WithMsg("calls type '%d', module '%s'", mod.Calls.Type.Int64(), mod.Name)
		}

		for _, callVariant := range callsType.Def.Variant.Variants {
			callIndex := types.CallIndex{
				SectionIndex: uint8(mod.Index),
				MethodIndex:  uint8(callVariant.Index),
			}

			callName := fmt.Sprintf("%s.%s", mod.Name, callVariant.Name)

			callFields, err := f.getTypeFields(meta, callVariant.Fields)

			if err != nil {
				return nil, ErrCallFieldsRetrieval.WithMsg(callName).Wrap(err)
			}

			callRegistry[callIndex] = &TypeDecoder{
				Name:   callName,
				Fields: callFields,
			}
		}
	}

	if err := f.resolveRecursiveDecoders(); err != nil {
		return nil, ErrRecursiveDecodersResolving.Wrap(err)
	}

	return callRegistry, nil
}

// CreateEventRegistry creates the registry that contains the types for events.
func (f *factory) CreateEventRegistry(meta *types.Metadata) (EventRegistry, error) {
	f.resetStorages()

	eventRegistry := make(map[types.EventID]*TypeDecoder)

	for _, mod := range meta.AsMetadataV14.Pallets {
		if !mod.HasEvents {
			continue
		}

		eventsType, ok := meta.AsMetadataV14.EfficientLookup[mod.Events.Type.Int64()]

		if !ok {
			return nil, ErrEventsTypeNotFound.WithMsg("events type '%d', module '%s'", mod.Events.Type.Int64(), mod.Name)
		}

		if !eventsType.Def.IsVariant {
			return nil, ErrEventsTypeNotVariant.WithMsg("events type '%d', module '%s'", mod.Events.Type.Int64(), mod.Name)
		}

		for _, eventVariant := range eventsType.Def.Variant.Variants {
			eventID := types.EventID{byte(mod.Index), byte(eventVariant.Index)}

			eventName := fmt.Sprintf("%s.%s", mod.Name, eventVariant.Name)

			eventFields, err := f.getTypeFields(meta, eventVariant.Fields)

			if err != nil {
				return nil, ErrEventFieldsRetrieval.WithMsg(eventName).Wrap(err)
			}

			eventRegistry[eventID] = &TypeDecoder{
				Name:   eventName,
				Fields: eventFields,
			}
		}
	}

	if err := f.resolveRecursiveDecoders(); err != nil {
		return nil, ErrRecursiveDecodersResolving.Wrap(err)
	}

	return eventRegistry, nil
}

// CreateExtrinsicDecoder creates an ExtrinsicDecoder based on the Extrinsic information provided in the metadata.
func (f *factory) CreateExtrinsicDecoder(meta *types.Metadata) (*ExtrinsicDecoder, error) {
	f.resetStorages()

	extrinsicLookupID := meta.AsMetadataV14.Extrinsic.Type

	extrinsicType := meta.AsMetadataV14.EfficientLookup[extrinsicLookupID.Int64()]

	extrinsicParams, err := extractExtrinsicParams(extrinsicType, meta)

	if err != nil {
		return nil, err
	}

	if err := validateExtrinsicParams(extrinsicParams); err != nil {
		return nil, err
	}

	extrinsicFields, err := f.getTypeParams(meta, extrinsicParams)

	if err != nil {
		return nil, ErrExtrinsicFieldRetrieval
	}

	if err := f.resolveRecursiveDecoders(); err != nil {
		return nil, ErrRecursiveDecodersResolving.Wrap(err)
	}

	return &ExtrinsicDecoder{
		Fields: extrinsicFields,
	}, nil
}

const (
	ExtrinsicAddressName   = "Address"
	ExtrinsicSignatureName = "Signature"
	ExtrinsicExtraName     = "Extra"
	ExtrinsicCallName      = "Call"
)

var expectedExtrinsicParams = map[string]struct{}{
	ExtrinsicAddressName:   {},
	ExtrinsicSignatureName: {},
	ExtrinsicExtraName:     {},
	ExtrinsicCallName:      {},
}

func validateExtrinsicParams(params []types.Si1TypeParameter) error {
	if len(params) != ExpectedExtrinsicParamsCount {
		return ErrInvalidExtrinsicParams
	}

	for _, param := range params {
		if _, ok := expectedExtrinsicParams[string(param.Name)]; !ok {
			return ErrUnexpectedExtrinsicParam.WithMsg("param - '%s'", param.Name)
		}
	}

	return nil
}

// resolveRecursiveDecoders resolves all recursive decoders with their according FieldDecoder.
// nolint:lll
func (f *factory) resolveRecursiveDecoders() error {
	for recursiveFieldLookupIndex, recursiveFieldDecoder := range f.recursiveFieldStorage {
		if recursiveFieldDecoder.FieldDecoder != nil {
			// Skip if the inner FieldDecoder is present, this could be an override.
			continue
		}

		fieldDecoder, ok := f.fieldStorage[recursiveFieldLookupIndex]

		if !ok {
			return ErrFieldDecoderForRecursiveFieldNotFound.
				WithMsg(
					"recursive field lookup index %d",
					recursiveFieldLookupIndex,
				)
		}

		if _, ok := fieldDecoder.(*RecursiveDecoder); ok {
			return ErrRecursiveFieldResolving.
				WithMsg(
					"recursive field lookup index %d",
					recursiveFieldLookupIndex,
				)
		}

		recursiveFieldDecoder.FieldDecoder = fieldDecoder
	}

	return nil
}

// getTypeFields returns a list of fields and their respective decoders from the provided parameters.
func (f *factory) getTypeParams(meta *types.Metadata, params []types.Si1TypeParameter) ([]*Field, error) {
	var typeFields []*Field

	for _, param := range params {
		paramType, ok := meta.AsMetadataV14.EfficientLookup[param.Type.Int64()]

		if !ok {
			return nil, ErrFieldTypeNotFound.WithMsg(string(param.Name))
		}

		paramName := string(param.Name)

		if storedFieldDecoder, ok := f.getStoredFieldDecoder(param.Type.Int64()); ok {
			typeFields = append(typeFields, &Field{
				Name:         paramName,
				FieldDecoder: storedFieldDecoder,
				LookupIndex:  param.Type.Int64(),
			})
			continue
		}

		paramTypeDef := paramType.Def

		fieldDecoder, err := f.getFieldDecoder(meta, paramName, paramTypeDef)

		if err != nil {
			return nil, ErrFieldDecoderRetrieval.WithMsg(paramName).Wrap(err)
		}

		f.fieldStorage[param.Type.Int64()] = fieldDecoder

		typeFields = append(typeFields, &Field{
			Name:         paramName,
			FieldDecoder: fieldDecoder,
			LookupIndex:  param.Type.Int64(),
		})
	}

	return typeFields, nil
}

// getTypeFields parses and returns all Field(s) for a type.
func (f *factory) getTypeFields(meta *types.Metadata, fields []types.Si1Field) ([]*Field, error) {
	var typeFields []*Field

	for _, field := range fields {
		fieldType, ok := meta.AsMetadataV14.EfficientLookup[field.Type.Int64()]

		if !ok {
			return nil, ErrFieldTypeNotFound.WithMsg(string(field.Name))
		}

		fieldName := getFullFieldName(field, fieldType)

		if storedFieldDecoder, ok := f.getStoredFieldDecoder(field.Type.Int64()); ok {
			typeFields = append(typeFields, &Field{
				Name:         fieldName,
				FieldDecoder: storedFieldDecoder,
				LookupIndex:  field.Type.Int64(),
			})
			continue
		}

		fieldTypeDef := fieldType.Def

		fieldDecoder, err := f.getFieldDecoder(meta, fieldName, fieldTypeDef)

		if err != nil {
			return nil, ErrFieldDecoderRetrieval.WithMsg(fieldName).Wrap(err)
		}

		f.fieldStorage[field.Type.Int64()] = fieldDecoder

		typeFields = append(typeFields, &Field{
			Name:         fieldName,
			FieldDecoder: fieldDecoder,
			LookupIndex:  field.Type.Int64(),
		})
	}

	return typeFields, nil
}

// getFieldDecoder returns the FieldDecoder based on the provided type definition.
// nolint:funlen
func (f *factory) getFieldDecoder(
	meta *types.Metadata,
	fieldName string,
	typeDef types.Si1TypeDef,
) (FieldDecoder, error) {
	switch {
	case typeDef.IsCompact:
		compactFieldType, ok := meta.AsMetadataV14.EfficientLookup[typeDef.Compact.Type.Int64()]

		if !ok {
			return nil, ErrCompactFieldTypeNotFound.WithMsg(fieldName)
		}

		return f.getCompactFieldDecoder(meta, fieldName, compactFieldType.Def)
	case typeDef.IsComposite:
		compositeDecoder := &CompositeDecoder{
			FieldName: fieldName,
		}

		fields, err := f.getTypeFields(meta, typeDef.Composite.Fields)

		if err != nil {
			return nil, ErrCompositeTypeFieldsRetrieval.WithMsg(fieldName).Wrap(err)
		}

		compositeDecoder.Fields = fields

		return compositeDecoder, nil
	case typeDef.IsVariant:
		return f.getVariantFieldDecoder(meta, typeDef)
	case typeDef.IsPrimitive:
		return getPrimitiveDecoder(typeDef.Primitive.Si0TypeDefPrimitive)
	case typeDef.IsArray:
		arrayFieldType, ok := meta.AsMetadataV14.EfficientLookup[typeDef.Array.Type.Int64()]

		if !ok {
			return nil, ErrArrayFieldTypeNotFound.WithMsg(fieldName)
		}

		return f.getArrayFieldDecoder(uint(typeDef.Array.Len), meta, fieldName, arrayFieldType.Def)
	case typeDef.IsSequence:
		vectorFieldType, ok := meta.AsMetadataV14.EfficientLookup[typeDef.Sequence.Type.Int64()]

		if !ok {
			return nil, ErrVectorFieldTypeNotFound.WithMsg(fieldName)
		}

		return f.getSliceFieldDecoder(meta, fieldName, vectorFieldType.Def)
	case typeDef.IsTuple:
		if typeDef.Tuple == nil {
			return &NoopDecoder{}, nil
		}

		return f.getTupleFieldDecoder(meta, fieldName, typeDef.Tuple)
	case typeDef.IsBitSequence:
		return f.getBitSequenceDecoder(meta, fieldName, typeDef.BitSequence)
	default:
		return nil, ErrFieldTypeDefinitionNotSupported.WithMsg(fieldName)
	}
}

// getVariantFieldDecoder parses a variant type definition and returns a VariantDecoder.
func (f *factory) getVariantFieldDecoder(meta *types.Metadata, typeDef types.Si1TypeDef) (FieldDecoder, error) {
	variantDecoder := &VariantDecoder{}

	fieldDecoderMap := make(map[byte]FieldDecoder)

	for _, variant := range typeDef.Variant.Variants {
		if len(variant.Fields) == 0 {
			fieldDecoderMap[byte(variant.Index)] = &NoopDecoder{}
			continue
		}

		variantName := getVariantName(variant)

		compositeDecoder := &CompositeDecoder{
			FieldName: variantName,
		}

		fields, err := f.getTypeFields(meta, variant.Fields)

		if err != nil {
			return nil, ErrVariantTypeFieldsRetrieval.WithMsg("variant '%d'", variant.Index).Wrap(err)
		}

		compositeDecoder.Fields = fields

		fieldDecoderMap[byte(variant.Index)] = compositeDecoder
	}

	variantDecoder.FieldDecoderMap = fieldDecoderMap

	return variantDecoder, nil
}

const (
	variantItemFieldNameFormat = "variant_item_%d"
)

func getVariantName(variant types.Si1Variant) string {
	if variant.Name != "" {
		return string(variant.Name)
	}

	return fmt.Sprintf(variantItemFieldNameFormat, variant.Index)
}

const (
	tupleItemFieldNameFormat = "tuple_item_%d"
)

// getCompactFieldDecoder parses a compact type definition and returns the according field decoder.
// nolint:funlen,lll
func (f *factory) getCompactFieldDecoder(meta *types.Metadata, fieldName string, typeDef types.Si1TypeDef) (FieldDecoder, error) {
	switch {
	case typeDef.IsPrimitive:
		return &ValueDecoder[types.UCompact]{}, nil
	case typeDef.IsTuple:
		if typeDef.Tuple == nil {
			return &NoopDecoder{}, nil
		}

		compositeDecoder := &CompositeDecoder{
			FieldName: fieldName,
		}

		for i, item := range typeDef.Tuple {
			itemTypeDef, ok := meta.AsMetadataV14.EfficientLookup[item.Int64()]

			if !ok {
				return nil, ErrCompactTupleItemTypeNotFound.WithMsg("tuple item '%d'", item.Int64())
			}

			fieldName := fmt.Sprintf(tupleItemFieldNameFormat, i)

			itemFieldDecoder, err := f.getCompactFieldDecoder(meta, fieldName, itemTypeDef.Def)

			if err != nil {
				return nil, ErrCompactTupleItemFieldDecoderRetrieval.
					WithMsg("tuple item '%d'", item.Int64()).
					Wrap(err)
			}

			compositeDecoder.Fields = append(compositeDecoder.Fields, &Field{
				Name:         fieldName,
				FieldDecoder: itemFieldDecoder,
				LookupIndex:  item.Int64(),
			})
		}

		return compositeDecoder, nil
	case typeDef.IsComposite:
		compactCompositeFields := typeDef.Composite.Fields

		compositeDecoder := &CompositeDecoder{
			FieldName: fieldName,
		}

		for _, compactCompositeField := range compactCompositeFields {
			compactCompositeFieldType, ok := meta.AsMetadataV14.EfficientLookup[compactCompositeField.Type.Int64()]

			if !ok {
				return nil, ErrCompactCompositeFieldTypeNotFound
			}

			compactFieldName := getFullFieldName(compactCompositeField, compactCompositeFieldType)

			compactCompositeDecoder, err := f.getCompactFieldDecoder(meta, compactFieldName, compactCompositeFieldType.Def)

			if err != nil {
				return nil, ErrCompactCompositeFieldDecoderRetrieval.Wrap(err)
			}

			compositeDecoder.Fields = append(compositeDecoder.Fields, &Field{
				Name:         compactFieldName,
				FieldDecoder: compactCompositeDecoder,
				LookupIndex:  compactCompositeField.Type.Int64(),
			})
		}

		return compositeDecoder, nil
	default:
		return nil, errors.New("unsupported compact field type")
	}
}

// getArrayFieldDecoder parses an array type definition and returns an ArrayDecoder.
// nolint:lll
func (f *factory) getArrayFieldDecoder(arrayLen uint, meta *types.Metadata, fieldName string, typeDef types.Si1TypeDef) (FieldDecoder, error) {
	itemFieldDecoder, err := f.getFieldDecoder(meta, fieldName, typeDef)

	if err != nil {
		return nil, ErrArrayItemFieldDecoderRetrieval.Wrap(err)
	}

	return &ArrayDecoder{Length: arrayLen, ItemDecoder: itemFieldDecoder}, nil
}

// getSliceFieldDecoder parses a slice type definition and returns an SliceDecoder.
func (f *factory) getSliceFieldDecoder(
	meta *types.Metadata,
	fieldName string,
	typeDef types.Si1TypeDef,
) (FieldDecoder, error) {
	itemFieldDecoder, err := f.getFieldDecoder(meta, fieldName, typeDef)

	if err != nil {
		return nil, ErrSliceItemFieldDecoderRetrieval.Wrap(err)
	}

	return &SliceDecoder{itemFieldDecoder}, nil
}

// getTupleFieldDecoder parses a tuple type definition and returns a CompositeDecoder.
func (f *factory) getTupleFieldDecoder(
	meta *types.Metadata,
	fieldName string,
	tuple types.Si1TypeDefTuple,
) (FieldDecoder, error) {
	compositeDecoder := &CompositeDecoder{
		FieldName: fieldName,
	}

	for i, item := range tuple {
		itemTypeDef, ok := meta.AsMetadataV14.EfficientLookup[item.Int64()]

		if !ok {
			return nil, ErrTupleItemTypeNotFound.WithMsg("tuple item '%d'", i)
		}

		tupleFieldName := fmt.Sprintf(tupleItemFieldNameFormat, i)

		itemFieldDecoder, err := f.getFieldDecoder(meta, tupleFieldName, itemTypeDef.Def)

		if err != nil {
			return nil, ErrTupleItemFieldDecoderRetrieval.Wrap(err)
		}

		compositeDecoder.Fields = append(compositeDecoder.Fields, &Field{
			Name:         tupleFieldName,
			FieldDecoder: itemFieldDecoder,
			LookupIndex:  item.Int64(),
		})
	}

	return compositeDecoder, nil
}

func (f *factory) getBitSequenceDecoder(
	meta *types.Metadata,
	fieldName string,
	bitSequenceTypeDef types.Si1TypeDefBitSequence,
) (FieldDecoder, error) {
	bitStoreType, ok := meta.AsMetadataV14.EfficientLookup[bitSequenceTypeDef.BitStoreType.Int64()]

	if !ok {
		return nil, ErrBitStoreTypeNotFound.WithMsg(fieldName)
	}

	if bitStoreType.Def.Primitive.Si0TypeDefPrimitive != types.IsU8 {
		return nil, ErrBitStoreTypeNotSupported.WithMsg(fieldName)
	}

	bitOrderType, ok := meta.AsMetadataV14.EfficientLookup[bitSequenceTypeDef.BitOrderType.Int64()]

	if !ok {
		return nil, ErrBitOrderTypeNotFound.WithMsg(fieldName)
	}

	bitOrder, err := types.NewBitOrderFromString(getBitOrderString(bitOrderType.Path))

	if err != nil {
		return nil, ErrBitOrderCreation.Wrap(err)
	}

	bitSequenceDecoder := &BitSequenceDecoder{
		FieldName: fieldName,
		BitOrder:  bitOrder,
	}

	return bitSequenceDecoder, nil
}

// getStoredFieldDecoder will attempt to return a FieldDecoder from storage,
// and perform an extra check for recursive decoders.
func (f *factory) getStoredFieldDecoder(fieldLookupIndex int64) (FieldDecoder, bool) {
	if ft, ok := f.fieldStorage[fieldLookupIndex]; ok {
		if rt, ok := ft.(*RecursiveDecoder); ok {
			f.recursiveFieldStorage[fieldLookupIndex] = rt
		}

		return ft, ok
	}

	// Ensure that a recursive type such as Xcm::TransferReserveAsset does not cause an infinite loop
	// by adding the RecursiveDecoder the first time the field is encountered.
	f.fieldStorage[fieldLookupIndex] = &RecursiveDecoder{}

	return nil, false
}

const (
	// ExpectedExtrinsicParamsCount is the count of generic params that we expect for a
	// generic Extrinsic type from the metadata.
	//
	// The parameters are expected to be in the following order:
	// 1. Address
	// 2. Call
	// 3. Signature
	// 4. Extra
	ExpectedExtrinsicParamsCount = 4
)

// genericExtrinsicPath represents the expected metadata path of a generic extrinsic.
var genericExtrinsicPath = types.Si1Path{
	"sp_runtime",
	"generic",
	"unchecked_extrinsic",
	"UncheckedExtrinsic",
}

// isGenericExtrinsic checks if the metadata path of the extrinsic path matches the one of the
// generic extrinsic.
func isGenericExtrinsic(path types.Si1Path) bool {
	if len(path) != len(genericExtrinsicPath) {
		return false
	}

	for i := range path {
		if path[i] != genericExtrinsicPath[i] {
			return false
		}
	}

	return true
}

// extractExtrinsicParams returns the extrinsic params if the provided extrinsic type is generic, otherwise,
// it extracts the generic extrinsic and then returns its params.
func extractExtrinsicParams(extrinsicType *types.Si1Type, meta *types.Metadata) ([]types.Si1TypeParameter, error) {
	if isGenericExtrinsic(extrinsicType.Path) {
		return extrinsicType.Params, nil
	}

	// If the metadata extrinsic type is not generic, its type is expected to be a composite with 1 field.
	if !extrinsicType.Def.IsComposite || len(extrinsicType.Def.Composite.Fields) != 1 {
		return nil, ErrInvalidExtrinsicType
	}

	// This composite field is the `sp_runtime::generic::unchecked_extrinsic::UncheckedExtrinsic`.
	genericUncheckedExtrinsic := extrinsicType.Def.Composite.Fields[0]

	genericUncheckedExtrinsicType := meta.AsMetadataV14.EfficientLookup[genericUncheckedExtrinsic.Type.Int64()]

	if !isGenericExtrinsic(genericUncheckedExtrinsicType.Path) {
		return nil, ErrInvalidGenericExtrinsicType
	}

	return genericUncheckedExtrinsicType.Params, nil
}

func getBitOrderString(path types.Si1Path) string {
	pathLen := len(path)

	if pathLen == 0 {
		return ""
	}

	return string(path[pathLen-1])
}

const (
	fieldSeparator    = "."
	lookupIndexFormat = "lookup_index_%d"
)

func getFieldPath(fieldType *types.Si1Type) string {
	var nameParts []string

	for _, pathEntry := range fieldType.Path {
		nameParts = append(nameParts, string(pathEntry))
	}

	return strings.Join(nameParts, fieldSeparator)
}

func getFullFieldName(field types.Si1Field, fieldType *types.Si1Type) string {
	fieldName := getFieldName(field)

	if fieldPath := getFieldPath(fieldType); fieldPath != "" {
		return fmt.Sprintf("%s%s%s", fieldPath, fieldSeparator, fieldName)
	}

	return getFieldName(field)
}

func getFieldName(field types.Si1Field) string {
	switch {
	case field.HasName:
		return string(field.Name)
	case field.HasTypeName:
		return string(field.TypeName)
	default:
		return fmt.Sprintf(lookupIndexFormat, field.Type.Int64())
	}
}
