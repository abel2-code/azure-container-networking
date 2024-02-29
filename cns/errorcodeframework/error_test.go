package errorcodeframework

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewErrorCode(t *testing.T) {
	err := NewErrorCode(http.StatusBadRequest, InvalidParameter, "Bad Request")
	if err == nil {
		t.Errorf("NewErrorCode returned nil error")
	}

	customErr, ok := err.(*ErrorCode)
	if !ok {
		t.Errorf("NewErrorCode did not return a custom error")
	}

	if customErr.HTTPErrorCode != http.StatusBadRequest {
		t.Errorf("HTTPErrorCode was not set correctly")
	}

	if customErr.SubCode != InvalidParameter {
		t.Errorf("SubCode was not set correctly")
	}

	if customErr.ErrorMessage != "Bad Request" {
		t.Errorf("ErrorMessage was not set correctly")
	}
}

func TestIsCustomErrorCode(t *testing.T) {
	err := errors.New("This is a regular error")
	if IsCustomErrorCode(err) {
		t.Errorf("IsCustomErrorCode returned true for a regular error")
	}

	customErr := NewErrorCode(http.StatusBadRequest, InvalidParameter, "Bad Request")
	if !IsCustomErrorCode(customErr) {
		t.Errorf("IsCustomErrorCode returned false for a custom error")
	}
}

func TestGetHTTPErrorCode(t *testing.T) {
	err := errors.New("This is a regular error")
	if GetHTTPErrorCode(err) != http.StatusInternalServerError {
		t.Errorf("GetHTTPErrorCode returned incorrect HTTP status code for a regular error")
	}

	customErr := NewErrorCode(http.StatusBadRequest, InvalidParameter, "Bad Request")
	if GetHTTPErrorCode(customErr) != http.StatusBadRequest {
		t.Errorf("GetHTTPErrorCode returned incorrect HTTP status code for a custom error")
	}
}

func TestHandleError(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewErrorCode(http.StatusBadRequest, InvalidParameter, "Bad Request")
	customErr := err.(*ErrorCode)
	http.Error(w, customErr.ErrorMessage, customErr.HTTPErrorCode)

	if w.Code != http.StatusBadRequest {
		t.Errorf("HandleError did not set the correct HTTP status code")
	}

	if w.Body.String() != "Bad Request\n" {
		t.Errorf("HandleError did not set the correct error message")
	}
}
