package medium

import (
	"reflect"
	"testing"
)

func TestSortMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
