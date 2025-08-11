package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty array",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "Single element array",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "Already sorted array",
			args: args{
				nums: []int{9, 1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5, 9},
		},
		{
			name: "Reverse sorted array",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Array with duplicates",
			args: args{
				nums: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			},
			want: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.nums)
			for i := range tt.args.nums {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("bubbleSort() = %v, want %v", tt.args.nums, tt.want)
					break
				}
			}
		})
	}
}
