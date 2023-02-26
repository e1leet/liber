package errors

import "errors"

var ErrSessionNotFound = errors.New("session not found")
var ErrSessionDeleteFailed = errors.New("session delete failed")
