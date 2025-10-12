package int_heap

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] } // max-int_heap
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// NewMaxHeap initializes and returns a pointer to a MaxHeap with the given data.
func NewMaxHeap(data []int) *MaxHeap {
	h := MaxHeap(data)
	heap.Init(&h)
	return &h
}
