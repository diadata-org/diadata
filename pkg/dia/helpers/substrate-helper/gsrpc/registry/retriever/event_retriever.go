package retriever

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/exec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/parser"
	regState "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/state"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

//nolint:lll
//go:generate mockery --name EventRetriever --structname EventRetrieverMock --filename event_retriever_mock.go --inpackage

// EventRetriever is the interface used for retrieving and decoding events.
type EventRetriever interface {
	GetEvents(blockHash types.Hash) ([]*parser.Event, error)
}

// eventRetriever implements the EventRetriever interface.
type eventRetriever struct {
	eventParser parser.EventParser

	eventProvider regState.EventProvider
	stateRPC      state.State

	registryFactory registry.Factory

	eventStorageExecutor exec.RetryableExecutor[*types.StorageDataRaw]
	eventParsingExecutor exec.RetryableExecutor[[]*parser.Event]

	eventRegistry registry.EventRegistry
	meta          *types.Metadata
}

// NewEventRetriever creates a new EventRetriever.
func NewEventRetriever(
	eventParser parser.EventParser,
	eventProvider regState.EventProvider,
	stateRPC state.State,
	registryFactory registry.Factory,
	eventStorageExecutor exec.RetryableExecutor[*types.StorageDataRaw],
	eventParsingExecutor exec.RetryableExecutor[[]*parser.Event],
) (EventRetriever, error) {
	retriever := &eventRetriever{
		eventParser:          eventParser,
		eventProvider:        eventProvider,
		stateRPC:             stateRPC,
		registryFactory:      registryFactory,
		eventStorageExecutor: eventStorageExecutor,
		eventParsingExecutor: eventParsingExecutor,
	}

	if err := retriever.updateInternalState(nil); err != nil {
		return nil, ErrInternalStateUpdate.Wrap(err)
	}

	return retriever, nil
}

// NewDefaultEventRetriever creates a new EventRetriever using defaults for:
//
// - parser.EventParser
// - registry.Factory
// - exec.RetryableExecutor - used for retrieving event storage data.
// - exec.RetryableExecutor - used for parsing events.
func NewDefaultEventRetriever(
	eventProvider regState.EventProvider,
	stateRPC state.State,
	fieldOverrides ...registry.FieldOverride,
) (EventRetriever, error) {
	eventParser := parser.NewEventParser()
	registryFactory := registry.NewFactory(fieldOverrides...)

	eventStorageExecutor := exec.NewRetryableExecutor[*types.StorageDataRaw](exec.WithRetryTimeout(1 * time.Second))
	eventParsingExecutor := exec.NewRetryableExecutor[[]*parser.Event](exec.WithMaxRetryCount(1))

	return NewEventRetriever(
		eventParser,
		eventProvider,
		stateRPC,
		registryFactory,
		eventStorageExecutor,
		eventParsingExecutor,
	)
}

// GetEvents retrieves the storage data for an Event and then parses it.
//
// Both the event storage data retrieval and the event parsing are handled via the exec.RetryableExecutor
// in order to ensure retries in case of network errors or parsing errors due to an outdated event registry.
func (e *eventRetriever) GetEvents(blockHash types.Hash) ([]*parser.Event, error) {
	storageEvents, err := e.eventStorageExecutor.ExecWithFallback(
		func() (*types.StorageDataRaw, error) {
			return e.eventProvider.GetStorageEvents(e.meta, blockHash)
		},
		func() error {
			return e.updateInternalState(&blockHash)
		},
	)

	if err != nil {
		return nil, ErrStorageEventRetrieval.Wrap(err)
	}

	events, err := e.eventParsingExecutor.ExecWithFallback(
		func() ([]*parser.Event, error) {
			return e.eventParser.ParseEvents(e.eventRegistry, storageEvents)
		},
		func() error {
			return e.updateInternalState(&blockHash)
		},
	)

	if err != nil {
		return nil, ErrEventParsing.Wrap(err)
	}

	return events, nil
}

// updateInternalState will retrieve the metadata at the provided blockHash, if provided,
// create an event registry based on this metadata and store both.
func (e *eventRetriever) updateInternalState(blockHash *types.Hash) error {
	var (
		meta *types.Metadata
		err  error
	)

	if blockHash == nil {
		meta, err = e.stateRPC.GetMetadataLatest()
	} else {
		meta, err = e.stateRPC.GetMetadata(*blockHash)
	}

	if err != nil {
		return ErrMetadataRetrieval.Wrap(err)
	}

	eventRegistry, err := e.registryFactory.CreateEventRegistry(meta)

	if err != nil {
		return ErrEventRegistryCreation.Wrap(err)
	}

	e.meta = meta
	e.eventRegistry = eventRegistry

	return nil
}
