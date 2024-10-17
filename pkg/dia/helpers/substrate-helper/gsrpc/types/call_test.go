package types_test

import (
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	c := Call{CallIndex{6, 1}, Args{0, 0, 0}}

	enc, err := codec.EncodeToHex(c)
	assert.NoError(t, err)
	assert.Equal(t, "0x0601000000", enc)
}

func TestNewCallV7(t *testing.T) {
	c, err := NewCall(&exampleMetadataV7, "Module2.my function", U8(3))
	assert.NoError(t, err)

	enc, err := codec.EncodeToHex(c)
	assert.NoError(t, err)

	assert.Equal(t, "0x010003", enc)
}

func TestNewCallV8(t *testing.T) {
	c, err := NewCall(&exampleMetadataV8, "Module2.my function", U8(3))
	assert.NoError(t, err)

	enc, err := codec.EncodeToHex(c)
	assert.NoError(t, err)

	assert.Equal(t, "0x010003", enc)
}
