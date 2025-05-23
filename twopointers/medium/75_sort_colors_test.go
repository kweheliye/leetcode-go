package medium

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	testCases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestCase1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		}, {
			name: "TestCase2",
			args: args{
				nums: []int{2, 0, 1},
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sortColors(tc.args.nums)
			if !reflect.DeepEqual(tc.args.nums, tc.want) {
				t.Errorf("expected: %v, want: %v", tc.args, tc.want)
			}
		})
	}
}
