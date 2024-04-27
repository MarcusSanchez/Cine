package tmdb

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal error")
)

func ErrorNotFound(msg string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, msg)
}

func ErrorInternal(msg string) error {
	return fmt.Errorf("%w: %s", ErrInternal, msg)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}
