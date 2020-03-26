package rest_errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, http.StatusNotFound, err.Status())

}

func TestNewError(t *testing.T) {
	err := NewError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Error())
}

func TestGetRestErrorInstance(t *testing.T) {
	assert.NotNil(t, GetRestErrorInstance())
}

func TestNewRestError(t *testing.T) {
	restErr := NewRestError("testing", http.StatusOK, "some testing", nil)
	assert.NotNil(t, restErr)
	assert.EqualValues(t, "testing", restErr.Message())
	assert.EqualValues(t, http.StatusOK, restErr.Status())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("testing")
	assert.NotNil(t, err)
	assert.EqualValues(t, "testing", err.Message())
}

func TestNewRestErrorFromBytes(t *testing.T) {
	resp := `{"message":"testing","status":400, "err":"some testing"}`
	rErr, err := NewRestErrorFromBytes([]byte(resp))
	assert.NotNil(t, rErr)
	assert.Nil(t, err)
	assert.EqualValues(t, "testing", rErr.Message())
	assert.EqualValues(t, http.StatusBadRequest, rErr.Status())
	assert.EqualValues(t, fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", rErr.Message(), rErr.Status(), rErr.(*restErr).ErrError, rErr.Causes()), rErr.Error())
}

func TestNewBadRequestErrorWithDecodeError(t *testing.T) {
	resp := `{"message":"testing","status":"400", "err":"some testing"}`
	rErr, err := NewRestErrorFromBytes([]byte(resp))
	assert.Nil(t, rErr)
	assert.NotNil(t, err)

}
