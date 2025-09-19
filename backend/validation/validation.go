package validation

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// Validator provides validation functions
type Validator struct{}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateString validates a string field
func (v *Validator) ValidateString(value, fieldName string, required bool, maxLength int) error {
	if required && strings.TrimSpace(value) == "" {
		return &ValidationError{Field: fieldName, Message: "is required"}
	}

	if maxLength > 0 && utf8.RuneCountInString(value) > maxLength {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be no more than %d characters", maxLength)}
	}

	return nil
}

// ValidateEmail validates an email address
func (v *Validator) ValidateEmail(email, fieldName string, required bool) error {
	if required && strings.TrimSpace(email) == "" {
		return &ValidationError{Field: fieldName, Message: "is required"}
	}

	if email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(email) {
			return &ValidationError{Field: fieldName, Message: "must be a valid email address"}
		}
	}

	return nil
}

// ValidateURL validates a URL
func (v *Validator) ValidateURL(urlStr, fieldName string, required bool) error {
	if required && strings.TrimSpace(urlStr) == "" {
		return &ValidationError{Field: fieldName, Message: "is required"}
	}

	if urlStr != "" {
		if _, err := url.Parse(urlStr); err != nil {
			return &ValidationError{Field: fieldName, Message: "must be a valid URL"}
		}
	}

	return nil
}

// ValidatePositiveInt validates a positive integer
func (v *Validator) ValidatePositiveInt(value int, fieldName string, required bool) error {
	if required && value == 0 {
		return &ValidationError{Field: fieldName, Message: "is required"}
	}

	if value < 0 {
		return &ValidationError{Field: fieldName, Message: "must be a positive number"}
	}

	return nil
}

// ValidateStringSlice validates a slice of strings
func (v *Validator) ValidateStringSlice(values []string, fieldName string, required bool, maxItems int) error {
	if required && len(values) == 0 {
		return &ValidationError{Field: fieldName, Message: "is required"}
	}

	if maxItems > 0 && len(values) > maxItems {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must have no more than %d items", maxItems)}
	}

	// Validate each string in the slice
	for i, value := range values {
		if err := v.ValidateString(value, fmt.Sprintf("%s[%d]", fieldName, i), true, 500); err != nil {
			return err
		}
	}

	return nil
}

// SanitizeString sanitizes a string by trimming whitespace and removing potentially dangerous characters
func (v *Validator) SanitizeString(input string) string {
	// Trim whitespace
	output := strings.TrimSpace(input)

	// Remove null bytes and control characters
	output = strings.ReplaceAll(output, "\x00", "")
	output = strings.ReplaceAll(output, "\r", "")
	output = strings.ReplaceAll(output, "\n", " ")
	output = strings.ReplaceAll(output, "\t", " ")

	// Normalize multiple spaces to single space
	spaceRegex := regexp.MustCompile(`\s+`)
	output = spaceRegex.ReplaceAllString(output, " ")

	return output
}

// SanitizeHTML sanitizes HTML content by removing potentially dangerous tags
func (v *Validator) SanitizeHTML(input string) string {
	// Remove script tags and their content
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)
	output := scriptRegex.ReplaceAllString(input, "")

	// Remove style tags and their content
	styleRegex := regexp.MustCompile(`(?i)<style[^>]*>.*?</style>`)
	output = styleRegex.ReplaceAllString(output, "")

	// Remove potentially dangerous attributes
	dangerousAttrRegex := regexp.MustCompile(`(?i)\s+(on\w+|javascript:|data:|vbscript:)[^=]*="[^"]*"`)
	output = dangerousAttrRegex.ReplaceAllString(output, "")

	return output
}
