package retriever

import (
	"errors"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/exec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/registry/test"
	chainMocks "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/chain/mocks"
	stateMocks "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state/mocks"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/block"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExtrinsicRetriever_New(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	chainExecMock := exec.NewRetryableExecutorMock[*block.SignedBlock](t)
	extrinsicDecodingExecMock := exec.NewRetryableExecutorMock[[]*registry.DecodedExtrinsic](t)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	extrinsicDecoder := &registry.ExtrinsicDecoder{}

	registryFactoryMock.On("CreateExtrinsicDecoder", latestMeta).
		Return(extrinsicDecoder, nil).
		Once()

	res, err := NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &extrinsicRetriever{}, res)
}

func TestExtrinsicRetriever_New_InternalStateUpdateError(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	chainExecMock := exec.NewRetryableExecutorMock[*block.SignedBlock](t)
	extrinsicDecodingExecMock := exec.NewRetryableExecutorMock[[]*registry.DecodedExtrinsic](t)

	latestMeta := &types.Metadata{}

	stateRPCMock.On("GetMetadataLatest").
		Return(nil, errors.New("error")).
		Once()

	res, err := NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.ErrorIs(t, err, ErrInternalStateUpdate)
	assert.Nil(t, res)

	stateRPCMock.On("GetMetadataLatest").
		Return(latestMeta, nil).
		Once()

	registryFactoryMock.On("CreateExtrinsicDecoder", latestMeta).
		Return(nil, errors.New("error")).
		Once()

	res, err = NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.ErrorIs(t, err, ErrInternalStateUpdate)
	assert.Nil(t, res)
}

func TestExtrinsicRetriever_NewDefault(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)

	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	stateRPCMock.On("GetMetadataLatest").
		Return(&meta, nil).
		Once()
	res, err := NewDefaultExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &extrinsicRetriever{}, res)
}

func TestExtrinsicRetriever_GetExtrinsics(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	chainExecMock := exec.NewRetryableExecutorMock[*block.SignedBlock](t)
	extrinsicDecodingExecMock := exec.NewRetryableExecutorMock[[]*registry.DecodedExtrinsic](t)

	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	stateRPCMock.On("GetMetadataLatest").
		Return(&meta, nil).
		Once()

	extrinsicDecoder, err := registry.NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	registryFactoryMock.On("CreateExtrinsicDecoder", &meta).
		Return(extrinsicDecoder, nil).
		Once()

	extRet, err := NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &extrinsicRetriever{}, extRet)

	blockHash := types.Hash{1, 2, 3, 4}

	// NOTE - The following test relies on the following data from the Centrifuge development chain:
	//
	// Metadata from block 517393
	// Centrifuge development block 480514 - 0xacd5bd66bb3144f559be2bed09d0a1ae36e650b31f9eb8f5833eeb2c545d07cd
	// System remark extrinsic - 0xb10184008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a480118346322ed93ad7d2583ab3e4b71acd66cc1fce77cb225624c8eb00977681468aec33b933606ed8c2eaa75b84278c42415d491f89c5e79db6910986c1b95f486e401e0000000000431

	encodedExtrinsic := "0xb10184008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a480118346322ed93ad7d2583ab3e4b71acd66cc1fce77cb225624c8eb00977681468aec33b933606ed8c2eaa75b84278c42415d491f89c5e79db6910986c1b95f486e401e0000000000431"

	testBlock := &block.SignedBlock{
		Block: block.Block{
			Header: types.Header{},
			Extrinsics: []string{
				encodedExtrinsic,
			},
		},
		Justification: nil,
	}

	chainRPCMock.On("GetBlock", blockHash).
		Return(testBlock, nil).
		Once()

	chainExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*block.SignedBlock, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, testBlock, execFnRes)
			},
		).Return(testBlock, nil)

	decodedExtrinsic, err := extrinsicDecoder.DecodeHex(encodedExtrinsic)
	assert.NoError(t, err)

	decodedExtrinsics := []*registry.DecodedExtrinsic{decodedExtrinsic}

	extrinsicDecodingExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() ([]*registry.DecodedExtrinsic, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, decodedExtrinsics, execFnRes)
			},
		).Return(decodedExtrinsics, nil)

	res, err := extRet.GetExtrinsics(blockHash)
	assert.NoError(t, err)
	assert.Equal(t, decodedExtrinsics, res)
}

func TestExtrinsicRetriever_GetExtrinsics_BlockRetrievalError(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	chainExecMock := exec.NewRetryableExecutorMock[*block.SignedBlock](t)
	extrinsicDecodingExecMock := exec.NewRetryableExecutorMock[[]*registry.DecodedExtrinsic](t)

	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	stateRPCMock.On("GetMetadataLatest").
		Return(&meta, nil).
		Once()

	extrinsicDecoder, err := registry.NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	registryFactoryMock.On("CreateExtrinsicDecoder", &meta).
		Return(extrinsicDecoder, nil).
		Once()

	extRet, err := NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &extrinsicRetriever{}, extRet)

	blockHash := types.Hash{1, 2, 3, 4}

	blockRetrievalError := errors.New("block retrieval error")

	chainRPCMock.On("GetBlock", blockHash).
		Return(nil, blockRetrievalError).
		Once()

	var signedBlock *block.SignedBlock

	chainExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*block.SignedBlock, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.ErrorIs(t, err, blockRetrievalError)
				assert.Nil(t, execFnRes)
			},
		).Return(signedBlock, blockRetrievalError)

	res, err := extRet.GetExtrinsics(blockHash)
	assert.ErrorIs(t, err, ErrBlockRetrieval)
	assert.Nil(t, res)
}

func TestExtrinsicRetriever_GetExtrinsics_ExtrinsicDecodeError(t *testing.T) {
	chainRPCMock := chainMocks.NewChain(t)
	stateRPCMock := stateMocks.NewState(t)
	registryFactoryMock := registry.NewFactoryMock(t)
	chainExecMock := exec.NewRetryableExecutorMock[*block.SignedBlock](t)
	extrinsicDecodingExecMock := exec.NewRetryableExecutorMock[[]*registry.DecodedExtrinsic](t)

	var meta types.Metadata

	err := codec.DecodeFromHex(test.CentrifugeMetadataHex, &meta)
	assert.NoError(t, err)

	stateRPCMock.On("GetMetadataLatest").
		Return(&meta, nil).
		Once()

	extrinsicDecoder, err := registry.NewFactory().CreateExtrinsicDecoder(&meta)
	assert.NoError(t, err)

	registryFactoryMock.On("CreateExtrinsicDecoder", &meta).
		Return(extrinsicDecoder, nil).
		Once()

	extRet, err := NewExtrinsicRetriever(
		chainRPCMock,
		stateRPCMock,
		registryFactoryMock,
		chainExecMock,
		extrinsicDecodingExecMock,
	)
	assert.NoError(t, err)
	assert.IsType(t, &extrinsicRetriever{}, extRet)

	blockHash := types.Hash{1, 2, 3, 4}

	// Invalid encoded extrinsic that should trigger an error.
	encodedExtrinsic := "0xb10184"

	testBlock := &block.SignedBlock{
		Block: block.Block{
			Header: types.Header{},
			Extrinsics: []string{
				encodedExtrinsic,
			},
		},
		Justification: nil,
	}

	chainRPCMock.On("GetBlock", blockHash).
		Return(testBlock, nil).
		Once()

	chainExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() (*block.SignedBlock, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.NoError(t, err)
				assert.Equal(t, testBlock, execFnRes)
			},
		).Return(testBlock, nil)

	extrinsicDecodeError := errors.New("extrinsic decode error")

	var decodedExtrinsics []*registry.DecodedExtrinsic

	extrinsicDecodingExecMock.On("ExecWithFallback", mock.Anything, mock.Anything).
		Run(
			func(args mock.Arguments) {
				execFn, ok := args.Get(0).(func() ([]*registry.DecodedExtrinsic, error))
				assert.True(t, ok)

				execFnRes, err := execFn()
				assert.Error(t, err)
				assert.Nil(t, execFnRes)
			},
		).Return(decodedExtrinsics, extrinsicDecodeError)

	res, err := extRet.GetExtrinsics(blockHash)
	assert.ErrorIs(t, err, ErrExtrinsicDecoding)
	assert.Nil(t, res)
}
