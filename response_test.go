package main

import (
	"testing"

	"github.com/dmitrio95/go-test-task/response"
)

// Ensure that we have at least one formatter available
func TestFormatterExists(t *testing.T) {
	const msg = "No response formatters available!"

	if len(response.Formats) < 1 {
		t.Error(msg)
	}
}

// Ensure that all formatters listed in response.Formats
// actually exist
func TestFormattersAvailable(t *testing.T) {
	for format, _ := range response.Formats {
		fmt := response.NewResponseFormatter(format)
		if fmt == nil {
			t.Error("Formatter is listed but not available:", format)
		}
	}
}

// Ensure that the formatter specified by
// response.DefaultFormat exists
func TestDefaultFormatterAvailable(t *testing.T) {
	fmt := response.NewResponseFormatter(response.DefaultFormat)
	if fmt == nil {
		t.Error("Default formatter is not available:", response.DefaultFormat)
	}
}
