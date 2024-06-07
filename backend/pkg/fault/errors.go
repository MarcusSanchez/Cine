package fault

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    Code
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.String(), e.Message)
}

func As(err error) (Error, bool) {
	var e Error
	if errors.As(err, &e) {
		return e, true
	}
	return e, false
}

func NotFound(message string) Error {
	return Error{Code: CodeNotFound, Message: message}
}

func Validation(message string) Error {
	return Error{Code: CodeValidation, Message: message}
}

func Conflict(message string) Error {
	return Error{Code: CodeConflict, Message: message}
}

func BadRequest(message string) Error {
	return Error{Code: CodeBadRequest, Message: message}
}

func Unauthorized(message string) Error {
	return Error{Code: CodeUnauthorized, Message: message}
}

func Forbidden(message string) Error {
	return Error{Code: CodeForbidden, Message: message}
}

func Internal(message string) Error {
	return Error{Code: CodeInternal, Message: message}
}
