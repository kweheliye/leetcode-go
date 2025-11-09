package easy

import "testing"

func TestCountOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				num1: 2,
				num2: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
