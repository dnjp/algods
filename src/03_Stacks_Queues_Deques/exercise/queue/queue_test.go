package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {

	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"the", "cult", "of", "vi"},
			expected: []string{"the", "cult", "of", "vi"},
		},
	}

	for _, test := range tests {
		q := New()

		for _, input := range test.input {
			q.Enqueue(input)
		}

		if q.IsEmpty() {
			t.Errorf("Expected queue to not be empty")
		}

		for _, expected := range test.expected {
			out := q.Dequeue()
			val := *out
			valStr := val.(string)
			if valStr != expected {
				t.Errorf("Expected %s to be %s", valStr, expected)
			}
		}
	}
}
