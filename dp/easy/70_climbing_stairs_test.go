package easy

import "testing"

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    4,
			expected: 5,
		},
		{
			name:     "Test Case 2",
			input:    5,
			expected: 8,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := climbStairsTopDown(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
		t.Run(testCase.name, func(t *testing.T) {
			result := climbStairsBottomUp(testCase.input)
			if result != testCase.expected {
			}
		})
	}
}
