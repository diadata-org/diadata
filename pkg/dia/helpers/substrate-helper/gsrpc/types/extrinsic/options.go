package extrinsic

import (
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic/extensions"
)

// SignedFieldValues is a map that holds a value for a particular SignedFieldName.
type SignedFieldValues map[SignedFieldName]any

// SigningOption is the type used for providing values to a SignedFieldValues map.
type SigningOption func(vals SignedFieldValues)

// WithEra returns a SigningOption that is used to add the era and block hash to a Payload.
func WithEra(era types.ExtrinsicEra, blockHash types.Hash) SigningOption {
	return func(vals SignedFieldValues) {
		vals[EraSignedField] = era
		vals[BlockHashSignedField] = blockHash
	}
}

// WithNonce returns a SigningOption that is used to add the nonce to a Payload.
func WithNonce(nonce types.UCompact) SigningOption {
	return func(vals SignedFieldValues) {
		vals[NonceSignedField] = nonce
	}
}

// WithMetadataMode returns a SigningOption that is used to add the check metadata mode and hash to a Payload.
func WithMetadataMode(mode extensions.CheckMetadataMode, metadataHash extensions.CheckMetadataHash) SigningOption {
	return func(vals SignedFieldValues) {
		vals[CheckMetadataHashModeSignedField] = mode
		vals[CheckMetadataHashSignedField] = metadataHash
	}
}

// WithTip returns a SigningOption that is used to add the tip to a Payload.
func WithTip(tip types.UCompact) SigningOption {
	return func(vals SignedFieldValues) {
		vals[TipSignedField] = tip
	}
}

// WithAssetID returns a SigningOption that is used to add the asset ID to a Payload.
func WithAssetID(assetID types.Option[types.AssetID]) SigningOption {
	return func(vals SignedFieldValues) {
		vals[AssetIDSignedField] = assetID
	}
}

// WithSpecVersion returns a SigningOption that is used to add the spec version to a Payload.
func WithSpecVersion(specVersion types.U32) SigningOption {
	return func(vals SignedFieldValues) {
		vals[SpecVersionSignedField] = specVersion
	}
}

// WithTransactionVersion returns a SigningOption that is used to add the transaction version to a Payload.
func WithTransactionVersion(transactionVersion types.U32) SigningOption {
	return func(vals SignedFieldValues) {
		vals[TransactionVersionSignedField] = transactionVersion
	}
}

// WithGenesisHash returns a SigningOption that is used to add the genesis hash to a Payload.
func WithGenesisHash(genesisHash types.Hash) SigningOption {
	return func(vals SignedFieldValues) {
		vals[GenesisHashSignedField] = genesisHash
	}
}
