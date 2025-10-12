package easy

import (
	"container/heap"

	"github.com/kweheliye/leetcode-go/heap/int_heap"
)

type KthLargest struct {
	k       int
	minHeap *int_heap.MinHeap
}

func Constructor(k int, nums []int) KthLargest {

	kLargest := KthLargest{
		k:       k,
		minHeap: int_heap.NewMinHeap([]int{}),
	}

	for i := 0; i < len(nums); i++ {
		heap.Push(kLargest.minHeap, nums[i])

		if kLargest.minHeap.Len() > k {
			heap.Pop(kLargest.minHeap)
		}
	}

	return kLargest

}

func (kLargest *KthLargest) Add(val int) int {
	heap.Push(kLargest.minHeap, val)

	if kLargest.minHeap.Len() > kLargest.k {
		heap.Pop(kLargest.minHeap)
	}

	return (*kLargest.minHeap)[0]
}
