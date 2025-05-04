package easy

import (
	"testing"
)

func TestNumEquivDominoPairs(test *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "TestCase 1",
			input:    [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
			expected: 1,
		},
		{
			name:     "TestCase 2",
			input:    [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}},
			expected: 3,
		},
	}

	for _, testCase := range testCases {

		test.Run(testCase.name, func(t *testing.T) {
			result := numEquivDominoPairs(testCase.input)

			if result != testCase.expected {
				t.Errorf("Expected: %v but got : %v", testCase.expected, result)
			}
		})
	}

}
