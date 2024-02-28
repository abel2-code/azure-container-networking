# Azure CNS HTTP Error Code Framework

## Introduction

This document describes the design of an HTTP error code framework for the Azure CNS server in AKS.  The framework aims to provide a consistent and standardized way of defining and returning errors.

## Data Model

The data model for the error framework consists of three main types: `HTTPErrorCode`, `SubCode`, and `ErrorMessage`.  The `HTTPErrorCode` type represents the HTTP status code that should be returned to the client, such as 400, 404, or 500.  The `SubCode` type represents a more specific subcategory of the error, such as `CNSServiceUnavailable`.

The following table shows some examples of the data model:

| HTTPErrorCode | SubCode                | ErrorMessage                                            |
|---------------|------------------------|---------------------------------------------------------|
| 503           | CNSServiceUnavailable  | The CNS service is not available or not ready.          |
| 507           | IPAMPoolInsufficient   | The IPAM pool does not have enough IPs to allocate.     |

## Error Definition

The error framework defines a custom error type that implements the `error` interface.  The custom error type has the following fields:

- `HTTPErrorCode`: the HTTP status code that should be returned to the client
- `SubCode`: the subcategory of the error
- `ErrorMessage`: A string that contains additional information or details about the error

The following code snippet shows how to define the custom error type:

```
type ErrorCode struct {
    HTTPErrorCode int
    SubCode       int
    ErrorMessage  string
}

func (e *errorCode) Error() string {
    return fmt.Sprintf("%d: %d (%s)",
        e.HTTPErrorCode, e.SubCode, e.ErrorMessage)
}
```

The error framework provides a set of functions and methods to create, handle, and return errors in the project.  The following are some of the main functions and methods:

- `NewErrorCode`: creates a new custom error with the given parameters
- `IsCustomErrorCode`: checks if a given error is a custom error
- `GetHTTPErrorCode`: returns the HTTP status code of a given error, or 500 if the error is not a custom error
- `HandleError`: handles a given error by logging it and returning the appropriate HTTP response to the client

The following code snippet shows how to use the error framework:

```
// Create a new custom error
err := NewErrorCode(http.StatusBadRequest, types.InvalidParameter, "Bad Request")

// Check if the error is a custom error
if IsCustomErrorCode(err) {
    // Get the HTTP status code of the error
    code := GetHTTPErrorCode(err)

    //Handle the error by logging it and returning the response
    err.HandleError(w,r)
} else {
    // Handle the error as a generic error
    log.Println(err)
    http.Error(w, http.StatusText(500), 500)
}
```

## Error Expansion

The error framework is designed to be easily expandable and maintainable.  To add new HTTP error codes and subcodes, add the new type and its corresponding value and destription to:

- the data model
- error.go
- error_test.go
- error_docs.md
- codes.go