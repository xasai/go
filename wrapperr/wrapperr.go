package wrapperr

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%s:%d -> %w", file, line, err)
}

func WrapOp(op string, err error) error {
	_, _, line, _ := runtime.Caller(1)
	return fmt.Errorf("%s:%d -> %w", op, line, err)
}

