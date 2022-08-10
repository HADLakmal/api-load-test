package errors

import "fmt"

type errorType int

const (
	DATA_ERROR errorType = iota
	MIDDLEWARE_ERROR
	SERVER_ERROR
	SERVICE_ERROR
	VALIDATION_ERROR
	DOMAIN_ERROR
	ADAPTER_ERROR
	UNKNOWN_ERROR
)

// Error is the type of errors thrown by the application.
type Error struct {
	Type    errorType
	Msg     string
	Code    int
	Details string
	Err     error
}

// WithPrevious creates a new error which holds the previous error information
func WithPrevious(previous error, errType errorType, code int, message string, details string) error {

	err := &Error{}

	err.Msg = message
	err.Code = code
	err.Details = details
	err.Type = errType
	err.Err = previous

	return err
}

// New creates a new DataError instance.
func New(errType errorType, code int, message string, details string) error {

	err := &Error{}

	err.Msg = message
	err.Code = code
	err.Details = details
	err.Type = errType
	err.Err = nil

	return err
}

// Error returns the DataError message.
func (e *Error) Error() string {
	return fmt.Sprintf("%s|%d|Error|%s", e.Msg, e.Code, e.Details)
}

func (e *Error) Unwrap() error { return e.Err }
