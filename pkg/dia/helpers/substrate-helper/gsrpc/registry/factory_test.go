package registry

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/test"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestFactory_CreateErrorRegistryWithLiveMetadata(t *testing.T) {
	var tests = []struct {
		Chain       string
		MetadataHex string
	}{
		{
			Chain:       "centrifuge",
			MetadataHex: test.CentrifugeMetadataHex,
		},
		{
			Chain:       "polkadot",
			MetadataHex: test.PolkadotMetadataHex,
		},
		{
			Chain:       "acala",
			MetadataHex: test.AcalaMetaHex,
		},
		{
			Chain:       "statemint",
			MetadataHex: test.StatemintMetaHex,
		},
		{
			Chain:       "moonbeam",
			MetadataHex: test.MoonbeamMetaHex,
		},
	}

	for _, test := range tests {
		t.Run(test.Chain, func(t *testing.T) {
			var meta types.Metadata

			err := codec.DecodeFromHex(test.MetadataHex, &meta)
			assert.NoError(t, err)

			t.Log("Metadata was decoded successfully")

			factory := NewFactory()

			reg, err := factory.CreateErrorRegistry(&meta)
			assert.NoError(t, err)

			t.Log("Error registry was created successfully")

			testAsserter := newTestAsserter()

			for _, pallet := range meta.AsMetadataV14.Pallets {
				if !pallet.HasErrors {
					continue
				}

				errorsType, ok := meta.AsMetadataV14.EfficientLookup[pallet.Errors.Type.Int64()]
				assert.True(t, ok, fmt.Sprintf("Error type %d not found", pallet.Errors.Type.Int64()))

				assert.True(t, errorsType.Def.IsVariant, fmt.Sprintf("Error type %d not a variant", pallet.Events.Type.Int64()))

				for _, errorVariant := range errorsType.Def.Variant.Variants {
					errorID := ErrorID{
						ModuleIndex: pallet.Index,
						ErrorIndex:  [4]types.U8{errorVariant.Index},
					}

					registryErrorType, ok := reg[errorID]
					assert.True(t, ok, fmt.Sprintf("Error '%v' not found in registry", errorID))

					testAsserter.assertRegistryItemContainsAllTypes(t, meta, registryErrorType.Fields, errorVariant.Fields)
				}
			}
		})
	}
}

func TestFactory_CreateErrorRegistry_NoPalletWithErrors(t *testing.T) {
	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					HasErrors: false,
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateErrorRegistry(testMeta)
	assert.NoError(t, err)
	assert.Empty(t, reg)
}

func TestFactory_CreateErrorRegistry_ErrorsTypeNotFound(t *testing.T) {
	errorLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasErrors: true,
					Errors: types.ErrorMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(errorLookupTypeID)),
						},
					},
				},
			},
			// EfficientLookup map is empty causing an error.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateErrorRegistry(testMeta)
	assert.ErrorIs(t, err, ErrErrorsTypeNotFound)
	assert.Empty(t, reg)
}

func TestFactory_CreateErrorRegistry_ErrorsTypeNotAVariant(t *testing.T) {
	errorLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasErrors: true,
					Errors: types.ErrorMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(errorLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(errorLookupTypeID): {
					Def: types.Si1TypeDef{
						// Error type definition not a variant causing an error.
						IsVariant: false,
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateErrorRegistry(testMeta)
	assert.ErrorIs(t, err, ErrErrorsTypeNotVariant)
	assert.Empty(t, reg)
}

func TestFactory_CreateErrorRegistry_GetTypeFieldsError(t *testing.T) {
	errorLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasErrors: true,
					Errors: types.ErrorMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(errorLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(errorLookupTypeID): {
					Def: types.Si1TypeDef{
						IsVariant: true,
						Variant: types.Si1TypeDefVariant{
							Variants: []types.Si1Variant{
								{
									Name: "ErrorVariant1",
									Fields: []types.Si1Field{
										{
											HasName: true,
											Name:    "ErrorVariant1Field",
											Type: types.Si1LookupTypeID{
												// This lookup type ID is not added in the lookup map which should
												// cause an error.
												UCompact: types.NewUCompactFromUInt(uint64(456)),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateErrorRegistry(testMeta)
	assert.ErrorIs(t, err, ErrErrorFieldsRetrieval)
	assert.Empty(t, reg)
}

func TestFactory_CreateCallRegistry_WithLiveMetadata(t *testing.T) {
	var tests = []struct {
		Chain       string
		MetadataHex string
	}{
		{
			Chain:       "centrifuge",
			MetadataHex: test.CentrifugeMetadataHex,
		},
		{
			Chain:       "polkadot",
			MetadataHex: test.PolkadotMetadataHex,
		},
		{
			Chain:       "acala",
			MetadataHex: test.AcalaMetaHex,
		},
		{
			Chain:       "statemint",
			MetadataHex: test.StatemintMetaHex,
		},
		{
			Chain:       "moonbeam",
			MetadataHex: test.MoonbeamMetaHex,
		},
	}

	for _, test := range tests {
		t.Run(test.Chain, func(t *testing.T) {
			var meta types.Metadata

			err := codec.DecodeFromHex(test.MetadataHex, &meta)
			assert.NoError(t, err)

			t.Log("Metadata was decoded successfully")

			factory := NewFactory()

			reg, err := factory.CreateCallRegistry(&meta)
			assert.NoError(t, err)

			t.Log("Call registry was created successfully")

			testAsserter := newTestAsserter()

			for _, pallet := range meta.AsMetadataV14.Pallets {
				if !pallet.HasCalls {
					continue
				}

				callsType, ok := meta.AsMetadataV14.EfficientLookup[pallet.Calls.Type.Int64()]
				assert.True(t, ok, fmt.Sprintf("Calls type %d not found", pallet.Calls.Type.Int64()))

				assert.True(t, callsType.Def.IsVariant, fmt.Sprintf("Calls type %d not a variant", pallet.Events.Type.Int64()))

				for _, callVariant := range callsType.Def.Variant.Variants {
					callIndex := types.CallIndex{
						SectionIndex: uint8(pallet.Index),
						MethodIndex:  uint8(callVariant.Index),
					}

					callName := fmt.Sprintf("%s.%s", pallet.Name, callVariant.Name)

					registryCallType, ok := reg[callIndex]
					assert.True(t, ok, fmt.Sprintf("Call '%s' not found in registry", callName))

					testAsserter.assertRegistryItemContainsAllTypes(t, meta, registryCallType.Fields, callVariant.Fields)
				}
			}
		})
	}
}

func TestFactory_CreateCallRegistry_Overrides(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	t.Log("Metadata was decoded successfully")

	// Lookup index for U64 in the test.CentrifugeMetadataHex
	targetLookupIndex := int64(11)

	f := NewFactory().(*factory)

	_, err = f.CreateCallRegistry(&meta)
	assert.NoError(t, err)

	assert.Equal(t, f.fieldStorage[targetLookupIndex], &ValueDecoder[types.U64]{})

	// With override

	fieldOverride := FieldOverride{
		FieldLookupIndex: targetLookupIndex,
		FieldDecoder:     &ValueDecoder[types.I64]{},
	}

	f = NewFactory(fieldOverride).(*factory)

	_, err = f.CreateCallRegistry(&meta)
	assert.NoError(t, err)

	assert.Equal(t, f.fieldStorage[targetLookupIndex], &ValueDecoder[types.I64]{})
}

func TestFactory_CreateCallRegistry_NoPalletWithCalls(t *testing.T) {
	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					HasCalls: false,
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateCallRegistry(testMeta)
	assert.NoError(t, err)
	assert.Empty(t, reg)
}

func TestFactory_CreateCallRegistry_CallsTypeNotFound(t *testing.T) {
	callLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:     "TestModule",
					HasCalls: true,
					Calls: types.FunctionMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(callLookupTypeID)),
						},
					},
				},
			},
			// EfficientLookup map is empty causing an error.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateCallRegistry(testMeta)
	assert.ErrorIs(t, err, ErrCallsTypeNotFound)
	assert.Empty(t, reg)
}

func TestFactory_CreateCallRegistry_CallTypeNotAVariant(t *testing.T) {
	callLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:     "TestModule",
					HasCalls: true,
					Calls: types.FunctionMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(callLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(callLookupTypeID): {
					Def: types.Si1TypeDef{
						// Calls type definition not a variant causing an error.
						IsVariant: false,
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateCallRegistry(testMeta)
	assert.ErrorIs(t, err, ErrCallsTypeNotVariant)
	assert.Empty(t, reg)
}

func TestFactory_CreateCallRegistry_GetTypeFieldsError(t *testing.T) {
	callLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:     "TestModule",
					HasCalls: true,
					Calls: types.FunctionMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(callLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(callLookupTypeID): {
					Def: types.Si1TypeDef{
						IsVariant: true,
						Variant: types.Si1TypeDefVariant{
							Variants: []types.Si1Variant{
								{
									Name: "CallVariant1",
									Fields: []types.Si1Field{
										{
											HasName: true,
											Name:    "CallVariant1Field",
											Type: types.Si1LookupTypeID{
												// This lookup type ID is not added in the lookup map which should
												// cause an error.
												UCompact: types.NewUCompactFromUInt(uint64(456)),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateCallRegistry(testMeta)
	assert.ErrorIs(t, err, ErrCallFieldsRetrieval)
	assert.Empty(t, reg)
}

func TestFactory_CreateEventRegistry_WithLiveMetadata(t *testing.T) {
	var tests = []struct {
		Chain       string
		MetadataHex string
	}{
		{
			Chain:       "centrifuge",
			MetadataHex: test.CentrifugeMetadataHex,
		},
		{
			Chain:       "polkadot",
			MetadataHex: test.PolkadotMetadataHex,
		},
		{
			Chain:       "acala",
			MetadataHex: test.AcalaMetaHex,
		},
		{
			Chain:       "statemint",
			MetadataHex: test.StatemintMetaHex,
		},
		{
			Chain:       "moonbeam",
			MetadataHex: test.MoonbeamMetaHex,
		},
	}

	for _, test := range tests {
		t.Run(test.Chain, func(t *testing.T) {
			var meta types.Metadata

			err := codec.DecodeFromHex(test.MetadataHex, &meta)
			assert.NoError(t, err)

			t.Log("Metadata was decoded successfully")

			factory := NewFactory()

			reg, err := factory.CreateEventRegistry(&meta)
			assert.NoError(t, err)

			t.Log("Event registry was created successfully")

			testAsserter := newTestAsserter()

			for _, pallet := range meta.AsMetadataV14.Pallets {
				if !pallet.HasEvents {
					continue
				}

				eventsType, ok := meta.AsMetadataV14.EfficientLookup[pallet.Events.Type.Int64()]
				assert.True(t, ok, fmt.Sprintf("Events type %d not found", pallet.Events.Type.Int64()))

				assert.True(t, eventsType.Def.IsVariant, fmt.Sprintf("Events type %d not a variant", pallet.Events.Type.Int64()))

				for _, eventVariant := range eventsType.Def.Variant.Variants {
					eventID := types.EventID{byte(pallet.Index), byte(eventVariant.Index)}

					registryEventType, ok := reg[eventID]
					assert.True(t, ok, fmt.Sprintf("Event with ID %v not found in registry", eventID))

					testAsserter.assertRegistryItemContainsAllTypes(t, meta, registryEventType.Fields, eventVariant.Fields)
				}
			}
		})
	}
}

func TestFactory_CreateEventRegistry_Overrides(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	t.Log("Metadata was decoded successfully")

	// Lookup index for DispatchInfo in the test.CentrifugeMetadataHex
	targetLookupIndex := int64(21)

	fieldOverride := FieldOverride{
		FieldLookupIndex: targetLookupIndex,
		FieldDecoder:     &ValueDecoder[types.DispatchInfo]{},
	}

	f := NewFactory(fieldOverride).(*factory)

	_, err = f.CreateEventRegistry(&meta)
	assert.NoError(t, err)

	assert.Equal(t, f.fieldStorage[targetLookupIndex], &ValueDecoder[types.DispatchInfo]{})
}

func TestFactory_CreateEventRegistry_NoPalletWithEvents(t *testing.T) {
	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					HasEvents: false,
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateEventRegistry(testMeta)
	assert.NoError(t, err)
	assert.Empty(t, reg)
}

func TestFactory_CreateEventRegistry_EventsTypeNotFound(t *testing.T) {
	eventLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasEvents: true,
					Events: types.EventMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(eventLookupTypeID)),
						},
					},
				},
			},
			// EfficientLookup map is empty causing an error.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateEventRegistry(testMeta)
	assert.ErrorIs(t, err, ErrEventsTypeNotFound)
	assert.Empty(t, reg)
}

func TestFactory_CreateEventRegistry_EventTypeNotAVariant(t *testing.T) {
	callLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasEvents: true,
					Events: types.EventMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(callLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(callLookupTypeID): {
					Def: types.Si1TypeDef{
						// Events type definition not a variant causing an error.
						IsVariant: false,
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateEventRegistry(testMeta)
	assert.ErrorIs(t, err, ErrEventsTypeNotVariant)
	assert.Empty(t, reg)
}

func TestFactory_CreateEventRegistry_GetTypeFieldError(t *testing.T) {
	callLookupTypeID := 123

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Pallets: []types.PalletMetadataV14{
				{
					Name:      "TestModule",
					HasEvents: true,
					Events: types.EventMetadataV14{
						Type: types.Si1LookupTypeID{
							UCompact: types.NewUCompactFromUInt(uint64(callLookupTypeID)),
						},
					},
				},
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(callLookupTypeID): {
					Def: types.Si1TypeDef{
						IsVariant: true,
						Variant: types.Si1TypeDefVariant{
							Variants: []types.Si1Variant{
								{
									Name: "EventVariant1",
									Fields: []types.Si1Field{
										{
											HasName: true,
											Name:    "EventVariant1Field",
											Type: types.Si1LookupTypeID{
												// This lookup type ID is not added in the lookup map which should
												// cause an error.
												UCompact: types.NewUCompactFromUInt(uint64(456)),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	reg, err := factory.CreateEventRegistry(testMeta)
	assert.ErrorIs(t, err, ErrEventFieldsRetrieval)
	assert.Empty(t, reg)
}

func TestFactory_CreateExtrinsicDecoder_WithLiveMetadata(t *testing.T) {
	var tests = []struct {
		Chain       string
		MetadataHex string
	}{
		{
			Chain:       "centrifuge",
			MetadataHex: test.CentrifugeMetadataHex,
		},
		{
			Chain:       "polkadot",
			MetadataHex: test.PolkadotMetadataHex,
		},
		{
			Chain:       "acala",
			MetadataHex: test.AcalaMetaHex,
		},
		{
			Chain:       "statemint",
			MetadataHex: test.StatemintMetaHex,
		},
		{
			Chain:       "moonbeam",
			MetadataHex: test.MoonbeamMetaHex,
		},
	}

	for _, test := range tests {
		t.Run(test.Chain, func(t *testing.T) {
			var meta types.Metadata

			err := codec.DecodeFromHex(test.MetadataHex, &meta)
			assert.NoError(t, err)

			factory := NewFactory()

			extDecoder, err := factory.CreateExtrinsicDecoder(&meta)
			assert.NoError(t, err)
			assert.NotNil(t, extDecoder)
		})
	}
}

func TestFactory_CreateExtrinsicDecoder_ExtrinsicParamsExtraction_InvalidExtrinsicTypeError(t *testing.T) {
	extrinsicLookupID := uint64(123)

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Extrinsic: types.ExtrinsicV14{
				Type:             types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(extrinsicLookupID)},
				Version:          0,
				SignedExtensions: nil,
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(extrinsicLookupID): {
					Def: types.Si1TypeDef{
						// `extractExtrinsicParams` expects a composite type with 1 field.
						IsPrimitive: true,
					},
				},
			},
		},
	}

	factory := NewFactory()

	res, err := factory.CreateExtrinsicDecoder(testMeta)
	assert.ErrorIs(t, err, ErrInvalidExtrinsicType)
	assert.Nil(t, res)
}

func TestFactory_CreateExtrinsicDecoder_ExtrinsicParamsExtraction_InvalidGenericExtrinsicTypeError(t *testing.T) {
	extrinsicLookupID := uint64(123)
	genericExtrinsicLookupID := uint64(456)

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Extrinsic: types.ExtrinsicV14{
				Type:             types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(extrinsicLookupID)},
				Version:          0,
				SignedExtensions: nil,
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(extrinsicLookupID): {
					Def: types.Si1TypeDef{
						IsComposite: true,
						Composite: types.Si1TypeDefComposite{
							Fields: []types.Si1Field{
								{
									Type: types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(genericExtrinsicLookupID)},
								},
							},
						},
					},
				},
				int64(genericExtrinsicLookupID): {
					// No path provided here will cause `isGenericExtrinsic` to returns false
					// on the second check from `extractExtrinsicParams`.
					Def: types.Si1TypeDef{},
				},
			},
		},
	}

	factory := NewFactory()

	res, err := factory.CreateExtrinsicDecoder(testMeta)
	assert.ErrorIs(t, err, ErrInvalidGenericExtrinsicType)
	assert.Nil(t, res)
}

func TestFactory_CreateExtrinsicDecoder_InvalidExtrinsicParams(t *testing.T) {
	extrinsicLookupID := uint64(123)

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Extrinsic: types.ExtrinsicV14{
				Type:             types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(extrinsicLookupID)},
				Version:          0,
				SignedExtensions: nil,
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(extrinsicLookupID): {
					Path: genericExtrinsicPath,
					Params: []types.Si1TypeParameter{
						{
							Name:    "param_1",
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	res, err := factory.CreateExtrinsicDecoder(testMeta)
	assert.ErrorIs(t, err, ErrInvalidExtrinsicParams)
	assert.Nil(t, res)
}

func TestFactory_CreateExtrinsicDecoder_UnexpectedExtrinsicParam(t *testing.T) {
	extrinsicLookupID := uint64(123)

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Extrinsic: types.ExtrinsicV14{
				Type:             types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(extrinsicLookupID)},
				Version:          0,
				SignedExtensions: nil,
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(extrinsicLookupID): {
					Path: genericExtrinsicPath,
					Params: []types.Si1TypeParameter{
						{
							Name:    "param_1",
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
						{
							Name:    "param_2",
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
						{
							Name:    "param_3",
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
						{
							Name:    "param_4",
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	res, err := factory.CreateExtrinsicDecoder(testMeta)
	assert.ErrorIs(t, err, ErrUnexpectedExtrinsicParam)
	assert.Nil(t, res)
}

func TestFactory_CreateExtrinsicDecoder_ExtrinsicFieldRetrievalError(t *testing.T) {
	extrinsicLookupID := uint64(123)

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			Extrinsic: types.ExtrinsicV14{
				Type:             types.Si1LookupTypeID{UCompact: types.NewUCompactFromUInt(extrinsicLookupID)},
				Version:          0,
				SignedExtensions: nil,
			},
			EfficientLookup: map[int64]*types.Si1Type{
				int64(extrinsicLookupID): {
					Path: genericExtrinsicPath,
					Params: []types.Si1TypeParameter{
						{
							Name:    ExtrinsicAddressName,
							HasType: false,
							Type: types.Si1LookupTypeID{
								// This lookup type ID is not added in the lookup map which should
								// cause an error.
								UCompact: types.NewUCompactFromUInt(uint64(456)),
							},
						},
						{
							Name:    ExtrinsicSignatureName,
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
						{
							Name:    ExtrinsicExtraName,
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
						{
							Name:    ExtrinsicCallName,
							HasType: false,
							Type:    types.Si1LookupTypeID{},
						},
					},
				},
			},
		},
	}

	factory := NewFactory()

	res, err := factory.CreateExtrinsicDecoder(testMeta)
	assert.ErrorIs(t, err, ErrExtrinsicFieldRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getTypeFields(t *testing.T) {
	fieldLookUpID := 123

	testFieldName := "TestFieldName"
	testFields := []types.Si1Field{
		{
			HasName: true,
			Name:    types.Text(testFieldName),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(fieldLookUpID)),
			},
		},
	}

	compactFieldTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsCompact: true,
		Compact: types.Si1TypeDefCompact{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compactFieldTypeLookupID)),
			},
		},
	}

	compactFieldTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(fieldLookUpID): {
					Def: testFieldTypeDef,
				},
				int64(compactFieldTypeLookupID): {
					Def: compactFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getTypeFields(testMeta, testFields)
	assert.NoError(t, err)
	assert.Len(t, res, 1)

	assert.Equal(t, testFieldName, res[0].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, res[0].FieldDecoder)
	assert.Equal(t, int64(fieldLookUpID), res[0].LookupIndex)
}

func TestFactory_getTypeParams(t *testing.T) {
	fieldLookUpID := 123

	testFieldName := "TestFieldName"
	testParams := []types.Si1TypeParameter{
		{
			Name: types.Text(testFieldName),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(fieldLookUpID)),
			},
		},
	}

	compactFieldTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsCompact: true,
		Compact: types.Si1TypeDefCompact{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compactFieldTypeLookupID)),
			},
		},
	}

	compactFieldTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(fieldLookUpID): {
					Def: testFieldTypeDef,
				},
				int64(compactFieldTypeLookupID): {
					Def: compactFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getTypeParams(testMeta, testParams)
	assert.NoError(t, err)
	assert.Len(t, res, 1)

	assert.Equal(t, testFieldName, res[0].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, res[0].FieldDecoder)
	assert.Equal(t, int64(fieldLookUpID), res[0].LookupIndex)
}

func TestFactory_getTypeFields_FieldDecoderRetrievalError(t *testing.T) {
	fieldLookUpID := 123

	testFieldName := "TestFieldName"
	testFields := []types.Si1Field{
		{
			HasName: true,
			Name:    types.Text(testFieldName),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(fieldLookUpID)),
			},
		},
	}

	compositeFieldTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					HasName: true,
					Name:    "CompositeField1",
					Type: types.Si1LookupTypeID{
						// This lookup ID is not in the efficient lookup map which should cause an error.
						UCompact: types.NewUCompactFromUInt(uint64(compositeFieldTypeLookupID)),
					},
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(fieldLookUpID): {
					Def: testFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getTypeFields(testMeta, testFields)
	assert.ErrorIs(t, err, ErrFieldDecoderRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getTypeFields_FieldTypeNotFoundError(t *testing.T) {
	fieldLookUpID := 123

	testFieldName := "TestFieldName"
	testFields := []types.Si1Field{
		{
			HasName: true,
			Name:    types.Text(testFieldName),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(fieldLookUpID)),
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			// Efficient lookup map is empty to ensure that an error is returned.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getTypeFields(testMeta, testFields)
	assert.ErrorIs(t, err, ErrFieldTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_UnsupportedTypeError(t *testing.T) {
	testFieldName := "TestFieldName"

	testFieldTypeDef := types.Si1TypeDef{
		IsHistoricMetaCompat: true,
	}

	testMeta := &types.Metadata{}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrFieldTypeDefinitionNotSupported)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_Compact(t *testing.T) {
	testFieldName := "TestFieldName"

	compactFieldTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsCompact: true,
		Compact: types.Si1TypeDefCompact{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compactFieldTypeLookupID)),
			},
		},
	}

	compactFieldTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(compactFieldTypeLookupID): {
					Def: compactFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, res)
}

func TestFactory_getFieldDecoder_Compact_TypeNotFoundError(t *testing.T) {
	testFieldName := "TestFieldName"

	compactFieldTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsCompact: true,
		Compact: types.Si1TypeDefCompact{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compactFieldTypeLookupID)),
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrCompactFieldTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_Composite(t *testing.T) {
	testFieldName := "TestFieldName"

	compositeFieldTypeLookupID1 := 123
	compositeFieldTypeLookupID2 := 456

	compositeFieldName1 := "CompositeFieldName1"
	compositeFieldName2 := "CompositeFieldName2"

	compositeFields := []types.Si1Field{
		{
			HasName: true,
			Name:    types.Text(compositeFieldName1),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compositeFieldTypeLookupID1)),
			},
		},
		{
			HasName: true,
			Name:    types.Text(compositeFieldName2),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compositeFieldTypeLookupID2)),
			},
		},
	}
	testFieldTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: compositeFields,
		},
	}

	compositeFieldTypeDef1 := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	compositeFieldTypeDef2 := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsI8),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(compositeFieldTypeLookupID1): {
					Def: compositeFieldTypeDef1,
				},
				int64(compositeFieldTypeLookupID2): {
					Def: compositeFieldTypeDef2,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	compositeFieldType, ok := res.(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeFieldType.Fields, 2)

	assert.Equal(t, &ValueDecoder[types.U8]{}, compositeFieldType.Fields[0].FieldDecoder)
	assert.Equal(t, compositeFieldName1, compositeFieldType.Fields[0].Name)
	assert.Equal(t, int64(compositeFieldTypeLookupID1), compositeFieldType.Fields[0].LookupIndex)

	assert.Equal(t, &ValueDecoder[types.I8]{}, compositeFieldType.Fields[1].FieldDecoder)
	assert.Equal(t, compositeFieldName2, compositeFieldType.Fields[1].Name)
	assert.Equal(t, int64(compositeFieldTypeLookupID2), compositeFieldType.Fields[1].LookupIndex)
}

func TestFactory_getFieldDecoder_Composite_FieldError(t *testing.T) {
	testFieldName := "TestFieldName"

	compositeFieldTypeLookupID1 := 123
	compositeFieldTypeLookupID2 := 456

	compositeFieldName1 := "CompositeFieldName1"
	compositeFieldName2 := "CompositeFieldName2"

	compositeFields := []types.Si1Field{
		{
			HasName: true,
			Name:    types.Text(compositeFieldName1),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compositeFieldTypeLookupID1)),
			},
		},
		{
			HasName: true,
			Name:    types.Text(compositeFieldName2),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(compositeFieldTypeLookupID2)),
			},
		},
	}
	testFieldTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: compositeFields,
		},
	}

	compositeFieldTypeDef1 := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(compositeFieldTypeLookupID1): {
					Def: compositeFieldTypeDef1,
				},
				// Omitting the entry for composite field 2 should cause and error when parsing the composite fields.
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrCompositeTypeFieldsRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_Variant(t *testing.T) {
	testFieldName := "TestField"

	variantName1 := "Variant1"
	variantName2 := "Variant2"

	variantFieldName := "VariantFieldName"
	variantFieldLookupID := 123
	variantFieldTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testFieldTypeDef := types.Si1TypeDef{
		IsVariant: true,
		Variant: types.Si1TypeDefVariant{
			Variants: []types.Si1Variant{
				{
					Name:   types.Text(variantName1),
					Fields: nil,
					Index:  0,
				},
				{
					Name: types.Text(variantName2),
					Fields: []types.Si1Field{
						{
							HasName: true,
							Name:    types.Text(variantFieldName),
							Type: types.Si1LookupTypeID{
								UCompact: types.NewUCompactFromUInt(uint64(variantFieldLookupID)),
							},
						},
					},
					Index: 1,
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(variantFieldLookupID): {
					Def: variantFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	variantFieldType, ok := res.(*VariantDecoder)
	assert.True(t, ok)
	assert.Len(t, variantFieldType.FieldDecoderMap, 2)

	assert.Equal(t, &NoopDecoder{}, variantFieldType.FieldDecoderMap[0])

	compositeVariant, ok := variantFieldType.FieldDecoderMap[1].(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeVariant.Fields, 1)

	assert.Equal(t, variantFieldName, compositeVariant.Fields[0].Name)
	assert.Equal(t, &ValueDecoder[types.U8]{}, compositeVariant.Fields[0].FieldDecoder)
	assert.Equal(t, int64(variantFieldLookupID), compositeVariant.Fields[0].LookupIndex)
}

func TestFactory_getFieldDecoder_Primitive(t *testing.T) {
	testFieldName := "TestFieldName"

	testFieldTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)
	assert.Equal(t, &ValueDecoder[types.U8]{}, res)
}

func TestFactory_getFieldDecoder_Array(t *testing.T) {
	testFieldName := "TestFieldName"

	arrayItemTypeLookupID := 456
	arrayLen := 32

	testFieldTypeDef := types.Si1TypeDef{
		IsArray: true,
		Array: types.Si1TypeDefArray{
			Len: types.U32(arrayLen),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(arrayItemTypeLookupID)),
			},
		},
	}

	arrayItemTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(arrayItemTypeLookupID): {
					Def: arrayItemTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	arrayFieldType, ok := res.(*ArrayDecoder)
	assert.True(t, ok)

	assert.Equal(t, uint(arrayLen), arrayFieldType.Length)
	assert.Equal(t, &ValueDecoder[types.U8]{}, arrayFieldType.ItemDecoder)
}

func TestFactory_getFieldDecoder_Array_TypeNotFoundError(t *testing.T) {
	testFieldName := "TestFieldName"

	arrayItemTypeLookupID := 456
	arrayLen := 32

	testFieldTypeDef := types.Si1TypeDef{
		IsArray: true,
		Array: types.Si1TypeDefArray{
			Len: types.U32(arrayLen),
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(arrayItemTypeLookupID)),
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			// The lookup map does not contain the array item type lookup ID which should cause an error.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrArrayFieldTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_Slice(t *testing.T) {
	testFieldName := "TestFieldName"

	sliceItemTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsSequence: true,
		Sequence: types.Si1TypeDefSequence{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(sliceItemTypeLookupID)),
			},
		},
	}

	sliceItemTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU256),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(sliceItemTypeLookupID): {
					Def: sliceItemTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	sliceFieldType, ok := res.(*SliceDecoder)
	assert.True(t, ok)

	assert.Equal(t, &ValueDecoder[types.U256]{}, sliceFieldType.ItemDecoder)
}

func TestFactory_getFieldDecoder_Slice_TypeNotFoundError(t *testing.T) {
	testFieldName := "TestFieldName"

	sliceItemTypeLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsSequence: true,
		Sequence: types.Si1TypeDefSequence{
			Type: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(sliceItemTypeLookupID)),
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			// The lookup map does not contain the array item type lookup ID which should cause an error.
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrVectorFieldTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_Tuple(t *testing.T) {
	testFieldName := "TestFieldName"

	tupleItemLookupID1 := 123
	tupleItemLookupID2 := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsTuple: true,
		Tuple: []types.Si1LookupTypeID{
			{
				UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID1)),
			},
			{
				UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID2)),
			},
		},
	}

	tupleItem1TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsChar),
		},
	}

	tupleItem2TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsI16),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(tupleItemLookupID1): {
					Def: tupleItem1TypeDef,
				},
				int64(tupleItemLookupID2): {
					Def: tupleItem2TypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	compositeFieldType, ok := res.(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeFieldType.Fields, 2)
	assert.Equal(t, testFieldName, compositeFieldType.FieldName)

	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 0), compositeFieldType.Fields[0].Name)
	assert.Equal(t, &ValueDecoder[byte]{}, compositeFieldType.Fields[0].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID1), compositeFieldType.Fields[0].LookupIndex)

	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 1), compositeFieldType.Fields[1].Name)
	assert.Equal(t, &ValueDecoder[types.I16]{}, compositeFieldType.Fields[1].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID2), compositeFieldType.Fields[1].LookupIndex)
}

func TestFactory_getFieldDecoder_Tuple_NilTuple(t *testing.T) {
	testFieldName := "TestFieldName"

	testFieldTypeDef := types.Si1TypeDef{
		IsTuple: true,
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)
	assert.Equal(t, &NoopDecoder{}, res)
}

func TestFactory_getFieldDecoder_BitSequence(t *testing.T) {
	testFieldName := "TestFieldName"

	bitStoreLookupID := 123
	bitOrderLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsBitSequence: true,
		BitSequence: types.Si1TypeDefBitSequence{
			BitStoreType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitStoreLookupID)),
			},
			BitOrderType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitOrderLookupID)),
			},
		},
	}

	bitStoreTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	bitOrderType := &types.Si1Type{
		Path: []types.Text{
			types.Text(types.BitOrderName[types.BitOrderLsb0]),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(bitStoreLookupID): {
					Def: bitStoreTypeDef,
				},
				int64(bitOrderLookupID): bitOrderType,
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.NoError(t, err)

	bitSequenceDecoder, ok := res.(*BitSequenceDecoder)
	assert.True(t, ok)

	assert.Equal(t, testFieldName, bitSequenceDecoder.FieldName)
	assert.Equal(t, types.BitOrderLsb0, bitSequenceDecoder.BitOrder)
}

func TestFactory_getFieldDecoder_BitSequence_BitStoreTypeNotFound(t *testing.T) {
	testFieldName := "TestFieldName"

	bitStoreLookupID := 123
	bitOrderLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsBitSequence: true,
		BitSequence: types.Si1TypeDefBitSequence{
			BitStoreType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitStoreLookupID)),
			},
			BitOrderType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitOrderLookupID)),
			},
		},
	}

	bitOrderTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsI256),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(bitOrderLookupID): {
					Def: bitOrderTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrBitStoreTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_BitSequence_BitStoreFieldTypeError(t *testing.T) {
	testFieldName := "TestFieldName"

	bitStoreLookupID := 123
	bitOrderLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsBitSequence: true,
		BitSequence: types.Si1TypeDefBitSequence{
			BitStoreType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitStoreLookupID)),
			},
			BitOrderType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitOrderLookupID)),
			},
		},
	}

	bitStoreTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU16),
		},
	}

	bitOrderTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsI256),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(bitStoreLookupID): {
					Def: bitStoreTypeDef,
				},
				int64(bitOrderLookupID): {
					Def: bitOrderTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrBitStoreTypeNotSupported)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_BitSequence_BitOrderTypeNotFound(t *testing.T) {
	testFieldName := "TestFieldName"

	bitStoreLookupID := 123
	bitOrderLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsBitSequence: true,
		BitSequence: types.Si1TypeDefBitSequence{
			BitStoreType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitStoreLookupID)),
			},
			BitOrderType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitOrderLookupID)),
			},
		},
	}

	bitStoreTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(bitStoreLookupID): {
					Def: bitStoreTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrBitOrderTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getFieldDecoder_BitSequence_BitOrderCreationError(t *testing.T) {
	testFieldName := "TestFieldName"

	bitStoreLookupID := 123
	bitOrderLookupID := 456

	testFieldTypeDef := types.Si1TypeDef{
		IsBitSequence: true,
		BitSequence: types.Si1TypeDefBitSequence{
			BitStoreType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitStoreLookupID)),
			},
			BitOrderType: types.Si1LookupTypeID{
				UCompact: types.NewUCompactFromUInt(uint64(bitOrderLookupID)),
			},
		},
	}

	bitStoreTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	bitOrderType := &types.Si1Type{
		Path: []types.Text{
			types.Text("unknown-order"),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(bitStoreLookupID): {
					Def: bitStoreTypeDef,
				},
				int64(bitOrderLookupID): bitOrderType,
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getFieldDecoder(testMeta, testFieldName, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrBitOrderCreation)
	assert.Nil(t, res)
}

func TestFactory_getVariantFieldType_CompositeVariantTypeFieldError(t *testing.T) {
	variantName1 := "Variant1"
	variantName2 := "Variant2"

	variantFieldName := "VariantFieldName"
	variantFieldLookupID := 123
	variantFieldTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					Name: "CompositeVariantField",
					Type: types.Si1LookupTypeID{
						// This lookup ID is not in the efficient lookup map which should cause an error.
						UCompact: types.NewUCompactFromUInt(uint64(456)),
					},
				},
			},
		},
	}

	testFieldTypeDef := types.Si1TypeDef{
		IsVariant: true,
		Variant: types.Si1TypeDefVariant{
			Variants: []types.Si1Variant{
				{
					Name:   types.Text(variantName1),
					Fields: nil,
					Index:  0,
				},
				{
					Name: types.Text(variantName2),
					Fields: []types.Si1Field{
						{
							HasName: true,
							Name:    types.Text(variantFieldName),
							Type: types.Si1LookupTypeID{
								UCompact: types.NewUCompactFromUInt(uint64(variantFieldLookupID)),
							},
						},
					},
					Index: 1,
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(variantFieldLookupID): {
					Def: variantFieldTypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)
	factory.resetStorages()

	res, err := factory.getVariantFieldDecoder(testMeta, testFieldTypeDef)
	assert.ErrorIs(t, err, ErrVariantTypeFieldsRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getCompactFieldType_CompactTuple(t *testing.T) {
	testFieldName := "TestFieldName"

	tupleItemLookupID1 := 111
	tupleItemLookupID2 := 222

	compactFieldTypeDef := types.Si1TypeDef{
		IsTuple: true,
		Tuple: []types.Si1LookupTypeID{
			{
				UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID1)),
			},
			{
				UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID2)),
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(tupleItemLookupID1): {
					Def: types.Si1TypeDef{
						IsPrimitive: true,
						Primitive: types.Si1TypeDefPrimitive{
							Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
						},
					},
				},
				int64(tupleItemLookupID2): {
					Def: types.Si1TypeDef{
						IsPrimitive: true,
						Primitive: types.Si1TypeDefPrimitive{
							Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU32),
						},
					},
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getCompactFieldDecoder(testMeta, testFieldName, compactFieldTypeDef)
	assert.NoError(t, err)

	compositeFieldType, ok := res.(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeFieldType.Fields, 2)

	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 0), compositeFieldType.Fields[0].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, compositeFieldType.Fields[0].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID1), compositeFieldType.Fields[0].LookupIndex)
	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 1), compositeFieldType.Fields[1].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, compositeFieldType.Fields[1].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID2), compositeFieldType.Fields[1].LookupIndex)
}

func TestFactory_getCompactFieldType_CompactComposite(t *testing.T) {
	testFieldName := "TestFieldName"

	compositeFieldName1 := "CompositeFieldName1"
	compositeFieldName2 := "CompositeFieldName2"

	compositeFieldLookupID1 := 111
	compositeFieldLookupID2 := 222

	compactFieldTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					HasName: true,
					Name:    types.Text(compositeFieldName1),
					Type: types.Si1LookupTypeID{
						UCompact: types.NewUCompactFromUInt(uint64(compositeFieldLookupID1)),
					},
				},
				{
					HasName: true,
					Name:    types.Text(compositeFieldName2),
					Type: types.Si1LookupTypeID{
						UCompact: types.NewUCompactFromUInt(uint64(compositeFieldLookupID2)),
					},
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(compositeFieldLookupID1): {
					Def: types.Si1TypeDef{
						IsPrimitive: true,
						Primitive: types.Si1TypeDefPrimitive{
							Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
						},
					},
				},
				int64(compositeFieldLookupID2): {
					Def: types.Si1TypeDef{
						IsPrimitive: true,
						Primitive: types.Si1TypeDefPrimitive{
							Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU32),
						},
					},
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getCompactFieldDecoder(testMeta, testFieldName, compactFieldTypeDef)
	assert.NoError(t, err)

	compositeFieldType, ok := res.(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeFieldType.Fields, 2)

	assert.Equal(t, compositeFieldName1, compositeFieldType.Fields[0].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, compositeFieldType.Fields[0].FieldDecoder)
	assert.Equal(t, int64(compositeFieldLookupID1), compositeFieldType.Fields[0].LookupIndex)
	assert.Equal(t, compositeFieldName2, compositeFieldType.Fields[1].Name)
	assert.Equal(t, &ValueDecoder[types.UCompact]{}, compositeFieldType.Fields[1].FieldDecoder)
	assert.Equal(t, int64(compositeFieldLookupID2), compositeFieldType.Fields[1].LookupIndex)
}

func TestFactory_getArrayFieldType(t *testing.T) {
	testFieldName := "TestFieldName"

	arrayLen := 32

	arrayItemTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{}

	factory := NewFactory().(*factory)

	res, err := factory.getArrayFieldDecoder(uint(arrayLen), testMeta, testFieldName, arrayItemTypeDef)
	assert.NoError(t, err)

	arrayFieldType, ok := res.(*ArrayDecoder)
	assert.True(t, ok)

	assert.Equal(t, uint(arrayLen), arrayFieldType.Length)
	assert.Equal(t, &ValueDecoder[types.U8]{}, arrayFieldType.ItemDecoder)
}

func TestFactory_getArrayFieldType_ItemFieldTypeError(t *testing.T) {
	testFieldName := "TestFieldName"

	arrayLen := 32

	compositeLookupID := 123

	arrayItemTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					HasName: true,
					Name:    "CompositeField1",
					Type: types.Si1LookupTypeID{
						// This lookup ID is not present in the efficient lookup map which should cause an error.
						UCompact: types.NewUCompactFromUInt(uint64(compositeLookupID)),
					},
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getArrayFieldDecoder(uint(arrayLen), testMeta, testFieldName, arrayItemTypeDef)
	assert.ErrorIs(t, err, ErrArrayItemFieldDecoderRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getSliceFieldType(t *testing.T) {
	testFieldName := "TestFieldName"

	sliceItemTypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{}

	factory := NewFactory().(*factory)

	res, err := factory.getSliceFieldDecoder(testMeta, testFieldName, sliceItemTypeDef)
	assert.NoError(t, err)

	sliceFieldType, ok := res.(*SliceDecoder)
	assert.True(t, ok)

	assert.Equal(t, &ValueDecoder[types.U8]{}, sliceFieldType.ItemDecoder)
}

func TestFactory_getSliceFieldType_ItemFieldTypeError(t *testing.T) {
	testFieldName := "TestFieldName"

	compositeLookupID := 123

	sliceItemTypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					HasName: true,
					Name:    "CompositeField1",
					Type: types.Si1LookupTypeID{
						// This lookup ID is not present in the efficient lookup map which should cause an error.
						UCompact: types.NewUCompactFromUInt(uint64(compositeLookupID)),
					},
				},
			},
		},
	}

	testMeta := &types.Metadata{}

	factory := NewFactory().(*factory)

	res, err := factory.getSliceFieldDecoder(testMeta, testFieldName, sliceItemTypeDef)
	assert.ErrorIs(t, err, ErrSliceItemFieldDecoderRetrieval)
	assert.Nil(t, res)
}

func TestFactory_getTupleType(t *testing.T) {
	testFieldName := "TestFieldName"

	tupleItemLookupID1 := 123
	tupleItemLookupID2 := 456

	tupleTypeDef := []types.Si1LookupTypeID{
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID1)),
		},
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID2)),
		},
	}

	tupleItem1TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	tupleItem2TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU32),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(tupleItemLookupID1): {
					Def: tupleItem1TypeDef,
				},
				int64(tupleItemLookupID2): {
					Def: tupleItem2TypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getTupleFieldDecoder(testMeta, testFieldName, tupleTypeDef)
	assert.NoError(t, err)

	compositeFieldType, ok := res.(*CompositeDecoder)
	assert.True(t, ok)
	assert.Len(t, compositeFieldType.Fields, 2)

	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 0), compositeFieldType.Fields[0].Name)
	assert.Equal(t, &ValueDecoder[types.U8]{}, compositeFieldType.Fields[0].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID1), compositeFieldType.Fields[0].LookupIndex)
	assert.Equal(t, fmt.Sprintf(tupleItemFieldNameFormat, 1), compositeFieldType.Fields[1].Name)
	assert.Equal(t, &ValueDecoder[types.U32]{}, compositeFieldType.Fields[1].FieldDecoder)
	assert.Equal(t, int64(tupleItemLookupID2), compositeFieldType.Fields[1].LookupIndex)
}

func TestFactory_getTupleType_TupleItemNotFound(t *testing.T) {
	testFieldName := "TestFieldName"

	tupleItemLookupID1 := 123
	tupleItemLookupID2 := 456

	tupleTypeDef := []types.Si1LookupTypeID{
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID1)),
		},
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID2)),
		},
	}

	tupleItem1TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(tupleItemLookupID1): {
					Def: tupleItem1TypeDef,
				},
				// Lookup ID for tuple item 2 is missing which should cause an error.
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getTupleFieldDecoder(testMeta, testFieldName, tupleTypeDef)
	assert.ErrorIs(t, err, ErrTupleItemTypeNotFound)
	assert.Nil(t, res)
}

func TestFactory_getTupleType_TupleItemFieldDecoderError(t *testing.T) {
	testFieldName := "TestFieldName"

	tupleItemLookupID1 := 123
	tupleItemLookupID2 := 456

	tupleTypeDef := []types.Si1LookupTypeID{
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID1)),
		},
		{
			UCompact: types.NewUCompactFromUInt(uint64(tupleItemLookupID2)),
		},
	}

	tupleItem1TypeDef := types.Si1TypeDef{
		IsPrimitive: true,
		Primitive: types.Si1TypeDefPrimitive{
			Si0TypeDefPrimitive: types.Si0TypeDefPrimitive(types.IsU8),
		},
	}

	tupleItem2TypeDef := types.Si1TypeDef{
		IsComposite: true,
		Composite: types.Si1TypeDefComposite{
			Fields: []types.Si1Field{
				{
					HasName: true,
					Name:    "CompositeField1",
					Type: types.Si1LookupTypeID{
						// This lookup ID is not in the efficient lookup map which should cause an error.
						UCompact: types.NewUCompactFromUInt(uint64(789)),
					},
				},
			},
		},
	}

	testMeta := &types.Metadata{
		AsMetadataV14: types.MetadataV14{
			EfficientLookup: map[int64]*types.Si1Type{
				int64(tupleItemLookupID1): {
					Def: tupleItem1TypeDef,
				},
				int64(tupleItemLookupID2): {
					Def: tupleItem2TypeDef,
				},
			},
		},
	}

	factory := NewFactory().(*factory)

	res, err := factory.getTupleFieldDecoder(testMeta, testFieldName, tupleTypeDef)
	assert.ErrorIs(t, err, ErrTupleItemFieldDecoderRetrieval)
	assert.Nil(t, res)
}

func TestFactory_Overrides(t *testing.T) {
	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	t.Log("Metadata was decoded successfully")

	// Lookup index for DispatchInfo in the test.CentrifugeMetadataHex
	targetLookupIndex := int64(22)

	fieldOverride := FieldOverride{
		FieldLookupIndex: targetLookupIndex,
		FieldDecoder:     &ValueDecoder[types.DispatchInfo]{},
	}

	f := NewFactory(fieldOverride).(*factory)

	reg, err := f.CreateEventRegistry(&meta)
	assert.NoError(t, err)

	assert.Equal(t, f.fieldStorage[targetLookupIndex], &ValueDecoder[types.DispatchInfo]{})

	// Event ID for System.ExtrinsicSuccess
	extrinsicSuccessEventID := types.EventID{0, 0}

	testDispatchInfo := types.DispatchInfo{
		Weight: types.Weight{
			RefTime:   types.NewUCompactFromUInt(1),
			ProofSize: types.NewUCompactFromUInt(2),
		},
		Class: types.DispatchClass{
			IsNormal: true,
		},
		PaysFee: types.Pays{
			IsYes: true,
		},
	}

	encodedTestDispatchInfo, err := codec.Encode(testDispatchInfo)
	assert.NoError(t, err)

	extrinsicSuccessEventDecoder, ok := reg[extrinsicSuccessEventID]
	assert.True(t, ok)

	res, err := extrinsicSuccessEventDecoder.Decode(scale.NewDecoder(bytes.NewReader(encodedTestDispatchInfo)))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	value, err := ProcessDecodedFieldValue[types.DispatchInfo](
		res,
		func(fieldIndex int, field *DecodedField) bool {
			return field.LookupIndex == targetLookupIndex
		},
		func(value any) (types.DispatchInfo, error) {
			dispatchInfo, ok := value.(types.DispatchInfo)
			assert.True(t, ok)

			return dispatchInfo, nil
		},
	)
	assert.Equal(t, testDispatchInfo, value)
}

func testDataToDecodedFields(data []any) DecodedFields {
	var res DecodedFields

	for i, datum := range data {
		res = append(res, &DecodedField{
			Name:        fmt.Sprintf("decoded_field_%d", i),
			Value:       datum,
			LookupIndex: int64(i),
		})
	}

	return res
}

func encodeTestData(data []any) ([]byte, error) {
	var res []byte

	for _, datum := range data {
		b, err := codec.Encode(datum)

		if err != nil {
			return nil, err
		}

		res = append(res, b...)
	}

	return res, nil
}

type testAsserter struct {
	recursiveTypeMap map[int64]struct{}
}

func newTestAsserter() *testAsserter {
	return &testAsserter{map[int64]struct{}{}}
}

func (a *testAsserter) assertRegistryItemContainsAllTypes(t *testing.T, meta types.Metadata, registryItemFields []*Field, metaItemFields []types.Si1Field) {
	for i, metaItemField := range metaItemFields {
		registryItemField := registryItemFields[i]
		registryItemFieldType := registryItemField.FieldDecoder
		metaLookupIndex := metaItemField.Type.Int64()

		if _, ok := a.recursiveTypeMap[metaLookupIndex]; ok {
			continue
		}

		if metaLookupIndex != registryItemField.LookupIndex {
			t.Fatalf("Field lookup index mismatch for field with index %d", i)
		}

		fieldType, ok := meta.AsMetadataV14.EfficientLookup[metaLookupIndex]
		assert.True(t, ok, "field type for field with type %d not found", metaItemField.Type.Int64())

		a.assertRegistryItemFieldIsCorrect(t, meta, registryItemFieldType, fieldType)

		if _, ok := registryItemField.FieldDecoder.(*RecursiveDecoder); ok {
			a.recursiveTypeMap[metaLookupIndex] = struct{}{}
		}
	}
}

func (a *testAsserter) assertRegistryItemFieldIsCorrect(t *testing.T, meta types.Metadata, registryItemFieldType FieldDecoder, metaFieldType *types.Si1Type) {
	metaFieldTypeDef := metaFieldType.Def

	switch {
	case metaFieldTypeDef.IsComposite:
		compositeRegistryFieldType, ok := registryItemFieldType.(*CompositeDecoder)

		if !ok {
			_, isRecursive := registryItemFieldType.(*RecursiveDecoder)
			assert.True(t, isRecursive, "expected recursive field")

			return
		}

		a.assertRegistryItemContainsAllTypes(t, meta, compositeRegistryFieldType.Fields, metaFieldTypeDef.Composite.Fields)
	case metaFieldTypeDef.IsVariant:
		variantRegistryFieldType, ok := registryItemFieldType.(*VariantDecoder)

		if !ok {
			_, isRecursive := registryItemFieldType.(*RecursiveDecoder)
			assert.True(t, isRecursive, "expected variant or recursive field")
			return
		}

		for _, variant := range metaFieldTypeDef.Variant.Variants {
			registryVariant, ok := variantRegistryFieldType.FieldDecoderMap[byte(variant.Index)]
			assert.True(t, ok, "expected registry variant")

			if len(variant.Fields) == 0 {
				_, ok = registryVariant.(*NoopDecoder)
				assert.True(t, ok, "expected noop decoder")
				continue
			}

			compositeRegistryField, ok := registryVariant.(*CompositeDecoder)
			assert.True(t, ok, "expected composite field type")

			a.assertRegistryItemContainsAllTypes(t, meta, compositeRegistryField.Fields, variant.Fields)
		}
	case metaFieldTypeDef.IsSequence:
		sliceRegistryField, ok := registryItemFieldType.(*SliceDecoder)

		if !ok {
			_, isRecursive := registryItemFieldType.(*RecursiveDecoder)
			assert.True(t, isRecursive, "expected recursive field")

			return
		}

		sequenceFieldType, ok := meta.AsMetadataV14.EfficientLookup[metaFieldTypeDef.Sequence.Type.Int64()]
		assert.True(t, ok, "couldn't get sequence field type")

		a.assertRegistryItemFieldIsCorrect(t, meta, sliceRegistryField.ItemDecoder, sequenceFieldType)
	case metaFieldTypeDef.IsArray:
		arrayRegistryField, ok := registryItemFieldType.(*ArrayDecoder)
		assert.True(t, ok, "expected array field type in registry")

		arrayFieldType, ok := meta.AsMetadataV14.EfficientLookup[metaFieldTypeDef.Array.Type.Int64()]
		assert.True(t, ok, "couldn't get array field type")

		a.assertRegistryItemFieldIsCorrect(t, meta, arrayRegistryField.ItemDecoder, arrayFieldType)
	case metaFieldTypeDef.IsTuple:
		if metaFieldTypeDef.Tuple == nil {
			_, ok := registryItemFieldType.(*NoopDecoder)
			assert.True(t, ok, "expected noop decoder")
			return
		}

		compositeRegistryFieldType, ok := registryItemFieldType.(*CompositeDecoder)

		if !ok {
			_, isRecursive := registryItemFieldType.(*RecursiveDecoder)
			assert.True(t, isRecursive, "expected composite or recursive field")
			return
		}

		for i, item := range metaFieldTypeDef.Tuple {
			itemTypeDef, ok := meta.AsMetadataV14.EfficientLookup[item.Int64()]
			assert.True(t, ok, "couldn't get tuple item field type")

			registryTupleItemFieldType := compositeRegistryFieldType.Fields[i].FieldDecoder

			a.assertRegistryItemFieldIsCorrect(t, meta, registryTupleItemFieldType, itemTypeDef)
		}
	case metaFieldTypeDef.IsPrimitive:
		primitiveFieldType, err := getPrimitiveDecoder(metaFieldTypeDef.Primitive.Si0TypeDefPrimitive)
		assert.NoError(t, err, "couldn't get primitive type")

		assert.Equal(t, primitiveFieldType, registryItemFieldType, "primitive field types should match")
	case metaFieldTypeDef.IsCompact:
		compactFieldType, ok := meta.AsMetadataV14.EfficientLookup[metaFieldTypeDef.Compact.Type.Int64()]
		assert.True(t, ok, "couldn't find compact field type")

		switch {
		case compactFieldType.Def.IsPrimitive:
			_, ok = registryItemFieldType.(*ValueDecoder[types.UCompact])
			assert.True(t, ok, "expected compact field type in registry")
		case compactFieldType.Def.IsTuple:
			if metaFieldTypeDef.Tuple == nil {
				_, ok := registryItemFieldType.(*NoopDecoder)
				assert.True(t, ok, "expected noop decoder")
				return
			}

			compositeRegistryField, ok := registryItemFieldType.(*CompositeDecoder)
			assert.True(t, ok, "expected composite field type in registry")

			for _, field := range compositeRegistryField.Fields {
				_, ok = field.FieldDecoder.(*ValueDecoder[types.UCompact])
				assert.True(t, ok, "expected compact field type in registry")
			}
		case compactFieldType.Def.IsComposite:
			compositeRegistryField, ok := registryItemFieldType.(*CompositeDecoder)
			assert.True(t, ok, "expected composite field type in registry")

			for _, field := range compositeRegistryField.Fields {
				_, ok = field.FieldDecoder.(*ValueDecoder[types.UCompact])
				assert.True(t, ok, "expected compact field type in registry")
			}
		default:
			t.Fatalf("unsupported compact field type")
		}
	case metaFieldTypeDef.IsBitSequence:
		bitSequenceDecoder, ok := registryItemFieldType.(*BitSequenceDecoder)
		assert.True(t, ok, "expected bit sequence field type in registry")

		bitOrderType, ok := meta.AsMetadataV14.EfficientLookup[metaFieldTypeDef.BitSequence.BitOrderType.Int64()]
		assert.True(t, ok, "expected bit order type")

		assert.Equal(t, types.BitOrderValue[getBitOrderString(bitOrderType.Path)], bitSequenceDecoder.BitOrder)
	case metaFieldTypeDef.IsHistoricMetaCompat:
		t.Fatalf("historic meta compat type not covered")
	}
}
