package medium

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{-4, 0, 7, -2, 3},
			},
			want: []int{-4, -2, 0, 3, 7},
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
