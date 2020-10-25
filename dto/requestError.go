package dto

import (
	"errors"
	"strings"
)

type RequestError struct {
	HTTPCode int
	Error    error
}

func InvalidRequestDataError() *RequestError {
	return &RequestError{
		HTTPCode: 400,
		Error:    errors.New("Invalid request data"),
	}
}

func ExternalServiceError() *RequestError {
	return &RequestError{
		HTTPCode: 500,
		Error:    errors.New("Error calling external service"),
	}
}

func OsrmServiceError(msg string) *RequestError {
	var sb strings.Builder
	sb.WriteString("Osrm service error: ")
	sb.WriteString(msg)

	return &RequestError{
		HTTPCode: 500,
		Error:    errors.New(sb.String()),
	}
}
