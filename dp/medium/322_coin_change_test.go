package medium

import "testing"

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		input2   int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    []int{1, 2, 5},
			input2:   11,
			expected: 3,
		},
		{
			name:     "Test Case 2",
			input:    []int{2},
			input2:   3,
			expected: -1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := coinChange(testCase.input, testCase.input2)

			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}

		})
	}
}
