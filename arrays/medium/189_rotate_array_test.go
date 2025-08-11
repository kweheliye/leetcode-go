package medium

import "testing"

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		want []int
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)

		})
	}
}
