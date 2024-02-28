package errorcodeframework_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-container-networking/cns/errorcodeframework"
	"github.com/Azure/azure-container-networking/cns/types"
)

func TestNewErrorCode(t *testing.T) {
	err := errorcodeframework.NewErrorCode(http.StatusBadRequest, types.InvalidParameter, "Bad Request")
	if err == nil {
		t.Errorf("NewErrorCode returned nil error")
	}

	customErr, ok := err.(*errorcodeframework.ErrorCode)
	if !ok {
		t.Errorf("NewErrorCode did not return a custom error")
	}

	if customErr.HTTPErrorCode != http.StatusBadRequest {
		t.Errorf("HTTPErrorCode was not set correctly")
	}

	if customErr.SubCode != types.InvalidParameter {
		t.Errorf("SubCode was not set correctly")
	}

	if customErr.ErrorMessage != "Bad Request" {
		t.Errorf("ErrorMessage was not set correctly")
	}
}

func TestIsCustomErrorCode(t *testing.T) {
	err := errors.New("This is a regular error")
	if errorcodeframework.IsCustomErrorCode(err) {
		t.Errorf("IsCustomErrorCode returned true for a regular error")
	}

	customErr := errorcodeframework.NewErrorCode(http.StatusBadRequest, types.InvalidParameter, "Bad Request")
	if !errorcodeframework.IsCustomErrorCode(customErr) {
		t.Errorf("IsCustomErrorCode returned false for a custom error")
	}
}

func TestGetHTTPErrorCode(t *testing.T) {
	err := errors.New("This is a regular error")
	if errorcodeframework.GetHTTPErrorCode(err) != http.StatusInternalServerError {
		t.Errorf("GetHTTPErrorCode returned incorrect HTTP status code for a regular error")
	}

	customErr := errorcodeframework.NewErrorCode(http.StatusBadRequest, types.InvalidParameter, "Bad Request")
	if errorcodeframework.GetHTTPErrorCode(customErr) != http.StatusBadRequest {
		t.Errorf("GetHTTPErrorCode returned incorrect HTTP status code for a custom error")
	}
}

func TestHandleError(t *testing.T) {
	w := httptest.NewRecorder()
	err := errorcodeframework.NewErrorCode(http.StatusBadRequest, types.InvalidParameter, "Bad Request")
	customErr := err.(*errorcodeframework.ErrorCode)
	http.Error(w, customErr.ErrorMessage, customErr.HTTPErrorCode)

	if w.Code != http.StatusBadRequest {
		t.Errorf("HandleError did not set the correct HTTP status code")
	}

	if w.Body.String() != "Bad Request\n" {
		t.Errorf("HandleError did not set the correct error message")
	}
}
