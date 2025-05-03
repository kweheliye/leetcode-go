package medium

import "testing"

func TestMinDominoRotations(t *testing.T) {
	testCases := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected int
	}{
		{
			name:     "Test Case 1",
			inputA:   []int{2, 1, 2, 4, 2, 2},
			inputB:   []int{5, 2, 6, 2, 3, 2},
			expected: 2,
		},
		{
			name:     "Test Case 2",
			inputA:   []int{3, 5, 1, 2, 3},
			inputB:   []int{3, 6, 3, 3, 4},
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := minDominoRotations(testCase.inputA, testCase.inputB)
			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}
