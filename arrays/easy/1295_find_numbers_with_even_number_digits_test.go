package easy

import (
	"testing"
)

func TestFindNumbersV1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    []int{12, 345, 2, 6, 7896},
			expected: 2,
		},
		{
			name:     "Test Case 2",
			input:    []int{555, 901, 482, 1771},
			expected: 1,
		},
	}

	t.Run("Version 1", func(t *testing.T) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				result := findNumbersV1(testCase.input)
				if result != testCase.expected {
					t.Errorf("Expected %v, but got %v", testCase.expected, result)
				}
			})
		}
	})
	
	t.Run("Version 2", func(t *testing.T) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				result := findNumbersV2(testCase.input)
				if result != testCase.expected {
					t.Errorf("Expected %v, but got %v", testCase.expected, result)
				}
			})
		}
	})

}
