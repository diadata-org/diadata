package retriever

import libErr "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/error"

const (
	ErrInternalStateUpdate      = libErr.Error("internal state update")
	ErrBlockRetrieval           = libErr.Error("block retrieval")
	ErrExtrinsicDecoding        = libErr.Error("extrinsic parsing")
	ErrMetadataRetrieval        = libErr.Error("metadata retrieval")
	ErrExtrinsicDecoderCreation = libErr.Error("extrinsic decoder creation")
	ErrStorageEventRetrieval    = libErr.Error("storage event retrieval")
	ErrEventParsing             = libErr.Error("event parsing")
	ErrEventRegistryCreation    = libErr.Error("event registry creation")
)
