package medium

import (
	"testing"
)

func TestMaximumTripletValueGreedy(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int64
	}{
		{
			name:     "Test Case 1",
			input:    []int{12, 6, 1, 2, 7},
			expected: 77, // Expected result
		},
		{
			name:     "Test Case 2",
			input:    []int{1, 10, 3, 4, 19},
			expected: 133, // Expected result
		},
		{
			name:     "Test Case 3",
			input:    []int{1, 2, 3},
			expected: 0, // Expected result
		},
		{
			name:     "Test Case with Negative Numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: 0, // Expected result
		},
		{
			name:     "Test Case with Mixed Numbers",
			input:    []int{-5, 10, -3, 4, -2},
			expected: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximumTripletValueGreedy(tt.input)

			if result != tt.expected {
				t.Errorf("MaximumTripletValueGreedy(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
