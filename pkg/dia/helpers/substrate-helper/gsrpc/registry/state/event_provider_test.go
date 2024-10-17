package state

import (
	"errors"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/rpc/state/mocks"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestProvider_GetStorageEvents(t *testing.T) {
	stateRPCMock := mocks.NewState(t)

	provider := NewEventProvider(stateRPCMock)

	testHash := types.Hash{}

	var testMeta types.Metadata

	// Empty metadata should cause an error when creating the storage key.
	res, err := provider.GetStorageEvents(&testMeta, testHash)
	assert.ErrorIs(t, err, ErrEventStorageKeyCreation)
	assert.Nil(t, res)

	err = codec.DecodeFromHex(types.MetadataV14Data, &testMeta)
	assert.NoError(t, err)

	storageKey, err := types.CreateStorageKey(&testMeta, storagePrefix, storageMethod, nil)
	assert.NoError(t, err)

	storageData := &types.StorageDataRaw{}

	stateRPCMock.On("GetStorageRaw", storageKey, testHash).
		Return(storageData, nil).
		Once()

	res, err = provider.GetStorageEvents(&testMeta, testHash)
	assert.NoError(t, err)
	assert.Equal(t, storageData, res)

	stateRPCError := errors.New("error")

	stateRPCMock.On("GetStorageRaw", storageKey, testHash).
		Return(nil, stateRPCError).
		Once()

	res, err = provider.GetStorageEvents(&testMeta, testHash)
	assert.ErrorIs(t, err, ErrEventStorageRetrieval)
	assert.Nil(t, res)
}
