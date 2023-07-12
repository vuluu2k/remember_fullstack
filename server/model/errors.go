package model

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization   Type = "AUTHORIZATION"
	BadRequest      Type = "BAD_REQUEST"
	Conflict        Type = "CONFLICT"
	Internal        Type = "INTERNAL"
	NotFound        Type = "NOT_FOUND"
	PayloadTooLarge Type = "PAYLOAD_TOO_LARGE"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad Request. Reason: %v", reason),
	}
}

func NewConflict(reason string) *Error {
	return &Error{
		Type:    Conflict,
		Message: reason,
	}
}

func NewInternal(reason string) *Error {
	return &Error{
		Type:    Internal,
		Message: reason,
	}
}

func NewNotFound(reason string) *Error {
	return &Error{
		Type:    NotFound,
		Message: reason,
	}
}

func NewPayloadTooLarge(reason string) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: reason,
	}
}