package reverselist

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"the", "bug", "count", "also", "rises"},
			expected: []string{"rises", "also", "count", "bug", "the"},
		},
		{
			input:    []string{"ed", "is", "the", "standard", "text", "editor"},
			expected: []string{"editor", "text", "standard", "the", "is", "ed"},
		},
	}

	for _, test := range tests {
		output := ReverseList(test.input)

		if len(output) != len(test.input) {
			t.Errorf("Expected output length %d to equal %d", len(output), len(test.input))
		}

		for idx, value := range output {
			if value != test.expected[idx] {
				t.Errorf("Expected output at index %d to be %s, but received %s", idx, test.expected[idx], value)
			}
		}
	}
}
