package easy_test

import (
	"testing"

	"github.com/kweheliye/leetcode-go/heap/easy"
)

func TestKthLargest_StaticK(t *testing.T) {
	k := 3
	nums := []int{4, 5, 8, 2}
	kLargest := easy.Constructor(k, nums)

	addOps := []int{3, 5, 10, 9, 4}
	expected := []int{4, 5, 5, 8, 8}

	for i, val := range addOps {
		got := kLargest.Add(val)
		if got != expected[i] {
			t.Errorf("Add(%d) = %d; want %d", val, got, expected[i])
		}
	}
}
