package medium

import "testing"

func TestPartitionArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{3, 6, 1, 2, 5},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("partitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
