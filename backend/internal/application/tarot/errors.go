package tarot

import "errors"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func validationError(message string) error {
	return ValidationError{Message: message}
}

func IsValidationError(err error) bool {
	var target ValidationError
	return errors.As(err, &target)
}
