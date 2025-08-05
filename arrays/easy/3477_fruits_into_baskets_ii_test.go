package easy

import "testing"

func Test_numOfUnplacedFruitsV1(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				fruits: []int{4, 2, 5},
				baskets: []int{
					3,
					5,
					4,
				},
			},
		},

		{
			name: "Test Case 2",
			args: args{
				fruits: []int{3, 6, 1},
				baskets: []int{
					6,
					4,
					7,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruitsV1(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruitsV1() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruitsV2(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruitsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
