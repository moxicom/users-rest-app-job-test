package utils

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"1111 111111", true},   // valid case
		{"1234 567890", true},   // valid case
		{"11111111111", false},  // missing space
		{"1111 11111", false},   // too short
		{"1111 1111111", false}, // too long
		{"abcd 123456", false},  // non-numeric characters in the first part
		{"1234 abcdef", false},  // non-numeric characters in the second part
		{"123 1234567", false},  // first part too short
		{"12345 123456", false}, // first part too long
		{"1234 12345a", false},  // non-numeric character at the end
	}

	for _, test := range tests {
		result := ValidatePassword(test.password)
		if result != test.expected {
			t.Errorf("ValidatePassword(%q) = %v; expected %v", test.password, result, test.expected)
		}
	}
}
