package main

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{
			input:    "((()))",
			expected: true,
		},
		{
			input:    "(()",
			expected: false,
		},
		{
			input:    "())",
			expected: false,
		},
	}
	for _, test := range tests {
		out := isBalanced(test.input)
		if out != test.expected {
			t.Errorf("Expected %s to be %t", test.input, test.expected)
		}
	}
}
