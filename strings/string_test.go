package strings_test

import (
	"bytes"
	"testing"

	"github.com/vinaycharlie01/godsapkg/strings"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "empty string",
			input:    []byte(""),
			expected: []byte(""),
		},
		{
			name:     "single character",
			input:    []byte("a"),
			expected: []byte("a"),
		},
		{
			name:     "two characters",
			input:    []byte("ab"),
			expected: []byte("ba"),
		},
		{
			name:     "word",
			input:    []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			name:     "palindrome",
			input:    []byte("madam"),
			expected: []byte("madam"),
		},
		{
			name:     "with spaces",
			input:    []byte("a b c"),
			expected: []byte("c b a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy input so we don’t modify the original test case slice
			got := append([]byte{}, tt.input...)
			strings.ReverseString(got)

			if !bytes.Equal(got, tt.expected) {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
