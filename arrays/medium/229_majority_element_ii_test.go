package medium

import (
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 1, 1, 2, 3},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
