package parser

import (
	"bytes"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

// Event holds all the information of a decoded storage event.
type Event struct {
	Name    string
	Fields  registry.DecodedFields
	EventID types.EventID
	Phase   *types.Phase
	Topics  []types.Hash
}

//go:generate mockery --name EventParser --structname EventParserMock --filename event_parser_mock.go --inpackage

// EventParser is the interface used for parsing event storage data into []*Event.
type EventParser interface {
	ParseEvents(eventRegistry registry.EventRegistry, sd *types.StorageDataRaw) ([]*Event, error)
}

// EventParserFn implements EventParser.
type EventParserFn func(eventRegistry registry.EventRegistry, sd *types.StorageDataRaw) ([]*Event, error)

// ParseEvents is the function required for satisfying the EventParser interface.
func (f EventParserFn) ParseEvents(eventRegistry registry.EventRegistry, sd *types.StorageDataRaw) ([]*Event, error) {
	return f(eventRegistry, sd)
}

// NewEventParser creates a new EventParser.
func NewEventParser() EventParser {
	// The EventParserFn provided here is decoding the total number of events from the storage data then attempts
	// to decode all the information for each event.
	return EventParserFn(func(eventRegistry registry.EventRegistry, sd *types.StorageDataRaw) ([]*Event, error) {
		decoder := scale.NewDecoder(bytes.NewReader(*sd))

		eventsCount, err := decoder.DecodeUintCompact()

		if err != nil {
			return nil, ErrEventsCountDecoding.Wrap(err)
		}

		var events []*Event

		for i := uint64(0); i < eventsCount.Uint64(); i++ {
			var phase types.Phase

			if err := decoder.Decode(&phase); err != nil {
				return nil, ErrEventPhaseDecoding.Wrap(fmt.Errorf("event #%d: %w", i, err))
			}

			var eventID types.EventID

			if err := decoder.Decode(&eventID); err != nil {
				return nil, ErrEventIDDecoding.Wrap(fmt.Errorf("event #%d: %w", i, err))
			}

			eventDecoder, ok := eventRegistry[eventID]

			if !ok {
				return nil, ErrEventDecoderNotFound.WithMsg("event #%d with ID: %v", i, eventID)
			}

			eventFields, err := eventDecoder.Decode(decoder)

			if err != nil {
				return nil, ErrEventFieldsDecoding.Wrap(fmt.Errorf("event #%d: %w", i, err))
			}

			var topics []types.Hash

			if err := decoder.Decode(&topics); err != nil {
				return nil, ErrEventTopicsDecoding.Wrap(fmt.Errorf("event #%d: %w", i, err))
			}

			event := &Event{
				Name:    eventDecoder.Name,
				Fields:  eventFields,
				EventID: eventID,
				Phase:   &phase,
				Topics:  topics,
			}

			events = append(events, event)
		}

		return events, nil
	})
}
