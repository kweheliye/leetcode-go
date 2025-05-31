package medium

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected bool
	}{
		// Test case 1: Strings are close (same length, same set of characters, same frequency distribution)
		{
			word1:    "abc",
			word2:    "bca",
			expected: true,
		},
		// Test case 2: Strings are close (different characters have same frequencies)
		{
			word1:    "cabbba",
			word2:    "abbccc",
			expected: true,
		},
		// Test case 3: Strings have different lengths
		{
			word1:    "abc",
			word2:    "abcd",
			expected: false,
		},
		// Test case 4: Strings have different sets of characters
		{
			word1:    "abc",
			word2:    "def",
			expected: false,
		},
		// Test case 5: Strings have same set of characters but different frequency distributions
		{
			word1:    "aabbc",
			word2:    "aaabc",
			expected: false,
		},
		// Test case 6: Strings have different sets of characters but same length
		{
			word1:    "abcde",
			word2:    "aacde",
			expected: false,
		},
	}

	for i, tc := range testCases {
		result := closeStrings(tc.word1, tc.word2)
		if result != tc.expected {
			t.Errorf("Test case %d: closeStrings(%q, %q) = %v, expected %v",
				i+1, tc.word1, tc.word2, result, tc.expected)
		}
	}
}