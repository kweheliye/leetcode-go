package easy

import (
	"testing"

	tree "github.com/kweheliye/leetcode-go/dfs"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *tree.TreeNode
		root2 *tree.TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same leaf sequence",
			args: args{
				root1: &tree.TreeNode{
					Val: 3,
					Left: &tree.TreeNode{
						Val:   5,
						Left:  &tree.TreeNode{Val: 6},
						Right: &tree.TreeNode{Val: 2},
					},
					Right: &tree.TreeNode{
						Val:   1,
						Left:  &tree.TreeNode{Val: 9},
						Right: &tree.TreeNode{Val: 8},
					},
				},
				root2: &tree.TreeNode{
					Val: 7,
					Left: &tree.TreeNode{
						Val: 4,
						Left: &tree.TreeNode{
							Val: 6,
						},
						Right: &tree.TreeNode{
							Val: 2,
						},
					},
					Right: &tree.TreeNode{
						Val: 1,
						Left: &tree.TreeNode{
							Val: 9,
						},
						Right: &tree.TreeNode{
							Val: 8,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "different leaf sequence",
			args: args{
				root1: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 3},
				},
				root2: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 4},
				},
			},
			want: false,
		},
		{
			name: "both nil",
			args: args{
				root1: nil,
				root2: nil,
			},
			want: true,
		},
		{
			name: "first nil",
			args: args{
				root1: nil,
				root2: &tree.TreeNode{Val: 1},
			},
			want: false,
		},
		{
			name: "second nil",
			args: args{
				root1: &tree.TreeNode{Val: 1},
				root2: nil,
			},
			want: false,
		},
		{
			name: "single nodes same value",
			args: args{
				root1: &tree.TreeNode{Val: 1},
				root2: &tree.TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "single nodes different value",
			args: args{
				root1: &tree.TreeNode{Val: 1},
				root2: &tree.TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "same leaves different root values",
			args: args{
				root1: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 3},
				},
				root2: &tree.TreeNode{
					Val:   100,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "same leaf repeated values",
			args: args{
				root1: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 2},
				},
				root2: &tree.TreeNode{
					Val:   10,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 2},
				},
			},
			want: true,
		},
		{
			name: "different order of leaves",
			args: args{
				root1: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 3},
				},
				root2: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 3},
					Right: &tree.TreeNode{Val: 2},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
