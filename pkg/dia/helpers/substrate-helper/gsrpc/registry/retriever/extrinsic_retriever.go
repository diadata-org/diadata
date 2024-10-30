package retriever

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/exec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/chain"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/block"
)

//nolint:lll
//go:generate mockery --name ExtrinsicRetriever --structname ExtrinsicRetrieverMock --filename extrinsic_retriever_mock.go --inpackage

// ExtrinsicRetriever is the interface used for retrieving and decoding extrinsic information
// from a particular block.
type ExtrinsicRetriever interface {
	GetExtrinsics(blockHash types.Hash) ([]*registry.DecodedExtrinsic, error)
}

// extrinsicRetriever implements the ExtrinsicRetriever interface.
type extrinsicRetriever struct {
	chainRPC chain.Chain
	stateRPC state.State

	registryFactory registry.Factory

	chainExecutor             exec.RetryableExecutor[*block.SignedBlock]
	extrinsicDecodingExecutor exec.RetryableExecutor[[]*registry.DecodedExtrinsic]

	extrinsicDecoder *registry.ExtrinsicDecoder
}

// NewExtrinsicRetriever creates a new ExtrinsicRetriever.
func NewExtrinsicRetriever(
	chainRPC chain.Chain,
	stateRPC state.State,
	registryFactory registry.Factory,
	chainExecutor exec.RetryableExecutor[*block.SignedBlock],
	extrinsicDecodingExecutor exec.RetryableExecutor[[]*registry.DecodedExtrinsic],
) (ExtrinsicRetriever, error) {
	retriever := &extrinsicRetriever{
		chainRPC:                  chainRPC,
		stateRPC:                  stateRPC,
		registryFactory:           registryFactory,
		chainExecutor:             chainExecutor,
		extrinsicDecodingExecutor: extrinsicDecodingExecutor,
	}

	if err := retriever.updateInternalState(nil); err != nil {
		return nil, ErrInternalStateUpdate.Wrap(err)
	}

	return retriever, nil
}

// NewDefaultExtrinsicRetriever returns an ExtrinsicRetriever with default values for the factory and executors.
func NewDefaultExtrinsicRetriever(
	chainRPC chain.Chain,
	stateRPC state.State,
	fieldOverrides ...registry.FieldOverride,
) (ExtrinsicRetriever, error) {
	registryFactory := registry.NewFactory(fieldOverrides...)

	chainExecutor := exec.NewRetryableExecutor[*block.SignedBlock](exec.WithRetryTimeout(1 * time.Second))
	extrinsicDecodingExecutor := exec.NewRetryableExecutor[[]*registry.DecodedExtrinsic](exec.WithMaxRetryCount(1))

	return NewExtrinsicRetriever(
		chainRPC,
		stateRPC,
		registryFactory,
		chainExecutor,
		extrinsicDecodingExecutor,
	)
}

// GetExtrinsics retrieves a generic.SignedBlock and then parses the extrinsics found in it.
//
// Both the block retrieval and the extrinsic parsing are handled via the exec.RetryableExecutor
// in order to ensure retries in case of network errors or parsing errors due to an outdated extrinsic decoder.
func (e *extrinsicRetriever) GetExtrinsics(blockHash types.Hash) ([]*registry.DecodedExtrinsic, error) {
	block, err := e.chainExecutor.ExecWithFallback(
		func() (*block.SignedBlock, error) {
			return e.chainRPC.GetBlock(blockHash)
		},
		func() error {
			return nil
		},
	)

	if err != nil {
		return nil, ErrBlockRetrieval.Wrap(err)
	}

	calls, err := e.extrinsicDecodingExecutor.ExecWithFallback(
		func() ([]*registry.DecodedExtrinsic, error) {
			return block.DecodeExtrinsics(e.extrinsicDecoder)
		},
		func() error {
			return e.updateInternalState(&blockHash)
		},
	)

	if err != nil {
		return nil, ErrExtrinsicDecoding.Wrap(err)
	}

	return calls, nil
}

// updateInternalState will retrieve the metadata at the provided blockHash, if provided,
// create an extrinsic decoder based on this metadata and store both.
func (e *extrinsicRetriever) updateInternalState(blockHash *types.Hash) error {
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

	extrinsicDecoder, err := e.registryFactory.CreateExtrinsicDecoder(meta)

	if err != nil {
		return ErrExtrinsicDecoderCreation.Wrap(err)
	}

	e.extrinsicDecoder = extrinsicDecoder

	return nil
}
