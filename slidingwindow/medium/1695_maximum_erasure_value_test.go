package medium

import "testing"

func TestMaximumUniqueSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{4, 2, 4, 5, 6},
			},
			want: 17,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
