package medium

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
