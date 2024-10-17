package exec

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//go:generate mockery --name RetryableExecutor --structname RetryableExecutorMock --filename exec_mock.go --inpackage

// RetryableExecutor is the interface used for executing a closure and its fallback if the initial execution fails.
//
// The interface is generic over type T which represents the return value of the closure.
type RetryableExecutor[T any] interface {
	ExecWithFallback(execFn func() (T, error), fallbackFn func() error) (T, error)
}

// retryableExecutor implements RetryableExecutor.
//
// It can be configured via the provided OptsFn(s).
type retryableExecutor[T any] struct {
	opts *Opts
}

// NewRetryableExecutor creates a new RetryableExecutor.
func NewRetryableExecutor[T any](opts ...OptsFn) RetryableExecutor[T] {
	execOpts := NewDefaultExecOpts()

	for _, opt := range opts {
		opt(execOpts)
	}

	return &retryableExecutor[T]{
		execOpts,
	}
}

// ExecWithFallback will attempt to execute the provided execFn and, in the case of failure, it will execute
// the fallbackFn and retry execution of execFn.
func (r *retryableExecutor[T]) ExecWithFallback(execFn func() (T, error), fallbackFn func() error) (res T, err error) {
	if execFn == nil {
		return res, ErrMissingExecFn
	}

	if fallbackFn == nil {
		return res, ErrMissingFallbackFn
	}

	execErr := &Error{}

	retryCount := uint(0)

	for {
		res, err = execFn()

		if err == nil {
			return res, nil
		}

		execErr.AddErr(fmt.Errorf("exec function error: %w", err))

		if retryCount == r.opts.maxRetryCount {
			return res, execErr
		}

		if err = fallbackFn(); err != nil && !r.opts.retryOnFallbackError {
			execErr.AddErr(fmt.Errorf("fallback function error: %w", err))

			return res, execErr
		}

		retryCount++

		time.Sleep(r.opts.retryTimeout)
	}
}

var (
	ErrMissingExecFn     = errors.New("no exec function provided")
	ErrMissingFallbackFn = errors.New("no fallback function provided")
)

const (
	defaultMaxRetryCount        = 3
	defaultErrTimeout           = 0 * time.Second
	defaultRetryOnFallbackError = true
)

// Opts holds the configurable options for a RetryableExecutor.
type Opts struct {
	// maxRetryCount holds maximum number of retries in the case of failure.
	maxRetryCount uint

	// retryTimeout holds the timeout between retries.
	retryTimeout time.Duration

	// retryOnFallbackError specifies whether a retry will be done in the case of
	// failure of the fallback function.
	retryOnFallbackError bool
}

// NewDefaultExecOpts creates the default Opts.
func NewDefaultExecOpts() *Opts {
	return &Opts{
		maxRetryCount:        defaultMaxRetryCount,
		retryTimeout:         defaultErrTimeout,
		retryOnFallbackError: defaultRetryOnFallbackError,
	}
}

// OptsFn is function that operate on Opts.
type OptsFn func(opts *Opts)

// WithMaxRetryCount sets the max retry count.
//
// Note that a default value is provided if the provided count is 0.
func WithMaxRetryCount(maxRetryCount uint) OptsFn {
	return func(opts *Opts) {
		if maxRetryCount == 0 {
			maxRetryCount = defaultMaxRetryCount
		}

		opts.maxRetryCount = maxRetryCount
	}
}

// WithRetryTimeout sets the retry timeout.
func WithRetryTimeout(retryTimeout time.Duration) OptsFn {
	return func(opts *Opts) {
		opts.retryTimeout = retryTimeout
	}
}

// WithRetryOnFallBackError sets the retryOnFallbackError flag.
func WithRetryOnFallBackError(retryOnFallbackError bool) OptsFn {
	return func(opts *Opts) {
		opts.retryOnFallbackError = retryOnFallbackError
	}
}

// Error holds none or multiple errors that can happen during execution.
type Error struct {
	errs []error
}

// AddErr appends an error to the error slice of Error.
func (e *Error) AddErr(err error) {
	e.errs = append(e.errs, err)
}

// Error implements the standard error interface.
func (e *Error) Error() string {
	sb := strings.Builder{}

	for i, err := range e.errs {
		sb.WriteString(fmt.Sprintf("error %d: %s\n", i, err))
	}

	return sb.String()
}
