package easy

import "testing"

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "TestCase1",
			input:    "aba",
			expected: true,
		},
		{
			name:     "TestCase2",
			input:    "abca",
			expected: true,
		},
		{
			name:     "TestCase3",
			input:    "abc",
			expected: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := validPalindrome(testCase.input)

			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}
