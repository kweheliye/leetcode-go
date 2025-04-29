package medium

import "testing"

func TestCountSubarrays(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		k        int
		expected int64
	}{
		{
			name:     "Test Case 1",
			input:    []int{1, 3, 2, 3, 3},
			k:        2,
			expected: 6,
		},
		{
			name:     "Test Case 2",
			input:    []int{1, 4, 2, 1},
			k:        3,
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := countSubarrays(testCase.input, testCase.k)

			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}
