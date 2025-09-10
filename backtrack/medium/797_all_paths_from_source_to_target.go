package medium

func allPathsSourceTarget(graph [][]int) [][]int {

	var result [][]int
	var path []int
	path = append(path, 0)
	backtrack(0, graph, path, &result)

	return result
}

func backtrack(currentNode int, graph [][]int, path []int, result *[][]int) {

	if currentNode == len(graph)-1 {
		// Make a copy of path and append
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for _, nextNode := range graph[currentNode] {
		path = append(path, nextNode)
		backtrack(nextNode, graph, path, result)
		path = path[:len(path)-1]
	}
}
