package easy

import (
	"container/heap"

	"github.com/kweheliye/leetcode-go/heap/int_heap"
)

func lastStoneWeight(stones []int) int {
	// Initialize a MaxHeap with no elements
	maxHeap := int_heap.NewMaxHeap([]int{})

	// Push all stones into the heap
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	// Pop two largest stones and push the difference if needed
	for maxHeap.Len() > 1 {
		stone1 := heap.Pop(maxHeap).(int)
		stone2 := heap.Pop(maxHeap).(int)

		if stone1 != stone2 {
			heap.Push(maxHeap, stone1-stone2)
		}
	}

	// Return the last stone or 0 if none left
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(maxHeap).(int)
}
