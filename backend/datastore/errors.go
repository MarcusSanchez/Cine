package datastore

import (
	"errors"
	"fmt"
)

var (
	ErrTxDone         = errors.New("transaction done")
	ErrTxNested       = errors.New("transaction nested")
	ErrNotFound       = errors.New("not found")
	ErrValidation     = errors.New("validation")
	ErrConstraint     = errors.New("constraint")
	ErrInternal       = errors.New("internal")
	ErrNotImplemented = errors.New("not implemented")
)

func Wrap(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

func IsTxDone(err error) bool {
	return errors.Is(err, ErrTxDone)
}

func IsTxNested(err error) bool {
	return errors.Is(err, ErrTxNested)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsValidation(err error) bool {
	return errors.Is(err, ErrValidation)
}

func IsConstraint(err error) bool {
	return errors.Is(err, ErrConstraint)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}

func IsNotImplemented(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}
