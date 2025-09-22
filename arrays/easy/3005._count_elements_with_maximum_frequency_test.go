package easy

import "testing"

func TestMaxFrequencyElements(t *testing.T) {
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
				nums: []int{1, 2, 2, 3, 1, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElementsV1(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElementsV1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElementsV2(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElementsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
