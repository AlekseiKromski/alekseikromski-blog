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

func (ie *JsonError) Error() string {
	return fmt.Sprintf("%s, cause: %v", ie.Message, ie.Error())
}
