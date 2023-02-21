package errors

import "github.com/e1leet/liber/pkg/errors"

var (
	ErrEmailNotSpecified    = errors.New("email not specified or incorrect type")
	ErrUsernameNotSpecified = errors.New("username not specified or incorrect type")
	ErrPasswordNotSpecified = errors.New("password not specified or incorrect type")
)
