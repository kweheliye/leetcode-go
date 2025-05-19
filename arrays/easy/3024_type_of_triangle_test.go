package easy

import "testing"

func TestTriangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{3, 3, 3},
			},
			want: "equilateral",
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{3, 4, 3},
			},
			want: "isosceles",
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{3, 4, 5},
			},
			want: "scalene",
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{9, 4, 3},
			},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
