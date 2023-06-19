package v1

import (
	"fmt"
	"net/http"
)

type ServerError interface {
	HttpError() int
}

// ClassifyError - return error header
func ClassifyError(err error) int {
	if e, ok := err.(ServerError); ok {
		return e.HttpError()
	}

	return http.StatusBadRequest
}

type InputError struct {
	Message string `json:"message"`
}

func NewInputError() *InputError {
	return &InputError{
		Message: "cannot read/write input",
	}
}

func (ie *InputError) HttpError() int {
	return http.StatusInternalServerError
}

func (ie *InputError) Error() string {
	return fmt.Sprintf("internal server error: %s, cause: %v", ie.Message, ie.Error())
}

type AuthError struct {
	Message string `json:"message"`
}

func NewAuthError(msg string) *AuthError {
	return &AuthError{
		Message: msg,
	}
}

func (au *AuthError) Error() string {
	return fmt.Sprintf("error: %s", au.Message)
}

type JsonError struct {
	Message string `json:"message"`
}

func NewDecodingError() *JsonError {
	return &JsonError{
		Message: "cannot encode or decode payload",
	}
}

func (ie *JsonError) HttpError() int {
	return http.StatusBadRequest
}

func (au *AuthError) HttpError() int {
	return http.StatusBadRequest
}

func (ie *JsonError) Error() string {
	return fmt.Sprintf("%s, cause: %v", ie.Message, ie.Error())
}

type ActionError struct {
	Message string `json:"message"`
}

func NewActionError(err error) *ActionError {
	return &ActionError{
		Message: fmt.Sprintf("cannot finish action, reason: %v", err),
	}
}

func (ie *ActionError) HttpError() int {
	return http.StatusBadRequest
}

func (ie *ActionError) Error() string {
	return fmt.Sprintf("Server error: %s, cause: %v", ie.Message, ie.Error())
}
