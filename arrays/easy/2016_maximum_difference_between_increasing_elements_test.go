package easy

import "testing"

func TestMaximumDifference(t *testing.T) {
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
				nums: []int{7, 1, 5, 4},
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
