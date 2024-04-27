package fault

import "net/http"

type Code int

const (
	CodeNotFound Code = iota + 1
	CodeValidation
	CodeConflict
	CodeBadRequest
	CodeUnauthorized
	CodeForbidden
	CodeInternal
	CodeNotImplemented
)

func (e Code) String() string {
	switch e {
	case CodeNotFound:
		return "not found"
	case CodeValidation:
		return "validation"
	case CodeConflict:
		return "conflict"
	case CodeBadRequest:
		return "bad request"
	case CodeUnauthorized:
		return "unauthorized"
	case CodeForbidden:
		return "forbidden"
	case CodeInternal:
		return "internal"
	case CodeNotImplemented:
		return "not implemented"
	default:
		return "unknown"
	}
}

func (e Code) Status() int {
	switch e {
	case CodeNotFound:
		return http.StatusNotFound
	case CodeBadRequest, CodeValidation:
		return http.StatusBadRequest
	case CodeConflict:
		return http.StatusConflict
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
