package easy

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Merge with overlap",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Nums2 smaller elements",
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Nums1 empty, only nums2",
			args: args{
				nums1: []int{0, 0, 0},
				m:     0,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			expected: []int{2, 5, 6},
		},
		{
			name: "Nums2 empty, only nums1",
			args: args{
				nums1: []int{1, 2, 3},
				m:     3,
				nums2: []int{},
				n:     0,
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "All elements equal",
			args: args{
				nums1: []int{2, 2, 2, 0, 0, 0},
				m:     3,
				nums2: []int{2, 2, 2},
				n:     3,
			},
			expected: []int{2, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.expected) {
				t.Errorf("merge() = %v, expected %v", tt.args.nums1, tt.expected)
			}
		})
	}
}
