package medium

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	testCases := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Test Case 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Test Case 2",
			intervals: [][]int{{1, 4}, {4, 5}, {5, 6}, {6, 8}, {2, 7}},
			expected:  [][]int{{1, 8}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := merge(testCase.intervals)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}

		})
	}

}
