package helpers

import "strings"

// IsDuplicateError checks if the error is a unique constraint (duplicate) error
func IsDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	// Check if the error contains "duplicate" or "UNIQUE constraint"
	return strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE constraint")
}
