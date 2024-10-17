package types_test

import (
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	optionFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *Option[U64], c fuzz.Continue) {
			if c.RandBool() {
				*o = NewEmptyOption[U64]()
				return
			}

			var u U64

			c.Fuzz(&u)

			*o = NewOption[U64](u)
		}),
	}
)

func TestOption_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Option[U64]](t, 100, optionFuzzOpts...)
	AssertEncodeEmptyObj[Option[U64]](t, 1)

	testOptionEncodeLen[U64](11, t)

	accountID := newTestAccountID()
	testOptionEncodeLen[AccountID](accountID, t)

	testOptionEncodeLen[ItemDetails](testInstanceDetails, t)
}

func testOptionEncodeLen[T any](testVal T, t *testing.T) {
	valEnc, err := Encode(testVal)
	assert.NoError(t, err)

	opt := NewOption[T](testVal)
	optEnc, err := Encode(opt)

	assert.Equal(t, len(optEnc), len(valEnc)+1)
}

func TestOption_OptionMethods(t *testing.T) {
	testOptionMethods[U64](11, t)

	accountID := newTestAccountID()

	testOptionMethods[*AccountID](&accountID, t)
	testOptionMethods[AccountID](accountID, t)

	testOptionMethods[*ItemDetails](&testInstanceDetails, t)
	testOptionMethods[ItemDetails](testInstanceDetails, t)
}

func testOptionMethods[T any](testVal T, t *testing.T) {
	o := NewEmptyOption[T]()

	var emptyVal T

	hasValue, value := o.Unwrap()
	assert.False(t, hasValue)
	assert.Equal(t, emptyVal, value)

	o.SetSome(testVal)

	hasValue, value = o.Unwrap()
	assert.True(t, hasValue)
	assert.Equal(t, testVal, value)

	o.SetNone()
	hasValue, value = o.Unwrap()
	assert.False(t, hasValue)
	assert.Equal(t, emptyVal, value)
}
