package extensions

import (
	"errors"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

type CheckMetadataMode byte

var (
	CheckMetadataModeDisabled CheckMetadataMode
	CheckMetadataModeEnabled  CheckMetadataMode = 1
)

func (m CheckMetadataMode) Encode(encoder scale.Encoder) error {
	switch m {
	case CheckMetadataModeDisabled:
		return encoder.PushByte(0)
	case CheckMetadataModeEnabled:
		return encoder.PushByte(1)
	default:
		return errors.New("unsupported check metadata mode")
	}
}

type CheckMetadataHash struct {
	Hash types.Option[types.H256]
}

func (c CheckMetadataHash) Encode(encoder *scale.Encoder) error {
	if c.Hash.HasValue() {
		_, hash := c.Hash.Unwrap()

		return encoder.Encode(hash)
	}

	return encoder.PushByte(0)
}
