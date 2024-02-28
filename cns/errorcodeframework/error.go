package errorcodeframework

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Azure/azure-container-networking/cns/types"
)

type ErrorCode struct {
	HTTPErrorCode int
	SubCode       types.ResponseCode
	ErrorMessage  string
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("%d: %d (%s)", e.HTTPErrorCode, e.SubCode, e.ErrorMessage)
}

func NewErrorCode(httpCode int, subCode types.ResponseCode, message string) error {
	return &ErrorCode{
		HTTPErrorCode: httpCode,
		SubCode:       subCode,
		ErrorMessage:  message,
	}
}

func IsCustomErrorCode(err error) bool {
	_, ok := err.(*ErrorCode)
	return ok
}

func GetHTTPErrorCode(err error) int {
	if customErr, ok := err.(*ErrorCode); ok {
		return customErr.HTTPErrorCode
	}

	return http.StatusInternalServerError
}

func (e *ErrorCode) HandleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, e.ErrorMessage, e.HTTPErrorCode)
	log.Printf("Error: %s", e.Error())
}
