package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
	Causes []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string, err error) *RestErr  {
	result := &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad request",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NewNotFoundError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not found",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal server error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
