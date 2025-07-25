package easy

import "testing"

func TestMaxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Multiple elements with duplicates",
			args: args{
				nums: []int{1, 2, -1, -2, 1, 0, -1},
			},
			want: 3,
		},
		{
			name: "Test Case 2: All unique elements",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "Test Case 3: Multiple duplicates",
			args: args{
				nums: []int{1, 1, 0, 1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
