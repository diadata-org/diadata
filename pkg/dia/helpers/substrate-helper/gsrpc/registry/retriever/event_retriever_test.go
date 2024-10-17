package retriever

import (
	"errors"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/exec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/state"
	stateMocks "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state/mocks"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventRetriever_New(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	registryFactoryMock.On("CreateEventRegistry", latestMeta).
		Return(eventRegistry, nil).
		Once()

	res, err := NewEventRetriever(
		eventParserMock,
		eventProviderMock,
		stateRPCMock,
		registryFactoryMock,
		storageExecMock,
		parsingExecMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &eventRetriever{}, res)
}

func TestEventRetriever_New_InternalStateUpdateError(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	metadataRetrievalError := errors.New("error")

	stateRPCMock.On("GetMetadataLatest").
		Return(nil, metadataRetrievalError).
		Once()

	res, err := NewEventRetriever(
		eventParserMock,
		eventProviderMock,
		stateRPCMock,
		registryFactoryMock,
		storageExecMock,
		parsingExecMock,
	)
	assert.ErrorIs(t, err, ErrInternalStateUpdate)
	assert.Nil(t, res)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	registryFactoryError := errors.New("error")

	registryFactoryMock.On("CreateEventRegistry", latestMeta).
		Return(nil, registryFactoryError).
		Once()

	res, err = NewEventRetriever(
		eventParserMock,
		eventProviderMock,
		stateRPCMock,
		registryFactoryMock,
		storageExecMock,
		parsingExecMock,
	)
	assert.ErrorIs(t, err, ErrInternalStateUpdate)
	assert.Nil(t, res)
}

func TestEventRetriever_NewDefault(t *testing.T) {
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	res, err := NewDefaultEventRetriever(eventProviderMock, stateRPCMock)
	assert.NoError(t, err)
	assert.IsType(t, &eventRetriever{}, res)

	retriever := res.(*eventRetriever)
	assert.IsType(t, parser.NewEventParser(), retriever.eventParser)
	assert.IsType(t, registry.NewFactory(), retriever.registryFactory)
	assert.IsType(t, exec.NewRetryableExecutor[*types.StorageDataRaw](), retriever.eventStorageExecutor)
	assert.IsType(t, exec.NewRetryableExecutor[[]*parser.Event](), retriever.eventParsingExecutor)
	assert.Equal(t, latestMeta, retriever.meta)
	assert.NotNil(t, retriever.eventRegistry)
}

func TestEventRetriever_GetEvents(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	testMeta := &types.Metadata{}

	eventRetriever.meta = testMeta

	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	eventRetriever.eventRegistry = eventRegistry

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	storageEvents := &types.StorageDataRaw{}

	eventProviderMock.On("GetStorageEvents", testMeta, blockHash).
		Return(storageEvents, nil).
		Once()

	storageExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*types.StorageDataRaw, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, storageEvents, execFnRes)
			},
		).Return(storageEvents, nil)

	parsedEvents := []*parser.Event{}

	eventParserMock.On("ParseEvents", eventRegistry, storageEvents).
		Return(parsedEvents, nil).
		Once()

	parsingExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() ([]*parser.Event, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, parsedEvents, execFnRes)
			},
		).Return(parsedEvents, nil)

	res, err := eventRetriever.GetEvents(blockHash)
	assert.NoError(t, err)
	assert.Equal(t, parsedEvents, res)
}

func TestEventRetriever_GetEvents_StorageRetrievalError(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	testMeta := &types.Metadata{}

	eventRetriever.meta = testMeta

	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	eventRetriever.eventRegistry = eventRegistry

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	storageRetrievalError := errors.New("error")

	eventProviderMock.On("GetStorageEvents", testMeta, blockHash).
		Return(nil, storageRetrievalError).
		Once()

	stateRPCMock.On("GetMetadata", blockHash).
		Return(testMeta, nil).
		Once()

	registryFactoryMock.On("CreateEventRegistry", testMeta).
		Return(eventRegistry, nil).
		Once()

	storageExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*types.StorageDataRaw, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.ErrorIs(t, err, storageRetrievalError)
				assert.Nil(t, execFnRes)

				fallbackFn, ok := args.Get(1).(func() error)
				assert.True(t, ok)

				err = fallbackFn()
				assert.NoError(t, err)
			},
		).Return(&types.StorageDataRaw{}, storageRetrievalError)

	res, err := eventRetriever.GetEvents(blockHash)
	assert.ErrorIs(t, err, ErrStorageEventRetrieval)
	assert.Nil(t, res)
}

func TestEventRetriever_GetEvents_EventParsingError(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	testMeta := &types.Metadata{}

	eventRetriever.meta = testMeta

	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	eventRetriever.eventRegistry = eventRegistry

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	storageEvents := &types.StorageDataRaw{}

	eventProviderMock.On("GetStorageEvents", testMeta, blockHash).
		Return(storageEvents, nil).
		Once()

	storageExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*types.StorageDataRaw, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, storageEvents, execFnRes)
			},
		).Return(storageEvents, nil)

	eventParsingError := errors.New("error")

	eventParserMock.On("ParseEvents", eventRegistry, storageEvents).
		Return(nil, eventParsingError).
		Once()

	stateRPCMock.On("GetMetadata", blockHash).
		Return(testMeta, nil).
		Once()

	registryFactoryMock.On("CreateEventRegistry", testMeta).
		Return(eventRegistry, nil).
		Once()

	parsingExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() ([]*parser.Event, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.ErrorIs(t, err, eventParsingError)
				assert.Nil(t, execFnRes)

				fallbackFn, ok := args.Get(1).(func() error)
				assert.True(t, ok)

				err = fallbackFn()
				assert.NoError(t, err)
			},
		).Return([]*parser.Event{}, eventParsingError)

	res, err := eventRetriever.GetEvents(blockHash)
	assert.ErrorIs(t, err, ErrEventParsing)
	assert.Nil(t, res)
}

func TestEventRetriever_updateInternalState(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	testMeta := &types.Metadata{}

	eventRegistry := registry.EventRegistry(map[types.EventID]*registry.TypeDecoder{})

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	stateRPCMock.On("GetMetadata", blockHash).
		Return(testMeta, nil).
		Once()

	registryFactoryMock.On("CreateEventRegistry", testMeta).
		Return(eventRegistry, nil).
		Once()

	err := eventRetriever.updateInternalState(&blockHash)
	assert.NoError(t, err)
	assert.Equal(t, testMeta, eventRetriever.meta)
	assert.Equal(t, eventRegistry, eventRetriever.eventRegistry)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	registryFactoryMock.On("CreateEventRegistry", latestMeta).
		Return(eventRegistry, nil).
		Once()

	err = eventRetriever.updateInternalState(nil)
	assert.NoError(t, err)
	assert.Equal(t, latestMeta, eventRetriever.meta)
	assert.Equal(t, eventRegistry, eventRetriever.eventRegistry)
}

func TestEventRetriever_updateInternalState_MetadataRetrievalError(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	metadataRetrievalError := errors.New("error")

	stateRPCMock.On("GetMetadata", blockHash).
		Return(nil, metadataRetrievalError).
		Once()

	err := eventRetriever.updateInternalState(&blockHash)
	assert.ErrorIs(t, err, ErrMetadataRetrieval)

	stateRPCMock.On("GetMetadataLatest").
		Return(nil, metadataRetrievalError).
		Once()

	err = eventRetriever.updateInternalState(nil)
	assert.ErrorIs(t, err, ErrMetadataRetrieval)
}

func TestEventRetriever_updateInternalState_RegistryFactoryError(t *testing.T) {
	eventParserMock := parser.NewEventParserMock(t)
	eventProviderMock := state.NewEventProviderMock(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	storageExecMock := exec.NewRetryableExecutorMock[*types.StorageDataRaw](t)
	parsingExecMock := exec.NewRetryableExecutorMock[[]*parser.Event](t)

	eventRetriever := &eventRetriever{
		eventParser:          eventParserMock,
		eventProvider:        eventProviderMock,
		stateRPC:             stateRPCMock,
		registryFactory:      registryFactoryMock,
		eventStorageExecutor: storageExecMock,
		eventParsingExecutor: parsingExecMock,
	}

	testMeta := &types.Metadata{}

	blockHash := types.NewHash([]byte{0, 1, 2, 3})

	stateRPCMock.On("GetMetadata", blockHash).
		Return(testMeta, nil).
		Once()

	registryFactoryError := errors.New("error")

	registryFactoryMock.On("CreateEventRegistry", testMeta).
		Return(nil, registryFactoryError).
		Once()

	err := eventRetriever.updateInternalState(&blockHash)
	assert.ErrorIs(t, err, ErrEventRegistryCreation)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	registryFactoryMock.On("CreateEventRegistry", latestMeta).
		Return(nil, registryFactoryError).
		Once()

	err = eventRetriever.updateInternalState(nil)
	assert.ErrorIs(t, err, ErrEventRegistryCreation)
}
