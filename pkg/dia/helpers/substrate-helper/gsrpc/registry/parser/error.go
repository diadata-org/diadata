package parser

import libErr "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/error"

const (
	ErrEventsCountDecoding  = libErr.Error("events count decoding")
	ErrEventPhaseDecoding   = libErr.Error("event phase decoding")
	ErrEventIDDecoding      = libErr.Error("event ID decoding")
	ErrEventDecoderNotFound = libErr.Error("event decoder not found")
	ErrEventFieldsDecoding  = libErr.Error("event fields decoding")
	ErrEventTopicsDecoding  = libErr.Error("event topics decoding")
)
