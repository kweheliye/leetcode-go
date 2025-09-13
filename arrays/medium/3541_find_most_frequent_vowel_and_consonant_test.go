package medium

import "testing"

func TestIsVowelByte(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				c: 'a',
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				c: 'b',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVowelByte(tt.args.c); got != tt.want {
				t.Errorf("isVowelByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFreqSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "successes",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreqSum(tt.args.s); got != tt.want {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
