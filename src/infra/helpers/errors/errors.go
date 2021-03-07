package errors

import "net/http"

type Error struct {
	Status  int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) StatusCode() int {
	return e.Status
}

func InternalServerError(msg string) Error {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return Error{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func NotFound(msg string) Error {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return Error{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

func Unauthorized(msg string) Error {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return Error{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

func Forbidden(msg string) Error {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return Error{
		Status:  http.StatusForbidden,
		Message: msg,
	}

}

func BadRequest(msg string) Error {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return Error{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

func Conflict(msg string) Error {
	if msg == "" {
		msg = "Your data have some conflicting."
	}
	return Error{
		Status:  http.StatusConflict,
		Message: msg,
	}
}
