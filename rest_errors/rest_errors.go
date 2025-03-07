package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e *restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e *restErr) Message() string {
	return e.ErrMessage
}

func (e *restErr) Status() int {
	return e.ErrStatus
}

func (e *restErr) Causes() []interface{} {
	return e.ErrCauses
}

func GetRestErrorInstance() *restErr {
	return &restErr{}
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr *restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) RestErr {
	result := &restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad request",
	}
	return result
}

func NewNotFoundError(message string) RestErr {
	result := &restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not found",
	}
	return result
}

func NewUnauthorizedError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "Unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := &restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal server error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
