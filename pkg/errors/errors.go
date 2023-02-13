package errors

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func New(msg string) error {
	return errors.New(msg)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(err error, msg string) error {
	if err == nil {
		return err
	}

	return Wrap(err, msg)
}

func Append(err error, errs ...error) *multierror.Error {
	return multierror.Append(err, errs...)
}

func Flatten(err error) error {
	return multierror.Flatten(err)
}

func Prefix(err error, prefix string) error {
	return multierror.Prefix(err, prefix)
}
