package github

import "fmt"

// RateError represents a case where the amount of GitHub url calls allowed per hour was exceeded
type RateError struct {
	message string
	error   error // Embed the standard error interface
}

// createRateError
/*
@brief:
	createRateError creates a new RateError object
@returns:
	error - the new error
*/
func createRateError() *RateError {
	message := "rate limit exceeded, please try again in an hour or open an issue manually here: https://placeholder.com"
	return &RateError{message: message, error: fmt.Errorf(message)}
}

// Error
/*
@brief:
	Error implementation of the error interface
@returns:
	string - the error itself
*/
func (e RateError) Error() string {
	return e.message
}

// Unwrap
/*
@brief:
	Unwrap returns the wrapped error, if any.
@returns:
	error
*/
func (e *RateError) Unwrap() error {
	return e.error
}
