package medium

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		n          int
		edges      [][]int
		restricted []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reachable nodes",
			args: args{
				n: 7,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{3, 1},
					{4, 0},
					{0, 5},
					{5, 6},
				},
				restricted: []int{4, 5},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachableNodes(tt.args.n, tt.args.edges, tt.args.restricted); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
