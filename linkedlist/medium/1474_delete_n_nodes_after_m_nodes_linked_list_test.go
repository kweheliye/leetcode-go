package medium

import (
	"github.com/kweheliye/leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestDeleteNodes(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				head: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val:  3,
							Next: nil}}},
				m: 1,
				n: 1,
			},
			want: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}),
				m:    2,
				n:    3,
			},
			want: createLinkedList([]int{1, 2, 6, 7, 11, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNodes(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to create a linked list from slice
func createLinkedList(nums []int) *linkedlist.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &linkedlist.ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &linkedlist.ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}
