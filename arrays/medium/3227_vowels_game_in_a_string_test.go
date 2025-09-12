package medium

import "testing"

func TestDoesAliceWin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "leetcode",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "bbcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesAliceWin(tt.args.s); got != tt.want {
				t.Errorf("doesAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
