package sorting

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				array: []int{3, 1, 5, 8, 2, 6, 7, 4, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
