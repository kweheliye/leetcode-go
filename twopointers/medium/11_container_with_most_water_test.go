package medium

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaV1(tt.args.height); got != tt.want {
				t.Errorf("maxAreaV1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaV2(tt.args.height); got != tt.want {
				t.Errorf("maxAreaV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
