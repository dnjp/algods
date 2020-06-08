package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "lsdkjfskf",
			expected: false,
		},
		{
			input:    "radar",
			expected: true,
		},
	}

	for _, test := range tests {
		out := isPalindrome(test.input)
		if out != test.expected {
			t.Errorf("Expected %t to be %t", out, test.expected)
		}
	}
}
