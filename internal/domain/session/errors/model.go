package errors

import "errors"

var (
	ErrTokenNotSpecified     = errors.New("token not specified")
	ErrUserIDNotSpecified    = errors.New("user id not specified")
	ErrExpiresInNotSpecified = errors.New("expires in not specified")
)
