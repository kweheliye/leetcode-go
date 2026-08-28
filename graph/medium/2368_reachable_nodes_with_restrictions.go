package medium

func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make(map[int][]int)
	visited := make([]bool, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var dfs func(node int) int

	dfs = func(node int) int {
		if isRestricted(node, restricted) {
			return 0
		}
		path := 1
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				path += dfs(neighbor)
			}
		}

		return path

	}

	return dfs(0)
}

func isRestricted(num int, restricted []int) bool {
	for _, res := range restricted {
		if res == num {
			return true
		}
	}

	return false
}
