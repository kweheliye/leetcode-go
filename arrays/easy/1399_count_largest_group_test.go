package easy

import "testing"

func Test_countLargestGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    13,
			expected: 4,
		},
		{
			name:     "Test Case 2",
			input:    2,
			expected: 2,
		},
		{
			name:     "Test Case 3",
			input:    15,
			expected: 6,
		},
		{
			name:     "Test Case 4",
			input:    24,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLargestGroup(tt.input); got != tt.expected {
				t.Errorf("countLargestGroup(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
