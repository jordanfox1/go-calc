package custom_error

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Message string
}

// Error returns the error message for CustomError.
func (e *CustomError) Error() string {
	return e.Message
}
