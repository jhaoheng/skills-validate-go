// Package skillerrors defines custom error types for skill-related operations.
package skillerrors

import "fmt"

// SkillError represents a generic skill-related error.
type SkillError struct {
	Message string
}

func (e *SkillError) Error() string {
	return e.Message
}

// ParseError represents an error during parsing of SKILL.md.
type ParseError struct {
	SkillError
}

// NewParseError creates a new ParseError with the given message.
func NewParseError(message string) *ParseError {
	return &ParseError{
		SkillError: SkillError{Message: message},
	}
}

// ValidationError represents validation errors for a skill.
type ValidationError struct {
	SkillError
	Errors []string
}

// NewValidationError creates a new ValidationError with a single error message.
func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		SkillError: SkillError{Message: message},
		Errors:     []string{message},
	}
}

// NewValidationErrors creates a new ValidationError with multiple error messages.
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
