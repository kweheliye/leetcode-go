package medium

import "testing"

func Test_minSum(t *testing.T) {

	testCases := []struct {
		name     string
		input1   []int
		input2   []int
		expected int64
	}{
		{
			name:     "Test Case 1",
			input1:   []int{3, 2, 0, 1, 0},
			input2:   []int{6, 5, 0},
			expected: 12,
		},
		{
			name:     "Test Case 2",
			input1:   []int{2, 0, 2, 0},
			input2:   []int{1, 4},
			expected: -1,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSum(tt.input1, tt.input2); got != tt.expected {
				t.Errorf("minSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}
