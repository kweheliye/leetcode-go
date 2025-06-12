package easy

import "testing"

func TestMaxAdjacentDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case1",
			args: args{
				nums: []int{1, 2, 4},
			},
			want: 3,
		},
		{
			name: "Test Case2",
			args: args{
				nums: []int{-5, -10, -5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAdjacentDistance(tt.args.nums); got != tt.want {
				t.Errorf("maxAdjacentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
