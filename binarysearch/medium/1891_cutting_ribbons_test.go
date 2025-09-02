package medium

import "testing"

func TestMaxLength(t *testing.T) {
	type args struct {
		ribbons []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				ribbons: []int{7, 5, 9},
				k:       4,
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				ribbons: []int{9, 7, 5},
				k:       3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.ribbons, tt.args.k); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
