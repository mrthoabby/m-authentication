package types

type customError struct {
	Message string
}

func (e *customError) Error() string {
	return e.Message
}

func NewCustomError(message string) error {
	return &customError{
		Message: message,
	}
}
