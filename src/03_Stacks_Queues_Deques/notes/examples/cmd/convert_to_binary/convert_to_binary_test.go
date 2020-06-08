package main

import (
	"testing"
)

func TestCovertToBinary(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{
			input:    42,
			expected: "101010",
		},
	}
	for _, test := range tests {
		out := convertToBinary(test.input)
		if out != test.expected {
			t.Errorf("Expected %s to be %s", out, test.expected)
		}
	}
}
