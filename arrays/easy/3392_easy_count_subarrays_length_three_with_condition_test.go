package easy

import "testing"

func TestCountSubarrays(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    []int{1, 2, 1, 4, 1},
			expected: 1,
		},
		{
			name:     "Test Case 2",
			input:    []int{1, 1, 1},
			expected: 0,
		},
	}

	// Test both implementations with the same test cases
	t.Run("Version 1", func(t *testing.T) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				result := countSubarraysV1(testCase.input)
				if result != testCase.expected {
					t.Errorf("Expected %v, but got %v", testCase.expected, result)
				}
			})
		}
	})

	t.Run("Version 2", func(t *testing.T) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				result := countSubarraysV2(testCase.input)
				if result != testCase.expected {
					t.Errorf("Expected %v, but got %v", testCase.expected, result)
				}
			})
		}
	})
}
