package medium

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "Test Case 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "Test Case 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := insert(tt.intervals, tt.newInterval)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("insert(%v, %v) = %v; want %v", tt.intervals, tt.newInterval, res, tt.expected)
			}
		})
	}
}
