package error

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testErr = Error("test error")
)

func TestError(t *testing.T) {
	newStdErr := errors.New("new std error")
	wrappedErr := testErr.Wrap(newStdErr)

	assert.True(t, errors.Is(wrappedErr, testErr))
	assert.True(t, errors.Is(wrappedErr, newStdErr))
	assert.Equal(t, fmt.Sprintf("%s: %s", testErr.Error(), newStdErr.Error()), wrappedErr.Error())

	newErr := Error("new error")
	newWrappedErr := newErr.Wrap(wrappedErr)

	assert.True(t, errors.Is(newWrappedErr, newErr))
	assert.True(t, errors.Is(newWrappedErr, testErr))
	assert.True(t, errors.Is(newWrappedErr, newStdErr))
	assert.Equal(t, fmt.Sprintf("%s: %s", newErr.Error(), wrappedErr.Error()), newWrappedErr.Error())

	err := testErr.WithMsg("%d", 1)
	assert.Equal(t, fmt.Sprintf("%s: 1", testErr), err.Error())

	err = testErr.WithMsg("test msg")
	assert.Equal(t, fmt.Sprintf("%s: test msg", testErr), err.Error())
}
