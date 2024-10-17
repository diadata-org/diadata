package error

import (
	"fmt"
	"strings"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

func (e Error) Is(err error) bool {
	return strings.Contains(string(e), err.Error())
}

func (e Error) Wrap(err error) Error {
	return Error(fmt.Errorf("%s: %w", e, err).Error())
}

func (e Error) WithMsg(msgFormat string, formatArgs ...any) Error {
	msg := fmt.Sprintf(msgFormat, formatArgs...)

	return Error(fmt.Sprintf("%s: %s", e, msg))
}
