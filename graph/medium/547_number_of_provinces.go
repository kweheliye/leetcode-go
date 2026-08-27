package medium

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)

	graph := make(map[int][]int, 0)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if isConnected[r][c] == 1 {
				graph[r] = append(graph[r], c)
			}
		}
	}

	var dfs func(city int, graph map[int][]int, visited []bool)

	dfs = func(city int, graph map[int][]int, visited []bool) {
		visited[city] = true

		for _, neighbor := range graph[city] {
			if !visited[neighbor] {
				dfs(neighbor, graph, visited)
			}
		}
	}

	provinces := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			dfs(i, graph, visited)
		}
	}

	return provinces
}
