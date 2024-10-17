package parser

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/stretchr/testify/assert"
)

func TestEventParserFn_ParseEvents(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
				{
					Name:  "byte_value",
					Value: byte(65),
				},
				{
					Name:  "string_value",
					Value: "test",
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
		{
			Name: "test_event_2",
			Phase: &types.Phase{
				IsFinalization: true,
			},
			EventID: types.EventID([2]byte{1, 0}),
			EventFields: []testField{
				{
					Name:  "u8_value",
					Value: types.NewU8(11),
				},
				{
					Name:  "u16_value",
					Value: types.NewU16(121),
				},
				{
					Name:  "u32_value",
					Value: types.NewU32(12134),
				},
				{
					Name:  "u64_value",
					Value: types.NewU64(128678),
				},
				{
					Name:  "u128_value",
					Value: types.NewU128(*big.NewInt(56346)),
				},
				{
					Name:  "u256_value",
					Value: types.NewU256(*big.NewInt(5674)),
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{3, 4, 5}),
			},
		},
		{
			Name: "test_event_3",
			Phase: &types.Phase{
				IsInitialization: true,
			},
			EventID: types.EventID([2]byte{1, 1}),
			EventFields: []testField{
				{
					Name:  "i8_value",
					Value: types.NewI8(45),
				},
				{
					Name:  "i16_value",
					Value: types.NewI16(445),
				},
				{
					Name:  "i32_value",
					Value: types.NewI32(545),
				},
				{
					Name:  "i64_value",
					Value: types.NewI64(4789),
				},
				{
					Name:  "i128_value",
					Value: types.NewI128(*big.NewInt(56747)),
				},
				{
					Name:  "i256_value",
					Value: types.NewI256(*big.NewInt(45356747)),
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{6, 7, 8}),
			},
		},
	}

	encodedEvents, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	res, err := eventParser.ParseEvents(reg, encodedEvents)
	assert.NoError(t, err)
	assert.Len(t, res, len(testEvents))

	for i, testEvent := range testEvents {
		assert.Equal(t, testEvent.Name, res[i].Name)
		assert.Equal(t, testEvent.EventID, res[i].EventID)
		assertEventFieldInformationIsCorrect(t, testEvent.EventFields, res[i])
		assert.Equal(t, testEvent.Phase, res[i].Phase)
		assert.Equal(t, testEvent.Topics, res[i].Topics)
	}
}

func TestEventParserFn_ParseEvents_EventCountDecodeError(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	// No storage data
	res, err := eventParser.ParseEvents(reg, &types.StorageDataRaw{})
	assert.ErrorIs(t, err, ErrEventsCountDecoding)
	assert.Nil(t, res)
}

func TestEventParserFn_ParseEvents_PhaseDecodeError(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	var b []byte

	buf := bytes.NewBuffer(b)

	encoder := scale.NewEncoder(buf)

	err = encoder.Encode(types.NewUCompactFromUInt(uint64(len(testEvents))))
	assert.NoError(t, err)

	// The storage data only contains the number of events, the phase decode step should fail in this case.
	storageData := types.StorageDataRaw(buf.Bytes())

	res, err := eventParser.ParseEvents(reg, &storageData)
	assert.ErrorIs(t, err, ErrEventPhaseDecoding)
	assert.Nil(t, res)
}

func TestEventParserFn_ParseEvents_EventIDDecodeError(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	var b []byte

	buf := bytes.NewBuffer(b)

	encoder := scale.NewEncoder(buf)

	err = encoder.Encode(types.NewUCompactFromUInt(uint64(len(testEvents))))
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].Phase)
	assert.NoError(t, err)

	// The storage data only contains:
	// - the number of events
	// - the phase
	// EventID decoding should fail in this case.
	storageData := types.StorageDataRaw(buf.Bytes())

	res, err := eventParser.ParseEvents(reg, &storageData)
	assert.ErrorIs(t, err, ErrEventIDDecoding)
	assert.Nil(t, res)
}

func TestEventParserFn_ParseEvents_EventFieldsDecodeError(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	var b []byte

	buf := bytes.NewBuffer(b)

	encoder := scale.NewEncoder(buf)

	err = encoder.Encode(types.NewUCompactFromUInt(uint64(len(testEvents))))
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].Phase)
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].EventID)
	assert.NoError(t, err)

	// The storage data only contains:
	// - the number of events
	// - the phase
	// - the event ID
	// Event field decoding should fail in this case.
	storageData := types.StorageDataRaw(buf.Bytes())

	res, err := eventParser.ParseEvents(reg, &storageData)
	assert.ErrorIs(t, err, ErrEventFieldsDecoding)
	assert.Nil(t, res)
}

func TestEventParserFn_ParseEvents_MissingEventDecoder(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, _, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	var b []byte

	buf := bytes.NewBuffer(b)

	encoder := scale.NewEncoder(buf)

	err = encoder.Encode(types.NewUCompactFromUInt(uint64(len(testEvents))))
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].Phase)
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].EventID)
	assert.NoError(t, err)

	storageData := types.StorageDataRaw(buf.Bytes())

	// Empty registry, decoding should fail.
	res, err := eventParser.ParseEvents(registry.EventRegistry{}, &storageData)
	assert.ErrorIs(t, err, ErrEventDecoderNotFound)
	assert.Nil(t, res)
}

func TestEventParserFn_ParseEvents_TopicsDecodeError(t *testing.T) {
	testEvents := []testEvent{
		{
			Name: "test_event_1",
			Phase: &types.Phase{
				IsApplyExtrinsic: true,
				AsApplyExtrinsic: 1,
			},
			EventID: types.EventID([2]byte{0, 1}),
			EventFields: []testField{
				{
					Name:  "bool_value",
					Value: true,
				},
			},
			Topics: []types.Hash{
				types.NewHash([]byte{0, 1, 2}),
			},
		},
	}

	_, reg, err := getEventParsingTestData(testEvents)
	assert.NoError(t, err)

	eventParser := NewEventParser()

	var b []byte

	buf := bytes.NewBuffer(b)

	encoder := scale.NewEncoder(buf)

	err = encoder.Encode(types.NewUCompactFromUInt(uint64(len(testEvents))))
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].Phase)
	assert.NoError(t, err)

	err = encoder.Encode(testEvents[0].EventID)
	assert.NoError(t, err)

	for _, field := range testEvents[0].EventFields {
		err := encoder.Encode(field.Value)
		assert.NoError(t, err)
	}

	// Set topics length to 1
	err = encoder.Encode(types.NewUCompactFromUInt(1))
	assert.NoError(t, err)

	// The storage data only contains:
	// - the number of events
	// - the phase
	// - the event ID
	// - the event fields
	// - the topics length
	// Topics decoding should fail in this case.
	storageData := types.StorageDataRaw(buf.Bytes())

	res, err := eventParser.ParseEvents(reg, &storageData)
	assert.ErrorIs(t, err, ErrEventTopicsDecoding)
	assert.Nil(t, res)
}

func assertEventFieldInformationIsCorrect(t *testing.T, testFields []testField, event *Event) {
	for testFieldIndex, testField := range testFields {
		assert.Equal(t, testField.Value, event.Fields[testFieldIndex].Value)
	}
}

func getEventParsingTestData(testEvents []testEvent) (*types.StorageDataRaw, registry.EventRegistry, error) {
	eventRegistry, err := getRegistryForTestEvents(testEvents)

	if err != nil {
		return nil, nil, err
	}

	encodedEventData, err := getEncodedEventData(testEvents)

	if err != nil {
		return nil, nil, err
	}

	storageData := types.StorageDataRaw(encodedEventData)

	return &storageData, eventRegistry, nil
}

func getEncodedEventData(testEvents []testEvent) ([]byte, error) {
	var encodedEventData []byte

	buf := bytes.NewBuffer(encodedEventData)

	encoder := scale.NewEncoder(buf)

	eventsLen := types.NewUCompactFromUInt(uint64(len(testEvents)))

	if err := encoder.Encode(eventsLen); err != nil {
		return nil, err
	}

	for _, testEvent := range testEvents {
		if err := testEvent.Encode(*encoder); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func getRegistryForTestEvents(testEvents []testEvent) (registry.EventRegistry, error) {
	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	for _, testEvent := range testEvents {
		regFields, err := getTestRegistryFields(testEvent.EventFields)

		if err != nil {
			return nil, err
		}

		eventRegistry[testEvent.EventID] = &registry.TypeDecoder{
			Name:   testEvent.Name,
			Fields: regFields,
		}
	}

	return eventRegistry, nil
}

func getTestRegistryFields(fields []testField) ([]*registry.Field, error) {
	var regFields []*registry.Field

	for _, field := range fields {
		regField, err := getTestRegistryField(field)

		if err != nil {
			return nil, err
		}

		regFields = append(regFields, regField)
	}

	return regFields, nil
}

func getTestRegistryField(field testField) (*registry.Field, error) {
	regField := &registry.Field{}
	regField.Name = field.Name

	switch field.Value.(type) {
	case bool:
		regField.FieldDecoder = &registry.ValueDecoder[bool]{}
	case byte:
		regField.FieldDecoder = &registry.ValueDecoder[byte]{}
	case string:
		regField.FieldDecoder = &registry.ValueDecoder[string]{}
	case types.U8:
		regField.FieldDecoder = &registry.ValueDecoder[types.U8]{}
	case types.U16:
		regField.FieldDecoder = &registry.ValueDecoder[types.U16]{}
	case types.U32:
		regField.FieldDecoder = &registry.ValueDecoder[types.U32]{}
	case types.U64:
		regField.FieldDecoder = &registry.ValueDecoder[types.U64]{}
	case types.U128:
		regField.FieldDecoder = &registry.ValueDecoder[types.U128]{}
	case types.U256:
		regField.FieldDecoder = &registry.ValueDecoder[types.U256]{}
	case types.I8:
		regField.FieldDecoder = &registry.ValueDecoder[types.I8]{}
	case types.I16:
		regField.FieldDecoder = &registry.ValueDecoder[types.I16]{}
	case types.I32:
		regField.FieldDecoder = &registry.ValueDecoder[types.I32]{}
	case types.I64:
		regField.FieldDecoder = &registry.ValueDecoder[types.I64]{}
	case types.I128:
		regField.FieldDecoder = &registry.ValueDecoder[types.I128]{}
	case types.I256:
		regField.FieldDecoder = &registry.ValueDecoder[types.I256]{}
	default:
		return nil, fmt.Errorf("type not supported - %v", field)
	}

	return regField, nil
}

type testField struct {
	Name  string
	Value any
}

type testEvent struct {
	Name        string
	Phase       *types.Phase
	EventID     types.EventID
	EventFields []testField
	Topics      []types.Hash
}

func (t testEvent) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(t.Phase); err != nil {
		return fmt.Errorf("couldn't encode phase: %w", err)
	}

	if err := encoder.Encode(t.EventID); err != nil {
		return fmt.Errorf("couldn't encode event ID: %w", err)
	}

	for _, field := range t.EventFields {
		if err := encoder.Encode(field.Value); err != nil {
			return fmt.Errorf("couldn't encode field %v: %w", field, err)
		}
	}

	if err := encoder.Encode(t.Topics); err != nil {
		return fmt.Errorf("couldn't encode topics: %w", err)
	}

	return nil
}
