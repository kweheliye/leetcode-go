package medium

import (
	"testing"
)

func TestCountCompleteSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    []int{1, 3, 1, 2, 2},
			expected: 4, // Subarrays [1,3,1,2], [1,3,1,2,2], [3,1,2], and [3,1,2,2]
		},
		{
			name:     "Test Case 2",
			input:    []int{5, 5, 5, 5},
			expected: 10, // All subarrays are complete because there's only 1 distinct element
		},
		{
			name:     "Test Case 3",
			input:    []int{1, 2, 3, 4},
			expected: 1, // Only the entire array is complete
		},
		{
			name:     "Test Case 4",
			input:    []int{1, 2, 1, 3, 4},
			expected: 2, // Subarrays [1,2,1,3,4] and [2,1,3,4]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCompleteSubarrays(tt.input)

			if result != tt.expected {
				t.Errorf("countCompleteSubarrays(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
