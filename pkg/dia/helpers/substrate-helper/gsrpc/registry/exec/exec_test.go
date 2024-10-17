package exec

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetryableExecutor_ExecWithFallback(t *testing.T) {
	exec := NewRetryableExecutor[int]()

	execFnCallCount := 0
	fallbackFnCallCount := 0

	execFnRes := 11

	res, err := exec.ExecWithFallback(func() (int, error) {
		execFnCallCount++

		return execFnRes, nil
	}, func() error {
		fallbackFnCallCount++

		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, execFnRes, res)
	assert.Equal(t, 1, execFnCallCount)
	assert.Equal(t, 0, fallbackFnCallCount)
}

func TestRetryableExecutor_ExecWithFallback_RetrySuccess(t *testing.T) {
	exec := NewRetryableExecutor[int]()

	execFnCallCount := 0
	fallbackFnCallCount := 0

	execFnRes := 11

	res, err := exec.ExecWithFallback(func() (int, error) {
		execFnCallCount++

		if execFnCallCount < 2 {
			return 0, errors.New("boom")
		}

		return execFnRes, nil
	}, func() error {
		fallbackFnCallCount++

		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, execFnRes, res)
	assert.Equal(t, 2, execFnCallCount)
	assert.Equal(t, 1, fallbackFnCallCount)
}

func TestRetryableExecutor_ExecWithFallback_NilFns(t *testing.T) {
	exec := NewRetryableExecutor[int]()

	res, err := exec.ExecWithFallback(nil, nil)
	assert.ErrorIs(t, err, ErrMissingExecFn)
	assert.Equal(t, 0, res)

	res, err = exec.ExecWithFallback(func() (int, error) {
		return 1, nil
	}, nil)
	assert.ErrorIs(t, err, ErrMissingFallbackFn)
	assert.Equal(t, 0, res)

}

func TestRetryableExecutor_ExecWithFallback_ExecFnError(t *testing.T) {
	retryCount := uint(5)

	exec := NewRetryableExecutor[int](
		WithMaxRetryCount(retryCount),
		WithRetryTimeout(100*time.Millisecond),
	)

	execFnCallCount := uint(0)
	fallbackFnCallCount := uint(0)

	res, err := exec.ExecWithFallback(func() (int, error) {
		execFnCallCount++

		return 0, errors.New("boom")
	}, func() error {
		fallbackFnCallCount++

		return nil
	})
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
	assert.Equal(t, retryCount+1, execFnCallCount)
	assert.Equal(t, retryCount, fallbackFnCallCount)

	execErr := err.(*Error)
	assert.Len(t, execErr.errs, int(retryCount+1))
}

func TestRetryableExecutor_ExecWithFallback_FallBackFnError(t *testing.T) {
	exec := NewRetryableExecutor[int]()

	execFnCallCount := 0
	fallbackFnCallCount := 0

	res, err := exec.ExecWithFallback(func() (int, error) {
		execFnCallCount++

		return 0, errors.New("boom")
	}, func() error {
		fallbackFnCallCount++

		return errors.New("boom")
	})
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
	assert.Equal(t, defaultMaxRetryCount+1, execFnCallCount)
	assert.Equal(t, defaultMaxRetryCount, fallbackFnCallCount)

	execErr := err.(*Error)
	assert.Len(t, execErr.errs, defaultMaxRetryCount+1)
}

func TestRetryableExecutor_ExecWithFallback_FallBackFnError_NoRetry(t *testing.T) {
	exec := NewRetryableExecutor[int](WithRetryOnFallBackError(false))

	execFnCallCount := 0
	fallbackFnCallCount := 0

	res, err := exec.ExecWithFallback(func() (int, error) {
		execFnCallCount++

		return 0, errors.New("boom")
	}, func() error {
		fallbackFnCallCount++

		return errors.New("boom")
	})
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
	assert.Equal(t, 1, execFnCallCount)
	assert.Equal(t, 1, fallbackFnCallCount)

	execErr := err.(*Error)
	assert.Len(t, execErr.errs, 2)
}
