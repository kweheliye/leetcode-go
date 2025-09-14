package medium

func countComponents(n int, edges [][]int) int {
	components := 0

	adjList := make([][]int, n)

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfsConnectedComponent(i, adjList, visited)
		}
	}
	return components
}

func dfsConnectedComponent(startNode int, list [][]int, visited []bool) {
	visited[startNode] = true

	for i := 0; i < len(list[startNode]); i++ {
		if !visited[list[startNode][i]] {
			dfsConnectedComponent(list[startNode][i], list, visited)
		}
	}
}
