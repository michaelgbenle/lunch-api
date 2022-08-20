package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// FieldError wraps around the validator error so it
// can be used and caught specifically
type FieldError struct {
	err validator.FieldError
}

func (q FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}

// NewFieldError returns a field error
func NewFieldError(err validator.FieldError) FieldError {
	return FieldError{err: err}
}

type InActiveUserError struct {
	err string
}

func (i InActiveUserError) Error() string {
	if i.err == "" {
		i.err = "user is inactive"
	}
	return i.err
}
