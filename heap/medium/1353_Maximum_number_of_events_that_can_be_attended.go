package medium

import (
	"container/heap"
	"github.com/kweheliye/leetcode-go/heap/intheap"
	"sort"
)

func maxEvents(events [][]int) int {
	n := len(events)

	var maxDay int
	for _, event := range events {
		maxDay = max(maxDay, event[1])
	}

	sort.Slice(events, func(a, b int) bool {
		return events[a][0] < events[b][0]
	})

	pq := intheap.NewIntHeap([]int{})
	ans := 0
	for i, j := 1, 0; i <= maxDay; i++ {
		for j < n && events[j][0] == i {
			heap.Push(pq, events[j][1])
			j++
		}

		for pq.Len() > 0 && (*pq)[0] < i {
			heap.Pop(pq)
		}

		if pq.Len() > 0 {
			heap.Pop(pq)
			ans++
		}
	}
	return ans
}
