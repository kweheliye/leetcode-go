package medium

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}
