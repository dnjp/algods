package main

import (
	"testing"
)

func TestHotPotato(t *testing.T) {
	tests := []struct {
		names    []string
		num      int
		expected string
	}{
		{
			names:    []string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"},
			num:      9,
			expected: "David",
		},
	}

	for _, test := range tests {
		out := hotPotato(test.names, test.num)
		if out != test.expected {
			t.Errorf("Expected %s to be %s", out, test.expected)
		}
	}
}
