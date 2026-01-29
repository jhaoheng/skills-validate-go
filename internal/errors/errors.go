// Package errors defines custom error types for skill-related operations.
package errors

import "fmt"

type SkillError struct {
	Message string
}

func (e *SkillError) Error() string {
	return e.Message
}

type ParseError struct {
	SkillError
}

func NewParseError(message string) *ParseError {
	return &ParseError{
		SkillError: SkillError{Message: message},
	}
}

type ValidationError struct {
	SkillError
	Errors []string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		SkillError: SkillError{Message: message},
		Errors:     []string{message},
	}
}

func NewValidationErrors(errors []string) *ValidationError {
	message := fmt.Sprintf("validation failed with %d error(s)", len(errors))
	if len(errors) > 0 {
		message = errors[0]
	}
	return &ValidationError{
		SkillError: SkillError{Message: message},
		Errors:     errors,
	}
}
