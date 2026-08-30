package medium

import "testing"

func Test_minimumDeletions2091(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 10, 7, 5, 4, 1, 8, 6},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions2091(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeletions2091() = %v, want %v", got, tt.want)
			}
		})
	}
}
