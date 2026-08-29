package easy

import (
	"slices"

	tree "github.com/kweheliye/leetcode-go/dfs"
)

func leafSimilar(root1 *tree.TreeNode, root2 *tree.TreeNode) bool {
	result1 := []int{}
	result2 := []int{}

	var dfs func(node *tree.TreeNode, result *[]int)

	dfs = func(node *tree.TreeNode, result *[]int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			*result = append(*result, node.Val)
		}

		dfs(node.Left, result)
		dfs(node.Right, result)

	}

	dfs(root1, &result1)
	dfs(root2, &result2)

	return slices.Equal(result1, result2)
}
